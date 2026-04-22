package config

import (
	"crypto/tls"
	"errors"

	"github.com/comfforts/comff-config/internal/auth"
	"github.com/comfforts/comff-config/internal/config"
)

const (
	ERR_MISSING_REQUIRED = "missing custom configuration details"
	ERR_UNDEFINED_TARGET = "target undefined"
)

var (
	ErrMissingRequired = errors.New(ERR_MISSING_REQUIRED)
	ErrUndefinedTarget = errors.New(ERR_UNDEFINED_TARGET)
)

type ConfigurationTarget string

const (
	SERVER               ConfigurationTarget = "SERVER"
	CLIENT               ConfigurationTarget = "CLIENT"
	GEO_CLIENT           ConfigurationTarget = "GEO_CLIENT"
	PROFILE_CLIENT       ConfigurationTarget = "PROFILE_CLIENT"
	SHOP_CLIENT          ConfigurationTarget = "SHOP_CLIENT"
	STORES_CLIENT        ConfigurationTarget = "STORES_CLIENT"
	COURIER_CLIENT       ConfigurationTarget = "COURIER_CLIENT"
	DELIVERY_CLIENT      ConfigurationTarget = "DELIVERY_CLIENT"
	BIZ_CLIENT           ConfigurationTarget = "BIZ_CLIENT"
	SCHEDULER_CLIENT     ConfigurationTarget = "SCHEDULER_CLIENT"
	OFFERS_CLIENT        ConfigurationTarget = "OFFERS_CLIENT"
	NOTIFICATIONS_CLIENT ConfigurationTarget = "NOTIFICATIONS_CLIENT"
	CUSTOM               ConfigurationTarget = "CUSTOM"
	NOBODY_CLIENT        ConfigurationTarget = "NOBODY_CLIENT"
)

