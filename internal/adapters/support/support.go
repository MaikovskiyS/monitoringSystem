package support

import (
	"diploma/internal/config"
	"diploma/internal/domain/model"
	"encoding/json"
	"net/http"

	"github.com/sirupsen/logrus"
)

type support struct {
	cfg *config.Support
	l   *logrus.Logger
}

func New(l *logrus.Logger, cfg *config.Support) *support {
	return &support{
		cfg: cfg,
		l:   l,
	}
}
func (s *support) GetData() ([]model.SupportData, error) {
	var input []model.SupportData

	r, err := http.Get(s.cfg.Url)
	if err != nil {
		s.l.Info(err)
		return nil, err
	}
	if r.StatusCode != 200 {
		s.l.Info(r.StatusCode)
		return nil, err
	}

	err = json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		s.l.Info("cant decode support request body", err)
		return nil, err
	}
	return input, nil
}
