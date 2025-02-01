# go-vgmdb

A Go library for scraping VGMdb.net.

## Pages implemented

### Home

- [ ] `/` - Home page

### Resources

- [ ] `/album/<id>` - Album page
- [ ] `/artist/<id>` - Artist page
- [ ] `/org/<id>` - Organization page
- [ ] `/product/<id>` - Product page
- [ ] `/event/<id>` - Event page
- [ ] `/role/<id>` - Role page

### Browse

- [ ] `/db/albums.php` - Browse all albums
- [ ] `/db/artists.php` - Browse all artists
- [ ] `/db/org.php` - Browse all organizations
- [ ] `/db/product.php` - Browse all products
- [ ] `/db/events.php` - Browse all events
- [ ] `/db/role.php` - Browse all roles

### Search

- [ ] `/search?q=<query>` - Search for all resources
- [ ] `/search?q=<query>&type=album` - Search for albums
- [ ] `/search?q=<query>&type=artist` - Search for artists
- [ ] `/search?q=<query>&type=org` - Search for organizations
- [ ] `/search?q=<query>&type=product` - Search for products

### User Lists

- [ ] `/db/collection.php?do=view&userid=<id>` - Collection
- [ ] `/db/marketplace.php?do=saleview&userid=<id>` - Sale List
- [ ] `/db/marketplace.php?do=wishview&userid=<id>` - Wish List
- [ ] `/db/user.php?do=submissions&id=<id>` - Submissions
- [ ] `/db/ratings.php?do=view&userid=<id>` - Ratings
- [ ] `/db/draft.php?do=view&userid=<id>` - Drafts

### Recent Updates

- [ ] `/db/recent.php?do=view_albums` - View recent album updates
- [ ] `/db/recent.php?do=view_media` - View recent media updates
- [ ] `/db/recent.php?do=view_tracklists` - View recent tracklist updates
- [ ] `/db/recent.php?do=view_scans` - View recent scan updates
- [ ] `/db/recent.php?do=view_artists` - View recent artist updates
- [ ] `/db/recent.php?do=view_credits` - View recent credits updates
- [ ] `/db/recent.php?do=view_drafts` - View recent draft updates
- [ ] `/db/recent.php?do=view_products` - View recent product updates
- [ ] `/db/recent.php?do=view_labels` - View recent organization updates
- [ ] `/db/recent.php?do=view_links` - View recent link updates
- [ ] `/db/recent.php?do=view_ratings` - View recent rating updates

### Other

- [ ] `/db/calendar.php?type=<resource>&year=2025&month=2` - Calendar
- [ ] `/db/marketplace.php?do=view` - Marketplace
- [ ] `/db/modq.php` - Album Moderation Queue
- [ ] `/db/modq.php?do=mod_artists` - Artist Moderation Queue
- [ ] `/db/statistics.php` - Site Statistics

## Installation

Inside your project directory, run:

```console
$ go get github.com/imkh/go-vgmdb
```

or import the module and run `go get` without parameters.

```go
import "github.com/imkh/go-vgmdb"
```

## Usage

```go
package main

import (
	"github.com/imkh/go-vgmdb"
)

func main() {
	scraper, err := vgmdb.NewScraper(
		//vgmdb.WithUserAgent("custom-user-agent")
	)
	if err != nil {
		log.Fatalf("Failed to create scraper: %v", err)
	}
}
```
