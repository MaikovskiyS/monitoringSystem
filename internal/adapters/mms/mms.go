package mms

import (
	"diploma/internal/config"
	"diploma/internal/domain/model"
	"encoding/json"
	"net/http"

	"github.com/sirupsen/logrus"
)

type mms struct {
	cfg *config.MMS
	l   *logrus.Logger
}

func New(l *logrus.Logger, cfg *config.MMS) *mms {
	return &mms{
		cfg: cfg,
		l:   l,
	}
}
func (m *mms) GetData() (mmses []model.MMSData, err error) {
	var input []Data

	r, err := http.Get(m.cfg.Url)
	if err != nil {
		m.l.Info(err)
		return nil, err
	}
	if r.StatusCode != 200 {
		m.l.Info(r.StatusCode)
		return mmses, err
	}

	err = json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		m.l.Info("cant decode request body", err)
		return mmses, err
	}
	for _, data := range input {
		err := data.Validate(m.cfg.CountryCodes)
		if err != nil {
			continue
		}
		mms := model.MMSData{Country: data.Country, Provider: data.Provider, Bandwidth: data.Bandwidth, ResponseTime: data.ResponseTime}
		mmses = append(mmses, mms)
	}

	return mmses, nil
}
