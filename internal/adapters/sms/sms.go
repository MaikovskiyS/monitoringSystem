package sms

import (
	"diploma/internal/config"
	"diploma/internal/domain/model"
	"encoding/csv"
	"io"
	"os"

	"github.com/sirupsen/logrus"
)

type sms struct {
	cfg        *config.SMS
	fileFields int
	l          *logrus.Logger
}

func New(l *logrus.Logger, cfg *config.SMS) *sms {
	return &sms{
		cfg:        cfg,
		fileFields: 4,
		l:          l,
	}

}

// Getdata from smsfile
func (s *sms) GetData() ([]model.SMSData, error) {
	var output []model.SMSData

	f, err := os.Open(s.cfg.FilePath)
	if err != nil {
		s.l.Info("read smsdata file err")
		return nil, err
	}
	defer f.Close()
	r := csv.NewReader(f)
	r.Comma = ';'
	r.FieldsPerRecord = s.fileFields

	for {
		w, err := r.Read()
		if err != nil {
			if err == io.EOF {
				break
			}
			s.l.Infof("sms file read rows err %s", err)
			continue
		}
		data, err := NewData(w, s.cfg.CountryCodes)
		if err != nil {
			s.l.Info("sms err: ", err)
			continue
		}
		smsData := model.SMSData{Country: data.Country, Bandwidth: data.Bandwidth, ResponseTime: data.ResponseTime, Provider: data.Provider}
		output = append(output, smsData)
	}
	return output, nil
}
