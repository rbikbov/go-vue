package views

import (
	soso "github.com/happierall/soso-server"
	"db"
	"models"
	"fmt"
	"net/http"
	"os"
	"io"
)

func init() {
	Routes.SEARCH("diplomas", Diplomas.search)
	Routes.Handle("diplomas", "random", Diplomas.random)
	Routes.Handle("diplomas", "get", Diplomas.get)
	Routes.DELETE("diplomas", Diplomas.remove)
	http.HandleFunc("/create_diploma", create_diploma)
}


var Diplomas = diplomas{}

type diplomas struct {}

func (d *diplomas) get(m *soso.Msg) {
	type Data struct {
		ID int64 `json:"id"`
	}

	data := &Data{}
	m.ReadData(data)

	obj, err := db.Diplomas.Get(data.ID)
	if err != nil {
		m.Error(400, soso.LevelError, err)
		return
	}

	m.Success(map[string]interface{}{
		"object": obj,
	})
}

func (d *diplomas) search(m *soso.Msg) {
	m.Success(map[string]interface{}{
		"list": db.Diplomas.List,
	})
}

func (d *diplomas) remove(m *soso.Msg) {
	if m.User.IsAuth {
		type Data struct {
			ID int64 `json:"id"`
		}
		data := &Data{}
		m.ReadData(data)
		db.Diplomas.Remove(data.ID)
		fmt.Println(data.ID)

		m.Success(map[string]interface{}{})
	}
}

func create_diploma(w http.ResponseWriter, r *http.Request) {

	if r.Method == "POST" {
		if r.FormValue("email") == "dribblerafis@gmail.com" {

			w.Header().Set("Access-Control-Allow-Origin", "*")
			r.ParseMultipartForm(32 << 20)
			file, handler, err := r.FormFile("uploadfile")
			if err != nil {
				fmt.Println(err)
				return
			}
			defer file.Close()
			fmt.Fprintf(w, "%v", handler.Header)
			f, err := os.OpenFile("./frontend/media/diplomas/" + handler.Filename, os.O_WRONLY | os.O_CREATE, 0666)
			if err != nil {
				fmt.Println(err)
				return
			}
			defer f.Close()
			io.Copy(f, file)

			title := r.FormValue("title")
			description := r.FormValue("description")

			diploma := &models.Diploma{
				Title: title,
				Description: description,
				Photo: "/media/diplomas/" + handler.Filename,
			}
			db.Diplomas.Create(diploma)

			fmt.Println(r.Method, r.FormValue("email"), "YES")

		}

	} else {
		fmt.Println(r.Method, r.FormValue("email"), "NO")
	}
}

func (d *diplomas) random(m *soso.Msg) {
	type Data struct {
		ID int64 `json:"id"`
	}

	data := &Data{}
	m.ReadData(data)

	list, err := db.Diplomas.Random(data.ID)
	if err != nil {
		m.Error(400, soso.LevelError, err)
		return
	}

	m.Success(map[string]interface{}{
		"list": list,
	})
}
