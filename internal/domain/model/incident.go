package model

type IncidentData struct {
	Topic  string
	Status string
}
type ChIncident struct {
	Data []IncidentData
	Err  error
}
type IncidentByStatus []IncidentData

func (b IncidentByStatus) Len() int {
	return len(b)
}
func (b IncidentByStatus) Less(i, j int) bool {
	return b[i].Status < b[j].Status

}
func (b IncidentByStatus) Swap(i, j int) {
	b[i], b[j] = b[j], b[i]
}
