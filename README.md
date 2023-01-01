# go-vgmdb

A Go library for accessing the VGMdb API (https://vgmdb.info/). Inspired by [go-gitlab](https://github.com/xanzy/go-gitlab).

## Coverage

### Searching

- [ ] `/search/<query>` - Search for matching items
- [ ] `/search?q=<query>` - Search for matching items
- [ ] `/search/albums/<query>` - Search for matching albums
- [ ] `/search/albums?q=<query>` - Search for matching albums
- [ ] `/search/artists/<query>` - Search for matching artists
- [ ] `/search/artists?q=<query>` - Search for matching artists
- [ ] `/search/orgs/<query>` - Search for matching organizations
- [ ] `/search/orgs?q=<query>` - Search for matching organizations
- [ ] `/search/products/<query>` - Search for matching products
- [ ] `/search/products?q=<query>` - Search for matching products

### Browse Lists

- [ ] `/albumlist/<letter>` - View all of the albums by letter
- [ ] `/artistlist/<letter>` - View all of the artists by letter
- [ ] `/orglist` - View all of the organizations
- [ ] `/orglist/<letter>` - View all of the organizations by letter
- [ ] `/eventlist` - View all of the events
- [ ] `/eventlist/<year>` - View all of the events by year
- [ ] `/productlist/<letter>` - View all of the products by letter

### Recent Changes

- [ ] `/recent/albums` - View recent album changes
- [ ] `/recent/media` - View recent album media changes
- [ ] `/recent/tracklists` - View recent album tracklist changes
- [ ] `/recent/scans` - View recent album scanned covers
- [ ] `/recent/artists` - View recent artist changes
- [ ] `/recent/products` - View recent product changes
- [ ] `/recent/labels` - View recent label organization changes
- [ ] `/recent/links` - View recent album and artist links changes
- [ ] `/recent/ratings` - View recent album rating changes

### Information Pages

- [x] `/album/<id>` - Album information
- [ ] `/artist/<id>` - Artist information
- [ ] `/org/<id>` - Organization information
- [ ] `/event/<id>` - Event information
- [x] `/product/<id>` - Product information

### Seller Information

- [ ] `/album/<id>/sellers` - Album sellers
- [ ] `/artist/<id>/sellers` - Artist sellers
- [ ] `/album/<id>/sellers?allow_partial=true` - Partial album sellers with Refresh header
- [ ] `/artist/<id>/sellers?allow_partial=true` - Partial artist sellers with Refresh header

## Installation

Inside your project directory, run:

```console
$ go get github.com/imkg/go-vgmdb
```

or import the module and run `go get` without parameters.

```go
import "github.com/imkg/go-vgmdb"
```

## Usage

```go
package main

import (
	"github.com/imkg/go-vgmdb"
)

func main() {
	client, err := vgmdb.NewClient()
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	// Get an album
	album, _, err := client.Albums.GetAlbum(75832)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(album.Name)
}
```