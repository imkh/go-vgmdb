package vgmdb

import (
	"fmt"
	"net/http"
)

// AlbumsService handles communication with the albums related methods
// of the VGMdb API.
type AlbumsService struct {
	client *Client
}

// Album represents a VGMdb album.
//
// VGMdb API schema: https://github.com/hufman/vgmdb/blob/80a491cb2eae0dd8da2c9a81de4777a812e1bf10/schema/album.json#L52
type Album struct {
	Arrangers      []*NamedItem      `json:"arrangers"`
	Barcode        string            `json:"barcode"`
	Catalog        string            `json:"catalog"`
	Categories     []string          `json:"categories,omitempty"`
	Category       string            `json:"category,omitempty"`
	Classification string            `json:"classification"`
	Composers      []*NamedItem      `json:"composers"`
	Covers         []*AlbumArt       `json:"covers"`
	Discs          []*Disc           `json:"discs"`
	Distributor    *NamedItem        `json:"distributor,omitempty"`
	Link           string            `json:"link"`
	Lyricists      []*NamedItem      `json:"lyricists"`
	MediaFormat    string            `json:"media_format"`
	Meta           *AlbumMeta        `json:"meta"`
	Name           string            `json:"name"`
	Names          map[string]string `json:"names"`
	Notes          string            `json:"notes"`
	Organizations  []*NamedItem      `json:"organizations,omitempty"`
	Performers     []*NamedItem      `json:"performers"`
	PictureFull    string            `json:"picture_full"`
	PictureSmall   string            `json:"picture_small"`
	PictureThumb   string            `json:"picture_thumb"`
	Platforms      []string          `json:"platforms,omitempty"`
	Products       []*NamedItem      `json:"products,omitempty"`
	PublishFormat  string            `json:"publish_format"`
	Publisher      *NamedItem        `json:"publisher,omitempty"`
	Rating         float64           `json:"rating,omitempty"`
	Related        []*LinkedAlbum    `json:"related,omitempty"`
	ReleaseDate    string            `json:"release_date,omitempty"`
	ReleasePrice   *ReleasePrice     `json:"release_price,omitempty"`
	Reprints       []*ReprintedAlbum `json:"reprints,omitempty"`
	Stores         []*Stores         `json:"stores,omitempty"`
	VgmdbLink      string            `json:"vgmdb_link"`
	Votes          int               `json:"votes"`
	Websites       ItemWebsites      `json:"websites,omitempty"`
}

// Track represents a track, with some translated names and a track length.
//
// VGMdb API schema: https://github.com/hufman/vgmdb/blob/80a491cb2eae0dd8da2c9a81de4777a812e1bf10/schema/album.json#L14
type Track struct {
	Names       map[string]string `json:"names"`
	TrackLength string            `json:"track_length"`
}

// Disc represents information about an audio disc.
//
// VGMdb API schema: https://github.com/hufman/vgmdb/blob/80a491cb2eae0dd8da2c9a81de4777a812e1bf10/schema/album.json#L24
type Disc struct {
	DiscLength string   `json:"disc_length"`
	Name       string   `json:"name"`
	Tracks     []*Track `json:"tracks"`
}

// VGMdb API schema: https://github.com/hufman/vgmdb/blob/80a491cb2eae0dd8da2c9a81de4777a812e1bf10/schema/album.json#L38
type AlbumMeta struct {
	Meta

	Freedb *int `json:"freedb,omitempty"`
}

// VGMdb API schema: https://github.com/hufman/vgmdb/blob/80a491cb2eae0dd8da2c9a81de4777a812e1bf10/schema/album.json#L84
type ReleasePrice struct {
	Currency string      `json:"currency,omitempty"`
	Price    interface{} `json:"price"`
}

// VGMdb API schema: https://github.com/hufman/vgmdb/blob/80a491cb2eae0dd8da2c9a81de4777a812e1bf10/schema/album.json#L99
type Stores struct {
	Link string `json:"link"`
	Name string `json:"name"`
}

// GetAlbum gets a specific album, identified by album ID.
//
// VGMdb API docs: https://github.com/hufman/vgmdb/blob/80a491cb2eae0dd8da2c9a81de4777a812e1bf10/raml/api.raml#L140
func (s *AlbumsService) GetAlbum(id int) (*Album, *http.Response, error) {
	u := fmt.Sprintf("album/%d", id)

	req, err := s.client.NewRequest(http.MethodGet, u)
	if err != nil {
		return nil, nil, err
	}

	a := new(Album)
	resp, err := s.client.Do(req, a)
	if err != nil {
		return nil, resp, err
	}

	return a, resp, err
}
