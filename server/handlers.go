package server

import (
	"encoding/json"
	"net/http"

	"gitea.k8s.attiss.xyz/attiss/rpi-switch/relaycontroller"
	"github.com/gorilla/mux"
	"go.uber.org/zap"
)

type requestHandler struct {
	relayController *relaycontroller.RelayController
	logger          *zap.Logger
}

func newRequestHandler(rc *relaycontroller.RelayController, logger *zap.Logger) requestHandler {
	return requestHandler{
		relayController: rc,
		logger:          logger,
	}
}

func (rh requestHandler) GetRouter() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/relay/{relay}", rh.SetState).Methods(http.MethodPost)
	router.HandleFunc("/relay/{relay}", rh.GetState).Methods(http.MethodGet)

	return router
}

func (rh requestHandler) SetState(w http.ResponseWriter, r *http.Request) {
	var rp relayProperties

	if err := json.NewDecoder(r.Body).Decode(&rp); err != nil {
		rh.logger.Error("failed to decode request body", zap.Error(err))
		if err := httpError(w, http.StatusBadRequest, err.Error()); err != nil {
			rh.logger.Error("failed to send response", zap.Error(err))
		}
		return
	}

	vars := mux.Vars(r)
	relay, ok := vars["relay"]
	if !ok {
		rh.logger.Error("relay identifier missing from path")
		if err := httpError(w, http.StatusBadRequest, "relay identifier missing from path"); err != nil {
			rh.logger.Error("failed to send response", zap.Error(err))
		}
		return
	}

	if err := rh.relayController.SetState(relay, rp.State); err != nil {
		rh.logger.Error("failed to state relay state", zap.Error(err))
		if err := httpError(w, http.StatusBadRequest, err.Error()); err != nil {
			rh.logger.Error("failed to send response", zap.Error(err))
		}
		return
	}

	if err := httpSuccess(w, rp); err != nil {
		rh.logger.Error("failed to send response", zap.Error(err))
	}
}

func (rh requestHandler) GetState(w http.ResponseWriter, r *http.Request) {
	var rp relayProperties

	vars := mux.Vars(r)
	relay, ok := vars["relay"]
	if !ok {
		rh.logger.Error("relay identifier missing from path")
		if err := httpError(w, http.StatusBadRequest, "relay identifier missing from path"); err != nil {
			rh.logger.Error("failed to send response", zap.Error(err))
		}
		return
	}

	var err error
	if rp.State, err = rh.relayController.GetState(relay); err != nil {
		rh.logger.Error("failed to state relay state", zap.Error(err))
		if err := httpError(w, http.StatusBadRequest, err.Error()); err != nil {
			rh.logger.Error("failed to send response", zap.Error(err))
		}
		return
	}

	if err := httpSuccess(w, rp); err != nil {
		rh.logger.Error("failed to send response", zap.Error(err))
	}
}
