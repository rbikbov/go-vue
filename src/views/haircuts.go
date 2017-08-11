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
	Routes.SEARCH("haircuts", Haircuts.search)
	Routes.Handle("haircuts", "random", Haircuts.random)
	Routes.Handle("haircuts", "get", Haircuts.get)
	Routes.DELETE("haircuts", Haircuts.remove)
	http.HandleFunc("/create_haircut", create_haircut)
}


var Haircuts = haircuts{}

type haircuts struct {}

func (h *haircuts) get(m *soso.Msg) {
	type Data struct {
		ID int64 `json:"id"`
	}

	data := &Data{}
	m.ReadData(data)

	obj, err := db.Haircuts.Get(data.ID)
	if err != nil {
		m.Error(400, soso.LevelError, err)
		return
	}

	m.Success(map[string]interface{}{
		"object": obj,
	})
}

func (h *haircuts) search(m *soso.Msg) {
	m.Success(map[string]interface{}{
		"list": db.Haircuts.List,
	})
}

func (h *haircuts) remove(m *soso.Msg) {
	if m.User.IsAuth {
		type Data struct {
			ID int64 `json:"id"`
		}
		data := &Data{}
		m.ReadData(data)
		db.Haircuts.Remove(data.ID)
		fmt.Println(data.ID)

		m.Success(map[string]interface{}{})
	}
}

func create_haircut(w http.ResponseWriter, r *http.Request) {

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
			f, err := os.OpenFile("./frontend/media/haircuts/" + handler.Filename, os.O_WRONLY | os.O_CREATE, 0666)
			if err != nil {
				fmt.Println(err)
				return
			}
			defer f.Close()
			io.Copy(f, file)

			title := r.FormValue("title")
			description := r.FormValue("description")
			price := r.FormValue("price")

			haircut := &models.Haircut{
				Title: title,
				Description: description,
				Price: price,
				Photo: "/media/haircuts/" + handler.Filename,
			}
			db.Haircuts.Create(haircut)

			fmt.Println(r.Method, r.FormValue("email"), "YES")

		}

	} else {
		fmt.Println(r.Method, r.FormValue("email"), "NO")
	}
}

func (h *haircuts) random(m *soso.Msg) {
	type Data struct {
		ID int64 `json:"id"`
	}

	data := &Data{}
	m.ReadData(data)

	list, err := db.Haircuts.Random(data.ID)
	if err != nil {
		m.Error(400, soso.LevelError, err)
		return
	}

	m.Success(map[string]interface{}{
		"list": list,
	})
}
