package hn_test

import (
	"github.com/davecgh/go-spew/spew"
	"github.com/stretchr/testify/assert"
	"goat/pkg/hn"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestTopStories(t *testing.T) {

	t.Run("status 500", func(t *testing.T) {
		h := func(w http.ResponseWriter, r *http.Request) {
			assert.Equal(t, "/topstories.json", r.URL.Path)
			w.WriteHeader(500)
		}
		ts := httptest.NewServer(http.HandlerFunc(h))
		defer ts.Close()

		ids := hn.NewClientFor(ts.URL).Top()
		assert.Empty(t, ids)
	})

	t.Run("status 200", func(t *testing.T) {
		h := func(w http.ResponseWriter, r *http.Request) {
			assert.Equal(t, "/topstories.json", r.URL.Path)
			w.WriteHeader(200)
			_, _ = io.WriteString(w, "[ 19216077, 19216315 ]")
		}
		ts := httptest.NewServer(http.HandlerFunc(h))
		defer ts.Close()

		ids := hn.NewClientFor(ts.URL).Top()
		assert.Equal(t, 2, len(ids))
		assert.Equal(t, 19216077, ids[0])
		assert.Equal(t, 19216315, ids[1])
	})
}

func TestTopStoriesIntegration(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}
	ids := hn.NewClient().Top()
	assert.NotEmpty(t, ids)
	spew.Dump(ids)
}

func TestItems(t *testing.T) {

	type deps struct {
		server *httptest.Server
		client hn.Client
	}

	setup := func(handler http.HandlerFunc) *deps {
		server := httptest.NewServer(handler)
		client := hn.NewClientFor(server.URL)
		return &deps{
			server: server,
			client: client,
		}
	}

	t.Run("status 500", func(t *testing.T) {
		s := setup(func(w http.ResponseWriter, r *http.Request) {
			assert.Equal(t, "/item/123.json", r.URL.Path)
			w.WriteHeader(500)
		})
		defer s.server.Close()

		item := s.client.Get(123)
		assert.Empty(t, item)
	})

	t.Run("status 200", func(t *testing.T) {
		s := setup(func(w http.ResponseWriter, r *http.Request) {
			assert.Equal(t, "/item/456.json", r.URL.Path)
			w.WriteHeader(200)
			buf, _ := ioutil.ReadFile("testdata/item-19216428.json")
			_, _ = w.Write(buf)
		})
		defer s.server.Close()

		item := s.client.Get(456)
		assert.Equal(t, "Show HN: ICONSVG – Customize and Generate Common SVG Icons", item.Title)
	})
}

func TestItemIntegration(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}
	item := hn.NewClient().Get(19216428)
	assert.Equal(t, "Show HN: ICONSVG – Customize and Generate Common SVG Icons", item.Title)
	spew.Dump(item)
}
