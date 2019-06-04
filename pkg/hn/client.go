package hn

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Client interface {
	Top() []int
	Get(id int) Item
}

func NewClient() Client {
	return NewClientFor("https://hacker-news.firebaseio.com/v0")
}

func NewClientFor(baseUrl string) Client {
	return &client{baseUrl: baseUrl}
}

type client struct {
	baseUrl string
}

func (c *client) Top() []int {
	response, err := http.Get(fmt.Sprintf("%s/topstories.json", c.baseUrl))
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
	response, err := http.Get(fmt.Sprintf("%s/item/%d.json", c.baseUrl, id))
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
