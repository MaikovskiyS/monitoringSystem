package service

import (
	"diploma/internal/domain/model"
	"diploma/internal/transport/http/dto"
	"sort"
	"sync"

	"github.com/sirupsen/logrus"
)

const (
	timeForTicket = 60 / 18
)

type Dependences struct {
	Email     Emailer
	Mms       MMSer
	Sms       SMSer
	Voicecall VoiceCaller
	Billing   Billinger
	Support   Supporter
	Incident  Incidenter
}
type service struct {
	smsChan      chan model.ChSms
	mmsChan      chan model.ChMms
	emailChan    chan model.ChEmail
	voiceChan    chan model.ChVoice
	billingChan  chan model.ChBilling
	supportChan  chan model.ChSupport
	incidentChan chan model.ChIncident

	deps Dependences
	l    *logrus.Logger
}

func New(l *logrus.Logger, d Dependences) *service {
	smsch := make(chan model.ChSms, 1)
	mmsch := make(chan model.ChMms, 1)
	emailch := make(chan model.ChEmail, 1)
	voicech := make(chan model.ChVoice, 1)
	billingch := make(chan model.ChBilling, 1)
	supportch := make(chan model.ChSupport, 1)
	incidentch := make(chan model.ChIncident, 1)

	return &service{
		smsChan:      smsch,
		mmsChan:      mmsch,
		emailChan:    emailch,
		voiceChan:    voicech,
		billingChan:  billingch,
		supportChan:  supportch,
		incidentChan: incidentch,

		deps: d,
		l:    l,
	}
}
func (svc *service) GetResultData() (dto.ResultT, error) {
	var r dto.ResultT
	result, err := svc.CollectData()
	if err != nil {
		svc.l.Info(err)
		r.Status = false
		r.Error = err.Error()
		return r, err
	}
	r.Status = true
	r.Data = result
	r.Error = "nil"
	return r, nil

}
func (svc *service) CollectData() (dto.ResultSetT, error) {
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		smses, err := svc.deps.Sms.GetData()
		if err != nil {
			svc.smsChan <- model.ChSms{
				Err: err,
			}
			return
		}
		smsResult, err := BuildSmsResult(smses)
		if err != nil {
			svc.smsChan <- model.ChSms{
				Err: err,
			}
			return
		}
		svc.smsChan <- model.ChSms{
			Data: smsResult,
			Err:  nil,
		}
	}()
	wg.Add(1)
	go func() {
		defer wg.Done()
		mmses, err := svc.deps.Mms.GetData()
		if err != nil {
			svc.mmsChan <- model.ChMms{
				Err: err,
			}
			return
		}
		mmsResult, err := BuildMmsResult(mmses)
		if err != nil {
			svc.mmsChan <- model.ChMms{
				Err: err,
			}
			return
		}
		svc.mmsChan <- model.ChMms{
			Data: mmsResult,
		}
	}()
	wg.Add(1)
	go func() {
		defer wg.Done()
		voicecalls, err := svc.deps.Voicecall.GetData()
		if err != nil {
			svc.voiceChan <- model.ChVoice{
				Err: err,
			}
			return
		}
		svc.voiceChan <- model.ChVoice{
			Data: voicecalls,
		}
	}()
	wg.Add(1)
	go func() {
		defer wg.Done()
		emails, err := svc.deps.Email.GetData()
		if err != nil {
			svc.emailChan <- model.ChEmail{
				Err: err,
			}
			return
		}
		emailResult, err := BuildEmailResult(emails)
		if err != nil {
			svc.emailChan <- model.ChEmail{
				Err: err,
			}
			return
		}
		svc.emailChan <- model.ChEmail{
			Data: emailResult,
		}
	}()
	wg.Add(1)
	go func() {
		defer wg.Done()
		billing, err := svc.deps.Billing.GetData()
		if err != nil {
			svc.billingChan <- model.ChBilling{
				Err: err,
			}
			return
		}
		svc.billingChan <- model.ChBilling{
			Data: billing,
		}
	}()
	wg.Add(1)
	go func() {
		defer wg.Done()
		supports, err := svc.deps.Support.GetData()
		if err != nil {
			svc.supportChan <- model.ChSupport{
				Err: err,
			}
			return
		}
		supportResult, err := BuildSupportResult(supports)
		if err != nil {
			svc.supportChan <- model.ChSupport{
				Err: err,
			}
			return
		}
		svc.supportChan <- model.ChSupport{
			Data: supportResult,
		}
	}()
	wg.Add(1)
	go func() {
		defer wg.Done()
		incidents, err := svc.deps.Incident.GetData()
		if err != nil {
			svc.incidentChan <- model.ChIncident{
				Err: err,
			}
			return
		}
		incidentResult := BuildIncidentResult(incidents)
		svc.incidentChan <- model.ChIncident{
			Data: incidentResult,
		}
	}()

	wg.Wait()
	result, err := svc.CreateResultDTO()
	if err != nil {
		return dto.ResultSetT{}, err
	}

	return result, nil
}

