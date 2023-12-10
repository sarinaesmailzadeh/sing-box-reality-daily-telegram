package subscribe

import (
	"fmt"
	"math/rand"
	"sing-box-telegram/service/builder"
	"sing-box-telegram/service/execute"
	"strconv"
)

type Subscribe struct {
	builder                       *builder.Builder
	AggregateSubscriptionNameLink string
	SubscriptionNameLink          string
	execute                       *execute.Execute
}

func NewSubscribe(builder *builder.Builder,
	executeInstance *execute.Execute) *Subscribe {

	sub := Subscribe{
		builder: builder,
		execute: executeInstance,
	}

	if len(sub.builder.Setting.AggregateSubscriptions) > 0 {
		sub.setAggregateSubscriptions()
	}

	err := sub.doSubscribe()
	if err != nil {
		fmt.Println("Error during subscribe ", err.Error())
	}

	return &sub
}

func (s *Subscribe) setAggregateSubscriptions() {

	AllConfigurations := s.builder.StringConfigAll

	for _, link := range s.builder.Setting.AggregateSubscriptions {

		fmt.Println("Curl URL: ", link)
		result, err := getURL(link)
		if err != nil {
			fmt.Println("Error getting aggregate subscription ", err.Error())
			continue
		}
		AllConfigurations = AllConfigurations + "\n" + result + "\n"
	}

	var subscriptionNameLink string

	if s.builder.Setting.DynamicSubscription {
		randomUUID := rand.Intn(10000)
		subscriptionNameLink = "aggregate." + strconv.Itoa(randomUUID) + ".txt"
	} else {
		subscriptionNameLink = "aggregate.txt"
	}

	SaveSubscribe("./"+subscriptionNameLink, AllConfigurations)

	s.AggregateSubscriptionNameLink = subscriptionNameLink

}

func (s *Subscribe) doSubscribe() error {

	if s.builder.Setting.DynamicSubscription {
		randomUUID := rand.Intn(10000)
		s.SubscriptionNameLink = "subscribe." + strconv.Itoa(randomUUID) + ".txt"
	} else {
		s.SubscriptionNameLink = "subscribe.txt"
	}

	SaveSubscribe("./"+s.SubscriptionNameLink, s.builder.StringConfigAll)

	_, err := s.execute.ExecuteCommand("./make_subscribe.sh")
	if err != nil {
		fmt.Printf("error make_subscribe %s", err)
		return err
	}

	return err

}
