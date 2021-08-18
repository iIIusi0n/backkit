package network

import (
	"encoding/json"
	"net"
)

func GetPublicIP() (string, error) {
	type IpResponse struct {
		IP string `json:"ip"`
	}

	ipApis := []string{
		"https://api.ipify.org?format=json",
		"https://api.myip.com",
		"https://api.ip.pe.kr/json",
	}
	var rawIpResp string
	var err error
	for _, apiAddr := range ipApis {
		rawIpResp, err = DownloadString(apiAddr)
		if err == nil {
			break
		}
	}
	if rawIpResp == "" {
		return "", err
	}
	var parsedIpResp IpResponse
	err = json.Unmarshal([]byte(rawIpResp), &parsedIpResp)
	if err != nil {
		return "", err
	}
	return parsedIpResp.IP, nil
}

func GetPrivateIP() (string, error) {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		return "", err
	}
	defer conn.Close()

	localAddr := conn.LocalAddr().(*net.UDPAddr)
	return localAddr.IP.String(), nil
}
