package main

import (
	"os"

	"github.com/propellerfactory/font/sfnt"
)

// Convert to otf
func Convert(font *sfnt.Font) error {
	_, err := font.WriteOTF(os.Stdout)
	return err
}
