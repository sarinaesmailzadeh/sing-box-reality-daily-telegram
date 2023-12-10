package telegram

import (
	"fmt"
	"net/http"
	"net/url"
	"os"
	"sing-box-telegram/entity"
	"sing-box-telegram/service/builder"
	"sing-box-telegram/service/subscribe"
	"time"
)

type TelegramClient struct {
	builder   *builder.Builder
	subscribe *subscribe.Subscribe
}

func NewTelegramClient(builder *builder.Builder,
	subscribe *subscribe.Subscribe) *TelegramClient {

	tel := TelegramClient{
		builder:   builder,
		subscribe: subscribe,
	}

	if len(tel.builder.Setting.BotToken) > 0 && len(tel.builder.Setting.ChatID) > 0 {
		tel.callFirst()
		tel.callAll()
		tel.callSend()
	}

	return &tel
}

func (c *TelegramClient) SendVNstat() {

	//send vnstat file
	if c.builder.Setting.SendVNstat {

		logByte, err := os.ReadFile("./log.txt")
		if err != nil {
			fmt.Println("error read log file ", err)
		}
		stringLogByte := string(logByte)
		err = CallTelegram(stringLogByte, c.builder.Setting)
		if err != nil {
			fmt.Println("error call log ", err)
		}
	}
}

func (c *TelegramClient) callFirst() {
	if c.builder.Setting.SendConfiguration == "first" {
		err := CallTelegram(c.builder.StringConfigZero, c.builder.Setting)
		if err != nil {
			fmt.Println("error first call ", err)
		}
	}
}

func (c *TelegramClient) callAll() {
	if c.builder.Setting.SendConfiguration == "all" {
		for _, currentConfig := range c.builder.SliceConfigAll {
			err := CallTelegram(currentConfig, c.builder.Setting)
			if err != nil {
				fmt.Println("error all call ", err)
			}
			time.Sleep(1 * time.Second)
		}
	}
}

func (c *TelegramClient) callSend() {

	if c.builder.Setting.SendSubscriptions {

		if len(c.builder.Setting.AggregateSubscriptions) > 0 {
			err := CallTelegram("Aggregate link is: \nhttp://"+c.builder.ServerIP+"/"+c.subscribe.AggregateSubscriptionNameLink, c.builder.Setting)
			if err != nil {
				fmt.Println("error Aggregate call ", err)
			}
		} else {
			err := CallTelegram("You can also use this link to subscribe to all configuration:\nhttp://"+c.builder.ServerIP+"/"+c.subscribe.SubscriptionNameLink, c.builder.Setting)
			if err != nil {
				fmt.Println("error subscribe call ", err)
			}
		}
	}

}

func CallTelegram(StringConfigZero string, setting entity.Setting) error {
	fmt.Println("curl Telegram...")
	fmt.Println(StringConfigZero)

	// make GET request to API to get user by ID
	telegramUrl := "https://api.telegram.org/bot" + setting.BotToken + "/sendMessage?chat_id=" + setting.ChatID + "&text=" + url.QueryEscape(StringConfigZero)

	// fmt.Println(telegramUrl)

	// Encode the URL
	encodedURL, err := url.Parse(telegramUrl)
	if err != nil {
		fmt.Println("Error parsing URL:", err)
		return err
	}

	// Make the GET request
	resp, err := http.Get(encodedURL.String())
	if err != nil {
		fmt.Println("Error making GET request:", err)
		return err
	}
	defer resp.Body.Close()

	return nil
}
