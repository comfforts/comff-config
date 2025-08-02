package auth

import (
	"errors"
	"fmt"
	"os"

	"github.com/casbin/casbin"
)

const (
	ERR_MISSING_REQUIRED = "missing custom configuration details"
	ERR_UNDEFINED_TARGET = "target undefined"
)

var (
	ErrMissingRequired = errors.New(ERR_MISSING_REQUIRED)
	ErrUndefinedTarget = errors.New(ERR_UNDEFINED_TARGET)
)

type Authorizer struct {
	enforcer *casbin.Enforcer
}

func NewAuthorizer(model, policy string) (*Authorizer, error) {
	_, err := os.Stat(model)
	if err != nil {
		if os.IsNotExist(err) {
			return nil, fmt.Errorf("file doesn't exist - %s", model)
		} else {
			return nil, fmt.Errorf("file inaccessible - %s: %w", model, err)
		}
	}

	_, err = os.Stat(policy)
	if err != nil {
		if os.IsNotExist(err) {
			return nil, fmt.Errorf("file doesn't exist - %s", model)
		} else {
			return nil, fmt.Errorf("file inaccessible - %s: %w", model, err)
		}
	}

	enforcer := casbin.NewEnforcer(model, policy)
	return &Authorizer{
		enforcer: enforcer,
	}, nil
}

func (a *Authorizer) Authorize(subject, object, action string) error {
	if !a.enforcer.Enforce(subject, object, action) {
		return fmt.Errorf("%s not permitted to %s to %s", subject, action, object)
	}
	return nil
}
