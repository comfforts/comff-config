package config

import (
	"crypto/tls"
	"crypto/x509"
	"errors"
	"os"
)

const (
	ERR_MISSING_CERT_FILE = "missing cert file"
	ERR_CERT_INACCESSIBLE = "cert file inaccessible"
	ERR_TLS_CFG_CREATION  = "error creating TLS config"
	ERR_PARSE_CERT        = "error parsing certificate file"
)

var (
	ErrMissingCertFile  = errors.New(ERR_MISSING_CERT_FILE)
	ErrCertInaccessible = errors.New(ERR_CERT_INACCESSIBLE)
	ErrTLSCfgCreation   = errors.New(ERR_TLS_CFG_CREATION)
	ErrParseCert        = errors.New(ERR_PARSE_CERT)
)

type TLSConfig struct {
	CertFile      string
	KeyFile       string
	CAFile        string
	ServerAddress string
	Server        bool
}

func SetupTLSConfig(cfg TLSConfig) (*tls.Config, error) {
	tlsConfig := &tls.Config{}

	if cfg.CertFile != "" && cfg.KeyFile != "" {
		_, err := os.Stat(cfg.CertFile)
		if err != nil {
			if os.IsNotExist(err) {
				return nil, ErrMissingCertFile
			} else {
				return nil, ErrCertInaccessible
			}
		}

		_, err = os.Stat(cfg.KeyFile)
		if err != nil {
			if os.IsNotExist(err) {
				return nil, ErrMissingCertFile
			} else {
				return nil, ErrCertInaccessible
			}
		}

		tlsConfig.Certificates = make([]tls.Certificate, 1)
		tlsConfig.Certificates[0], err = tls.LoadX509KeyPair(cfg.CertFile, cfg.KeyFile)
		if err != nil {
			return nil, ErrTLSCfgCreation
		}
	}
	if cfg.CAFile != "" {
		b, err := os.ReadFile(cfg.CAFile)
		if err != nil {
			return nil, ErrMissingCertFile
		}
		ca := x509.NewCertPool()
		ok := ca.AppendCertsFromPEM([]byte(b))
		if !ok {
			return nil, ErrParseCert
		}
		if cfg.Server {
			tlsConfig.ClientCAs = ca
			tlsConfig.ClientAuth = tls.RequireAndVerifyClientCert
		} else {
			tlsConfig.RootCAs = ca
		}
		tlsConfig.ServerName = cfg.ServerAddress
	}

	return tlsConfig, nil
}
