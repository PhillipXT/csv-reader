package parser

import (
	"bytes"
	"fmt"
	"log"
	"strings"
)

type dataRow struct {
	columns     []string
	columnCount int
	isValid     bool
}

func (r dataRow) Print() {
	for i, col := range r.columns {
		output := strings.ReplaceAll(col, "\n", "\\n")
		fmt.Printf("Column %d: %s\n", i, output)
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

	output := bytes.ReplaceAll(data, []byte("\n"), []byte("\\n"))
	log.Printf("Parsing data: %s\n", string(output))

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
		} else if !(i == len(data)-1 && ch == byte('\n')) {
			content += string(ch)
		}
		i += 1
		if i >= len(data) {
			row.columnCount += 1
			row.columns = append(row.columns, content)
			break
		}
	}

	return row
}
