package models

import (
	"errors"
	"time"
	"os"
)

type Diploma struct {
	ID     int64 `json:"id"`
	Title  string `json:"title"`
	Description  string `json:"description"`
	Date   string `json:"date"`
	Photo  string  `json:"photo"`
}

type Diplomas struct {
	List []*Diploma

	LastID int64
}

func (d *Diplomas) Get(id int64) (*Diploma, error) {
	for _, diploma := range d.List {
		if diploma.ID == id {
			return diploma, nil
		}
	}
	return nil, errors.New("Diploma not found")
}

func (d *Diplomas) Create(diploma *Diploma) {
	d.LastID++
	diploma.ID = d.LastID
	diploma.Date = time.Now().Format("02/01/2006 15:04:05")

	d.List = append(d.List, diploma)
}

func (d *Diplomas) Remove(id int64) {
	for key, item := range d.List {
		if item.ID == id {
			var err = os.Remove("./frontend"+item.Photo)
			checkError(err)
			d.List = append(d.List[:key], d.List[key+1:]...)
		}
	}
}

func (d *Diplomas) Random(id int64) ([]*Diploma, error) {
	var RandomDiplomas []*Diploma

	for {
		var x = randInt(0, len(d.List))

		if !(diplomaInDiplomas(d.List[x], RandomDiplomas)) && d.List[x].ID != id {
			RandomDiplomas = append(RandomDiplomas, d.List[x])
		}

		if len(RandomDiplomas) == 4 || len(RandomDiplomas) == len(d.List) -1 {
			break
		} else if len(d.List) == 0 {
			return nil, errors.New("Diplomas not found")
		}
	}

	return RandomDiplomas, nil
}

func diplomaInDiplomas(a *Diploma, list []*Diploma) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}