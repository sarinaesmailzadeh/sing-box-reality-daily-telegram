package main

func Block(newReality RealityJson) RealityJson {

	newReality.Log.Loglevel = "warning"
	newReality.Routing.DomainStrategy = "IPOnDemand"
	newReality.Routing.Rules = make([]Rule, 2)
	newReality.Routing.Rules[0] = Rule{
		Type:        "field",
		IP:          []string{"geoip:cn", "geoip:ir"},
		OutboundTag: "block",
	}

	newReality.Routing.Rules[1] = Rule{
		Type:        "field",
		Domain:      []string{"geosite:category-porn"},
		OutboundTag: "block",
	}

	newReality.Outbounds = make([]Outbound, 1)
	newReality.Outbounds[0] = Outbound{
		Tag:      "direct",
		Protocol: "freedom",
	}

	return newReality

}
