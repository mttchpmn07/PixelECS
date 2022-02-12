package messenger

import ecs "github.com/mttchpmn07/PixelECS/core"

type Messenger interface {
	Subscribe(messageKey string, sys *ecs.System)
	Broadcast(messageKey string, content interface{})
}

type messenger struct{}

type messageQue struct {
	messages map[string][]*ecs.System
}

var que messageQue

func init() {
	que = messageQue{
		messages: map[string][]*ecs.System{},
	}
}

func NewMessenger() Messenger {
	return &messenger{}
}

func (mq *messageQue) subscribe(messageKey string, sys *ecs.System) {
	if _, ok := que.messages[messageKey]; ok {
		que.messages[messageKey] = append(que.messages[messageKey], sys)
		return
	}
	que.messages[messageKey] = []*ecs.System{sys}
}

func (mq *messageQue) broadcast(messageKey string, content interface{}) {
	if systems, ok := que.messages[messageKey]; ok {
		for _, sys := range systems {
			sys.Callback(messageKey, content)
		}
	}
}

func (m *messenger) Subscribe(messageKey string, sys *ecs.System) {
	que.subscribe(messageKey, sys)
}

func (m *messenger) Broadcast(messageKey string, content interface{}) {
	que.broadcast(messageKey, content)
}
