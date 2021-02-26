/*
 * Twilio - Api
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.10.0
 * Contact: support@twilio.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// CreateParticipantRequest struct for CreateParticipantRequest
type CreateParticipantRequest struct {
	// Whether to play a notification beep to the conference when the participant joins. Can be: `true`, `false`, `onEnter`, or `onExit`. The default value is `true`.
	Beep string `json:"Beep,omitempty"`
	// The SID of a BYOC (Bring Your Own Carrier) trunk to route this call with. Note that `byoc` is only meaningful when `to` is a phone number; it will otherwise be ignored. (Beta)
	Byoc string `json:"Byoc,omitempty"`
	// The Reason for the outgoing call. Use it to specify the purpose of the call that is presented on the called party's phone. (Branded Calls Beta)
	CallReason string `json:"CallReason,omitempty"`
	// The SID of the participant who is being `coached`. The participant being coached is the only participant who can hear the participant who is `coaching`.
	CallSidToCoach string `json:"CallSidToCoach,omitempty"`
	// The phone number, Client identifier, or username portion of SIP address that made this call. Phone numbers are in [E.164](https://www.twilio.com/docs/glossary/what-e164) format (e.g., +16175551212). Client identifiers are formatted `client:name`. If using a phone number, it must be a Twilio number or a Verified [outgoing caller id](https://www.twilio.com/docs/voice/api/outgoing-caller-ids) for your account. If the `to` parameter is a phone number, `callerId` must also be a phone number. If `to` is sip address, this value of `callerId` should be a username portion to be used to populate the From header that is passed to the SIP endpoint.
	CallerId string `json:"CallerId,omitempty"`
	// Whether the participant is coaching another call. Can be: `true` or `false`. If not present, defaults to `false` unless `call_sid_to_coach` is defined. If `true`, `call_sid_to_coach` must be defined.
	Coaching bool `json:"Coaching,omitempty"`
	// Whether to record the conference the participant is joining. Can be: `true`, `false`, `record-from-start`, and `do-not-record`. The default value is `false`.
	ConferenceRecord string `json:"ConferenceRecord,omitempty"`
	// The URL we should call using the `conference_recording_status_callback_method` when the conference recording is available.
	ConferenceRecordingStatusCallback string `json:"ConferenceRecordingStatusCallback,omitempty"`
	// The conference recording state changes that generate a call to `conference_recording_status_callback`. Can be: `in-progress`, `completed`, and `failed`. Separate multiple values with a space. The default value is `in-progress completed failed`.
	ConferenceRecordingStatusCallbackEvent []string `json:"ConferenceRecordingStatusCallbackEvent,omitempty"`
	// The HTTP method we should use to call `conference_recording_status_callback`. Can be: `GET` or `POST` and defaults to `POST`.
	ConferenceRecordingStatusCallbackMethod string `json:"ConferenceRecordingStatusCallbackMethod,omitempty"`
	// The URL we should call using the `conference_status_callback_method` when the conference events in `conference_status_callback_event` occur. Only the value set by the first participant to join the conference is used. Subsequent `conference_status_callback` values are ignored.
	ConferenceStatusCallback string `json:"ConferenceStatusCallback,omitempty"`
	// The conference state changes that should generate a call to `conference_status_callback`. Can be: `start`, `end`, `join`, `leave`, `mute`, `hold`, and `speaker`. Separate multiple values with a space. Defaults to `start end`.
	ConferenceStatusCallbackEvent []string `json:"ConferenceStatusCallbackEvent,omitempty"`
	// The HTTP method we should use to call `conference_status_callback`. Can be: `GET` or `POST` and defaults to `POST`.
	ConferenceStatusCallbackMethod string `json:"ConferenceStatusCallbackMethod,omitempty"`
	// Whether to trim leading and trailing silence from your recorded conference audio files. Can be: `trim-silence` or `do-not-trim` and defaults to `trim-silence`.
	ConferenceTrim string `json:"ConferenceTrim,omitempty"`
	// Whether to allow an agent to hear the state of the outbound call, including ringing or disconnect messages. Can be: `true` or `false` and defaults to `true`.
	EarlyMedia bool `json:"EarlyMedia,omitempty"`
	// Whether to end the conference when the participant leaves. Can be: `true` or `false` and defaults to `false`.
	EndConferenceOnExit bool `json:"EndConferenceOnExit,omitempty"`
	// The phone number, Client identifier, or username portion of SIP address that made this call. Phone numbers are in [E.164](https://www.twilio.com/docs/glossary/what-e164) format (e.g., +16175551212). Client identifiers are formatted `client:name`. If using a phone number, it must be a Twilio number or a Verified [outgoing caller id](https://www.twilio.com/docs/voice/api/outgoing-caller-ids) for your account. If the `to` parameter is a phone number, `from` must also be a phone number. If `to` is sip address, this value of `from` should be a username portion to be used to populate the P-Asserted-Identity header that is passed to the SIP endpoint.
	From string `json:"From"`
	// Jitter buffer size for the connecting participant. Twilio will use this setting to apply Jitter Buffer before participant's audio is mixed into the conference. Can be: `off`, `small`, `medium`, and `large`. Default to `large`.
	JitterBufferSize string `json:"JitterBufferSize,omitempty"`
	// A label for this participant. If one is supplied, it may subsequently be used to fetch, update or delete the participant.
	Label string `json:"Label,omitempty"`
	// The maximum number of participants in the conference. Can be a positive integer from `2` to `250`. The default value is `250`.
	MaxParticipants int32 `json:"MaxParticipants,omitempty"`
	// Whether the agent is muted in the conference. Can be `true` or `false` and the default is `false`.
	Muted bool `json:"Muted,omitempty"`
	// Whether to record the participant and their conferences, including the time between conferences. Can be `true` or `false` and the default is `false`.
	Record bool `json:"Record,omitempty"`
	// The recording channels for the final recording. Can be: `mono` or `dual` and the default is `mono`.
	RecordingChannels string `json:"RecordingChannels,omitempty"`
	// The URL that we should call using the `recording_status_callback_method` when the recording status changes.
	RecordingStatusCallback string `json:"RecordingStatusCallback,omitempty"`
	// The recording state changes that should generate a call to `recording_status_callback`. Can be: `in-progress`, `completed`, and `failed`. Separate multiple values with a space. The default value is `in-progress completed failed`.
	RecordingStatusCallbackEvent []string `json:"RecordingStatusCallbackEvent,omitempty"`
	// The HTTP method we should use when we call `recording_status_callback`. Can be: `GET` or `POST` and defaults to `POST`.
	RecordingStatusCallbackMethod string `json:"RecordingStatusCallbackMethod,omitempty"`
	// The audio track to record for the call. Can be: `inbound`, `outbound` or `both`. The default is `both`. `inbound` records the audio that is received by Twilio. `outbound` records the audio that is sent from Twilio. `both` records the audio that is received and sent by Twilio.
	RecordingTrack string `json:"RecordingTrack,omitempty"`
	// The [region](https://support.twilio.com/hc/en-us/articles/223132167-How-global-low-latency-routing-and-region-selection-work-for-conferences-and-Client-calls) where we should mix the recorded audio. Can be:`us1`, `ie1`, `de1`, `sg1`, `br1`, `au1`, or `jp1`.
	Region string `json:"Region,omitempty"`
	// The SIP password for authentication.
	SipAuthPassword string `json:"SipAuthPassword,omitempty"`
	// The SIP username used for authentication.
	SipAuthUsername string `json:"SipAuthUsername,omitempty"`
	// Whether to start the conference when the participant joins, if it has not already started. Can be: `true` or `false` and the default is `true`. If `false` and the conference has not started, the participant is muted and hears background music until another participant starts the conference.
	StartConferenceOnEnter bool `json:"StartConferenceOnEnter,omitempty"`
	// The URL we should call using the `status_callback_method` to send status information to your application.
	StatusCallback string `json:"StatusCallback,omitempty"`
	// The conference state changes that should generate a call to `status_callback`. Can be: `initiated`, `ringing`, `answered`, and `completed`. Separate multiple values with a space. The default value is `completed`.
	StatusCallbackEvent []string `json:"StatusCallbackEvent,omitempty"`
	// The HTTP method we should use to call `status_callback`. Can be: `GET` and `POST` and defaults to `POST`.
	StatusCallbackMethod string `json:"StatusCallbackMethod,omitempty"`
	// The number of seconds that we should allow the phone to ring before assuming there is no answer. Can be an integer between `5` and `600`, inclusive. The default value is `60`. We always add a 5-second timeout buffer to outgoing calls, so  value of 10 would result in an actual timeout that was closer to 15 seconds.
	Timeout int32 `json:"Timeout,omitempty"`
	// The phone number, SIP address, or Client identifier that received this call. Phone numbers are in [E.164](https://www.twilio.com/docs/glossary/what-e164) format (e.g., +16175551212). SIP addresses are formatted as `sip:name@company.com`. Client identifiers are formatted `client:name`. [Custom parameters](https://www.twilio.com/docs/voice/api/conference-participant-resource#custom-parameters) may also be specified.
	To string `json:"To"`
	// The HTTP method we should use to call `wait_url`. Can be `GET` or `POST` and the default is `POST`. When using a static audio file, this should be `GET` so that we can cache the file.
	WaitMethod string `json:"WaitMethod,omitempty"`
	// The URL we should call using the `wait_method` for the music to play while participants are waiting for the conference to start. The default value is the URL of our standard hold music. [Learn more about hold music](https://www.twilio.com/labs/twimlets/holdmusic).
	WaitUrl string `json:"WaitUrl,omitempty"`
}