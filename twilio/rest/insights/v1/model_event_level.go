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

// EventLevel the model 'EventLevel'
type EventLevel string

// List of event_level
const (
	EVENTLEVEL_UNKNOWN EventLevel = "UNKNOWN"
	EVENTLEVEL_DEBUG   EventLevel = "DEBUG"
	EVENTLEVEL_INFO    EventLevel = "INFO"
	EVENTLEVEL_WARNING EventLevel = "WARNING"
	EVENTLEVEL_ERROR   EventLevel = "ERROR"
)