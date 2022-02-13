package messenger

import "fmt"

type Subscriber interface {
	HandleBroadcast(key string, content interface{})
}

type Messenger interface {
	Subscribe(messageKey string, sub Subscriber)
	Broadcast(messageKey string, content interface{}) error
}

type messageQue struct {
	messages map[string][]Subscriber
}

func NewMessenger() Messenger {
	return &messageQue{
		messages: map[string][]Subscriber{},
	}
}

func (que *messageQue) Subscribe(messageKey string, sub Subscriber) {
	if _, ok := que.messages[messageKey]; ok {
		que.messages[messageKey] = append(que.messages[messageKey], sub)
		return
	}
	que.messages[messageKey] = []Subscriber{sub}
}

func (que *messageQue) Broadcast(messageKey string, content interface{}) error {
	if subscribers, ok := que.messages[messageKey]; ok {
		for _, sys := range subscribers {
			sys.HandleBroadcast(messageKey, content)
		}
		return nil
	}
	return fmt.Errorf("no subscribers to '%s'", messageKey)
}
