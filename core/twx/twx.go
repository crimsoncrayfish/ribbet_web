package twx

import (
	"io"
	"text/tabwriter"
)

func New(destination io.Writer) *tabwriter.Writer {
	return tabwriter.NewWriter(destination, 0, 0, 3, ' ', 0)
}
