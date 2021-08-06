package entities

type Notification struct {
	Type        string `json:"type"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
	Text        string `json:"text"`
	PurchaseID  string `json:"purchase_id"`
}
