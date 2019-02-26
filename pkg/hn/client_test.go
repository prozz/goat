package hn

import (
	"github.com/davecgh/go-spew/spew"
	"github.com/stretchr/testify/assert"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestStories(t *testing.T) {

	type deps struct {
		server *httptest.Server
		client *client
	}

	setup := func(handler http.HandlerFunc) *deps {
		server := httptest.NewServer(handler)
		client := &client{
			topUrl: func() string {
				return server.URL
			},
		}
		return &deps{
			server: server,
			client: client,
		}
	}

	t.Run("status 500", func(t *testing.T) {
		s := setup(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(500)
		})
		defer s.server.Close()

		ids := s.client.Top()
		assert.Empty(t, ids)
	})

	t.Run("status 200", func(t *testing.T) {
		s := setup(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(200)
			io.WriteString(w, "[ 19216077, 19216315 ]")
		})
		defer s.server.Close()

		ids := s.client.Top()
		assert.Equal(t, 2, len(ids))
		assert.Equal(t, 19216077, ids[0])
		assert.Equal(t, 19216315, ids[1])
	})
}

func TestTopStoriesIntegration(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}
	ids := NewClient().Top()
	assert.NotEmpty(t, ids)
	spew.Dump(ids)
}

func TestItems(t *testing.T) {

	type deps struct {
		server *httptest.Server
		client *client
	}

	itemID := 0

	setup := func(handler http.HandlerFunc) *deps {
		server := httptest.NewServer(handler)
		client := &client{
			itemUrl: func(id int) string {
				itemID = id
				return server.URL
			},
		}
		return &deps{
			server: server,
			client: client,
		}
	}

	t.Run("status 500", func(t *testing.T) {
		s := setup(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(500)
		})
		defer s.server.Close()

		item := s.client.Get(123)
		assert.Empty(t, item)
		assert.Equal(t, 123, itemID)
	})

	t.Run("status 200", func(t *testing.T) {
		s := setup(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(200)
			buf, _ := ioutil.ReadFile("testdata/item-19216428.json")
			w.Write(buf)
		})
		defer s.server.Close()

		item := s.client.Get(456)
		assert.Equal(t, "Show HN: ICONSVG – Customize and Generate Common SVG Icons", item.Title)
		assert.Equal(t, 456, itemID)
	})
}

func TestItemIntegration(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}
	item := NewClient().Get(19216428)
	assert.Equal(t, "Show HN: ICONSVG – Customize and Generate Common SVG Icons", item.Title)
	spew.Dump(item)
}