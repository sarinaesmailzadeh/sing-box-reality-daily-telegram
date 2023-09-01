package main

import (
	"strconv"

	"github.com/google/uuid"
)

func RenewConfigurations(setting Setting, serverIP string, newReality RealityJson) (
	StringConfigZero string,
	StringConfigAll string,
	outReality RealityJson,
	SliceConfigAll []string) {

	privateKey, publicKey := getPublicKeyAndPrivateKey()

	newReality.Inbounds = make([]Inbound, len(setting.Domains))

	shortID := GenerateRandomString(16)

	for counter := 0; counter < len(setting.Domains); counter++ {

		var inbound Inbound
		inbound.Listen = "0.0.0.0"
		inbound.Port = setting.Ports[counter]
		inbound.Protocol = "vless"
		inbound.Settings.Clients = make([]Client, 1)
		inbound.Settings.Clients[0].Flow = ""
		inbound.Settings.Clients[0].ID = uuid.New().String()
		inbound.Settings.Decryption = "none"
		inbound.StreamSettings.Network = "tcp"
		inbound.StreamSettings.Security = "reality"
		inbound.StreamSettings.RealitySettings.Show = false
		inbound.StreamSettings.RealitySettings.Dest = setting.Domains[counter] + ":" + strconv.Itoa(setting.Ports[counter])
		inbound.StreamSettings.RealitySettings.Xver = 0
		inbound.StreamSettings.RealitySettings.ServerNames = []string{setting.Domains[counter]}
		inbound.StreamSettings.RealitySettings.PrivateKey = privateKey
		inbound.StreamSettings.RealitySettings.MinClientVer = ""
		inbound.StreamSettings.RealitySettings.MaxClientVer = ""
		inbound.StreamSettings.RealitySettings.MaxTimeDiff = 0
		inbound.StreamSettings.RealitySettings.ShortIds = []string{shortID}
		inbound.StreamSettings.RealitySettings.SpiderX = "/doggo"

		name := setting.ChannelName + "-" + setting.Domains[counter]
		newReality.Inbounds[counter] = inbound

		//GRPC setting
		if setting.GRPC[counter] {
			inbound.StreamSettings.Network = "grpc"
		}

		//capture setting

		StringConfig := "vless://" + inbound.Users[0].UUID + "@" + serverIP + ":" + strconv.Itoa(setting.Ports[counter]) +
			"?encryption=none&flow=xtls-rprx-vision&security=reality&sni=" + setting.Domains[counter] +
			"&fp=chrome&pbk=" + publicKey + "&sid=" + inbound.TLS.Reality.ShortID[0] + "&type=tcp&headerType=none#" + name

		if counter == 0 {
			StringConfigZero = StringConfig
		}

		StringConfigAll += StringConfig + "\n"

		SliceConfigAll = append(SliceConfigAll, StringConfig)

	}

	outReality = newReality
	return StringConfigZero, StringConfigAll, outReality, SliceConfigAll

}
