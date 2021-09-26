package data

type EventType string

type Subscriber func([]interface{})

type SubscriberStorage struct {
	store map[EventType][]Subscriber
}

func (store SubscriberStorage) Get(key EventType) []Subscriber {
	_, ok := store.store[key]
	if !ok {
		return []Subscriber{}
	}
	return store.store[key]
}

func (store *SubscriberStorage) Add(key EventType, handler Subscriber) {
	_, ok := store.store[key]
	if !ok {
		store.store[key] = []Subscriber{}
	}
	store.store[key] = append(store.store[key], handler)
}

func NewSubscriberStorage() *SubscriberStorage {
	store := make(map[EventType][]Subscriber)
	return &SubscriberStorage{store}
}

type Emitter struct {
	handlers *SubscriberStorage
}

func (e Emitter) Subscribe(event EventType, handler Subscriber) {
	e.handlers.Add(event, handler)
}

func (e Emitter) Publish(event EventType, args ...interface{}) {
	for _, handler := range e.handlers.Get(event) {
		handler(args)
	}
}

func NewEmitter() Emitter {
	store := NewSubscriberStorage()
	return Emitter{store}
}
