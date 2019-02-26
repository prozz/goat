package hn

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const baseUrl = "https://hacker-news.firebaseio.com/v0"

type client struct {
	topUrl  func() string
	itemUrl func(id int) string
}

type Client interface {
	Top() []int
	Get(id int) Item
}

func NewClient() Client {
	return &client{
		topUrl: func() string {
			return baseUrl + "/topstories.json"
		},
		itemUrl: func(id int) string {
			return baseUrl + fmt.Sprintf("/item/%d.json", id)
		},
	}
}

func (c *client) Top() []int {
	response, err := http.Get(c.topUrl())
	if err != nil || response.StatusCode != 200 {
		return nil
	}

	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil
	}

	var ids []int
	err = json.Unmarshal(body, &ids)
	if err != nil {
		return nil
	}

	return ids
}

type Item struct {
	Title string `json:"title"`
}

func (c *client) Get(id int) Item {
	response, err := http.Get(c.itemUrl(id))
	if err != nil || response.StatusCode != 200 {
		return Item{}
	}

	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return Item{}
	}

	var item Item
	err = json.Unmarshal(body, &item)
	if err != nil {
		return Item{}
	}

	return item
}