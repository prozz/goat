package ai

import (
	"goat/pkg/hn"
	"io"
)

type Dumper struct {
	client hn.Client
}

func NewDumper(client hn.Client) *Dumper {
	return &Dumper{
		client: client,
	}
}

func (d *Dumper) Dump(w io.Writer) error {
	ids := d.client.Top()
	for _, id := range ids {
		item := d.client.Get(id)
		_, err := w.Write([]byte(item.Title + "\n"))
		if err != nil {
			return err
		}
	}
	return nil
}
