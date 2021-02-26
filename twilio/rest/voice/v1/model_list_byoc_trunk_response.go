/*
 * Twilio - Voice
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.10.0
 * Contact: support@twilio.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// ListByocTrunkResponse struct for ListByocTrunkResponse
type ListByocTrunkResponse struct {
	ByocTrunks []VoiceV1ByocTrunk        `json:"ByocTrunks,omitempty"`
	Meta       ListByocTrunkResponseMeta `json:"Meta,omitempty"`
}