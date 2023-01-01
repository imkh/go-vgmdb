package vgmdb

import (
	"fmt"
	"net/http"
)

// ProductsService handles communication with the products related methods
// of the VGMdb API.
type ProductsService struct {
	client *Client
}

// Product represents a VGMdb product.
//
// VGMdb API schema: https://github.com/hufman/vgmdb/blob/80a491cb2eae0dd8da2c9a81de4777a812e1bf10/schema/product.json#L21
type Product struct {
	Link          string            `json:"link"`
	Type          ProductType       `json:"type"`
	Name          string            `json:"name"`
	NameReal      string            `json:"name_real,omitempty"`
	Description   string            `json:"description"`
	ReleaseDate   string            `json:"release_date,omitempty"`
	Meta          *Meta             `json:"meta"`
	Albums        []*ProductAlbum   `json:"albums"`
	Franchises    []*NamedItem      `json:"franchises,omitempty"`
	Superproduct  *NamedItem        `json:"superproduct,omitempty"`
	Superproducts [][]*ProductTitle `json:"superproducts,omitempty"`
	Titles        []*ProductTitle   `json:"titles,omitempty"`
	Organizations []*NamedItem      `json:"organizations,omitempty"`
	VgmdbLink     string            `json:"vgmdb_link"`
	Websites      *ItemWebsites     `json:"websites"`
}

// ProductTitle represents a sub product in a franchise.
//
// VGMdb API schema: https://github.com/hufman/vgmdb/blob/80a491cb2eae0dd8da2c9a81de4777a812e1bf10/schema/product.json#L8
type ProductTitle struct {
	Date  string      `json:"date,omitempty"`
	Link  string      `json:"link,omitempty"`
	Names Names       `json:"names"`
	Type  ProductType `json:"type,omitempty"`
}

// GetProduct gets a specific product, identified by product ID.
//
// VGMdb API docs: https://github.com/hufman/vgmdb/blob/80a491cb2eae0dd8da2c9a81de4777a812e1bf10/raml/api.raml#L171
func (s *ProductsService) GetProduct(id int) (*Product, *http.Response, error) {
	u := fmt.Sprintf("product/%d", id)

	req, err := s.client.NewRequest(http.MethodGet, u)
	if err != nil {
		return nil, nil, err
	}

	a := new(Product)
	resp, err := s.client.Do(req, a)
	if err != nil {
		return nil, resp, err
	}

	return a, resp, err
}
