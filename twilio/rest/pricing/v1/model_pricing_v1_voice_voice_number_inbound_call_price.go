/*
 * Twilio - Pricing
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.12.0
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

// PricingV1VoiceVoiceNumberInboundCallPrice The InboundCallPrice record
type PricingV1VoiceVoiceNumberInboundCallPrice struct {
	BasePrice    float32 `json:"BasePrice,omitempty"`
	CurrentPrice float32 `json:"CurrentPrice,omitempty"`
	NumberType   string  `json:"NumberType,omitempty"`
}