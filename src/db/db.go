package db

import (
	mb "github.com/happierall/membase"
	"models"
)

var Hairstyles models.Hairstyles
var Haircuts models.Haircuts
var Diplomas models.Diplomas
var Contacts models.Contacts

func init() {
	mb.Default.StoreFolder = "./membase/"
	mb.Listen(&Hairstyles)
	mb.Listen(&Haircuts)
	mb.Listen(&Diplomas)
	mb.Listen(&Contacts)
	go mb.Run()
}
