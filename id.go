package backkit

import (
	"crypto/sha256"
	"fmt"
	"os"

	"github.com/denisbrodbeck/machineid"
)

func GenerateUniqueID() (string, error) {
	publicIP, err := GetPublicIP()
	if err != nil {
		return "<error>", err
	}
	id, err := machineid.ID()
	if err != nil {
		return "<error>", err
	}
	host, err := os.Hostname()
	if err != nil {
		return "<error>", err
	}
	hashHexString := ByteToHexString(sha256.Sum256([]byte(fmt.Sprintf("%s|%s|%s", publicIP, id, host))))
	return hashHexString, nil
}
