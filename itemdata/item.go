package itemdata

import (
	"io/ioutil"
	"log"
	"strings"
)

const ColumnSeperator = "\t"

type ItemLine struct {
	Fields []string
}

type ItemData struct {
	Items []ItemLine
}

func ReadItemData(filename string) ItemData {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	dataString := string(data)

	lines := strings.Split(dataString, "\n")
	itemLines := make([]ItemLine, 0)

	for _, v := range lines {
		if len(v) > 1 { // last line is empty but strangely has a character
			itemLine := ItemLine{Fields: strings.Split(v, ColumnSeperator)}
			itemLines = append(itemLines, itemLine)
		}
	}

	log.Printf("Read %d lines", len(itemLines))
	return ItemData{Items: itemLines}
}
