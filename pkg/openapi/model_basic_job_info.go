// Copyright 2022 Cisco Systems, Inc. and its affiliates
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// SPDX-License-Identifier: Apache-2.0
/*
 * Flame REST API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// BasicJobInfo - Basic Job specification
type BasicJobInfo struct {
	Id string `json:"id,omitempty"`

	UserId string `json:"userId,omitempty"`

	DesignId string `json:"designId"`

	SchemaVersion string `json:"schemaVersion"`

	CodeVersion string `json:"codeVersion"`

	Priority JobPriority `json:"priority,omitempty"`

	MaxRunTime int32 `json:"maxRunTime,omitempty"`
}

// AssertBasicJobInfoRequired checks if the required fields are not zero-ed
func AssertBasicJobInfoRequired(obj BasicJobInfo) error {
	elements := map[string]interface{}{
		"designId":      obj.DesignId,
		"schemaVersion": obj.SchemaVersion,
		"codeVersion":   obj.CodeVersion,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	return nil
}

// AssertRecurseBasicJobInfoRequired recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of BasicJobInfo (e.g. [][]BasicJobInfo), otherwise ErrTypeAssertionError is thrown.
func AssertRecurseBasicJobInfoRequired(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aBasicJobInfo, ok := obj.(BasicJobInfo)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertBasicJobInfoRequired(aBasicJobInfo)
	})
}
