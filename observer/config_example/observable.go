package config_example

type Observable interface {
	Add(o Observer)
	Remove(o Observer)
	Notify()
	UpdateCAFile(newFile string) 
}