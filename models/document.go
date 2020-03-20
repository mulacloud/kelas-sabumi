package models

type Document struct {
	ID string `json:"id"`
	Title string `json:"title"`
	Body string `json:"body"`
}

func GetAllDocument() []Document {
	
   var docs []Document
   docs = append(docs, Document{ID: "1-asf-asd-fa-sd-fa-dsfa-s", Title: "Dokumen pertama", Body: "Test test tes"})
   return docs
}
