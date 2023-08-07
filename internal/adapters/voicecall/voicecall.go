package voicecall

import (
	"diploma/internal/config"
	"diploma/internal/domain/model"
	"encoding/csv"
	"io"
	"os"

	"github.com/sirupsen/logrus"
)

type voiceCall struct {
	cfg        *config.VoiceCall
	fileFields int
	l          *logrus.Logger
}

func New(l *logrus.Logger, cfg *config.VoiceCall) *voiceCall {
	return &voiceCall{
		cfg:        cfg,
		fileFields: 8,
		l:          l,
	}
}

func (v *voiceCall) GetData() ([]model.VoiceCallData, error) {
	var output []model.VoiceCallData
	f, err := os.Open(v.cfg.FilePath)
	if err != nil {
		v.l.Info("read voicecalldata file err")
		return nil, err
	}
	defer f.Close()
	r := csv.NewReader(f)
	r.Comma = ';'
	r.FieldsPerRecord = v.fileFields

	for {
		w, err := r.Read()
		if err != nil {
			if err == io.EOF {
				break
			}
			v.l.Infof("email file read rows err %s", err)
			continue
		}
		data, err := NewData(w)
		data.Validate(v.cfg.CountryCodes)
		if err != nil {
			v.l.Info("voicecall err validation: ", err)
			continue
		}
		voicecall := model.VoiceCallData{
			Country:             data.Country,
			Provider:            data.Provider,
			Bandwidth:           data.Bandwidth,
			ResponseTime:        data.ResponseTime,
			ConnectionStability: data.ConnectionStability,
			TTFB:                data.TTFB,
			MedianOfCallsTime:   data.MedianOfCallsTime,
			VoicePurity:         data.VoicePurity,
		}

		output = append(output, voicecall)
	}
	return output, nil
}