func BuildSmsResult(input []model.SMSData) (output [][]model.SMSData, err error) {
	var byCountry model.SMSByCountry
	var byProvider model.SMSByProvider
	for _, v := range input {
		byCountry = append(byCountry, v)
		byProvider = append(byProvider, v)
	}
	sort.Sort(byCountry)
	sort.Sort(byProvider)
	output = append(output, byCountry, byProvider)
	return
}
func BuildMmsResult(input []model.MMSData) (output [][]model.MMSData, err error) {
	var byCountry model.MMSByCountry
	var byProvider model.MMSByProvider
	for _, v := range input {
		byCountry = append(byCountry, v)
		byProvider = append(byProvider, v)
	}
	sort.Sort(byCountry)
	sort.Sort(byProvider)
	output = append(output, byCountry, byProvider)
	return
}
func BuildEmailResult(input []model.EmailData) (map[string][][]model.EmailData, error) {
	output := make(map[string][][]model.EmailData)
	var emails model.EmailByDeliveryTime
	for _, v := range input {
		emails = append(emails, v)
	}
	sort.Sort(emails)
	for _, v := range emails {
		matrix, ok := output[v.Country]
		if ok {
			matrix[0] = append(matrix[0], v)
			continue

		}
		arr1 := make([]model.EmailData, 0)
		matrix1 := make([][]model.EmailData, 2)
		arr1 = append(arr1, v)
		matrix1[0] = arr1
		output[v.Country] = matrix1
	}
	var forsort model.EmailByDeliveryTime
	for _, matrix := range output {
		arr := matrix[0]
		forsort = arr
		sort.Sort(forsort)
		slow := forsort[:3]
		fast := forsort[len(forsort)-3:]
		matrix[0] = fast
		matrix[1] = slow
	}
	return output, nil
}
func BuildSupportResult(input []model.SupportData) (output [2]int, err error) {
	var tasks int
	for _, v := range input {
		tasks += v.ActiveTickets
		if tasks < 9 {
			output[0] = 1
		} else if tasks > 16 {
			output[0] = 3
		} else {
			output[0] = 2
		}
		output[1] = tasks * timeForTicket
	}
	return output, nil
}
func BuildIncidentResult(input model.IncidentByStatus) []model.IncidentData {
	sort.Sort(input)
	return input
}
func (s *service) CreateResultDTO() (dto.ResultSetT, error) {
	var result dto.ResultT
	sms := <-s.smsChan
	if sms.Err != nil {
		return dto.ResultSetT{}, sms.Err
	}
	result.Data.SMS = sms.Data
	mms := <-s.mmsChan
	if mms.Err != nil {
		return dto.ResultSetT{}, mms.Err
	}
	result.Data.MMS = mms.Data

	email := <-s.emailChan
	if email.Err != nil {
		return dto.ResultSetT{}, email.Err
	}
	result.Data.Email = email.Data

	voice := <-s.voiceChan
	if voice.Err != nil {
		return dto.ResultSetT{}, voice.Err
	}
	result.Data.VoiceCall = voice.Data

	billing := <-s.billingChan
	if billing.Err != nil {
		return dto.ResultSetT{}, billing.Err
	}
	result.Data.Billing = billing.Data

	support := <-s.supportChan
	if support.Err != nil {
		return dto.ResultSetT{}, support.Err
	}
	result.Data.Support = support.Data

	incident := <-s.incidentChan
	if incident.Err != nil {
		return dto.ResultSetT{}, incident.Err
	}
	result.Data.Incidents = incident.Data
	return result.Data, nil
}
