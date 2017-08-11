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
	Routes.SEARCH("hairstyles", Hairstyles.search)
	Routes.Handle("hairstyles", "random", Hairstyles.random)
	Routes.Handle("hairstyles", "get", Hairstyles.get)
	Routes.DELETE("hairstyles", Hairstyles.remove)
	http.HandleFunc("/create_hairstyle", create_hairstyle)
}


var Hairstyles = hairstyles{}

type hairstyles struct {}

func (h *hairstyles) get(m *soso.Msg) {
	type Data struct {
		ID int64 `json:"id"`
	}

	data := &Data{}
	m.ReadData(data)

	obj, err := db.Hairstyles.Get(data.ID)
	if err != nil {
		m.Error(400, soso.LevelError, err)
		return
	}

	m.Success(map[string]interface{}{
		"object": obj,
	})
}

func (h *hairstyles) search(m *soso.Msg) {
	m.Success(map[string]interface{}{
		"list": db.Hairstyles.List,
	})
}

func (h *hairstyles) remove(m *soso.Msg) {
	if m.User.IsAuth {
		type Data struct {
			ID int64 `json:"id"`
		}
		data := &Data{}
		m.ReadData(data)
		db.Hairstyles.Remove(data.ID)
		fmt.Println(data.ID)

		m.Success(map[string]interface{}{})
	}
}

func create_hairstyle(w http.ResponseWriter, r *http.Request) {

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

			f, err := os.OpenFile("./frontend/media/hairstyles/" + handler.Filename, os.O_WRONLY | os.O_CREATE, 0666)
			if err != nil {
				fmt.Println(err)
				return
			}
			defer f.Close()
			io.Copy(f, file)

			title := r.FormValue("title")
			description := r.FormValue("description")
			price := r.FormValue("price")

			hairstyle := &models.Hairstyle{
				Title: title,
				Description: description,
				Price: price,
				Photo: "/media/hairstyles/" + handler.Filename,
			}
			db.Hairstyles.Create(hairstyle)

			fmt.Println(r.Method, r.URL, r.FormValue("email"), "YES")

		}

	} else {
		fmt.Println(r.Method, r.URL, r.FormValue("email"), "NO")
	}
}

func (h *hairstyles) random(m *soso.Msg) {
	type Data struct {
		ID int64 `json:"id"`
	}

	data := &Data{}
	m.ReadData(data)

	list, err := db.Hairstyles.Random(data.ID)
	if err != nil {
		m.Error(400, soso.LevelError, err)
		return
	}

	m.Success(map[string]interface{}{
		"list": list,
	})
}
