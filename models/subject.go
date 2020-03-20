package models

type Subject struct {
	ID string `json:"id"`
	Title string `json:"title"`
	Desc string `json:"desc"`
}

func GetAllSubject() []Subject {
	
   var docs []Subject
   docs = append(docs, Subject{ID: "222", Title: "Belajar makan", Desc: "Belajar makan yuk"})
   return docs
}
