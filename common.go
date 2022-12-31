package vgmdb

// AlbumType represents the different types an album can be.
//
// VGMdb API schema: https://github.com/hufman/vgmdb/blob/9f00fcd437a9392fd3b8891b6d0a8cef9e0932ea/schema/common.json#L12
type AlbumType string

// List of available visibility levels.
//
// VGMdb API schema: https://github.com/hufman/vgmdb/blob/9f00fcd437a9392fd3b8891b6d0a8cef9e0932ea/schema/common.json#L12
const (
	GameAlbumType    AlbumType = "game"
	AnimeAlbumType   AlbumType = "anime"
	PrintAlbumType   AlbumType = "print"
	DramaAlbumType   AlbumType = "drama"
	DemoAlbumType    AlbumType = "demo"
	WorksAlbumType   AlbumType = "works"
	BonusAlbumType   AlbumType = "bonus"
	DoujinAlbumType  AlbumType = "doujin"
	CancelAlbumType  AlbumType = "cancel"
	BootlegAlbumType AlbumType = "bootleg"
	MultAlbumType    AlbumType = "mult"
	LiveAlbumType    AlbumType = "live"
	TokuAlbumType    AlbumType = "toku"
)

// NamedItem represents a data item that has names and might have a link.
//
// VGMdb API schema: https://github.com/hufman/vgmdb/blob/9f00fcd437a9392fd3b8891b6d0a8cef9e0932ea/schema/common.json#L56
type NamedItem struct {
	Link  string            `json:"link,omitempty"`
	Names map[string]string `json:"names"`
}

// LinkedAlbum represents a link object to an album.
//
// VGMdb API schema: https://github.com/hufman/vgmdb/blob/9f00fcd437a9392fd3b8891b6d0a8cef9e0932ea/schema/common.json#L70
type LinkedAlbum struct {
	NamedItem

	Catalog string    `json:"catalog"`
	Type    AlbumType `json:"type"`
}

// ReprintedAlbum represents a link object to a reprint of an album.
//
// VGMdb API schema: https://github.com/hufman/vgmdb/blob/9f00fcd437a9392fd3b8891b6d0a8cef9e0932ea/schema/common.json#L129
type ReprintedAlbum struct {
	Catalog string `json:"catalog"`
	Link    string `json:"link"`
	Note    string `json:"note"`
}

// Meta represents basic information about the information, like last updated time.
//
// VGMdb API schema: https://github.com/hufman/vgmdb/blob/9f00fcd437a9392fd3b8891b6d0a8cef9e0932ea/schema/common.json#L162
type Meta struct {
	AddedDate   string `json:"added_date"`
	EditedDate  string `json:"edited_date"`
	FetchedDate string `json:"fetched_date,omitempty"`
	Ttl         int    `json:"ttl"`
	Visitors    int    `json:"visitors"`
}

// AlbumArt represents various album art pictures.
//
// VGMdb API schema: https://github.com/hufman/vgmdb/blob/9f00fcd437a9392fd3b8891b6d0a8cef9e0932ea/schema/common.json#L174
type AlbumArt struct {
	Full   string `json:"full"`
	Medium string `json:"medium"`
	Name   string `json:"name"`
	Thumb  string `json:"thumb"`
}

// ExternalWebsite
//
// VGMdb API schema: https://github.com/hufman/vgmdb/blob/9f00fcd437a9392fd3b8891b6d0a8cef9e0932ea/schema/common.json#L186
type ExternalWebsite struct {
	Name string `json:"name"`
	Link string `json:"link"`
}

// ItemWebsites represents a collection of external websites about the item.
//
// VGMdb API schema: https://github.com/hufman/vgmdb/blob/9f00fcd437a9392fd3b8891b6d0a8cef9e0932ea/schema/common.json#L195
type ItemWebsites map[string][]ExternalWebsite
