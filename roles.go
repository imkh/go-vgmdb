package vgmdb

import (
	"errors"
	"fmt"
	"strings"

	"github.com/gocolly/colly/v2"
)

// List a couple of standard errors.
var (
	ErrUnknown          = errors.New("An error has occurred!")
	ErrRoleUnauthorized = errors.New("Only registered members can view role information.")
	ErrRoleNotFound     = errors.New("Role not found!")
)

// RolesService handles scraping the roles related pages of the VGMdb.net website.
type RolesService struct {
	scraper *Scraper
}

// Role represents a role entry on VGMdb.
type Role struct {
	ID      int      `json:"id"`
	Name    string   `json:"name"`
	Aliases []string `json:"aliases"`
	Notes   string   `json:"notes"`
	URL     string   `json:"url"`
	Image   *Image   `json:"image"`
}

// Image represents a role's image.
type Image struct {
	ThumbURL string `json:"thumbUrl"`
	FullURL  string `json:"fullUrl"`
	Caption  string `json:"caption"`
}

// GetRole retrieves a role by its ID.
//
// Scraped page: https://vgmdb.net/role/<id>
func (s *RolesService) GetRole(id int) (*Role, error) {
	var vgmdbError error
	role := &Role{
		ID:  id,
		URL: fmt.Sprintf("%s/role/%d", baseURL, id),
	}

	// Check for error page
	s.scraper.collector.OnHTML(`table[cellpadding="0"][cellspacing="0"]:has(img[src="/db/img/banner-error.gif"])`, func(e *colly.HTMLElement) {
		tableText := strings.TrimSpace(e.Text)
		if strings.Contains(tableText, ErrRoleUnauthorized.Error()) {
			vgmdbError = ErrRoleUnauthorized
		} else if strings.Contains(tableText, ErrRoleNotFound.Error()) {
			vgmdbError = ErrRoleNotFound
		} else {
			vgmdbError = ErrUnknown
		}
	})

	// Handle the main content
	s.scraper.collector.OnHTML("#innermain", func(e *colly.HTMLElement) {
		role.Name = e.ChildText(fmt.Sprintf(`a[href="/role/%d?alias=0"]`, id))

		role.Aliases = e.ChildTexts(fmt.Sprintf(`#leftfloat a[href^="/role/%d?alias="]`, id))

		if thumbURL := e.ChildAttr(`#leftfloat img[src*="thumb-media.vgm.io"]`, "src"); thumbURL != "" {
			role.Image = &Image{
				ThumbURL: thumbURL,
				FullURL:  e.ChildAttr(`#leftfloat a[href*="media.vgm.io"]`, "href"),
				Caption:  e.ChildText("#leftfloat div.highslide-caption"),
			}
		}

		notes := e.ChildText(`#rightfloat > div[style="background-color: #2F364F;"] > div.smallfont`)
		if notes != "No notes available." {
			role.Notes = notes
		}
	})

	// Visit the page
	err := s.scraper.collector.Visit(role.URL)
	if err != nil {
		return nil, fmt.Errorf("unable to visit role page: %w", err)
	}

	return role, vgmdbError
}
