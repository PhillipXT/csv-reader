package loader

import (
	"fmt"
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

func parseLine(line []byte) dataRow {
	row := dataRow{
		columnCount: 0,
		columns:     []string{},
	}

	quoteDelim := byte('"')
	colDelim := byte(',')

	data := ""
	quotes := false
	i := 0

	for {
		ch := line[i]
		if ch == quoteDelim {
			isDelim := true
			if i < len(line) && line[i+1] == quoteDelim {
				isDelim = false
				i += 1
			}
			if isDelim {
				if !quotes {
					quotes = true
				} else {
					quotes = false
					row.columnCount += 1
					row.columns = append(row.columns, data)
					data = ""
				}
			}
			data += string(quoteDelim)
		} else if ch == colDelim {
			if quotes == true {
				data += string(ch)
			} else {
				row.columnCount += 1
				row.columns = append(row.columns, data)
				data = ""
			}
		} else {
			data += string(ch)
		}
		i += 1
		if i >= len(line) {
			row.columnCount += 1
			row.columns = append(row.columns, data)
			break
		}
	}

	return row
}
