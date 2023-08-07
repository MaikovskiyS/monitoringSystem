package incident

import (
	"diploma/internal/config"
	"diploma/internal/domain/model"
	"encoding/json"
	"net/http"

	"github.com/sirupsen/logrus"
)

type incident struct {
	cfg *config.Incident
	l   *logrus.Logger
}

func New(l *logrus.Logger, cfg *config.Incident) *incident {
	return &incident{
		cfg: cfg,
		l:   l,
	}
}
func (i *incident) GetData() ([]model.IncidentData, error) {
	var data []model.IncidentData
	r, err := http.Get(i.cfg.Url)
	if err != nil {
		i.l.Info(err)
		return []model.IncidentData{}, err
	}
	if r.StatusCode != 200 {
		i.l.Info(r.StatusCode)
		return []model.IncidentData{}, err
	}

	err = json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		i.l.Info("cant decode request body", err)
		return []model.IncidentData{}, err
	}

	return data, nil
}
