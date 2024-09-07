package model

import (
	"github.com/diogoalvesf/meu-primeiro-crud-go/src/configuration/logger"
	"github.com/diogoalvesf/meu-primeiro-crud-go/src/configuration/rest_err"
	"go.uber.org/zap"
)

func (ud *UserDomain) CreateUser() *rest_err.RestErr {
	logger.Info("Init createUser model", zap.String("jorney", "createUser"))

	ud.EncryptPassword()
	return nil
}
