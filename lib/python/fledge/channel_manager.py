import asyncio
import atexit

from .backends import backend_provider
from .channel import Channel
from .common.constants import SOCK_OP_WAIT_TIME, BackendEvent
from .common.util import background_thread_loop, run_async
from .registry_agents import registry_agent_provider


class ChannelManager(object):
    '''
    ChannelManager only allows a singleton instance
    '''
    _instance = None

    _job = None
    _role = None
    _channels_roles = None

    _channels = None

    _loop = None

    _backend = None
    _registry_agent = None

    def __new__(cls):
        if cls._instance is None:
            print('creating a ChannelManager instance')
            cls._instance = super(ChannelManager, cls).__new__(cls)
        return cls._instance

    def __call__(self, backend, registry_agent, job, role, channels_roles):
        self._job = job
        self._role = role
        self._channels_roles = channels_roles

        self._channels = {}

        with background_thread_loop() as loop:
            self._loop = loop

        self._backend = backend_provider.get(backend)
        self._registry_agent = registry_agent_provider.get(registry_agent)

        async def inner():
            # create a coroutine task
            coro = self._backend_eventq_task(self._backend.eventq())
            _ = asyncio.create_task(coro)

        _ = run_async(inner(), self._backend.loop())

        atexit.register(self.cleanup)

    async def _backend_eventq_task(self, eventq):
        while True:
            (event_type, info) = await eventq.get()

            if event_type == BackendEvent.DISCONNECT:
                for _, channel in self._channels.items():
                    await channel.remove(info)

    def join(self, name):
        '''
        joins a channel
        '''
        if self.is_joined(name):
            return True

        coro = self._registry_agent.connect()
        _, status = run_async(coro, self._loop, SOCK_OP_WAIT_TIME)
        if not status:
            return False

        coro = self._registry_agent.register(
            self._job, name, self._role, self._backend.uid(),
            self._backend.endpoint()
        )
        _, status = run_async(coro, self._loop, SOCK_OP_WAIT_TIME)
        if status:
            self._channels[name] = Channel(name, self._backend)
            self._backend.add_channel(self._channels[name])
        else:
            return False

        # role_tuples should have at most two entries
        role_tuples = self._channels_roles[name]

        coro = self._registry_agent.get(self._job, name)
        channel_info, status = run_async(coro, self._loop, SOCK_OP_WAIT_TIME)
        if not status:
            return False

        for role, end_id, endpoint in channel_info:
            # the same backend id; skip
            if end_id is self._backend.uid():
                continue

            proceed = False
            for from_role, to_role in role_tuples:
                proceed = self._role == from_role and role == to_role
                if proceed:
                    break
            if not proceed:
                continue

            # connect to endpoint
            self._backend.connect(end_id, endpoint)

            # notify end_id of the channel handled by the backend
            self._backend.notify(end_id, name)

            # update channel
            coro = self._channels[name].add(end_id)
            _ = run_async(coro, self._backend.loop())

        coro = self._registry_agent.close()
        _ = run_async(coro, self._loop, SOCK_OP_WAIT_TIME)

        return True

    def leave(self, name):
        '''
        leave a channel
        '''
        if not self.is_joined(name):
            return

        coro = self._registry_agent.reset_channel(
            self._job, name, self._role, self._backend.uid()
        )

        _, status = run_async(coro, self._loop, SOCK_OP_WAIT_TIME)
        if status:
            del self._channels[name]

        return status

    def get(self, name):
        '''
        returns a channel object in the given channel
        '''
        if not self.is_joined(name):
            # didn't join the channel yet
            return None

        return self._channels[name]

    def is_joined(self, name):
        '''
        check if node joined a channel or not
        '''
        if name in self._channels:
            return True
        else:
            return False

    def cleanup(self):
        for _, channel in self._channels.items():
            channel.cleanup()
