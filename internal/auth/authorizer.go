package auth

import (
	"fmt"
	"os"

	"github.com/casbin/casbin"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/comfforts/errors"
	"github.com/comfforts/logger"
)

type Authorizer struct {
	enforcer *casbin.Enforcer
	logger   logger.AppLogger
}

func NewAuthorizer(model, policy string, logger logger.AppLogger) (*Authorizer, error) {
	_, err := os.Stat(model)
	if err != nil {
		if os.IsNotExist(err) {
			return nil, errors.NewAppError("file doesn't exist - %s", model)
		} else {
			return nil, errors.WrapError(err, "file inaccessible", model)
		}
	}

	_, err = os.Stat(policy)
	if err != nil {
		if os.IsNotExist(err) {
			return nil, errors.NewAppError("file doesn't exist - %s", model)
		} else {
			return nil, errors.WrapError(err, "file inaccessible", model)
		}
	}

	enforcer := casbin.NewEnforcer(model, policy)
	return &Authorizer{
		enforcer: enforcer,
		logger:   logger,
	}, nil
}

func (a *Authorizer) Authorize(subject, object, action string) error {
	if !a.enforcer.Enforce(subject, object, action) {
		msg := fmt.Sprintf("%s not permitted to %s to %s", subject, action, object)
		a.logger.Error(msg)
		st := status.New(codes.PermissionDenied, msg)
		return st.Err()
	}
	return nil
}
