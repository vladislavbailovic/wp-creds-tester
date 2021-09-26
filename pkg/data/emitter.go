package data

type Subscriber func([]interface{})

type SubscriberStorage struct {
	store map[string][]Subscriber
}

func (store SubscriberStorage) Get(key string) []Subscriber {
	_, ok := store.store[key]
	if !ok {
		return []Subscriber{}
	}
	return store.store[key]
}

func (store *SubscriberStorage) Add(key string, handler Subscriber) {
	_, ok := store.store[key]
	if !ok {
		store.store[key] = []Subscriber{}
	}
	store.store[key] = append(store.store[key], handler)
}

func NewSubscriberStorage() *SubscriberStorage {
	store := make(map[string][]Subscriber)
	return &SubscriberStorage{store}
}

type Emitter struct {
	handlers *SubscriberStorage
}

func (e Emitter) Subscribe(event string, handler Subscriber) {
	e.handlers.Add(event, handler)
}

func (e Emitter) Publish(event string, args ...interface{}) {
	for _, handler := range e.handlers.Get(event) {
		handler(args)
	}
}

func NewEmitter() Emitter {
	store := NewSubscriberStorage()
	return Emitter{store}
}
