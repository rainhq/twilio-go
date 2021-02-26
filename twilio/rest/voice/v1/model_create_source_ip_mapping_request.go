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

// CreateSourceIpMappingRequest struct for CreateSourceIpMappingRequest
type CreateSourceIpMappingRequest struct {
	// The Twilio-provided string that uniquely identifies the IP Record resource to map from.
	IpRecordSid string `json:"IpRecordSid"`
	// The SID of the SIP Domain that the IP Record should be mapped to.
	SipDomainSid string `json:"SipDomainSid"`
}