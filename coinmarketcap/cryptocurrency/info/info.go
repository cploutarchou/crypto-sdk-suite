package info

import "C"
import (
	"fmt"
	c "github.com/cploutarchou/crypto-sdk-suite/coinmarketcap/client"
)

const (
	Urls        Aux = "urls"
	Logo        Aux = "logo"
	Description Aux = "description"
	Tags        Aux = "tags"
	Platform    Aux = "platform"
	DateAdded   Aux = "date_added"
	Notice      Aux = "notice"
)

type Aux string
type Metadata struct {
	client *c.Client
}

func New(c *c.Client) *Metadata {
	return &Metadata{
		client: c,
	}
}
func (m *Metadata) setDefaultAux() string {
	return fmt.Sprintf("%s,%s,%s,%s,%s,%s,%s", Urls, Logo, Description, Tags, Platform, DateAdded, Notice)
}
