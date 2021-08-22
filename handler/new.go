package handler

import "errors"

type Handler struct {
	Host     string
	Port     int
	Interval int
	Tor      bool // TODO

	Commands map[string]func(...string)
}

// Create new handler and return it.
func NewHandler(host string, port int) Handler {
	return Handler{
		Host:     host,
		Port:     port,
		Interval: 60,
		Commands: make(map[string]func(...string)),
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
