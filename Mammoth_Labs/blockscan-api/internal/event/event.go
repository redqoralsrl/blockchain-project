package event

import "sync"

// Event - 기본 이벤트 인터페이스
type Event struct {
	Type string
	Data interface{}
}

// Handler - 이벤트 핸들러 함수 타입
type Handler func(event Event)

// EventBus - 이벤트 버스 구조
type EventBus struct {
	mu          sync.RWMutex
	subscribers map[string][]Handler
}

// NewEventBus - EventBus 생성자
func NewEventBus() *EventBus {
	return &EventBus{
		subscribers: make(map[string][]Handler),
	}
}

// Subscribe - 이벤트 구독
func (eb *EventBus) Subscribe(eventType string, handler Handler) {
	eb.mu.Lock()
	defer eb.mu.Unlock()

	eb.subscribers[eventType] = append(eb.subscribers[eventType], handler)
}

// Publish - 이벤트 발행
func (eb *EventBus) Publish(eventType string, event Event) {
	eb.mu.RLock()
	defer eb.mu.RUnlock()

	handlers, ok := eb.subscribers[eventType]
	if !ok {
		return
	}

	for _, handler := range handlers {
		go handler(event)
	}
}
