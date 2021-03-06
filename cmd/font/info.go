package main

import (
	"fmt"
	"strconv"

	"github.com/propellerfactory/font/sfnt"
)

// Info prints the name table (contains metadata).
func Info(font *sfnt.Font) error {
	if font.HasTable(sfnt.TagName) {
		name, err := font.NameTable()
		if err != nil {
			return err
		}

		for _, entry := range name.List() {
			ids := " (" + strconv.Itoa(int(entry.PlatformID)) + "," + strconv.Itoa(int(entry.EncodingID)) + "," + strconv.Itoa(int(entry.LanguageID)) + "," + strconv.Itoa(int(entry.NameID)) + ") "
			fmt.Println(entry.Platform() + ids + entry.Label() + ": " + entry.String())
		}
	}
	return nil
}
