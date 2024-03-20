package info

import "C"
import (
	"errors"
	"fmt"
	c "github.com/cploutarchou/crypto-sdk-suite/coinmarketcap/client"
)

const (
	Urls             Aux    = "urls"
	Logo             Aux    = "logo"
	Description      Aux    = "description"
	Tags             Aux    = "tags"
	Platform         Aux    = "platform"
	DateAdded        Aux    = "date_added"
	Notice           Aux    = "notice"
	MetadataEndpoint string = "/v2/cryptocurrency/info"
)

type Aux string
type Metadata struct {
	client *c.Client
	Params *Params
	aux    []Aux
}

func New(c *c.Client) *Metadata {
	return &Metadata{
		client: c,
	}
}
func (m *Metadata) setDefaultAux() string {
	return fmt.Sprintf("%s,%s,%s,%s,%s,%s,%s", Urls, Logo, Description, Tags, Platform, DateAdded, Notice)
}

// Params represents the query parameters for fetching data.
type Params struct {
	Id          *string  `json:"id,omitempty"`
	Slugs       []string `json:"slugs,omitempty"`
	Symbols     []string `json:"symbols,omitempty"`
	Address     *string  `json:"address,omitempty"`
	SkipInvalid *bool    `json:"skip_invalid,omitempty"`
	Aux         []string `json:"aux,omitempty"`
}

func (m *Metadata) GetMetadata(params *Params) (map[string]CryptoCurrency, error) {
	if params.Id == nil && len(params.Slugs) == 0 && len(params.Symbols) == 0 {
		return nil, errors.New("missing required params! |id : " +
			"the coinmarketcap token id |slugs : like bitcoin,ethereum |" +
			" symbols : like BTC,ETH}")
	}
	queryParams := make(map[string]string)
	if params == nil {
		params = &Params{}
	}
	if params.Id != nil {
		queryParams["limit"] = *params.Id
	}
	if params.Symbols != nil {
		str, err := SliceToString(params.Symbols, ",")
		if err != nil {
			return nil, err
		}
		queryParams["symbol"] = str
	}
	if params.Slugs != nil {
		str, err := SliceToString(params.Slugs, ",")
		if err != nil {
			return nil, err
		}
		queryParams["slug"] = str
	}
	if params.Address != nil {
		queryParams["address"] = *params.Address
	}
	if params.SkipInvalid != nil {
		queryParams["skip_invalid"] = BoolToString(*params.SkipInvalid)
	}
	if params.Aux != nil {
		var theParams []string
		for _, i := range params.Aux {
			theParams = append(theParams, fmt.Sprintf("%s", i))
		}
		str, err := SliceToString(theParams, ",")
		if err != nil {
			return nil, err
		}
		queryParams["aux"] = str
	} else {
		m.setDefaultAux()
	}

	resp, err := m.client.Get(MetadataEndpoint, queryParams)
	if err != nil {
		return nil, err
	}

	var data Response
	if err := resp.Unmarshal(&data); err != nil {
		return nil, err
	}

	return data.Data, nil
}
