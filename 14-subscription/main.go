package main

import "time"

// Item represents an item with
type Item struct{ Tile, Channel, GUID string }

// Fetcher fetches Items, returning the time when the next fetch
// should be attempted returns error if any
type Fetcher interface {
	Fetch() (item []Item, next time.Time, err error)
}

// Subscription delivers Items over a channel.
// Close() cancels the subscription, closes the Updates channel,
// and returns the last fetch error if one exists
type Subscription interface {
	Updates() <-chan Item
	Close() error
}

type sub struct {
	fetcher Fetcher
	updates chan Item
	closing chan chan error
}

func Subscribe(fetcher Fetcher) Subscription {
	s := &sub{
		fetcher: fetcher,
		updates: make(chan Item),
		closing: make(chan chan error),
	}
	return s
}

func (s *sub) Updates() <-chan Item {
	return s.updates
}

func (s *sub) Close() error {
	errc := make(chan error)
	s.closing <- errc
	return <-errc
}

func main() {

}
