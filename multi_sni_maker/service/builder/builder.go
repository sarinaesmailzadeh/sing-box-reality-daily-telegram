package builder

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"sing-box-telegram/entity"
	"strconv"
	"strings"

	"github.com/google/uuid"
)

// SetServerIP  Get preferred outbound ip of this machine
func (b *Builder) SetServerIP() *Builder {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		fmt.Println("error during the SetServerIP ", err)
		return nil
	}
	defer conn.Close()

	localAddr := conn.LocalAddr().(*net.UDPAddr)

	b.ServerIP = localAddr.IP.String()
	return b
}

// SetSettingsFile returns the settings file
func (b *Builder) SetSettingsFile() *Builder {

	// Open our jsonFile
	jsonFile, err := os.Open("./setting.json")
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println("error open setting file", err)
		return nil
	}

	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	byteValue, err := io.ReadAll(jsonFile)
	if err != nil {
		fmt.Println("error ReadAll json file", err)
		return nil
	}

	setting := entity.Setting{}
	// we unmarshal our byteArray which contains our
	err = json.Unmarshal(byteValue, &setting)
	if err != nil {
		fmt.Println("error unmarshal setting", err)
		return nil
	}

	if len(setting.Ports) != len(setting.Domains) || len(setting.GRPC) != len(setting.Ports) {
		fmt.Println("length ports and domain is not equals")
		return nil
	}

	b.Setting = setting
	return b

}

// SetConfigurations sets the xray configuration
func (b *Builder) SetConfigurations() *Builder {

	var StringConfigAll string
	var SliceConfigAll []string

	if len(b.Setting.Domains) == 0 || len(b.Setting.Ports) == 0 || len(b.Setting.GRPC) == 0 {
		fmt.Println("length ports or domain or grpc is zero")
		return nil
	}

	if b.privateKey == "" || b.publicKey == "" {
		fmt.Println("private key or public key is empty")
		return nil
	}

	b.newReality.Inbounds = make([]entity.Inbound, len(b.Setting.Domains))

	shortID := GenerateRandomString(4)

	for counter := 0; counter < len(b.Setting.Domains); counter++ {

		var inbound entity.Inbound
		inbound.Listen = "0.0.0.0"
		inbound.Port = b.Setting.Ports[counter]
		inbound.Protocol = "vless"
		inbound.Settings.Clients = make([]entity.Client, 1)
		inbound.Settings.Clients[0].Flow = ""
		inbound.Settings.Clients[0].ID = uuid.New().String()
		inbound.Settings.Decryption = "none"
		inbound.StreamSettings.Network = "tcp"
		inbound.StreamSettings.Security = "reality"
		inbound.StreamSettings.RealitySettings.Show = false
		inbound.StreamSettings.RealitySettings.Dest = b.Setting.Domains[counter] + ":" + strconv.Itoa(b.Setting.Ports[counter])
		inbound.StreamSettings.RealitySettings.Xver = 0
		inbound.StreamSettings.RealitySettings.ServerNames = []string{b.Setting.Domains[counter]}
		inbound.StreamSettings.RealitySettings.PrivateKey = b.privateKey
		inbound.StreamSettings.RealitySettings.MinClientVer = ""
		inbound.StreamSettings.RealitySettings.MaxClientVer = ""
		inbound.StreamSettings.RealitySettings.MaxTimeDiff = 0
		inbound.StreamSettings.RealitySettings.ShortIds = []string{shortID}
		inbound.StreamSettings.RealitySettings.SpiderX = "/doggo"

		name := b.Setting.ChannelName + "-" + b.Setting.Domains[counter]
		//GRPC setting
		if b.Setting.GRPC[counter] {
			inbound.StreamSettings.Network = "grpc"
		}

		b.newReality.Inbounds[counter] = inbound

		//capture setting
		headerType := ""
		if inbound.StreamSettings.Network == "tcp" {
			headerType = "&headerType=none"
		} else {
			headerType = ""
		}

		StringConfig := "vless://" + inbound.Settings.Clients[0].ID + "@" + b.ServerIP + ":" + strconv.Itoa(b.Setting.Ports[counter]) +
			"?encryption=none&security=reality&sni=" + b.Setting.Domains[counter] +
			"&fp=chrome&pbk=" + b.publicKey + "&sid=" + inbound.StreamSettings.RealitySettings.ShortIds[0] + "&spx=%2Fdoggo&type=" + inbound.StreamSettings.Network + headerType + "#" + name

		if counter == 0 {
			b.StringConfigZero = StringConfig
		}

		StringConfigAll += StringConfig + "\n"

		SliceConfigAll = append(SliceConfigAll, StringConfig)

	}

	b.StringConfigAll = StringConfigAll
	b.SliceConfigAll = SliceConfigAll

	return b

}

// SetPublicKeyAndPrivateKey read public and private key from key pair
func (b *Builder) SetPublicKeyAndPrivateKey() *Builder {

	dat, err := os.ReadFile("./key_pair.txt")
	if err != nil {
		fmt.Println("error during the key pair")
		return nil
	}
	allData := string(dat)

	allData = strings.TrimSpace(allData)
	allData = strings.ReplaceAll(allData, " ", "")

	privateKeyFirst := RemoveRightPart(allData, "Publickey:")
	privateKey := RemoveLeftPart(privateKeyFirst, "Privatekey:")

	pubAns := strings.SplitAfter(allData, "Publickey:")
	publicKey := pubAns[1]

	b.privateKey = privateKey
	b.publicKey = publicKey

	return b

}

// SetBlock block Iranian and Chinese and porn websites
func (b *Builder) SetBlock() *Builder {

	b.newReality.Log.Loglevel = "warning"
	b.newReality.Routing.DomainStrategy = "IPOnDemand"
	b.newReality.Routing.Rules = make([]entity.Rule, 2)
	b.newReality.Routing.Rules[0] = entity.Rule{
		Type:        "field",
		IP:          []string{"geoip:cn", "geoip:ir"},
		OutboundTag: "block",
	}

	b.newReality.Routing.Rules[1] = entity.Rule{
		Type:        "field",
		Domain:      []string{"geosite:category-porn"},
		OutboundTag: "block",
	}

	b.newReality.Outbounds = make([]entity.Outbound, 2)
	b.newReality.Outbounds[0] = entity.Outbound{
		Tag:      "direct",
		Protocol: "freedom",
	}
	b.newReality.Outbounds[1] = entity.Outbound{
		Tag:      "block",
		Protocol: "blackhole",
	}

	return b

}

func (b *Builder) Save() *Builder {

	if b.StringConfigAll == "" || b.StringConfigZero == "" || len(b.SliceConfigAll) == 0 {
		fmt.Println("string config all or string config zero or slice config all is empty")
		return nil
	}

	//save new Reality in file
	err := WriteFile("./config.json", b.newReality)
	if err != nil {
		log.Fatal("error during the WriteFile ", err)
		return nil
	}

	return b

}
