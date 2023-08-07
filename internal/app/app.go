package app

import (
	"diploma/internal/adapters/billing"
	"diploma/internal/adapters/email"
	"diploma/internal/adapters/incident"
	"diploma/internal/adapters/mms"
	"diploma/internal/adapters/sms"
	"diploma/internal/adapters/support"
	"diploma/internal/adapters/voicecall"
	"diploma/internal/cache"
	"diploma/internal/config"
	"diploma/internal/domain/service"
	"diploma/internal/transport/http/handler"
	"diploma/internal/transport/http/server"
	"diploma/simulator"

	"github.com/sirupsen/logrus"
)

func Run() {
	go simulator.StartSimulator()
	l := logrus.New()
	cfg, err := config.NewConfig(l)
	if err != nil {
		l.Fatal(err)
	}
	sms := sms.New(l, cfg.Adapters.SMS)
	mms := mms.New(l, cfg.Adapters.MMS)
	email := email.New(l, cfg.Adapters.Email)
	billing := billing.New(l, cfg.Adapters.Billing)
	voiceCall := voicecall.New(l, cfg.Adapters.VoiceCall)
	support := support.New(l, cfg.Adapters.Support)
	incident := incident.New(l, cfg.Adapters.Incident)
	deps := service.Dependences{
		Email:     email,
		Sms:       sms,
		Mms:       mms,
		Billing:   billing,
		Voicecall: voiceCall,
		Support:   support,
		Incident:  incident,
	}
	cache := cache.New(l)
	svc := service.New(l, deps)
	handler := handler.New(l, svc, cache)
	handler.RegisterRoutes()
	server := server.New(handler)
	server.Start()

}
