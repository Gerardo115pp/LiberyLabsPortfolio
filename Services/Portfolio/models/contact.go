package models

type ContactInfo struct {
	Email     string `json:"email"`
	Name      string `json:"name"`
	Budget    string `json:"budget"`
	Company   string `json:"company"`
	StartDate string `json:"start_date"`
	EndDate   string `json:"end_date"`
	Subject   string `json:"subject"`
	SenderIP  string `json:"sender_ip"`
}

func (c *ContactInfo) FormatTextMessage() string {
	var text_message string = "*New contact submission*\n\n"

	text_message += "Name: " + c.Name + "\n"
	text_message += "Email: " + c.Email + "\n"
	text_message += "Budget: " + c.Budget + "\n"
	text_message += "Company: " + c.Company + "\n"
	text_message += "IP: " + c.SenderIP + "\n\n"
	text_message += "Start Date: " + c.StartDate + "\n"
	text_message += "End Date: " + c.EndDate + "\n\n"
	text_message += c.Subject

	return text_message
}
