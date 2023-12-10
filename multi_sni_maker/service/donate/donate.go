package donate

import (
	"fmt"
	"net/http"
	"net/url"
	"sing-box-telegram/entity"
	"sing-box-telegram/service/builder"
	"sing-box-telegram/service/subscribe"
)

type Donate struct {
	builder   *builder.Builder
	subscribe *subscribe.Subscribe
}

func NewDonate(builderInstance *builder.Builder,
	subscribe *subscribe.Subscribe) *Donate {
	return &Donate{
		builder:   builderInstance,
		subscribe: subscribe,
	}
}

func (d *Donate) CallDonate() {
	//call donate endpoint
	if len(d.builder.Setting.DonateURL) > 4 {

		if len(d.builder.Setting.AggregateSubscriptions) > 0 {

			subscriptionNameLinkFull := "http://" + d.builder.ServerIP + "/" + d.subscribe.AggregateSubscriptionNameLink
			CallDonate(subscriptionNameLinkFull, d.builder.Setting)

		} else {

			subscriptionNameLinkFull := "http://" + d.builder.ServerIP + "/" + d.subscribe.SubscriptionNameLink
			CallDonate(subscriptionNameLinkFull, d.builder.Setting)
		}

	}
}

func CallDonate(subscriptionLink string, setting entity.Setting) {
	fmt.Println("curl Donate...")

	var address string
	var donateURL string

	if setting.DonateURL == "stop" {
		address = "https://yebekhe.000webhostapp.com/donate/"
		donateURL = address + "?remove=true" + "&username=" + setting.ChannelName
	} else if setting.DonateURL == "yebekhe" {
		address = "https://yebekhe.000webhostapp.com/donate/"
		donateURL = address + "?url=" + url.QueryEscape(subscriptionLink) + "&username=" + setting.ChannelName
	}

	fmt.Println(donateURL)

	err := curlFunc(donateURL)
	if err != nil {
		fmt.Println(err)
	}

}

func curlFunc(address string) error {

	// Encode the URL
	encodedURL, err := url.Parse(address)
	if err != nil {
		fmt.Println("Error parsing URL:", err)
	}

	// Make the GET request
	resp, err := http.Get(encodedURL.String())
	if err != nil {
		fmt.Println("Error making GET request:", err)
		return nil
	}
	defer resp.Body.Close()

	return nil

}
