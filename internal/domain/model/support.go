package model

type SupportData struct {
	Topic         string `json:"topic"`
	ActiveTickets int    `json:"active_tickets"`
}
type ChSupport struct {
	Data [2]int
	Err  error
}
