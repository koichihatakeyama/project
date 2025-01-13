package event

import (
	"context"
	"sync"
)

type Event interface {
	Name() string
	Entity() interface{}
}

type EventDispatcher struct {
	listeners map[string][]EventListener
	mutex     sync.RWMutex
}

type EventListener interface {
	Handle(ctx context.Context, event Event) error
}

func NewEventDispatcher() *EventDispatcher {
	return &EventDispatcher{
		listeners: make(map[string][]EventListener),
	}
}

func (d *EventDispatcher) AddListener(eventName string, listener EventListener) {
	d.mutex.Lock()
	defer d.mutex.Unlock()
	d.listeners[eventName] = append(d.listeners[eventName], listener)
}

func (d *EventDispatcher) Dispatch(ctx context.Context, event Event) error {
	d.mutex.RLock()
	listeners := d.listeners[event.Name()]
	d.mutex.RUnlock()

	for _, listener := range listeners {
		if err := listener.Handle(ctx, event); err != nil {
			return err
		}
	}
	return nil
}
