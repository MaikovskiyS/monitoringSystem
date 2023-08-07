package model

type SMSData struct {
	Country      string
	Bandwidth    string
	ResponseTime string
	Provider     string
}
type ChSms struct {
	Data [][]SMSData
	Err  error
}
type SMSByProvider []SMSData

func (s SMSByProvider) Len() int {
	return len(s)
}
func (s SMSByProvider) Less(i, j int) bool {
	return s[i].Provider < s[j].Provider
}
func (s SMSByProvider) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

type SMSByCountry []SMSData

func (c SMSByCountry) Len() int {
	return len(c)
}
func (c SMSByCountry) Less(i, j int) bool {
	return c[i].Country < c[j].Country
}
func (c SMSByCountry) Swap(i, j int) {
	c[i], c[j] = c[j], c[i]
}
