package main

import (
  "fmt"
  "html/template"
	"net/http"
)

type Todo struct {
	Title string
	Done  bool
}

type TodoPageData struct {
	PageTitle string
	Todos     []Todo
}

type ContactDetails struct {
	Email   string
	Subject string
	Message string
}

func main() {

  tmpl := template.Must(template.ParseFiles("layout.html"));
  tmpl_form := template.Must(template.ParseFiles("form.html"));

	http.HandleFunc("/", func (w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Welcome to my website!")
  })
  
  http.HandleFunc("/test", func (w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, r.URL.Query().Get("token"))
  })
  
  http.HandleFunc("/home", func (w http.ResponseWriter, r *http.Request) {
		data := TodoPageData{
			PageTitle: "My TODO list",
			Todos: []Todo{
				{Title: "Task 1", Done: false},
				{Title: "Task 2", Done: true},
				{Title: "Task 3", Done: true},
			},
		}
		tmpl.Execute(w, data)
  })
  
  http.HandleFunc("/register",func (w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
			tmpl_form.Execute(w, nil)
			return
    }
    details := ContactDetails{
			Email:   r.FormValue("email"),
			Subject: r.FormValue("subject"),
			Message: r.FormValue("message"),
    }
    _ = details
    tmpl_form.Execute(w, struct{ Success bool }{true})
  })

  http.HandleFunc("/postregister",func (w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w,r.FormValue("email_address"));
  })
  
  http.ListenAndServe(":8080", nil)

}
