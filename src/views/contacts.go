package views

import (
	soso "github.com/happierall/soso-server"
	"db"
	"models"
	"fmt"
)

func init() {
	Routes.SEARCH("contacts", Contacts.search)
	Routes.CREATE("contacts", Contacts.create)
	Routes.DELETE("contacts", Contacts.remove)
}


var Contacts = contacts{}

type contacts struct {}


func (c *contacts) search(m *soso.Msg) {
	m.Success(map[string]interface{}{
		"list": db.Contacts.List,
	})
}

func (c *contacts) create(m *soso.Msg) {
	if m.User.IsAuth {
		type Data struct {
			Title       string `json:"title"`
			Description string `json:"description"`
			Icon        string `json:"icon"`
			Prefix      string `json:"prefix"`
		}
		data := &Data{}
		m.ReadData(data)

		contact := &models.Contact{
			Title: data.Title,
			Description: data.Description,
			Icon: data.Icon,
			Prefix: data.Prefix,
		}
		db.Contacts.Create(contact)

		m.Success(map[string]interface{}{
			"ID": 1,
		})
	}
}

func (c *contacts) remove(m *soso.Msg) {
	if m.User.IsAuth {
		type Data struct {
			Id int64 `json:"id"`
		}
		data := &Data{}
		m.ReadData(data)
		db.Contacts.Remove(data.Id)
		fmt.Println(data.Id)
	}
}