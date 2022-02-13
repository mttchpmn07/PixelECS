package messenger

import "log"

type Subscriber interface {
	HandleBroadcast(key string, content interface{})
}

type Messenger interface {
	Subscribe(messageKey string, sub Subscriber)
	Broadcast(messageKey string, content interface{})
}

type messenger struct{}

type messageQue struct {
	messages map[string][]Subscriber
}

var que messageQue

func init() {
	que = messageQue{
		messages: map[string][]Subscriber{},
	}
}

func NewMessenger() Messenger {
	return &messenger{}
}

func (mq *messageQue) subscribe(messageKey string, sub Subscriber) {
	if _, ok := que.messages[messageKey]; ok {
		que.messages[messageKey] = append(que.messages[messageKey], sub)
		return
	}
	que.messages[messageKey] = []Subscriber{sub}
}

func (mq *messageQue) broadcast(messageKey string, content interface{}) {
	if subscribers, ok := que.messages[messageKey]; ok {
		for _, sys := range subscribers {
			sys.HandleBroadcast(messageKey, content)
		}
		return
	}
	log.Printf("No subscribers to %s", messageKey)
}

func (m *messenger) Subscribe(messageKey string, sub Subscriber) {
	que.subscribe(messageKey, sub)
}

func (m *messenger) Broadcast(messageKey string, content interface{}) {
	que.broadcast(messageKey, content)
}
