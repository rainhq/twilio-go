/*
 * Twilio - Trusthub
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.16.0
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

// TrusthubV1Policies struct for TrusthubV1Policies
type TrusthubV1Policies struct {
	// A human-readable description of the Policy resource
	FriendlyName *string `json:"friendly_name,omitempty"`
	// The sid of a Policy object that dictates requirements
	Requirements *map[string]interface{} `json:"requirements,omitempty"`
	// The unique string that identifies the Policy resource
	Sid *string `json:"sid,omitempty"`
	// The absolute URL of the Policy resource
	Url *string `json:"url,omitempty"`
}