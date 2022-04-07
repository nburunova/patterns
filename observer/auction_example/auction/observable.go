package auction

import (
	"fmt"
	"patterns/observer/auction_example"
	"sync"
)

type Auction struct {
	observers []auction_example.Observer
	observersMu sync.Mutex
	currentBidderID int
	currentBid int
}

func NewAuction() *Auction {
	return &Auction{
		observers: make([]auction_example.Observer, 0),
	}
}

func (a *Auction) GetCurrentBid() int {
	return a.currentBid
}

func (a *Auction) GetBidderID() int {
	return a.currentBidderID
}

func (a *Auction) AcceptBid(id, bid int) {
	fmt.Printf("Bid offer: ID %v %v\n", id, bid)
	if bid > a.currentBid {
		a.currentBidderID = id
		a.currentBid = bid
	}
}

func (a *Auction) Add(o auction_example.Observer) {
	a.observersMu.Lock()
	defer a.observersMu.Unlock()
	a.observers = append(a.observers, o)
}

func (a *Auction) Remove(o auction_example.Observer) {
	a.observersMu.Lock()
	defer a.observersMu.Unlock()
	for i, obs := range a.observers {
		if obs.ID() != o.ID() {
			continue
		}
		a.observers = append(a.observers[:i], a.observers[i+1:]...)
		break
	}
}

func (a *Auction) Notify() {
	a.observersMu.Lock()
	defer a.observersMu.Unlock()
	for _, obs := range a.observers {
		obs.Update()
	}
}