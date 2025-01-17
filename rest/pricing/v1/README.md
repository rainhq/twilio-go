# Go API client for openapi

This is the public Twilio REST API.

## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project from the OpenAPI specs located at [twilio/twilio-oai](https://github.com/twilio/twilio-oai/tree/main/spec).  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: 1.17.0
- Package version: 1.0.0
- Build package: com.twilio.oai.TwilioGoGenerator
For more information, please visit [https://support.twilio.com](https://support.twilio.com)

## Installation

Install the following dependencies:

```shell
go get github.com/stretchr/testify/assert
go get golang.org/x/net/context
```

Put the package under your project folder and add the following in import:

```golang
import "./openapi"
```

## Documentation for API Endpoints

All URIs are relative to *https://pricing.twilio.com*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*DefaultApi* | [**FetchMessagingCountry**](docs/DefaultApi.md#fetchmessagingcountry) | **Get** /v1/Messaging/Countries/{IsoCountry} | 
*DefaultApi* | [**FetchPhoneNumberCountry**](docs/DefaultApi.md#fetchphonenumbercountry) | **Get** /v1/PhoneNumbers/Countries/{IsoCountry} | 
*DefaultApi* | [**FetchVoiceCountry**](docs/DefaultApi.md#fetchvoicecountry) | **Get** /v1/Voice/Countries/{IsoCountry} | 
*DefaultApi* | [**FetchVoiceNumber**](docs/DefaultApi.md#fetchvoicenumber) | **Get** /v1/Voice/Numbers/{Number} | 
*DefaultApi* | [**ListMessagingCountry**](docs/DefaultApi.md#listmessagingcountry) | **Get** /v1/Messaging/Countries | 
*DefaultApi* | [**ListPhoneNumberCountry**](docs/DefaultApi.md#listphonenumbercountry) | **Get** /v1/PhoneNumbers/Countries | 
*DefaultApi* | [**ListVoiceCountry**](docs/DefaultApi.md#listvoicecountry) | **Get** /v1/Voice/Countries | 


## Documentation For Models

 - [ListMessagingCountryResponse](docs/ListMessagingCountryResponse.md)
 - [ListMessagingCountryResponseMeta](docs/ListMessagingCountryResponseMeta.md)
 - [ListPhoneNumberCountryResponse](docs/ListPhoneNumberCountryResponse.md)
 - [ListVoiceCountryResponse](docs/ListVoiceCountryResponse.md)
 - [PricingV1Messaging](docs/PricingV1Messaging.md)
 - [PricingV1MessagingMessagingCountry](docs/PricingV1MessagingMessagingCountry.md)
 - [PricingV1MessagingMessagingCountryInstance](docs/PricingV1MessagingMessagingCountryInstance.md)
 - [PricingV1MessagingMessagingCountryInstanceInboundSmsPrices](docs/PricingV1MessagingMessagingCountryInstanceInboundSmsPrices.md)
 - [PricingV1MessagingMessagingCountryInstanceOutboundSmsPrices](docs/PricingV1MessagingMessagingCountryInstanceOutboundSmsPrices.md)
 - [PricingV1PhoneNumber](docs/PricingV1PhoneNumber.md)
 - [PricingV1PhoneNumberPhoneNumberCountry](docs/PricingV1PhoneNumberPhoneNumberCountry.md)
 - [PricingV1PhoneNumberPhoneNumberCountryInstance](docs/PricingV1PhoneNumberPhoneNumberCountryInstance.md)
 - [PricingV1Voice](docs/PricingV1Voice.md)
 - [PricingV1VoiceVoiceCountry](docs/PricingV1VoiceVoiceCountry.md)
 - [PricingV1VoiceVoiceCountryInstance](docs/PricingV1VoiceVoiceCountryInstance.md)
 - [PricingV1VoiceVoiceCountryInstanceOutboundPrefixPrices](docs/PricingV1VoiceVoiceCountryInstanceOutboundPrefixPrices.md)
 - [PricingV1VoiceVoiceNumber](docs/PricingV1VoiceVoiceNumber.md)
 - [PricingV1VoiceVoiceNumberInboundCallPrice](docs/PricingV1VoiceVoiceNumberInboundCallPrice.md)
 - [PricingV1VoiceVoiceNumberOutboundCallPrice](docs/PricingV1VoiceVoiceNumberOutboundCallPrice.md)


## Documentation For Authorization



## accountSid_authToken

- **Type**: HTTP basic authentication

Example

```golang
auth := context.WithValue(context.Background(), sw.ContextBasicAuth, sw.BasicAuth{
    UserName: "username",
    Password: "password",
})
r, err := client.Service.Operation(auth, args)
```


## Author

support@twilio.com

