package types

type CMDWeb struct {
	HTTPServer *HTTPServer `json:"http_server"`
	Database   *Database   `json:"database"`
	WebFs      *WebFs      `json:"-"`
}

func NewDefaultCMDWeb() *CMDWeb {
	return &CMDWeb{
		HTTPServer: NewDefaultHTTPServer(),
		Database:   NewDefaultDatabase(),
	}
}
