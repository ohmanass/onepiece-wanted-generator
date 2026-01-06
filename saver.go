package main

type Saver interface {
	Save(p Pirate) error
}