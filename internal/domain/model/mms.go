package model

type MMSData struct {
	Country      string
	Provider     string
	Bandwidth    string
	ResponseTime string
}
type ChMms struct {
	Data [][]MMSData
	Err  error
}
type MMSByProvider []MMSData

func (s MMSByProvider) Len() int {
	return len(s)
}
func (s MMSByProvider) Less(i, j int) bool {
	return s[i].Provider < s[j].Provider
}
func (s MMSByProvider) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

type MMSByCountry []MMSData

func (c MMSByCountry) Len() int {
	return len(c)
}
func (c MMSByCountry) Less(i, j int) bool {
	return c[i].Country < c[j].Country
}
func (c MMSByCountry) Swap(i, j int) {
	c[i], c[j] = c[j], c[i]
}
