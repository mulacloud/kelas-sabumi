package controllers

import (
	"net/http"
	"encoding/json"
	"github.com/mulacloud/kelas-sabumi/models"
)

func ListDocument(w http.ResponseWriter, req *http.Request){

   docs := models.GetAllDocument()

   w.Header().Set("Content-Type", "application/json")
   json.NewEncoder(w).Encode(docs)
}

