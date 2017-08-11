package models

import (
	"errors"
	"time"
	"os"
)

type Hairstyle struct {
	ID     int64 `json:"id"`
	Title  string `json:"title"`
	Description  string `json:"description"`
	Date   string `json:"date"`
	Price   string `json:"price"`
	Photo  string  `json:"photo"`
}

type Hairstyles struct {
	List []*Hairstyle

	LastID int64
}

func (h *Hairstyles) Get(id int64) (*Hairstyle, error) {
	for _, hairstyle := range h.List {
		if hairstyle.ID == id {
			return hairstyle, nil
		}
	}
	return nil, errors.New("Hairstyle not found")
}

func (h *Hairstyles) Create(hairstyle *Hairstyle) {
	h.LastID++
	hairstyle.ID = h.LastID
	hairstyle.Date = time.Now().Format("02/01/2006 15:04:05")

	h.List = append(h.List, hairstyle)
}

func (h *Hairstyles) Remove(id int64) {
	for key, item := range h.List {
		if item.ID == id {
			var err = os.Remove("./frontend"+item.Photo)
			checkError(err)
			h.List = append(h.List[:key], h.List[key+1:]...)
		}
	}
}

func (h *Hairstyles) Random(id int64) ([]*Hairstyle, error) {
	var RandomHairstyles []*Hairstyle

	for {
		var x = randInt(0, len(h.List))

		if !(hairstyleInHairstyles(h.List[x], RandomHairstyles)) && h.List[x].ID != id {
			RandomHairstyles = append(RandomHairstyles, h.List[x])
		}

		if len(RandomHairstyles) == 4 || len(RandomHairstyles) == len(h.List) -1 {
			break
		} else if len(h.List) == 0 {
			return nil, errors.New("Hairstyles not found")
		}
	}

	return RandomHairstyles, nil
}

func hairstyleInHairstyles(a *Hairstyle, list []*Hairstyle) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}