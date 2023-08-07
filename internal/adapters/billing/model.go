package billing

import (
	"math"
	"strconv"
)

type Data struct {
	CreateCustomer bool
	Purchase       bool
	Payout         bool
	Recurring      bool
	FraudControl   bool
	CheckoutPage   bool
}

func NewData(w []string) *Data {
	str := w[0]
	arr := make([]bool, len(str))

	for i := 0; i < len(str)-1; i++ {
		num, _ := strconv.ParseUint(string(str[i]), 2, 32)

		switch {
		case num == 0:
			arr[i] = false
		case num == 1:
			arr[i] = true
		}
	}
	return &Data{
		CreateCustomer: arr[0],
		Purchase:       arr[1],
		Payout:         arr[2],
		Recurring:      arr[3],
		FraudControl:   arr[4],
		CheckoutPage:   arr[5],
	}
}
func (d *Data) ValidateTODO(w []string) (float64, error) {
	q := w[0]
	var sum float64
	var j uint8
	for i := len(w[0]) - 1; i >= 0; i-- {
		str := (string(q[i]))
		u, err := strconv.ParseUint(str, 2, 32)
		if err != nil {
			return 0.0, err
		}
		if u == 0 {
			j++
			continue
		}
		num := math.Pow(2, float64(j))
		sum += num
		j++

	}
	//Todo задание 5 в Billing в описании дипломной работы
	//file:///C:/Users/User/Downloads/Финальная%20работа%20курса%20«Go-разработчик»%20(1).pdf
	return sum, nil
}
