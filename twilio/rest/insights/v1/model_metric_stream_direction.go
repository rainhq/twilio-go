/*
 * Twilio - Insights
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.10.0
 * Contact: support@twilio.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// MetricStreamDirection the model 'MetricStreamDirection'
type MetricStreamDirection string

// List of metric_stream_direction
const (
	METRICSTREAMDIRECTION_UNKNOWN  MetricStreamDirection = "unknown"
	METRICSTREAMDIRECTION_INBOUND  MetricStreamDirection = "inbound"
	METRICSTREAMDIRECTION_OUTBOUND MetricStreamDirection = "outbound"
	METRICSTREAMDIRECTION_BOTH     MetricStreamDirection = "both"
)