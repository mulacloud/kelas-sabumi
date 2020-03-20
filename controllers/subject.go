package controllers

import (
	"net/http"
	"encoding/json"
	"github.com/mulacloud/kelas-sabumi/models"
)

func ListSubject(w http.ResponseWriter, req *http.Request){

   docs := models.GetAllSubject()

   w.Header().Set("Content-Type", "application/json")
   json.NewEncoder(w).Encode(docs)
}

