package messenger

type Messenger interface {
	Subscribe(messageKey string, callback func(interface{}))
	Broadcast(messageKey string, message interface{})
}

type messenger struct{}

type messageQue struct {
	messages map[string]func(content interface{})
}

var que messageQue

func init() {
	que = messageQue{
		messages: map[string]func(content interface{}){},
	}
}

func NewMessenger() Messenger {
	return &messenger{}
}

func (mq *messageQue) subscribe(messageKey string, callback func(content interface{})) {
	que.messages[messageKey] = callback
}

func (mq *messageQue) broadcast(messageKey string, message interface{}) {
	que.messages[messageKey](message)
}

func (m *messenger) Subscribe(messageKey string, callback func(interface{})) {
	que.subscribe(messageKey, callback)
}

func (m *messenger) Broadcast(messageKey string, message interface{}) {
	que.broadcast(messageKey, message)
}
