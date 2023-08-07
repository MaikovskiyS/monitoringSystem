package dto

import "diploma/internal/domain/model"

type ResultT struct {
	Status bool       `json:"status"`
	Data   ResultSetT `json:"data"`
	Error  string     `json:"error"`
}
type ResultSetT struct {
	SMS       [][]model.SMSData              `json:"sms"`
	MMS       [][]model.MMSData              `json:"mms"`
	VoiceCall []model.VoiceCallData          `json:"voice_call"`
	Email     map[string][][]model.EmailData `json:"email"`
	Billing   model.BillingData              `json:"billing"`
	Support   [2]int                         `json:"support"`
	Incidents []model.IncidentData           `json:"incident"`
}
