package controller

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"webhookapi/config"
	"webhookapi/models"
	"webhookapi/utils"
)

func ReminderWebhook(w http.ResponseWriter, r *http.Request) {
	DB := config.Connect()
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	var responseRequest models.ResponseStruct
	err = json.Unmarshal(body, &responseRequest)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	fmt.Println(responseRequest.Record.Status, "responseRequest")

	var query = `SELECT * FROM forms WHERE id = $1`

	templatesJSON := models.CareerForm{}
	DBerr := DB.Get(&templatesJSON, query, responseRequest.Record.FormID)
	if DBerr != nil {
		utils.ReturnError(DBerr.Error())
		return
	}
	var tempId = ""
	switch templatesJSON.Related_to {

	case "career":
		tempId = GetTemplateIdCareers(templatesJSON, responseRequest.Record.Status)
	}

	dstruct := models.DataStruct{
		Name:  responseRequest.Record.Data.Name,
		Email: responseRequest.Record.Data.Email,
		Phone: responseRequest.Record.Data.Phone,
	}

	Json, er := json.Marshal(dstruct)
	if er != nil {
		utils.ReturnError(er.Error())
	}
	dataJson := string(Json)
	tx := DB.MustBegin()
	tx.MustExec("INSERT INTO reminders (template_id,data,status,related_to) values ($1,$2,$3,$4)", tempId, dataJson, responseRequest.Record.Status, templatesJSON.Related_to)
	tx.Commit()
	DB.Close()
}

func GetTemplateIdCareers(formData models.CareerForm, status string) string {
	tempvar := models.CareerTemplates{}
	Unmarerr := json.Unmarshal(formData.Templates, &tempvar)
	if Unmarerr != nil {
		utils.ReturnError(Unmarerr.Error())
		return ""
	}
	switch status {
	case "p":
		return tempvar.On_applied
	case "a":
		return tempvar.Accepted
	case "r":
		return tempvar.Rejected
	default:
		return ""
	}

}
