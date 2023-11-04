package info

type Response struct {
	Data   map[string]CryptoCurrency `json:"data"`
	Status StatusData                `json:"status"`
}

type CryptoCurrency struct {
	Urls         UrlData  `json:"urls"`
	Logo         string   `json:"logo"`
	ID           int      `json:"id"`
	Name         string   `json:"name"`
	Symbol       string   `json:"symbol"`
	Slug         string   `json:"slug"`
	Description  string   `json:"description"`
	DateAdded    string   `json:"date_added"`
	DateLaunched string   `json:"date_launched"`
	Tags         []string `json:"tags"`
	Platform     string   `json:"platform"`
	Category     string   `json:"category"`
	Notice       string   `json:"notice"`
	// You might need to add more fields if they are available
}

type UrlData struct {
	Website      []string `json:"website"`
	TechnicalDoc []string `json:"technical_doc"`
	Twitter      []string `json:"twitter"`
	Reddit       []string `json:"reddit"`
	MessageBoard []string `json:"message_board"`
	Announcement []string `json:"announcement"`
	Chat         []string `json:"chat"`
	Explorer     []string `json:"explorer"`
	SourceCode   []string `json:"source_code"`
	// Add more fields if needed
}

type StatusData struct {
	Timestamp    string `json:"timestamp"`
	ErrorCode    int    `json:"error_code"`
	ErrorMessage string `json:"error_message"`
	Elapsed      int    `json:"elapsed"`
	CreditCount  int    `json:"credit_count"`
	Notice       string `json:"notice"`
}

//
//// Params represents the query parameters for fetching data.
//type Params struct {
//	Id          *string `json:"id,omitempty"`
//	Slug        *string `json:"slug,omitempty"`
//	Symbol      *string `json:"symbol,omitempty"`
//	Address     *string `json:"address,omitempty"`
//	SkipInvalid *bool   `json:"skip_invalid,omitempty"`
//	Aux         []Aux
//}
