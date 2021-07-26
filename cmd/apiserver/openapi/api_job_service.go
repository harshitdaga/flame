/*
 * Job REST API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

import (
	"context"
	"errors"
	"net/http"

	"go.uber.org/zap"

	"wwwin-github.cisco.com/eti/fledge/pkg/objects"
	"wwwin-github.cisco.com/eti/fledge/pkg/util"
)

// JobApiService is a service that implents the logic for the JobApiServicer
// This service should implement the business logic for every endpoint for the JobApi API.
// Include any external packages or services that will be required by this service.
type JobApiService struct {
}

// NewJobApiService creates a default api service
func NewJobApiService() JobApiServicer {
	return &JobApiService{}
}

// DeleteJob - Delete job by id.
func (s *JobApiService) DeleteJob(ctx context.Context, user string, jobId string) (ImplResponse, error) {
	// TODO - update DeleteJob with the required logic for this service method.
	// Add api_job_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	//TODO: Uncomment the next line to return response Response(200, {}) or use other options such as http.Ok ...
	//return Response(200, nil),nil

	//TODO: Uncomment the next line to return response Response(404, {}) or use other options such as http.Ok ...
	//return Response(404, nil),nil

	//TODO: Uncomment the next line to return response Response(401, {}) or use other options such as http.Ok ...
	//return Response(401, nil),nil

	//TODO: Uncomment the next line to return response Response(0, Error{}) or use other options such as http.Ok ...
	//return Response(0, Error{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("DeleteJob method not implemented")
}

// GetJob - Get job detail.
func (s *JobApiService) GetJob(ctx context.Context, user string, jobId string) (ImplResponse, error) {
	//TODO - validate the input
	zap.S().Debugf("Get job details recieved for user: %s | jobId: %v", user, jobId)

	//create controller request
	uriMap := map[string]string{
		"user":  user,
		"jobId": jobId,
	}
	url := CreateURI(util.GetJobEndPoint, uriMap)

	//send get request
	responseBody, err := util.HTTPGet(url)

	//response to the user
	if err != nil {
		return Response(http.StatusInternalServerError, nil), errors.New("get job details request failed")
	}
	resp := objects.JobInfo{}
	err = util.ByteToStruct(responseBody, &resp)
	return Response(http.StatusOK, resp), err
}

// SubmitJob - Submit a new job.
func (s *JobApiService) SubmitJob(ctx context.Context, user string, jobInfo objects.JobInfo) (ImplResponse, error) {
	//TODO - validate the input
	zap.S().Debugf("New job request recieved for user: %s | jobInfo: %v", user, jobInfo)

	//create controller request
	jobInfo.UserId = user
	uriMap := map[string]string{
		"user": user,
	}
	url := CreateURI(util.SubmitJobEndPoint, uriMap)

	//send post request
	code, responseBody, err := util.HTTPPost(url, jobInfo, "application/json")

	//response to the user
	if err != nil {
		return Response(http.StatusInternalServerError, nil), errors.New("submit new job request failed")
	}

	if code == http.StatusMultiStatus {
		resp := map[string]interface{}{}
		err := util.ByteToStruct(responseBody, &resp)
		if err != nil {
			zap.S().Errorf("error sending out mutlistatus response %v", err)
		}
		return Response(http.StatusMultiStatus, resp), nil
	}
	return Response(http.StatusCreated, responseBody), nil
}

// UpdateJob - Update job by id.
func (s *JobApiService) UpdateJob(ctx context.Context, user string, jobId string, jobInfo objects.JobInfo) (ImplResponse, error) {
	// TODO - update UpdateJobEndPoint with the required logic for this service method.
	// Add api_job_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	//TODO: Uncomment the next line to return response Response(200, {}) or use other options such as http.Ok ...
	//return Response(200, nil),nil

	//TODO: Uncomment the next line to return response Response(0, Error{}) or use other options such as http.Ok ...
	//return Response(0, Error{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("UpdateJobEndPoint method not implemented")
}