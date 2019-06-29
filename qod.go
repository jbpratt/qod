package qod

import (
	"encoding/json"
	"net/url"
)

// Response returns a slice of quotes
type Response struct {
	Success struct {
		Total int `json:"total"`
	} `json:"success"`
	Contents struct {
		Quotes []struct {
			Quote      string   `json:"quote"`
			Length     string   `json:"length"`
			Author     string   `json:"author"`
			Tags       []string `json:"tags"`
			Category   string   `json:"category"`
			Date       string   `json:"date"`
			Title      string   `json:"title"`
			Background string   `json:"background"`
			ID         string   `json:"id"`
		} `json:"quotes"`
	} `json:"contents"`
}

// CategoriesResponse contains all of the QOD categories
type CategoriesResponse struct {
	Success struct {
		Total int `json:"total"`
	} `json:"success"`
	Contents struct {
		Categories struct {
			Inspire    string `json:"inspire"`
			Management string `json:"management"`
			Sports     string `json:"sports"`
			Life       string `json:"life"`
			Funny      string `json:"funny"`
			Love       string `json:"love"`
			Art        string `json:"art"`
			Students   string `json:"students"`
		} `json:"categories"`
		Copyright string `json:"copyright"`
	} `json:"contents"`
}

// CategoryResponse is a slice of quotes by category
type CategoryResponse struct {
	Success struct {
		Total int `json:"total"`
	} `json:"success"`
	Contents struct {
		Quotes []struct {
			Quote    string      `json:"quote"`
			Author   string      `json:"author"`
			Length   string      `json:"length"`
			Tags     []string    `json:"tags"`
			Category string      `json:"category"`
			Title    string      `json:"title"`
			Date     string      `json:"date"`
			ID       interface{} `json:"id"`
		} `json:"quotes"`
		Copyright string `json:"copyright"`
	} `json:"contents"`
}

// GetQuoteOfTheDay requests the quote of the day
func (c *Client) GetQuoteOfTheDay() (*Response, error) {

	res, err := c.get(uri+"qod", nil)
	if err != nil {
		return nil, err
	}

	var r Response
	dec := json.NewDecoder(res.Body)
	err = dec.Decode(&r)
	if err != nil {
		return nil, err
	}

	return &r, nil
}

// GetCategories requests all of the QOD categories
func (c *Client) GetCategories() (*CategoriesResponse, error) {

	res, err := c.get(uri+"qod/categories", nil)
	if err != nil {
		return nil, err
	}
	var r CategoriesResponse
	dec := json.NewDecoder(res.Body)
	err = dec.Decode(&r)
	if err != nil {
		return nil, err
	}

	return &r, nil
}

// GetByCategory requests the QOD by category
func (c *Client) GetByCategory(category string) (*CategoryResponse, error) {

	u, err := url.Parse(uri + "qod")
	if err != nil {
		return nil, err
	}

	q := u.Query()
	q.Set("category", category)
	u.RawQuery = q.Encode()

	res, err := c.get(u.String(), nil)
	if err != nil {
		return nil, err
	}

	var r CategoryResponse
	dec := json.NewDecoder(res.Body)
	err = dec.Decode(&r)
	if err != nil {
		return nil, err
	}

	return &r, nil
}
