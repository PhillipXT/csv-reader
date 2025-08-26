package parser

import (
	"fmt"
	"log"
)

type dataRow struct {
	columns     []string
	columnCount int
	isValid     bool
}

func (r dataRow) Print() {
	for i, col := range r.columns {
		fmt.Printf("Column %d: %s\n", i, col)
	}
}

func parseColumns(data []byte) dataRow {
	row := dataRow{
		columnCount: 0,
		columns:     []string{},
	}

	quoteDelim := byte('"')
	colDelim := byte(',')

	content := ""
	quotes := false
	i := 0

	log.Println("Parsing data:", string(data))

	for {
		ch := data[i]
		if ch == quoteDelim {
			isDelim := true
			if quotes == true && i < len(data)-1 && data[i+1] == quoteDelim {
				isDelim = false
				i += 1
				content += string(quoteDelim)
			}
			if isDelim {
				if !quotes {
					quotes = true
				} else {
					quotes = false
				}
			}
		} else if ch == colDelim {
			if quotes == true {
				content += string(ch)
			} else {
				row.columnCount += 1
				row.columns = append(row.columns, content)
				content = ""
			}
		} else {
			content += string(ch)
		}
		i += 1
		if i >= len(data) {
			if content != "" {
				row.columnCount += 1
				row.columns = append(row.columns, content)
			}
			break
		}
	}

	return row
}
