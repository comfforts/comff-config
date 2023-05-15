package config

import (
	"crypto/tls"

	"github.com/comfforts/errors"
	"github.com/comfforts/logger"

	"github.com/comfforts/comff-config/internal/auth"
	"github.com/comfforts/comff-config/internal/config"
)

const (
	ERR_MISSING_REQUIRED = "missing custom configuration details"
	ERR_UNDEFINED_TARGET = "target undefined"
)

var (
	ErrMissingRequired = errors.NewAppError(ERR_MISSING_REQUIRED)
	ErrUndefinedTarget = errors.NewAppError(ERR_UNDEFINED_TARGET)
)

type ConfigurationTarget string

const (
	SERVER          ConfigurationTarget = "SERVER"
	CLIENT          ConfigurationTarget = "CLIENT"
	GEO_CLIENT      ConfigurationTarget = "GEO_CLIENT"
	PROFILE_CLIENT  ConfigurationTarget = "PROFILE_CLIENT"
	SHOP_CLIENT     ConfigurationTarget = "SHOP_CLIENT"
	COURIER_CLIENT  ConfigurationTarget = "COURIER_CLIENT"
	DELIVERY_CLIENT ConfigurationTarget = "DELIVERY_CLIENT"
	BIZ_CLIENT      ConfigurationTarget = "BIZ_CLIENT"
	OFFERS_CLIENT   ConfigurationTarget = "OFFERS_CLIENT"
	CUSTOM          ConfigurationTarget = "CUSTOM"
	NOBODY_CLIENT   ConfigurationTarget = "NOBODY_CLIENT"
)

type CustomOpts struct {
	CertFilePath string
	KeyFilePath  string
	IsServer     bool
}

type ConfigOpts struct {
	Addr   string
	Target ConfigurationTarget
	Opts   *CustomOpts
}

type Authorizer interface {
	Authorize(subject, object, action string) error
}

func SetupAuthorizer(logger logger.AppLogger) (*auth.Authorizer, error) {
	return auth.NewAuthorizer(config.PolicyFile(config.ACLModelFile), config.PolicyFile(config.ACLPolicyFile), logger)
}

func SetupTLSConfig(opts *ConfigOpts) (*tls.Config, error) {
	switch opts.Target {
	case SERVER:
		if opts.Addr == "" {
			return nil, ErrMissingRequired
		}
		return config.SetupTLSConfig(config.TLSConfig{
			CertFile:      config.CertFile(config.ServerCertFile),
			KeyFile:       config.CertFile(config.ServerKeyFile),
			CAFile:        config.CertFile(config.CAFile),
			ServerAddress: opts.Addr,
			Server:        true,
		})
	case CLIENT:
		return config.SetupTLSConfig(config.TLSConfig{
			CertFile: config.CertFile(config.ClientCertFile),
			KeyFile:  config.CertFile(config.ClientKeyFile),
			CAFile:   config.CertFile(config.CAFile),
			Server:   false,
		})
	case GEO_CLIENT:
		return config.SetupTLSConfig(config.TLSConfig{
			CertFile: config.CertFile(config.GeoClientCertFile),
			KeyFile:  config.CertFile(config.GeoClientKeyFile),
			CAFile:   config.CertFile(config.CAFile),
			Server:   false,
		})
	case PROFILE_CLIENT:
		return config.SetupTLSConfig(config.TLSConfig{
			CertFile: config.CertFile(config.ProfileClientCertFile),
			KeyFile:  config.CertFile(config.ProfileClientKeyFile),
			CAFile:   config.CertFile(config.CAFile),
			Server:   false,
		})
	case SHOP_CLIENT:
		return config.SetupTLSConfig(config.TLSConfig{
			CertFile: config.CertFile(config.ShopClientCertFile),
			KeyFile:  config.CertFile(config.ShopClientKeyFile),
			CAFile:   config.CertFile(config.CAFile),
			Server:   false,
		})
	case COURIER_CLIENT:
		return config.SetupTLSConfig(config.TLSConfig{
			CertFile: config.CertFile(config.CourierClientCertFile),
			KeyFile:  config.CertFile(config.CourierClientKeyFile),
			CAFile:   config.CertFile(config.CAFile),
			Server:   false,
		})
	case DELIVERY_CLIENT:
		return config.SetupTLSConfig(config.TLSConfig{
			CertFile: config.CertFile(config.DeliveryClientCertFile),
			KeyFile:  config.CertFile(config.DeliveryClientKeyFile),
			CAFile:   config.CertFile(config.CAFile),
			Server:   false,
		})
	case OFFERS_CLIENT:
		return config.SetupTLSConfig(config.TLSConfig{
			CertFile: config.CertFile(config.OffersClientCertFile),
			KeyFile:  config.CertFile(config.OffersClientKeyFile),
			CAFile:   config.CertFile(config.CAFile),
			Server:   false,
		})
	case BIZ_CLIENT:
		return config.SetupTLSConfig(config.TLSConfig{
			CertFile: config.CertFile(config.BusinessClientCertFile),
			KeyFile:  config.CertFile(config.BusinessClientKeyFile),
			CAFile:   config.CertFile(config.CAFile),
			Server:   false,
		})
	case NOBODY_CLIENT:
		return config.SetupTLSConfig(config.TLSConfig{
			CertFile: config.CertFile(config.NobodyClientCertFile),
			KeyFile:  config.CertFile(config.NobodyClientKeyFile),
			CAFile:   config.CertFile(config.CAFile),
			Server:   false,
		})
	case CUSTOM:
		if opts.Opts == nil || opts.Opts.CertFilePath == "" || opts.Opts.KeyFilePath == "" {
			return nil, ErrMissingRequired
		}

		addr := ""
		if opts.Opts.IsServer {
			if opts.Addr == "" {
				return nil, ErrMissingRequired
			}
			addr = opts.Addr
		}
		return config.SetupTLSConfig(config.TLSConfig{
			CertFile:      opts.Opts.CertFilePath,
			KeyFile:       opts.Opts.KeyFilePath,
			CAFile:        config.CertFile(config.CAFile),
			Server:        opts.Opts.IsServer,
			ServerAddress: addr,
		})
	default:
		return nil, ErrUndefinedTarget
	}
}
