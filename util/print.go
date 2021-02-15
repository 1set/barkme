package util

import (
	"strings"

	"github.com/olekukonko/tablewriter"
)

// RenderTableString renders the rows as table and returns as string for console.
func RenderTableString(header []string, rows [][]string) string {
	var emptyStr string
	s := strings.Builder{}
	table := tablewriter.NewWriter(&s)
	table.SetHeader(header)
	table.SetHeaderAlignment(tablewriter.ALIGN_LEFT)
	table.SetHeaderLine(false)
	table.SetTablePadding("\t")
	table.SetAlignment(tablewriter.ALIGN_LEFT)
	table.SetAutoWrapText(false)
	table.SetAutoFormatHeaders(true)
	table.SetBorders(tablewriter.Border{Left: true, Top: false, Right: true, Bottom: false})
	table.SetCenterSeparator(emptyStr)
	table.SetColumnSeparator(emptyStr)
	table.SetRowSeparator(emptyStr)

	for _, r := range rows {
		table.Append(r)
	}

	table.Render()
	return s.String()
}
