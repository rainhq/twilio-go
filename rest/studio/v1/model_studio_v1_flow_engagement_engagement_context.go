/*
 * Twilio - Studio
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.17.0
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

// StudioV1FlowEngagementEngagementContext struct for StudioV1FlowEngagementEngagementContext
type StudioV1FlowEngagementEngagementContext struct {
	// Account SID
	AccountSid *string `json:"account_sid,omitempty"`
	// Flow state
	Context *map[string]interface{} `json:"context,omitempty"`
	// Engagement SID
	EngagementSid *string `json:"engagement_sid,omitempty"`
	// Flow SID
	FlowSid *string `json:"flow_sid,omitempty"`
	// The URL of the resource
	Url *string `json:"url,omitempty"`
}
