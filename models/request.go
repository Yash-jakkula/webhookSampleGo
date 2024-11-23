package models

type DataStruct struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Phone string `json:"phone"`
}

type RecordStruct struct {
	ID        string     `json:"id"`
	Status    string     `json:"status"`
	Data      DataStruct `json:"data"`
	FormID    string     `json:"form_id"`
	CreatedAt string     `json:"created_at"`
}

type ResponseStruct struct {
	Type   string       `json:"type"`
	Table  string       `json:"table"`
	Record RecordStruct `json:"record"`
}
