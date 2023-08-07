package model

type EmailData struct {
	Country      string
	Provider     string
	DeliveryTime int
}
type ChEmail struct {
	Data map[string][][]EmailData
	Err  error
}
type EmailByDeliveryTime []EmailData

func (b EmailByDeliveryTime) Len() int {
	return len(b)
}
func (b EmailByDeliveryTime) Less(i, j int) bool {
	return b[i].DeliveryTime < b[j].DeliveryTime

}
func (b EmailByDeliveryTime) Swap(i, j int) {
	b[i], b[j] = b[j], b[i]
}
