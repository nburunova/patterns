package auction_example

type Observer interface{
	ID() int
	Update()
}