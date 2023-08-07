package email

import (
	"diploma/internal/config"
	"diploma/internal/domain/model"
	"encoding/csv"
	"io"
	"os"

	"github.com/sirupsen/logrus"
)

type email struct {
	cfg        *config.Email
	l          *logrus.Logger
	fileFields int
}

func New(l *logrus.Logger, cfg *config.Email) *email {
	return &email{
		cfg:        cfg,
		l:          l,
		fileFields: 3,
	}
}
func (e *email) GetData() ([]model.EmailData, error) {
	var emails []model.EmailData
	f, err := os.Open(e.cfg.FilePath)
	if err != nil {
		e.l.Info("read emaildata file err")
		return nil, err
	}
	defer f.Close()
	r := csv.NewReader(f)
	r.Comma = ';'
	r.FieldsPerRecord = e.fileFields

	for {
		w, err := r.Read()
		if err != nil {
			if err == io.EOF {
				break
			}
			e.l.Infof("email file read rows err %s", err)
			continue
		}
		data, err := NewData(w)
		if err != nil {
			e.l.Info("cant creae data")
		}
		err = data.Validate(e.cfg.CountryCodes, e.cfg.EmailProviders)
		if err != nil {
			e.l.Info("err in validation: ", err)
			continue
		}
		email := model.EmailData{Country: data.Country, Provider: data.Provider, DeliveryTime: data.DeliveryTime}

		emails = append(emails, email)

	}

	return emails, nil
}
