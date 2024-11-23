package models

type CareerForm struct {
	Id         string `db:"id"`
	Created_at string `db:"created_at"`
	Name       string `db:"name"`
	Related_to string `db:"related_to"`
	Templates  []byte `db:"templates"`
}

type CareerTemplates struct {
	On_applied string `json:"on_applied" db:"on_applied"`
	Accepted   string `json:"Accepted" db:"Accepted"`
	Rejected   string `json:"Rejected" db:"Rejected"`
}

type CourseTemplates struct {
	On_boarding    string `json:"on_boarding" db:"on_boarding"`
	Course_Payment string `json:"course_payment" db:"course_payment"`
	Installments   string `json:"installments" db:"installments"`
}

type EventsTemplates struct {
	Ticket        string `json:"ticket" db:"ticket"`
	Event_Payment string `json:"event_payment" db:"event_payment"`
}

type ServicesTemplates struct {
	Response string `json:"response" db:"response"`
}

type GlobalTemplates struct {
	Email_Response string `json:"email_response" db:"email_response"`
}
