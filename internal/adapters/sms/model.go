package sms

import (
	"errors"
	"strconv"
)

type Data struct {
	Country      string
	Bandwidth    string
	ResponseTime string
	Provider     string
}

func NewData(data []string, codes map[string]string) (*Data, error) {
	d := &Data{
		Country:      data[0],
		Bandwidth:    data[1],
		ResponseTime: data[2],
		Provider:     data[3],
	}
	err := d.Validate(codes)
	if err != nil {
		return &Data{}, err
	}

	return d, nil
}
func (o *Data) Validate(codes map[string]string) error {
	if _, ok := codes[o.Country]; !ok {
		return errors.New("invalid country code")
	}
	band, err := strconv.Atoi(o.Bandwidth)
	if err != nil {
		return err
	}
	if band < 0 || band > 100 {
		return errors.New("wrong bandwidth value")
	}
	if o.Provider != "Topolo" && o.Provider != "Rond" && o.Provider != "Kildy" {
		return errors.New("wrong provaider name")
	}
	return nil
}
