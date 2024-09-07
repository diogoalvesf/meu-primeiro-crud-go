package service

import (
	"github.com/diogoalvesf/meu-primeiro-crud-go/src/configuration/logger"
	"github.com/diogoalvesf/meu-primeiro-crud-go/src/configuration/rest_err"
	"github.com/diogoalvesf/meu-primeiro-crud-go/src/model"
	"go.uber.org/zap"
)

func (ud *userDomainService) CreateUser(
	userDomain model.UserDomainInterface,
) *rest_err.RestErr {
	logger.Info("Init createUser model", zap.String("jorney", "createUser"))

	userDomain.EncryptPassword()
	return nil
}
