package backkit

import (
	"encoding/json"
	"io"
	"io/ioutil"
)

type command struct {
	Type string   `json:"type"`
	Args []string `json:"args"`
}

type commands struct {
	Commands []command `json:"commands"`
}

func (h *Handler) parseReceivedCommand(body io.ReadCloser) error {
	bodyBytes, err := ioutil.ReadAll(body)
	if err != nil {
		return err
	}
	bodyString := string(bodyBytes)
	decodedBody := HexStringToByte(bodyString)

	var receivedCommand commands
	err = json.Unmarshal(decodedBody, &receivedCommand)
	if err != nil {
		return err
	}
	for _, command := range receivedCommand.Commands {
		h.Commands[command.Type](command.Args...)
	}
	return nil
}
