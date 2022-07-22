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

package controller

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/cisco-open/flame/cmd/controller/app/database"
	"github.com/cisco-open/flame/pkg/openapi"
)

// ComputesApiService is a service that implements the logic for the ComputesApiServicer
// This service should implement the business logic for every endpoint for the ComputesApi API.
// Include any external packages or services that will be required by this service.
type ComputesApiService struct {
	dbService database.DBService
}

// NewComputesApiService creates a default api service
func NewComputesApiService(dbService database.DBService) openapi.ComputesApiServicer {
	return &ComputesApiService{
		dbService: dbService,
	}
}

// AddDeploymentStatus - Add the deployment status for a job on a compute cluster
func (s *ComputesApiService) AddDeploymentStatus(ctx context.Context, computeId string,
	jobId string, deploymentStatus openapi.DeploymentStatus) (openapi.ImplResponse, error) {
	// TODO - update AddDeploymentStatus with the required logic for this service method.
	// Add api_computes_service.go to the .openapi-generator-ignore to avoid overwriting
	// this service implementation when updating open api generation.

	//TODO: Uncomment the next line to return response Response(200, {}) or use other options such as http.Ok ...
	//return Response(200, nil),nil

	//TODO: Uncomment the next line to return response Response(401, {}) or use other options such as http.Ok ...
	//return Response(401, nil),nil

	//TODO: Uncomment the next line to return response Response(0, Error{}) or use other options such as http.Ok ...
	//return Response(0, Error{}), nil

	return openapi.Response(http.StatusNotImplemented, nil), errors.New("AddDeploymentStatus method not implemented")
}

// DeleteCompute - Delete compute cluster specification
func (s *ComputesApiService) DeleteCompute(ctx context.Context, computeId string) (openapi.ImplResponse, error) {
	// TODO - update DeleteCompute with the required logic for this service method.
	// Add api_computes_service.go to the .openapi-generator-ignore to avoid overwriting
	// this service implementation when updating open api generation.

	//TODO: Uncomment the next line to return response Response(200, {}) or use other options such as http.Ok ...
	//return Response(200, nil),nil

	//TODO: Uncomment the next line to return response Response(404, {}) or use other options such as http.Ok ...
	//return Response(404, nil),nil

	//TODO: Uncomment the next line to return response Response(401, {}) or use other options such as http.Ok ...
	//return Response(401, nil),nil

	//TODO: Uncomment the next line to return response Response(0, Error{}) or use other options such as http.Ok ...
	//return Response(0, Error{}), nil

	return openapi.Response(http.StatusNotImplemented, nil), errors.New("DeleteCompute method not implemented")
}

// GetComputeConfig - Get configuration for a compute cluster
func (s *ComputesApiService) GetComputeConfig(ctx context.Context, computeId string) (openapi.ImplResponse, error) {
	// TODO - update GetComputeConfig with the required logic for this service method.
	// Add api_computes_service.go to the .openapi-generator-ignore to avoid overwriting
	// this service implementation when updating open api generation.

	//TODO: Uncomment the next line to return response Response(200, ComputeSpec{}) or use other options such as http.Ok ...
	//return Response(200, ComputeSpec{}), nil

	//TODO: Uncomment the next line to return response Response(0, Error{}) or use other options such as http.Ok ...
	//return Response(0, Error{}), nil

	return openapi.Response(http.StatusNotImplemented, nil), errors.New("GetComputeConfig method not implemented")
}

// GetComputeStatus - Get status of a given compute cluster
func (s *ComputesApiService) GetComputeStatus(ctx context.Context, computeId string) (openapi.ImplResponse, error) {
	// TODO - update GetComputeStatus with the required logic for this service method.
	// Add api_computes_service.go to the .openapi-generator-ignore to avoid overwriting
	// this service implementation when updating open api generation.

	//TODO: Uncomment the next line to return response Response(200, ComputeStatus{}) or use other options such as http.Ok ...
	//return Response(200, ComputeStatus{}), nil

	//TODO: Uncomment the next line to return response Response(0, Error{}) or use other options such as http.Ok ...
	//return Response(0, Error{}), nil

	return openapi.Response(http.StatusNotImplemented, nil), errors.New("GetComputeStatus method not implemented")
}

