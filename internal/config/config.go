package config

import (
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"os"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

type Envs struct {
	sms        string
	mms        string
	email      string
	billing    string
	voicecall  string
	support    string
	incident   string
	contryCode string
}
type Config struct {
	Adapters *Adapters
}
type Adapters struct {
	SMS       *SMS
	MMS       *MMS
	VoiceCall *VoiceCall
	Email     *Email
	Billing   *Billing
	Support   *Support
	Incident  *Incident
}

type SMS struct {
	CountryCodes map[string]string
	FilePath     string `env-required:"true"  env:"SMS_PATH"`
}
type MMS struct {
	CountryCodes map[string]string
	Url          string `env-required:"true"  env:"MMS_URL"`
}
type VoiceCall struct {
	CountryCodes map[string]string
	FilePath     string `env-required:"true"  env:"VOICECALL_PATH"`
}
type Email struct {
	CountryCodes   map[string]string
	EmailProviders map[string]bool
	FilePath       string `env-required:"true"  env:"EMAIL_PATH"`
}
type Billing struct {
	CountryCodes map[string]string
	FilePath     string `env-required:"true"  env:"BILLING_PATH"`
}
type Support struct {
	Url string `env-required:"true"  env:"SUPPORT_URL"`
}
type Incident struct {
	Url string `env-required:"true"  env:"INCIDENT_URL"`
}
type CountryCode struct {
	CountryCodes map[string]string
	FilePath     string `env-required:"true"  env:"COUNTRY_CODE"`
}

func NewConfig(l *logrus.Logger) (*Config, error) {
	fmt.Println("in config")
	if err := godotenv.Load(); err != nil {
		l.Fatal("No .env file found")
	}
	envs, err := GetEnvs()
	if err != nil {
		l.Info("get envs err in config", err)
		return nil, err
	}
	c := make(map[string]string, 250)
	f, err := os.Open(envs.contryCode)
	if err != nil {
		l.Info(err)
	}
	r := csv.NewReader(f)
	r.Comma = ';'
	r.FieldsPerRecord = 2
	for {
		code, err := r.Read()
		if err != nil {
			if err == io.EOF {
				break
			}
			l.Info("cofig read country code file err:  ", err)
			break
		}
		c[code[1]] = code[0]
	}
	providers := []string{"Gmail", "Yahoo", "Hotmail", "MSN", "Orange", "Comcast", "AOL", "Live", "RediffMail", "GMX", "Protonmail", "Yandex", "Mail.ru"}
	emailproviders := make(map[string]bool)
	for _, v := range providers {
		emailproviders[v] = true
	}
	return &Config{
		Adapters: &Adapters{
			SMS: &SMS{
				FilePath:     envs.sms,
				CountryCodes: c,
			},
			VoiceCall: &VoiceCall{
				FilePath:     envs.voicecall,
				CountryCodes: c,
			},
			Email: &Email{
				FilePath:       envs.email,
				EmailProviders: emailproviders,
				CountryCodes:   c,
			},
			Billing: &Billing{
				FilePath:     envs.billing,
				CountryCodes: c,
			},
			MMS: &MMS{
				CountryCodes: c,
				Url:          envs.mms,
			},
			Support: &Support{
				Url: envs.support,
			},
			Incident: &Incident{
				Url: envs.incident,
			},
		},
	}, nil
}
func GetEnvs() (*Envs, error) {

	smsPath, ok := os.LookupEnv("SMS_PATH")
	if !ok {
		return nil, errors.New("sms_path env error")
	}
	emailPath, ok := os.LookupEnv("EMAIL_PATH")
	if !ok {
		return nil, errors.New("EMAIL_PATH env error")
	}
	voiceCallPath, ok := os.LookupEnv("VOICECALL_PATH")
	if !ok {
		return nil, errors.New("VOICECALL_PATH env error")
	}
	billingPath, ok := os.LookupEnv("BILLING_PATH")
	if !ok {
		return nil, errors.New("BILLING_PATH env error")
	}
	mmsUrl, ok := os.LookupEnv("MMS_URL")
	if !ok {
		return nil, errors.New("MMS_URL env error")
	}
	supportUrl, ok := os.LookupEnv("SUPPORT_URL")
	if !ok {
		return nil, errors.New("SUPPORT_URL env error")
	}
	IncidentUrl, ok := os.LookupEnv("INCIDENT_URL")
	if !ok {
		return nil, errors.New("INCIDENT_URL env error")
	}
	countryCode, ok := os.LookupEnv("COUNTRY_CODE")
	if !ok {
		return nil, errors.New("COUNTRY_CODE env error")
	}
	return &Envs{
		sms:        smsPath,
		mms:        mmsUrl,
		email:      emailPath,
		billing:    billingPath,
		voicecall:  voiceCallPath,
		support:    supportUrl,
		incident:   IncidentUrl,
		contryCode: countryCode,
	}, nil
}
