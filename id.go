package backkit

import (
	"crypto/sha256"
	"fmt"

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
	hashHexString := ByteToHexString(sha256.Sum256([]byte(fmt.Sprintf("%s|%s", publicIP, id))))
	return hashHexString, nil
}
