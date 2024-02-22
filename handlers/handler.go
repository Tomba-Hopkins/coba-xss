package handlers

import (
	"coba-xss/views"
	"fmt"
	"html/template"
	"net/http"
)

func DashboardHanlder(w http.ResponseWriter, r *http.Request) {

	// coba coba
	if r.Method == "POST"{
		dataPost := map[string]any{
			"Data":r.FormValue("send"),
		}
		views.MyTemplates.ExecuteTemplate(w, "index.html", dataPost)
		fmt.Fprintf(w, `<script>alert("Stage 1 berhasil"); window.location.href = '/parameter';</script>`) //pakai redirect golang malah ribet
		return
	}


	// xss post data
	dataPost := map[string]any{
		"Data":r.FormValue("send"),
	}
	views.MyTemplates.ExecuteTemplate(w, "index.html", dataPost)

}

func ParameterHandler(w http.ResponseWriter, r *http.Request) {
	// xss parameter

	data := map[string]any{
		"Data": template.HTML(r.URL.Query().Get("data")), // harus ada {{.Data}} di parameternya baru bisa
	}
	views.MyTemplates.ExecuteTemplate(w, "parameter.html", data) 
}