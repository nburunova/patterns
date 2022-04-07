package main

import (
	"fmt"
	"patterns/observer/auction_example/auction"
	"time"
)

type Observable interface {
	Add()
}

func main() {
	auctionObservable := auction.NewAuction()
	auctionObservers := make([]*auction.AuctionObserver, 0)
	for i := 1; i <= 4; i++ {
		observer := auction.NewAuctionObserver(i, auctionObservable)
		auctionObservers = append(auctionObservers, observer)
		auctionObservable.Add(observer)
	}
	
	for j := 0; j < 3; j++ {
		fmt.Printf("Round %v\n", j + 1)
		for _, obs := range auctionObservers {
			auctionObservable.AcceptBid(obs.ID(), obs.OfferBid())
		}
		fmt.Printf("Accepted bid: ID %v %v\n", auctionObservable.GetBidderID(), auctionObservable.GetCurrentBid())
		auctionObservable.Notify()
		time.Sleep(1 * time.Second)
	}
}