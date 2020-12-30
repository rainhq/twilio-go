/*
 * Twilio - Verify
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.0.1
 * Contact: support@twilio.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package twilio

import (
    twilio "github.com/twilio/twilio-go/client"
)

type ApiService struct {
    GAApi *GAApiService
    PreviewApi *PreviewApiService
    
}

func NewApiService(client *twilio.Client) *ApiService {
    return &ApiService{
        GAApi: NewGAApiService(client),
        PreviewApi: NewPreviewApiService(client),
        
    }
}