/*
 * Fledge REST API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

import (
	"context"
	"net/http"
	"errors"
)

// DesignsApiService is a service that implents the logic for the DesignsApiServicer
// This service should implement the business logic for every endpoint for the DesignsApi API. 
// Include any external packages or services that will be required by this service.
type DesignsApiService struct {
}

// NewDesignsApiService creates a default api service
func NewDesignsApiService() DesignsApiServicer {
	return &DesignsApiService{}
}

// GetDesigns - Get list of all the designs created by the user.
func (s *DesignsApiService) GetDesigns(ctx context.Context, user string, limit int32) (ImplResponse, error) {
	// TODO - update GetDesigns with the required logic for this service method.
	// Add api_designs_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	//TODO: Uncomment the next line to return response Response(200, []DesignInfo{}) or use other options such as http.Ok ...
	//return Response(200, []DesignInfo{}), nil

	//TODO: Uncomment the next line to return response Response(0, Error{}) or use other options such as http.Ok ...
	//return Response(0, Error{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("GetDesigns method not implemented")
}

