package billing

import (
	"diploma/internal/config"
	"diploma/internal/domain/model"
	"encoding/csv"
	"io"
	"os"

	"github.com/sirupsen/logrus"
)

type billing struct {
	cfg *config.Billing
	l   *logrus.Logger
}

func New(l *logrus.Logger, cfg *config.Billing) *billing {
	return &billing{
		cfg: cfg,
		l:   l,
	}
}
func (b *billing) GetData() (model.BillingData, error) {
	var billing model.BillingData
	f, err := os.Open(b.cfg.FilePath)
	if err != nil {
		b.l.Info("read billingdata file err")
		return model.BillingData{}, err
	}
	defer f.Close()
	r := csv.NewReader(f)
	r.Comma = ';'

	for {
		w, err := r.Read()
		if err != nil {
			if err == io.EOF {
				break
			}
			b.l.Infof("billing file read rows err %s", err)
			continue
		}
		data := NewData(w)
		billing = model.BillingData{
			CreateCustomer: data.CreateCustomer,
			Purchase:       data.Purchase,
			Payout:         data.Payout,
			Recurring:      data.Recurring,
			FraudControl:   data.FraudControl,
			CheckoutPage:   data.CheckoutPage,
		}
	}
	return billing, nil
}
