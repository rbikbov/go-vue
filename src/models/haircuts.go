package models

import (
	"errors"
	"time"
	"os"
)

type Haircut struct {
	ID     int64 `json:"id"`
	Title  string `json:"title"`
	Description  string `json:"description"`
	Date   string `json:"date"`
	Price   string `json:"price"`
	Photo  string  `json:"photo"`
}

type Haircuts struct {
	List []*Haircut

	LastID int64
}

func (h *Haircuts) Get(id int64) (*Haircut, error) {
	for _, haircut := range h.List {
		if haircut.ID == id {
			return haircut, nil
		}
	}
	return nil, errors.New("Haircut not found")
}

func (h *Haircuts) Create(haircut *Haircut) {
	h.LastID++
	haircut.ID = h.LastID
	haircut.Date = time.Now().Format("02/01/2006 15:04:05")

	h.List = append(h.List, haircut)
}

func (h *Haircuts) Remove(id int64) {
	for key, item := range h.List {
		if item.ID == id {
			var err = os.Remove("./frontend"+item.Photo)
			checkError(err)
			h.List = append(h.List[:key], h.List[key+1:]...)
		}
	}
}

func (h *Haircuts) Random(id int64) ([]*Haircut, error) {
	var RandomHaircuts []*Haircut

	for {
		var x = randInt(0, len(h.List))

		if !(haircutInHaircuts(h.List[x], RandomHaircuts)) && h.List[x].ID != id {
			RandomHaircuts = append(RandomHaircuts, h.List[x])
		}

		if len(RandomHaircuts) == 4 || len(RandomHaircuts) == len(h.List) -1 {
			break
		} else if len(h.List) == 0 {
			return nil, errors.New("Haircuts not found")
		}
	}

	return RandomHaircuts, nil
}

func haircutInHaircuts(a *Haircut, list []*Haircut) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}