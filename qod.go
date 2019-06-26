package qod

import (
	"encoding/json"
	"net/url"
)

type QODResponse struct {
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

type QODCategoriesResponse struct {
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

type QODCategoryResponse struct {
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

// GetQuoteOfTheDay requests the quote of the
// day and returns it in a struct
func (c *Client) GetQuoteOfTheDay() (*QODResponse, error) {

	res, err := c.Get(uri+"qod", nil)
	if err != nil {
		return nil, err
	}

	var r QODResponse
	dec := json.NewDecoder(res.Body)
	err = dec.Decode(&r)
	if err != nil {
		return nil, err
	}

	return &r, nil
}

func (c *Client) GetQODCategories() (*QODCategoriesResponse, error) {

	res, err := c.Get(uri+"qod/categories", nil)
	if err != nil {
		return nil, err
	}
	var r QODCategoriesResponse
	dec := json.NewDecoder(res.Body)
	err = dec.Decode(&r)
	if err != nil {
		return nil, err
	}

	return &r, nil
}

func (c *Client) GetQODByCategory(category string) (*QODCategoryResponse, error) {

	u, err := url.Parse(uri + "qod")
	if err != nil {
		return nil, err
	}

	q := u.Query()
	q.Set("category", category)
	u.RawQuery = q.Encode()

	res, err := c.Get(u.String(), nil)
	if err != nil {
		return nil, err
	}

	var r QODCategoryResponse
	dec := json.NewDecoder(res.Body)
	err = dec.Decode(&r)
	if err != nil {
		return nil, err
	}

	return &r, nil
}
