package serviceimpl

import (
	"github.com/tbe-team/raybot/internal/config"
	"github.com/tbe-team/raybot/internal/repository"
	"github.com/tbe-team/raybot/internal/service"
	"github.com/tbe-team/raybot/pkg/validator"
)

type serviceImpl struct {
	robotService  *RobotService
	systemService *SystemService
}

func New(cfgManager *config.Manager, repo repository.Repository, validator validator.Validator) service.Service {
	return &serviceImpl{
		robotService:  NewRobotService(repo.RobotState(), validator),
		systemService: NewSystemService(cfgManager),
	}
}

func (s serviceImpl) RobotService() service.RobotService {
	return s.robotService
}

func (s serviceImpl) SystemService() service.SystemService {
	return s.systemService
}
