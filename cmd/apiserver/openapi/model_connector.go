/*
 * Fledge REST API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

type Connector struct {

	Name string `json:"name"`

	Description string `json:"description,omitempty"`

	Connection map[string]interface{} `json:"connection"`
}
