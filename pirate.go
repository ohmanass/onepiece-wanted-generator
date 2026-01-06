package main

import (
	"errors"
	"strings"
)

type Pirate struct {
	Name  string
	Prime string
	Img   string
}

func New(name, prime, img string) (*Pirate, error) {
	
	// Gestion des erreures
	if name == "" {
		return nil, errors.New("le nom du pirate est vide")
	}

	if name != strings.ToUpper(name) {
		return nil, errors.New("le nom du pirate doit être en MAJUSCULES")
	}

	if prime == "" {
		return nil, errors.New("la prime du pirate est vide")
	}

	if img == "" {
		return nil, errors.New("le chemin de l image est vide")
	}

	// Creation de notre pirate
	p := &Pirate{
		Name:  name,
		Prime: prime,
		Img:   img,
	}
	return p, nil
}