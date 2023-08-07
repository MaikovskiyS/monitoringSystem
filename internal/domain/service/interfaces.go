package service

import "diploma/internal/domain/model"

type SMSer interface {
	GetData() ([]model.SMSData, error)
}
type MMSer interface {
	GetData() ([]model.MMSData, error)
}
type VoiceCaller interface {
	GetData() ([]model.VoiceCallData, error)
}
type Emailer interface {
	GetData() ([]model.EmailData, error)
}
type Billinger interface {
	GetData() (model.BillingData, error)
}
type Supporter interface {
	GetData() ([]model.SupportData, error)
}
type Incidenter interface {
	GetData() ([]model.IncidentData, error)
}