type CustomOpts struct {
	CertFilePath string
	KeyFilePath  string
	CAFilePath   string
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

func SetupAuthorizer() (*auth.Authorizer, error) {
	return auth.NewAuthorizer(config.PolicyFile(config.ACLModelFile), config.PolicyFile(config.ACLPolicyFile))
}

func SetupTLSConfig(opts *ConfigOpts) (*tls.Config, error) {
	var caFilePath string
	var certFilePath string
	var keyFilePath string

	if opts.Opts != nil {
		if opts.Opts.CAFilePath != "" {
			caFilePath = opts.Opts.CAFilePath
		}
		if opts.Opts.CertFilePath != "" {
			certFilePath = opts.Opts.CertFilePath
		}
		if opts.Opts.KeyFilePath != "" {
			keyFilePath = opts.Opts.KeyFilePath
		}
	}

	if caFilePath == "" {
		caFilePath = config.CertFile(config.CAFile)
	}

	switch opts.Target {
	case SERVER:
		if opts.Addr == "" {
			return nil, ErrMissingRequired
		}

		if certFilePath == "" {
			certFilePath = config.CertFile(config.ServerCertFile)
		}

		if keyFilePath == "" {
			keyFilePath = config.CertFile(config.ServerKeyFile)
		}

		return config.SetupTLSConfig(config.TLSConfig{
			CertFile:      certFilePath,
			KeyFile:       keyFilePath,
			CAFile:        caFilePath,
			ServerAddress: opts.Addr,
			Server:        true,
		})
	case CLIENT:
		if certFilePath == "" {
			certFilePath = config.CertFile(config.ClientCertFile)
		}

		if keyFilePath == "" {
			keyFilePath = config.CertFile(config.ClientKeyFile)
		}

		return config.SetupTLSConfig(config.TLSConfig{
			CertFile: certFilePath,
			KeyFile:  keyFilePath,
			CAFile:   caFilePath,
			Server:   false,
		})
	case GEO_CLIENT:
		if certFilePath == "" {
			certFilePath = config.CertFile(config.GeoClientCertFile)
		}

		if keyFilePath == "" {
			keyFilePath = config.CertFile(config.GeoClientKeyFile)
		}

		return config.SetupTLSConfig(config.TLSConfig{
			CertFile: certFilePath,
			KeyFile:  keyFilePath,
			CAFile:   caFilePath,
			Server:   false,
		})
	case PROFILE_CLIENT:
		return config.SetupTLSConfig(config.TLSConfig{
			CertFile: config.CertFile(config.ProfileClientCertFile),
			KeyFile:  config.CertFile(config.ProfileClientKeyFile),
			CAFile:   caFilePath,
			Server:   false,
		})
	case SHOP_CLIENT:
		return config.SetupTLSConfig(config.TLSConfig{
			CertFile: config.CertFile(config.ShopClientCertFile),
			KeyFile:  config.CertFile(config.ShopClientKeyFile),
			CAFile:   caFilePath,
			Server:   false,
		})
	case STORES_CLIENT:
		if certFilePath == "" {
			certFilePath = config.CertFile(config.StoresClientCertFile)
		}

		if keyFilePath == "" {
			keyFilePath = config.CertFile(config.StoresClientKeyFile)
		}

		return config.SetupTLSConfig(config.TLSConfig{
			CertFile: certFilePath,
			KeyFile:  keyFilePath,
			CAFile:   caFilePath,
			Server:   false,
		})
	case COURIER_CLIENT:
		return config.SetupTLSConfig(config.TLSConfig{
			CertFile: config.CertFile(config.CourierClientCertFile),
			KeyFile:  config.CertFile(config.CourierClientKeyFile),
			CAFile:   caFilePath,
			Server:   false,
		})
	case DELIVERY_CLIENT:
		return config.SetupTLSConfig(config.TLSConfig{
			CertFile: config.CertFile(config.DeliveryClientCertFile),
			KeyFile:  config.CertFile(config.DeliveryClientKeyFile),
			CAFile:   caFilePath,
			Server:   false,
		})
	case OFFERS_CLIENT:
		return config.SetupTLSConfig(config.TLSConfig{
			CertFile: config.CertFile(config.OffersClientCertFile),
			KeyFile:  config.CertFile(config.OffersClientKeyFile),
			CAFile:   caFilePath,
			Server:   false,
		})
	case NOTIFICATIONS_CLIENT:
		return config.SetupTLSConfig(config.TLSConfig{
			CertFile: config.CertFile(config.NotificationsClientCertFile),
			KeyFile:  config.CertFile(config.NotificationsClientKeyFile),
			CAFile:   caFilePath,
			Server:   false,
		})
	case BIZ_CLIENT:
		return config.SetupTLSConfig(config.TLSConfig{
			CertFile: config.CertFile(config.BusinessClientCertFile),
			KeyFile:  config.CertFile(config.BusinessClientKeyFile),
			CAFile:   caFilePath,
			Server:   false,
		})
	case SCHEDULER_CLIENT:
		return config.SetupTLSConfig(config.TLSConfig{
			CertFile: config.CertFile(config.SchedulerClientCertFile),
			KeyFile:  config.CertFile(config.SchedulerClientKeyFile),
			CAFile:   caFilePath,
			Server:   false,
		})
	case NOBODY_CLIENT:
		if certFilePath == "" {
			certFilePath = config.CertFile(config.NobodyClientCertFile)
		}

		if keyFilePath == "" {
			keyFilePath = config.CertFile(config.NobodyClientKeyFile)
		}

		return config.SetupTLSConfig(config.TLSConfig{
			CertFile: certFilePath,
			KeyFile:  keyFilePath,
			CAFile:   caFilePath,
			Server:   false,
		})
	case CUSTOM:
		if opts.Opts == nil || opts.Opts.CertFilePath == "" || opts.Opts.KeyFilePath == "" {
			return nil, ErrMissingRequired
		}

		caFilePath := opts.Opts.CAFilePath
		if caFilePath == "" {
			caFilePath = config.CertFile(config.CAFile)
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
			CAFile:        caFilePath,
			Server:        opts.Opts.IsServer,
			ServerAddress: addr,
		})
	default:
		return nil, ErrUndefinedTarget
	}
}
