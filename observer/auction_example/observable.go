package auction_example

type Observable interface{
	Add(o Observer)
	Remove(o Observer)
	Notify()
}