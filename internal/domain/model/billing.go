package model

type BillingData struct {
	CreateCustomer bool
	Purchase       bool
	Payout         bool
	Recurring      bool
	FraudControl   bool
	CheckoutPage   bool
}
type ChBilling struct {
	Data BillingData
	Err  error
}
