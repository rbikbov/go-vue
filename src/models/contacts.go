package models

type Contact struct {
	ID     int64 `json:"id"`
	Title  string `json:"title"`
	Description  string `json:"description"`
	Icon   string `json:"icon"`
	Prefix   string `json:"prefix"`
}

type Contacts struct {
	List []*Contact

	LastID int64
}

func (c *Contacts) Create(contact *Contact) {
	c.LastID++
	contact.ID = c.LastID

	c.List = append(c.List, contact)
}

func (c *Contacts) Remove(id int64) {
	for key, item := range c.List {
		if item.ID == id {
			c.List = append(c.List[:key], c.List[key+1:]...)
		}
	}
}