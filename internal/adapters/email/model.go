package email

import (
	"errors"
	"strconv"
)

type Data struct {
	Country      string
	Provider     string
	DeliveryTime int
}

func NewData(w []string) (*Data, error) {
	time, err := strconv.Atoi(w[2])
	if err != nil {
		return &Data{}, err
	}

	return &Data{
		Country:      w[0],
		Provider:     w[1],
		DeliveryTime: time,
	}, nil
}
func (d *Data) Validate(codes map[string]string, providers map[string]bool) error {
	if _, ok := codes[d.Country]; !ok {
		return errors.New("invalid country code")
	}

	if ok := providers[d.Provider]; !ok {
		return errors.New("wrong provaider name")
	}
	return nil
}
