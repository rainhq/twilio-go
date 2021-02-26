/*
 * Twilio - Numbers
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.10.0
 * Contact: support@twilio.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// ListBundleResponse struct for ListBundleResponse
type ListBundleResponse struct {
	Meta    ListBundleResponseMeta                `json:"Meta,omitempty"`
	Results []NumbersV2RegulatoryComplianceBundle `json:"Results,omitempty"`
}