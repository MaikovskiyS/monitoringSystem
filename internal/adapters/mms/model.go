package mms

import "errors"

type Data struct {
	Country      string `json:"country"`
	Provider     string `json:"provider"`
	Bandwidth    string `json:"bandwidth"`
	ResponseTime string `json:"response_time"`
}

func (d *Data) Validate(codes map[string]string) error {
	if _, ok := codes[d.Country]; !ok {
		return errors.New("invalid country code")
	}
	d.Country = codes[d.Country]
	if d.Provider != "Topolo" && d.Provider != "Rond" && d.Provider != "Kildy" {
		return errors.New("wrong provaider name")
	}
	return nil
}
