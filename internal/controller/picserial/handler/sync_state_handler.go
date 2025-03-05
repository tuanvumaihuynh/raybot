package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"log/slog"
	"strconv"

	"github.com/tbe-team/raybot/internal/model"
	"github.com/tbe-team/raybot/internal/service"
)

// syncStateType is the type of sync state received from the PIC
type syncStateType uint8

// UnmarshalJSON implements the json.Unmarshaler interface.
func (s *syncStateType) UnmarshalJSON(data []byte) error {
	n, err := strconv.ParseUint(string(data), 10, 8)
	if err != nil {
		return fmt.Errorf("parse uint8: %w", err)
	}

	switch n {
	case 0:
		*s = syncStateTypeBattery
	case 1:
		*s = syncStateTypeCharge
	case 2:
		*s = syncStateTypeDischarge
	case 3:
		*s = syncStateTypeDistanceSensor
	case 4:
		*s = syncStateTypeLiftMotor
	case 5:
		*s = syncStateTypeDriveMotor
	default:
		return fmt.Errorf("invalid sync state type: %s", string(data))
	}
	return nil
}

const (
	syncStateTypeBattery syncStateType = iota
	syncStateTypeCharge
	syncStateTypeDischarge
	syncStateTypeDistanceSensor
	syncStateTypeLiftMotor
	syncStateTypeDriveMotor
)

type SyncStateMessage struct {
	StateType syncStateType   `json:"state_type"`
	Data      json.RawMessage `json:"data"`
}

type SyncStateHandler struct {
	robotService service.RobotService
	log          *slog.Logger
}

func NewSyncStateHandler(robotService service.RobotService) *SyncStateHandler {
	return &SyncStateHandler{
		robotService: robotService,
		log: slog.With(
			slog.String("module", "pic"),
			slog.String("handler", "SyncStateHandler"),
		),
	}
}

func (h *SyncStateHandler) Handle(ctx context.Context, msg SyncStateMessage) {
	params := service.UpdateRobotStateParams{}
	switch msg.StateType {
	case syncStateTypeBattery:
		var temp struct {
			Current      uint16   `json:"current"`
			Temp         uint8    `json:"temp"`
			Voltage      uint16   `json:"voltage"`
			CellVoltages []uint16 `json:"cell_voltages"`
			Percent      uint8    `json:"percent"`
			Fault        uint8    `json:"fault"`
			Health       uint8    `json:"health"`
			Status       uint8    `json:"status"`
		}
		if err := json.Unmarshal(msg.Data, &temp); err != nil {
			h.log.Error("failed to unmarshal battery data", slog.Any("error", err), slog.Any("data", msg.Data))
			return
		}

		params.SetBattery = true
		params.Battery = service.BatteryParams{
			Current:      temp.Current,
			Temp:         temp.Temp,
			Voltage:      temp.Voltage,
			CellVoltages: temp.CellVoltages,
			Percent:      temp.Percent,
			Fault:        temp.Fault,
			Health:       temp.Health,
			Status:       temp.Status,
		}

	case syncStateTypeCharge:
		var temp struct {
			CurrentLimit uint16 `json:"current_limit"`
			Enabled      uint8  `json:"enabled"`
		}
		if err := json.Unmarshal(msg.Data, &temp); err != nil {
			h.log.Error("failed to unmarshal charge data", slog.Any("error", err), slog.Any("data", msg.Data))
			return
		}

		params.SetCharge = true
		params.Charge = service.ChargeParams{
			CurrentLimit: temp.CurrentLimit,
			Enabled:      temp.Enabled == 1,
		}

	case syncStateTypeDischarge:
		var temp struct {
			CurrentLimit uint16 `json:"current_limit"`
			Enabled      uint8  `json:"enabled"`
		}
		if err := json.Unmarshal(msg.Data, &temp); err != nil {
			h.log.Error("failed to unmarshal discharge data", slog.Any("error", err), slog.Any("data", msg.Data))
			return
		}

		params.SetDischarge = true
		params.Discharge = service.DischargeParams{
			CurrentLimit: temp.CurrentLimit,
			Enabled:      temp.Enabled == 1,
		}

	case syncStateTypeDistanceSensor:
		var temp struct {
			FrontDistance uint16 `json:"front_distance"`
			BackDistance  uint16 `json:"back_distance"`
			DownDistance  uint16 `json:"down_distance"`
		}
		if err := json.Unmarshal(msg.Data, &temp); err != nil {
			h.log.Error("failed to unmarshal distance sensor data", slog.Any("error", err), slog.Any("data", msg.Data))
			return
		}

		params.SetDistanceSensor = true
		params.DistanceSensor = service.DistanceSensorParams{
			FrontDistance: temp.FrontDistance,
			BackDistance:  temp.BackDistance,
			DownDistance:  temp.DownDistance,
		}

	case syncStateTypeLiftMotor:
		var temp struct {
			CurrentPosition uint16 `json:"current_position"`
			TargetPosition  uint16 `json:"target_position"`
			IsRunning       uint8  `json:"is_running"`
			Enabled         uint8  `json:"enabled"`
		}
		if err := json.Unmarshal(msg.Data, &temp); err != nil {
			h.log.Error("failed to unmarshal lift motor data", slog.Any("error", err), slog.Any("data", msg.Data))
			return
		}

		params.SetLiftMotor = true
		params.LiftMotor = service.LiftMotorParams{
			CurrentPosition: temp.CurrentPosition,
			TargetPosition:  temp.TargetPosition,
			IsRunning:       temp.IsRunning == 1,
			Enabled:         temp.Enabled == 1,
		}

	case syncStateTypeDriveMotor:
		var temp struct {
			Direction model.DriveMotorDirection `json:"direction"`
			Speed     uint8                     `json:"speed"`
			IsRunning uint8                     `json:"is_running"`
			Enabled   uint8                     `json:"enabled"`
		}
		if err := json.Unmarshal(msg.Data, &temp); err != nil {
			h.log.Error("failed to unmarshal drive motor data", slog.Any("error", err), slog.Any("data", msg.Data))
			return
		}

		params.SetDriveMotor = true
		params.DriveMotor = service.DriveMotorParams{
			Direction: temp.Direction,
			Speed:     temp.Speed,
			IsRunning: temp.IsRunning == 1,
			Enabled:   temp.Enabled == 1,
		}

	default:
		h.log.Error("unknown state type", slog.Int("type", int(msg.StateType)))
		return
	}

	_, err := h.robotService.UpdateRobotState(ctx, params)
	if err != nil {
		h.log.Error("failed to update robot state", slog.Any("error", err))
	}
}
