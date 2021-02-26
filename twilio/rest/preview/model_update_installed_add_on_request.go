/*
 * Twilio - Preview
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.10.0
 * Contact: support@twilio.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// UpdateInstalledAddOnRequest struct for UpdateInstalledAddOnRequest
type UpdateInstalledAddOnRequest struct {
	// Valid JSON object that conform to the configuration schema exposed by the associated AvailableAddOn resource. This is only required by Add-ons that need to be configured
	Configuration map[string]interface{} `json:"Configuration,omitempty"`
	// An application-defined string that uniquely identifies the resource. This value must be unique within the Account.
	UniqueName string `json:"UniqueName,omitempty"`
}