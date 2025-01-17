/*
 * Twilio - Ip_messaging
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.17.0
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

// ListInviteResponse struct for ListInviteResponse
type ListInviteResponse struct {
	Invites []IpMessagingV2ServiceChannelInvite `json:"invites,omitempty"`
	Meta    ListCredentialResponseMeta          `json:"meta,omitempty"`
}
