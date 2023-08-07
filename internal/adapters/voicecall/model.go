package voicecall

import (
	"errors"
	"strconv"
)

type Data struct {
	Country             string
	Bandwidth           string
	ResponseTime        string
	Provider            string
	ConnectionStability float32
	TTFB                int
	VoicePurity         int
	MedianOfCallsTime   int
}

func NewData(data []string) (*Data, error) {
	con, err := strconv.ParseFloat(data[4], 32)
	if err != nil {
		return &Data{}, err
	}
	ttfb, err := strconv.Atoi(data[5])
	if err != nil {
		return &Data{}, err
	}
	pir, err := strconv.Atoi(data[6])
	if err != nil {
		return &Data{}, err
	}
	time, err := strconv.Atoi(data[7])
	if err != nil {
		return &Data{}, err
	}
	d := &Data{
		Country:             data[0],
		Bandwidth:           data[1],
		ResponseTime:        data[2],
		Provider:            data[3],
		ConnectionStability: float32(con),
		TTFB:                ttfb,
		VoicePurity:         pir,
		MedianOfCallsTime:   time,
	}

	return d, nil
}
func (d *Data) Validate(codes map[string]string) error {
	if _, ok := codes[d.Country]; !ok {
		return errors.New("invalid country code")
	}

	band, err := strconv.Atoi(d.Bandwidth)
	if err != nil {
		return err
	}
	if band < 0 || band > 100 {
		return errors.New("wrong bandwidth value")
	}

	if d.Provider != "TransparentCalls" && d.Provider != "E-Voice" && d.Provider != "JustPhone" {
		return errors.New("wrong provaider name")
	}
	return nil
}
