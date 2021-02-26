/*
 * Twilio - Supersim
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.10.0
 * Contact: support@twilio.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strings"
	"time"

	twilio "github.com/twilio/twilio-go/client"
)

type DefaultApiService struct {
	baseURL string
	client  *twilio.Client
}

func NewDefaultApiService(client *twilio.Client) *DefaultApiService {
	return &DefaultApiService{
		client:  client,
		baseURL: fmt.Sprintf("https://supersim.twilio.com"),
	}
}

// CreateCommandParams Optional parameters for the method 'CreateCommand'
type CreateCommandParams struct {
	CallbackMethod *string `json:"CallbackMethod,omitempty"`
	CallbackUrl    *string `json:"CallbackUrl,omitempty"`
	Command        *string `json:"Command,omitempty"`
	Sim            *string `json:"Sim,omitempty"`
}

/*
CreateCommand Method for CreateCommand
Send a Command to a Sim.
 * @param optional nil or *CreateCommandOpts - Optional Parameters:
 * @param "CallbackMethod" (string) - The HTTP method we should use to call `callback_url`. Can be: `GET` or `POST` and the default is POST.
 * @param "CallbackUrl" (string) - The URL we should call using the `callback_method` after we have sent the command.
 * @param "Command" (string) - The message body of the command.
 * @param "Sim" (string) - The `sid` or `unique_name` of the [SIM](https://www.twilio.com/docs/wireless/api/sim-resource) to send the Command to.
@return SupersimV1Command
*/
func (c *DefaultApiService) CreateCommand(params *CreateCommandParams) (*SupersimV1Command, error) {
	path := "/v1/Commands"

	data := url.Values{}
	headers := 0

	if params != nil && params.CallbackMethod != nil {
		data.Set("CallbackMethod", *params.CallbackMethod)
	}
	if params != nil && params.CallbackUrl != nil {
		data.Set("CallbackUrl", *params.CallbackUrl)
	}
	if params != nil && params.Command != nil {
		data.Set("Command", *params.Command)
	}
	if params != nil && params.Sim != nil {
		data.Set("Sim", *params.Sim)
	}

	resp, err := c.client.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &SupersimV1Command{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// CreateFleetParams Optional parameters for the method 'CreateFleet'
type CreateFleetParams struct {
	CommandsEnabled      *bool   `json:"CommandsEnabled,omitempty"`
	CommandsMethod       *string `json:"CommandsMethod,omitempty"`
	CommandsUrl          *string `json:"CommandsUrl,omitempty"`
	DataEnabled          *bool   `json:"DataEnabled,omitempty"`
	DataLimit            *int32  `json:"DataLimit,omitempty"`
	NetworkAccessProfile *string `json:"NetworkAccessProfile,omitempty"`
	UniqueName           *string `json:"UniqueName,omitempty"`
}

/*
CreateFleet Method for CreateFleet
Create a Fleet
 * @param optional nil or *CreateFleetOpts - Optional Parameters:
 * @param "CommandsEnabled" (bool) - Defines whether SIMs in the Fleet are capable of sending and receiving machine-to-machine SMS via Commands. Defaults to `true`.
 * @param "CommandsMethod" (string) - A string representing the HTTP method to use when making a request to `commands_url`. Can be one of `POST` or `GET`. Defaults to `POST`.
 * @param "CommandsUrl" (string) - The URL that will receive a webhook when a Super SIM in the Fleet is used to send an SMS from your device to the Commands number. Your server should respond with an HTTP status code in the 200 range; any response body will be ignored.
 * @param "DataEnabled" (bool) - Defines whether SIMs in the Fleet are capable of using 2G/3G/4G/LTE/CAT-M data connectivity. Defaults to `true`.
 * @param "DataLimit" (int32) - The total data usage (download and upload combined) in Megabytes that each Sim resource assigned to the Fleet resource can consume during a billing period (normally one month). Value must be between 1MB (1) and 2TB (2,000,000). Defaults to 1GB (1,000).
 * @param "NetworkAccessProfile" (string) - The SID or unique name of the Network Access Profile that will control which cellular networks the Fleet's SIMs can connect to.
 * @param "UniqueName" (string) - An application-defined string that uniquely identifies the resource. It can be used in place of the resource's `sid` in the URL to address the resource.
@return SupersimV1Fleet
*/
func (c *DefaultApiService) CreateFleet(params *CreateFleetParams) (*SupersimV1Fleet, error) {
	path := "/v1/Fleets"

	data := url.Values{}
	headers := 0

	if params != nil && params.CommandsEnabled != nil {
		data.Set("CommandsEnabled", fmt.Sprint(*params.CommandsEnabled))
	}
	if params != nil && params.CommandsMethod != nil {
		data.Set("CommandsMethod", *params.CommandsMethod)
	}
	if params != nil && params.CommandsUrl != nil {
		data.Set("CommandsUrl", *params.CommandsUrl)
	}
	if params != nil && params.DataEnabled != nil {
		data.Set("DataEnabled", fmt.Sprint(*params.DataEnabled))
	}
	if params != nil && params.DataLimit != nil {
		data.Set("DataLimit", fmt.Sprint(*params.DataLimit))
	}
	if params != nil && params.NetworkAccessProfile != nil {
		data.Set("NetworkAccessProfile", *params.NetworkAccessProfile)
	}
	if params != nil && params.UniqueName != nil {
		data.Set("UniqueName", *params.UniqueName)
	}

	resp, err := c.client.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &SupersimV1Fleet{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// CreateNetworkAccessProfileParams Optional parameters for the method 'CreateNetworkAccessProfile'
type CreateNetworkAccessProfileParams struct {
	Networks   *[]string `json:"Networks,omitempty"`
	UniqueName *string   `json:"UniqueName,omitempty"`
}

/*
CreateNetworkAccessProfile Method for CreateNetworkAccessProfile
Create a new Network Access Profile
 * @param optional nil or *CreateNetworkAccessProfileOpts - Optional Parameters:
 * @param "Networks" ([]string) - List of Network SIDs that this Network Access Profile will allow connections to.
 * @param "UniqueName" (string) - An application-defined string that uniquely identifies the resource. It can be used in place of the resource's `sid` in the URL to address the resource.
@return SupersimV1NetworkAccessProfile
*/
func (c *DefaultApiService) CreateNetworkAccessProfile(params *CreateNetworkAccessProfileParams) (*SupersimV1NetworkAccessProfile, error) {
	path := "/v1/NetworkAccessProfiles"

	data := url.Values{}
	headers := 0

	if params != nil && params.Networks != nil {
		data.Set("Networks", strings.Join(*params.Networks, ","))
	}
	if params != nil && params.UniqueName != nil {
		data.Set("UniqueName", *params.UniqueName)
	}

	resp, err := c.client.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &SupersimV1NetworkAccessProfile{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// CreateNetworkAccessProfileNetworkParams Optional parameters for the method 'CreateNetworkAccessProfileNetwork'
type CreateNetworkAccessProfileNetworkParams struct {
	Network *string `json:"Network,omitempty"`
}

/*
CreateNetworkAccessProfileNetwork Method for CreateNetworkAccessProfileNetwork
Add a Network resource to the Network Access Profile resource.
 * @param NetworkAccessProfileSid The unique string that identifies the Network Access Profile resource.
 * @param optional nil or *CreateNetworkAccessProfileNetworkOpts - Optional Parameters:
 * @param "Network" (string) - The SID of the Network resource to be added to the Network Access Profile resource.
@return SupersimV1NetworkAccessProfileNetworkAccessProfileNetwork
*/
func (c *DefaultApiService) CreateNetworkAccessProfileNetwork(NetworkAccessProfileSid string, params *CreateNetworkAccessProfileNetworkParams) (*SupersimV1NetworkAccessProfileNetworkAccessProfileNetwork, error) {
	path := "/v1/NetworkAccessProfiles/{NetworkAccessProfileSid}/Networks"
	path = strings.Replace(path, "{"+"NetworkAccessProfileSid"+"}", NetworkAccessProfileSid, -1)

	data := url.Values{}
	headers := 0

	if params != nil && params.Network != nil {
		data.Set("Network", *params.Network)
	}

	resp, err := c.client.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &SupersimV1NetworkAccessProfileNetworkAccessProfileNetwork{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

/*
DeleteNetworkAccessProfileNetwork Method for DeleteNetworkAccessProfileNetwork
Remove a Network resource from the Network Access Profile resource&#39;s.
 * @param NetworkAccessProfileSid The unique string that identifies the Network Access Profile resource.
 * @param Sid The SID of the Network resource to be removed from the Network Access Profile resource.
*/
func (c *DefaultApiService) DeleteNetworkAccessProfileNetwork(NetworkAccessProfileSid string, Sid string) error {
	path := "/v1/NetworkAccessProfiles/{NetworkAccessProfileSid}/Networks/{Sid}"
	path = strings.Replace(path, "{"+"NetworkAccessProfileSid"+"}", NetworkAccessProfileSid, -1)
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := 0

	resp, err := c.client.Delete(c.baseURL+path, data, headers)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	return nil
}

/*
FetchCommand Method for FetchCommand
Fetch a Command instance from your account.
 * @param Sid The SID of the Command resource to fetch.
@return SupersimV1Command
*/
func (c *DefaultApiService) FetchCommand(Sid string) (*SupersimV1Command, error) {
	path := "/v1/Commands/{Sid}"
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := 0

	resp, err := c.client.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &SupersimV1Command{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

/*
FetchFleet Method for FetchFleet
Fetch a Fleet instance from your account.
 * @param Sid The SID of the Fleet resource to fetch.
@return SupersimV1Fleet
*/
func (c *DefaultApiService) FetchFleet(Sid string) (*SupersimV1Fleet, error) {
	path := "/v1/Fleets/{Sid}"
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := 0

	resp, err := c.client.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &SupersimV1Fleet{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

/*
FetchNetwork Method for FetchNetwork
Fetch a Network resource.
 * @param Sid The SID of the Network resource to fetch.
@return SupersimV1Network
*/
func (c *DefaultApiService) FetchNetwork(Sid string) (*SupersimV1Network, error) {
	path := "/v1/Networks/{Sid}"
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := 0

	resp, err := c.client.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &SupersimV1Network{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

/*
FetchNetworkAccessProfile Method for FetchNetworkAccessProfile
Fetch a Network Access Profile instance from your account.
 * @param Sid The SID of the Network Access Profile resource to fetch.
@return SupersimV1NetworkAccessProfile
*/
func (c *DefaultApiService) FetchNetworkAccessProfile(Sid string) (*SupersimV1NetworkAccessProfile, error) {
	path := "/v1/NetworkAccessProfiles/{Sid}"
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := 0

	resp, err := c.client.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &SupersimV1NetworkAccessProfile{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

/*
FetchNetworkAccessProfileNetwork Method for FetchNetworkAccessProfileNetwork
Fetch a Network Access Profile resource&#39;s Network resource.
 * @param NetworkAccessProfileSid The unique string that identifies the Network Access Profile resource.
 * @param Sid The SID of the Network resource to fetch.
@return SupersimV1NetworkAccessProfileNetworkAccessProfileNetwork
*/
func (c *DefaultApiService) FetchNetworkAccessProfileNetwork(NetworkAccessProfileSid string, Sid string) (*SupersimV1NetworkAccessProfileNetworkAccessProfileNetwork, error) {
	path := "/v1/NetworkAccessProfiles/{NetworkAccessProfileSid}/Networks/{Sid}"
	path = strings.Replace(path, "{"+"NetworkAccessProfileSid"+"}", NetworkAccessProfileSid, -1)
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := 0

	resp, err := c.client.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &SupersimV1NetworkAccessProfileNetworkAccessProfileNetwork{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

/*
FetchSim Method for FetchSim
Fetch a Super SIM instance from your account.
 * @param Sid The SID of the Sim resource to fetch.
@return SupersimV1Sim
*/
func (c *DefaultApiService) FetchSim(Sid string) (*SupersimV1Sim, error) {
	path := "/v1/Sims/{Sid}"
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := 0

	resp, err := c.client.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &SupersimV1Sim{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// ListCommandParams Optional parameters for the method 'ListCommand'
type ListCommandParams struct {
	Sim       *string `json:"Sim,omitempty"`
	Status    *string `json:"Status,omitempty"`
	Direction *string `json:"Direction,omitempty"`
	PageSize  *int32  `json:"PageSize,omitempty"`
}

/*
ListCommand Method for ListCommand
Retrieve a list of Commands from your account.
 * @param optional nil or *ListCommandOpts - Optional Parameters:
 * @param "Sim" (string) - The SID or unique name of the Sim that Command was sent to or from.
 * @param "Status" (string) - The status of the Command. Can be: `queued`, `sent`, `delivered`, `received` or `failed`. See the [Command Status Values](https://www.twilio.com/docs/wireless/api/command-resource#status-values) for a description of each.
 * @param "Direction" (string) - The direction of the Command. Can be `to_sim` or `from_sim`. The value of `to_sim` is synonymous with the term `mobile terminated`, and `from_sim` is synonymous with the term `mobile originated`.
 * @param "PageSize" (int32) - How many resources to return in each list page. The default is 50, and the maximum is 1000.
@return ListCommandResponse
*/
func (c *DefaultApiService) ListCommand(params *ListCommandParams) (*ListCommandResponse, error) {
	path := "/v1/Commands"

	data := url.Values{}
	headers := 0

	if params != nil && params.Sim != nil {
		data.Set("Sim", *params.Sim)
	}
	if params != nil && params.Status != nil {
		data.Set("Status", *params.Status)
	}
	if params != nil && params.Direction != nil {
		data.Set("Direction", *params.Direction)
	}
	if params != nil && params.PageSize != nil {
		data.Set("PageSize", fmt.Sprint(*params.PageSize))
	}

	resp, err := c.client.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListCommandResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// ListFleetParams Optional parameters for the method 'ListFleet'
type ListFleetParams struct {
	NetworkAccessProfile *string `json:"NetworkAccessProfile,omitempty"`
	PageSize             *int32  `json:"PageSize,omitempty"`
}

/*
ListFleet Method for ListFleet
Retrieve a list of Fleets from your account.
 * @param optional nil or *ListFleetOpts - Optional Parameters:
 * @param "NetworkAccessProfile" (string) - The SID or unique name of the Network Access Profile that controls which cellular networks the Fleet's SIMs can connect to.
 * @param "PageSize" (int32) - How many resources to return in each list page. The default is 50, and the maximum is 1000.
@return ListFleetResponse
*/
func (c *DefaultApiService) ListFleet(params *ListFleetParams) (*ListFleetResponse, error) {
	path := "/v1/Fleets"

	data := url.Values{}
	headers := 0

	if params != nil && params.NetworkAccessProfile != nil {
		data.Set("NetworkAccessProfile", *params.NetworkAccessProfile)
	}
	if params != nil && params.PageSize != nil {
		data.Set("PageSize", fmt.Sprint(*params.PageSize))
	}

	resp, err := c.client.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListFleetResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// ListNetworkParams Optional parameters for the method 'ListNetwork'
type ListNetworkParams struct {
	IsoCountry *string `json:"IsoCountry,omitempty"`
	Mcc        *string `json:"Mcc,omitempty"`
	Mnc        *string `json:"Mnc,omitempty"`
	PageSize   *int32  `json:"PageSize,omitempty"`
}

/*
ListNetwork Method for ListNetwork
Retrieve a list of Network resources.
 * @param optional nil or *ListNetworkOpts - Optional Parameters:
 * @param "IsoCountry" (string) - The [ISO country code](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) of the Network resources to read.
 * @param "Mcc" (string) - The 'mobile country code' of a country. Network resources with this `mcc` in their `identifiers` will be read.
 * @param "Mnc" (string) - The 'mobile network code' of a mobile operator network. Network resources with this `mnc` in their `identifiers` will be read.
 * @param "PageSize" (int32) - How many resources to return in each list page. The default is 50, and the maximum is 1000.
@return ListNetworkResponse
*/
func (c *DefaultApiService) ListNetwork(params *ListNetworkParams) (*ListNetworkResponse, error) {
	path := "/v1/Networks"

	data := url.Values{}
	headers := 0

	if params != nil && params.IsoCountry != nil {
		data.Set("IsoCountry", *params.IsoCountry)
	}
	if params != nil && params.Mcc != nil {
		data.Set("Mcc", *params.Mcc)
	}
	if params != nil && params.Mnc != nil {
		data.Set("Mnc", *params.Mnc)
	}
	if params != nil && params.PageSize != nil {
		data.Set("PageSize", fmt.Sprint(*params.PageSize))
	}

	resp, err := c.client.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListNetworkResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// ListNetworkAccessProfileParams Optional parameters for the method 'ListNetworkAccessProfile'
type ListNetworkAccessProfileParams struct {
	PageSize *int32 `json:"PageSize,omitempty"`
}

/*
ListNetworkAccessProfile Method for ListNetworkAccessProfile
Retrieve a list of Network Access Profiles from your account.
 * @param optional nil or *ListNetworkAccessProfileOpts - Optional Parameters:
 * @param "PageSize" (int32) - How many resources to return in each list page. The default is 50, and the maximum is 1000.
@return ListNetworkAccessProfileResponse
*/
func (c *DefaultApiService) ListNetworkAccessProfile(params *ListNetworkAccessProfileParams) (*ListNetworkAccessProfileResponse, error) {
	path := "/v1/NetworkAccessProfiles"

	data := url.Values{}
	headers := 0

	if params != nil && params.PageSize != nil {
		data.Set("PageSize", fmt.Sprint(*params.PageSize))
	}

	resp, err := c.client.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListNetworkAccessProfileResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// ListNetworkAccessProfileNetworkParams Optional parameters for the method 'ListNetworkAccessProfileNetwork'
type ListNetworkAccessProfileNetworkParams struct {
	PageSize *int32 `json:"PageSize,omitempty"`
}

/*
ListNetworkAccessProfileNetwork Method for ListNetworkAccessProfileNetwork
Retrieve a list of Network Access Profile resource&#39;s Network resource.
 * @param NetworkAccessProfileSid The unique string that identifies the Network Access Profile resource.
 * @param optional nil or *ListNetworkAccessProfileNetworkOpts - Optional Parameters:
 * @param "PageSize" (int32) - How many resources to return in each list page. The default is 50, and the maximum is 1000.
@return ListNetworkAccessProfileNetworkResponse
*/
func (c *DefaultApiService) ListNetworkAccessProfileNetwork(NetworkAccessProfileSid string, params *ListNetworkAccessProfileNetworkParams) (*ListNetworkAccessProfileNetworkResponse, error) {
	path := "/v1/NetworkAccessProfiles/{NetworkAccessProfileSid}/Networks"
	path = strings.Replace(path, "{"+"NetworkAccessProfileSid"+"}", NetworkAccessProfileSid, -1)

	data := url.Values{}
	headers := 0

	if params != nil && params.PageSize != nil {
		data.Set("PageSize", fmt.Sprint(*params.PageSize))
	}

	resp, err := c.client.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListNetworkAccessProfileNetworkResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// ListSimParams Optional parameters for the method 'ListSim'
type ListSimParams struct {
	Status   *string `json:"Status,omitempty"`
	Fleet    *string `json:"Fleet,omitempty"`
	Iccid    *string `json:"Iccid,omitempty"`
	PageSize *int32  `json:"PageSize,omitempty"`
}

/*
ListSim Method for ListSim
Retrieve a list of Super SIMs from your account.
 * @param optional nil or *ListSimOpts - Optional Parameters:
 * @param "Status" (string) - The status of the Sim resources to read. Can be `new`, `ready`, `active`, `inactive`, or `scheduled`.
 * @param "Fleet" (string) - The SID or unique name of the Fleet to which a list of Sims are assigned.
 * @param "Iccid" (string) - The [ICCID](https://en.wikipedia.org/wiki/Subscriber_identity_module#ICCID) associated with a Super SIM to filter the list by. Passing this parameter will always return a list containing zero or one SIMs.
 * @param "PageSize" (int32) - How many resources to return in each list page. The default is 50, and the maximum is 1000.
@return ListSimResponse
*/
func (c *DefaultApiService) ListSim(params *ListSimParams) (*ListSimResponse, error) {
	path := "/v1/Sims"

	data := url.Values{}
	headers := 0

	if params != nil && params.Status != nil {
		data.Set("Status", *params.Status)
	}
	if params != nil && params.Fleet != nil {
		data.Set("Fleet", *params.Fleet)
	}
	if params != nil && params.Iccid != nil {
		data.Set("Iccid", *params.Iccid)
	}
	if params != nil && params.PageSize != nil {
		data.Set("PageSize", fmt.Sprint(*params.PageSize))
	}

	resp, err := c.client.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListSimResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// ListUsageRecordParams Optional parameters for the method 'ListUsageRecord'
type ListUsageRecordParams struct {
	Sim         *string    `json:"Sim,omitempty"`
	Fleet       *string    `json:"Fleet,omitempty"`
	Network     *string    `json:"Network,omitempty"`
	IsoCountry  *string    `json:"IsoCountry,omitempty"`
	Group       *string    `json:"Group,omitempty"`
	Granularity *string    `json:"Granularity,omitempty"`
	StartTime   *time.Time `json:"StartTime,omitempty"`
	EndTime     *time.Time `json:"EndTime,omitempty"`
	PageSize    *int32     `json:"PageSize,omitempty"`
}

/*
ListUsageRecord Method for ListUsageRecord
List UsageRecords
 * @param optional nil or *ListUsageRecordOpts - Optional Parameters:
 * @param "Sim" (string) - SID or unique name of a Sim resource. Only show UsageRecords representing usage incurred by this Super SIM.
 * @param "Fleet" (string) - SID or unique name of a Fleet resource. Only show UsageRecords representing usage for Super SIMs belonging to this Fleet resource at the time the usage occurred.
 * @param "Network" (string) - SID of a Network resource. Only show UsageRecords representing usage on this network.
 * @param "IsoCountry" (string) - Alpha-2 ISO Country Code. Only show UsageRecords representing usage in this country.
 * @param "Group" (string) - Dimension over which to aggregate usage records. Can be: `sim`, `fleet`, `network`, `isoCountry`. Default is to not aggregate across any of these dimensions, UsageRecords will be aggregated into the time buckets described by the `Granularity` parameter.
 * @param "Granularity" (string) - Time-based grouping that UsageRecords should be aggregated by. Can be: `hour`, `day`, or `all`. Default is `all`. `all` returns one UsageRecord that describes the usage for the entire period.
 * @param "StartTime" (time.Time) - Only include usage that occurred at or after this time, specified in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format. Default is one month before the `end_time`.
 * @param "EndTime" (time.Time) - Only include usage that occurred before this time, specified in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format. Default is the current time.
 * @param "PageSize" (int32) - How many resources to return in each list page. The default is 50, and the maximum is 1000.
@return ListUsageRecordResponse
*/
func (c *DefaultApiService) ListUsageRecord(params *ListUsageRecordParams) (*ListUsageRecordResponse, error) {
	path := "/v1/UsageRecords"

	data := url.Values{}
	headers := 0

	if params != nil && params.Sim != nil {
		data.Set("Sim", *params.Sim)
	}
	if params != nil && params.Fleet != nil {
		data.Set("Fleet", *params.Fleet)
	}
	if params != nil && params.Network != nil {
		data.Set("Network", *params.Network)
	}
	if params != nil && params.IsoCountry != nil {
		data.Set("IsoCountry", *params.IsoCountry)
	}
	if params != nil && params.Group != nil {
		data.Set("Group", *params.Group)
	}
	if params != nil && params.Granularity != nil {
		data.Set("Granularity", *params.Granularity)
	}
	if params != nil && params.StartTime != nil {
		data.Set("StartTime", fmt.Sprint(*params.StartTime))
	}
	if params != nil && params.EndTime != nil {
		data.Set("EndTime", fmt.Sprint(*params.EndTime))
	}
	if params != nil && params.PageSize != nil {
		data.Set("PageSize", fmt.Sprint(*params.PageSize))
	}

	resp, err := c.client.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListUsageRecordResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// UpdateFleetParams Optional parameters for the method 'UpdateFleet'
type UpdateFleetParams struct {
	CommandsMethod       *string `json:"CommandsMethod,omitempty"`
	CommandsUrl          *string `json:"CommandsUrl,omitempty"`
	NetworkAccessProfile *string `json:"NetworkAccessProfile,omitempty"`
	UniqueName           *string `json:"UniqueName,omitempty"`
}

/*
UpdateFleet Method for UpdateFleet
Updates the given properties of a Super SIM Fleet instance from your account.
 * @param Sid The SID of the Fleet resource to update.
 * @param optional nil or *UpdateFleetOpts - Optional Parameters:
 * @param "CommandsMethod" (string) - A string representing the HTTP method to use when making a request to `commands_url`. Can be one of `POST` or `GET`. Defaults to `POST`.
 * @param "CommandsUrl" (string) - The URL that will receive a webhook when a Super SIM in the Fleet is used to send an SMS from your device to the Commands number. Your server should respond with an HTTP status code in the 200 range; any response body will be ignored.
 * @param "NetworkAccessProfile" (string) - The SID or unique name of the Network Access Profile that will control which cellular networks the Fleet's SIMs can connect to.
 * @param "UniqueName" (string) - An application-defined string that uniquely identifies the resource. It can be used in place of the resource's `sid` in the URL to address the resource.
@return SupersimV1Fleet
*/
func (c *DefaultApiService) UpdateFleet(Sid string, params *UpdateFleetParams) (*SupersimV1Fleet, error) {
	path := "/v1/Fleets/{Sid}"
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := 0

	if params != nil && params.CommandsMethod != nil {
		data.Set("CommandsMethod", *params.CommandsMethod)
	}
	if params != nil && params.CommandsUrl != nil {
		data.Set("CommandsUrl", *params.CommandsUrl)
	}
	if params != nil && params.NetworkAccessProfile != nil {
		data.Set("NetworkAccessProfile", *params.NetworkAccessProfile)
	}
	if params != nil && params.UniqueName != nil {
		data.Set("UniqueName", *params.UniqueName)
	}

	resp, err := c.client.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &SupersimV1Fleet{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// UpdateNetworkAccessProfileParams Optional parameters for the method 'UpdateNetworkAccessProfile'
type UpdateNetworkAccessProfileParams struct {
	UniqueName *string `json:"UniqueName,omitempty"`
}

/*
UpdateNetworkAccessProfile Method for UpdateNetworkAccessProfile
Updates the given properties of a Network Access Profile in your account.
 * @param Sid The SID of the Network Access Profile to update.
 * @param optional nil or *UpdateNetworkAccessProfileOpts - Optional Parameters:
 * @param "UniqueName" (string) - The new unique name of the Network Access Profile.
@return SupersimV1NetworkAccessProfile
*/
func (c *DefaultApiService) UpdateNetworkAccessProfile(Sid string, params *UpdateNetworkAccessProfileParams) (*SupersimV1NetworkAccessProfile, error) {
	path := "/v1/NetworkAccessProfiles/{Sid}"
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := 0

	if params != nil && params.UniqueName != nil {
		data.Set("UniqueName", *params.UniqueName)
	}

	resp, err := c.client.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &SupersimV1NetworkAccessProfile{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// UpdateSimParams Optional parameters for the method 'UpdateSim'
type UpdateSimParams struct {
	AccountSid     *string `json:"AccountSid,omitempty"`
	CallbackMethod *string `json:"CallbackMethod,omitempty"`
	CallbackUrl    *string `json:"CallbackUrl,omitempty"`
	Fleet          *string `json:"Fleet,omitempty"`
	Status         *string `json:"Status,omitempty"`
	UniqueName     *string `json:"UniqueName,omitempty"`
}

/*
UpdateSim Method for UpdateSim
Updates the given properties of a Super SIM instance from your account.
 * @param Sid The SID of the Sim resource to update.
 * @param optional nil or *UpdateSimOpts - Optional Parameters:
 * @param "AccountSid" (string) - The SID of the Account to which the Sim resource should belong. The Account SID can only be that of the requesting Account or that of a Subaccount of the requesting Account. Only valid when the Sim resource's status is new.
 * @param "CallbackMethod" (string) - The HTTP method we should use to call `callback_url`. Can be: `GET` or `POST` and the default is POST.
 * @param "CallbackUrl" (string) - The URL we should call using the `callback_method` after an asynchronous update has finished.
 * @param "Fleet" (string) - The SID or unique name of the Fleet to which the SIM resource should be assigned.
 * @param "Status" (string) - The new status of the resource. Can be: `ready`, `active`, or `inactive`. See the [Super SIM Status Values](https://www.twilio.com/docs/iot/supersim/api/sim-resource#status-values) for more info.
 * @param "UniqueName" (string) - An application-defined string that uniquely identifies the resource. It can be used in place of the resource's `sid` in the URL to address the resource.
@return SupersimV1Sim
*/
func (c *DefaultApiService) UpdateSim(Sid string, params *UpdateSimParams) (*SupersimV1Sim, error) {
	path := "/v1/Sims/{Sid}"
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := 0

	if params != nil && params.AccountSid != nil {
		data.Set("AccountSid", *params.AccountSid)
	}
	if params != nil && params.CallbackMethod != nil {
		data.Set("CallbackMethod", *params.CallbackMethod)
	}
	if params != nil && params.CallbackUrl != nil {
		data.Set("CallbackUrl", *params.CallbackUrl)
	}
	if params != nil && params.Fleet != nil {
		data.Set("Fleet", *params.Fleet)
	}
	if params != nil && params.Status != nil {
		data.Set("Status", *params.Status)
	}
	if params != nil && params.UniqueName != nil {
		data.Set("UniqueName", *params.UniqueName)
	}

	resp, err := c.client.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &SupersimV1Sim{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}