// GetDeploymentConfig - Get the deployment config for a job for a compute cluster
func (s *ComputesApiService) GetDeploymentConfig(ctx context.Context, computeId string, jobId string) (openapi.ImplResponse, error) {
	// TODO - update GetDeploymentConfig with the required logic for this service method.
	// Add api_computes_service.go to the .openapi-generator-ignore to avoid overwriting
	// this service implementation when updating open api generation.

	//TODO: Uncomment the next line to return response Response(200, DeploymentConfig{}) or use other options such as http.Ok ...
	//return Response(200, DeploymentConfig{}), nil

	//TODO: Uncomment the next line to return response Response(0, Error{}) or use other options such as http.Ok ...
	//return Response(0, Error{}), nil

	return openapi.Response(http.StatusNotImplemented, nil), errors.New("GetDeploymentConfig method not implemented")
}

// GetDeploymentStatus - Get the deployment status for a job on a compute cluster
func (s *ComputesApiService) GetDeploymentStatus(ctx context.Context, computeId string, jobId string) (openapi.ImplResponse, error) {
	// TODO - update GetDeploymentStatus with the required logic for this service method.
	// Add api_computes_service.go to the .openapi-generator-ignore to avoid overwriting
	// this service implementation when updating open api generation.

	//TODO: Uncomment the next line to return response Response(200, DeploymentStatus{}) or use other options such as http.Ok ...
	//return Response(200, DeploymentStatus{}), nil

	//TODO: Uncomment the next line to return response Response(0, Error{}) or use other options such as http.Ok ...
	//return Response(0, Error{}), nil

	return openapi.Response(http.StatusNotImplemented, nil), errors.New("GetDeploymentStatus method not implemented")
}

// GetDeployments - Get all deployments within a compute cluster
func (s *ComputesApiService) GetDeployments(ctx context.Context, computeId string) (openapi.ImplResponse, error) {
	// TODO - update GetDeployments with the required logic for this service method.
	// Add api_computes_service.go to the .openapi-generator-ignore to avoid overwriting
	// this service implementation when updating open api generation.

	//TODO: Uncomment the next line to return response Response(200, []DeploymentStatus{}) or use other options such as http.Ok ...
	//return Response(200, []DeploymentStatus{}), nil

	//TODO: Uncomment the next line to return response Response(0, Error{}) or use other options such as http.Ok ...
	//return Response(0, Error{}), nil

	return openapi.Response(http.StatusNotImplemented, nil), errors.New("GetDeployments method not implemented")
}

// RegisterCompute - Register a new compute cluster
func (s *ComputesApiService) RegisterCompute(ctx context.Context, computeSpec openapi.ComputeSpec) (openapi.ImplResponse, error) {
	computeStatus, err := s.dbService.RegisterCompute(computeSpec)
	if err != nil {
		errMsg := fmt.Errorf("failed to register a new compute: %v", err)
		return errMsgFunc(errMsg)
	}
	return openapi.Response(http.StatusCreated, computeStatus), nil
}

// UpdateCompute - Update a compute cluster&#39;s specification
func (s *ComputesApiService) UpdateCompute(ctx context.Context, computeId string,
	computeSpec openapi.ComputeSpec) (openapi.ImplResponse, error) {
	// TODO - update UpdateCompute with the required logic for this service method.
	// Add api_computes_service.go to the .openapi-generator-ignore to avoid overwriting
	// this service implementation when updating open api generation.

	//TODO: Uncomment the next line to return response Response(200, {}) or use other options such as http.Ok ...
	//return Response(200, nil),nil

	//TODO: Uncomment the next line to return response Response(401, {}) or use other options such as http.Ok ...
	//return Response(401, nil),nil

	//TODO: Uncomment the next line to return response Response(0, Error{}) or use other options such as http.Ok ...
	//return Response(0, Error{}), nil

	return openapi.Response(http.StatusNotImplemented, nil), errors.New("UpdateCompute method not implemented")
}

// UpdateDeploymentStatus - Update the deployment status for a job on a compute cluster
func (s *ComputesApiService) UpdateDeploymentStatus(ctx context.Context, computeId string,
	jobId string, deploymentStatus openapi.DeploymentStatus) (openapi.ImplResponse, error) {
	// TODO - update UpdateDeploymentStatus with the required logic for this service method.
	// Add api_computes_service.go to the .openapi-generator-ignore to avoid overwriting
	// this service implementation when updating open api generation.

	//TODO: Uncomment the next line to return response Response(200, {}) or use other options such as http.Ok ...
	//return Response(200, nil),nil

	//TODO: Uncomment the next line to return response Response(401, {}) or use other options such as http.Ok ...
	//return Response(401, nil),nil

	//TODO: Uncomment the next line to return response Response(0, Error{}) or use other options such as http.Ok ...
	//return Response(0, Error{}), nil

	return openapi.Response(http.StatusNotImplemented, nil), errors.New("UpdateDeploymentStatus method not implemented")
}
