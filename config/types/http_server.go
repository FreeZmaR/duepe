package types

import (
	"crypto/tls"
	"duepe/internal/lib/util"
)

type HTTPServer struct {
	Host               *string        `yaml:"host"`
	Port               *string        `yaml:"port"`
	ReadTimeoutSec     *uint          `yaml:"read_timeout_sec"`
	WriteTimeoutSec    *uint          `yaml:"write_timeout_sec"`
	IdleTimeoutSec     *uint          `yaml:"idle_timeout_sec"`
	CSRFKey            *string        `yaml:"csrf_key"`
	CSRFMaxAgeSec      *int           `yaml:"csrf_max_age_sec"`
	ShutdownTimeoutSec *int           `yaml:"shutdown_timeout_sec"`
	TLS                *HTTPServerTLS `yaml:"tls"`
}

type HTTPServerTLS struct {
	KeyFilepath  *string       `yaml:"key_filepath"`
	CertFilepath *string       `yaml:"cert_filepath"`
	MinVersion   *uint16       `yaml:"min_version"`
	MaxVersion   *uint16       `yaml:"max_version"`
	Curves       []tls.CurveID `yaml:"curves"`
	Ciphers      []uint16      `yaml:"ciphers"`
}

func NewDefaultHTTPServer() *HTTPServer {
	return &HTTPServer{
		Host:               util.WithPointer("localhost"),
		Port:               util.WithPointer("8080"),
		ReadTimeoutSec:     util.WithPointer(uint(5)),
		WriteTimeoutSec:    util.WithPointer(uint(5)),
		IdleTimeoutSec:     util.WithPointer(uint(5)),
		CSRFKey:            util.WithPointer("csrf_key"),
		CSRFMaxAgeSec:      util.WithPointer(3600),
		ShutdownTimeoutSec: util.WithPointer(5),
	}
}
