package auction

import (
	"math/rand"
)

type AuctionObserver struct {
	id int
	lastBid int
	auction *Auction
}

func NewAuctionObserver(id int, auction *Auction) *AuctionObserver {
	return &AuctionObserver{
		id: id,
		auction: auction,
	}
}

func (c *AuctionObserver) ID() int {
	return c.id
}

func (c *AuctionObserver) Update() {
	c.lastBid = c.auction.GetCurrentBid()
}

func (c *AuctionObserver) OfferBid() int {
	offer := c.lastBid + rand.Intn(10)
	return offer
} 