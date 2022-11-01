package relaycontroller

import (
	"errors"

	"gitea.k8s.attiss.xyz/attiss/rpi-switch/config"
	"github.com/stianeikeland/go-rpio"
	"go.uber.org/zap"
)

const (
	StateOff = iota
	StateOn
)

type RelayController struct {
	pinAssignment config.PinAssignment
	logger        *zap.Logger
}

func New(config config.Config, logger *zap.Logger) (RelayController, error) {
	return RelayController{
		pinAssignment: config.PinAssignment,
		logger:        logger,
	}, nil
}

func (rc RelayController) Init() error {
	if err := rpio.Open(); err != nil {
		rc.logger.Error("failed to open and map gpio memory", zap.Error(err))
		return err
	}

	for relay, pin := range rc.pinAssignment {
		rpiPin := rpio.Pin(pin)
		rpiPin.Output()
		rpiPin.High()
		rc.logger.Debug("successfully initialized pin", zap.String("relay", relay), zap.Uint8("pin", pin))
	}

	rc.logger.Info("successfully initialized pins")
	return nil
}

func (rc RelayController) SetState(relay string, state int) error {
	logger := rc.logger.With(zap.String("relay", relay), zap.Int("state", state))

	pin, defined := rc.pinAssignment[relay]
	if !defined {
		logger.Error("relay not defined")
		return errors.New("relay not defined")
	}

	rpiPin := rpio.Pin(pin)
	rpiPin.Output()

	switch state {
	case StateOn:
		rpiPin.Low()
	case StateOff:
		rpiPin.High()
	default:
		logger.Error("invalid state")
		return errors.New("invalid state")
	}

	logger.Info("successfully set relay state")
	return nil
}

func (rc RelayController) Toggle(relay string) error {
	logger := rc.logger.With(zap.String("relay", relay))

	pin, defined := rc.pinAssignment[relay]
	if !defined {
		logger.Error("relay not defined")
		return errors.New("relay not defined")
	}

	rpiPin := rpio.Pin(pin)
	rpiPin.Output()
	rpiPin.Toggle()

	logger.Info("successfully toggled relay state")
	return nil
}

func (rc RelayController) GetState(relay string) (int, error) {
	logger := rc.logger.With(zap.String("relay", relay))

	pin, defined := rc.pinAssignment[relay]
	if !defined {
		logger.Error("relay not defined")
		return -1, errors.New("relay not defined")
	}

	rpiPin := rpio.Pin(pin)

	var state int
	switch rpiPin.Read() {
	case rpio.Low:
		state = StateOn
	case rpio.High:
		state = StateOff
	}

	return state, nil
}

func (rc RelayController) Close() {
	rc.logger.Info("unmapping gpio memory")
	rpio.Close()
}
