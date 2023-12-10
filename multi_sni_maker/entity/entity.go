package entity

type RealityJson struct {
	Log struct {
		Loglevel string `json:"loglevel"`
	} `json:"log"`
	Inbounds []Inbound `json:"inbounds"`
	Routing  struct {
		DomainStrategy string `json:"domainStrategy"`
		Rules          []Rule `json:"rules"`
	} `json:"routing"`
	Outbounds []Outbound `json:"outbounds"`
}

type Rule struct {
	Type        string   `json:"type"`
	IP          []string `json:"ip,omitempty"`
	OutboundTag string   `json:"outboundTag"`
	Domain      []string `json:"domain,omitempty"`
}

type Outbound struct {
	Tag      string `json:"tag"`
	Protocol string `json:"protocol"`
}

type Inbound struct {
	Listen   string `json:"listen"`
	Port     int    `json:"port"`
	Protocol string `json:"protocol"`
	Settings struct {
		Clients    []Client `json:"clients"`
		Decryption string   `json:"decryption"`
	} `json:"settings"`
	StreamSettings struct {
		Network         string `json:"network"`
		Security        string `json:"security"`
		RealitySettings struct {
			Show         bool     `json:"show"`
			Dest         string   `json:"dest"`
			Xver         int      `json:"xver"`
			ServerNames  []string `json:"serverNames"`
			PrivateKey   string   `json:"privateKey"`
			MinClientVer string   `json:"minClientVer"`
			MaxClientVer string   `json:"maxClientVer"`
			MaxTimeDiff  int      `json:"maxTimeDiff"`
			ShortIds     []string `json:"shortIds"`
			SpiderX      string   `json:"spiderX"`
		} `json:"realitySettings"`
	} `json:"streamSettings"`
}

type Client struct {
	ID   string `json:"id"`
	Flow string `json:"flow"`
}

type Setting struct {
	Ports                  []int    `json:"ports"`
	Domains                []string `json:"domains"`
	GRPC                   []bool   `json:"grpc"`
	BotToken               string   `json:"bot_token"`
	ChatID                 string   `json:"chat_id"`
	DonateURL              string   `json:"donate_url"`
	DynamicSubscription    bool     `json:"dynamic_subscription"`
	ChannelName            string   `json:"channel_name"`
	SendVNstat             bool     `json:"send_vnstat"`
	AggregateSubscriptions []string `json:"aggregate_subscriptions"`
	SendSubscriptions      bool     `json:"send_subscriptions"`
	SendConfiguration      string   `json:"send_configuration"`
}
