package vgmdb

import (
	"errors"
	"fmt"
	"net/url"
	"regexp"
	"strconv"
	"strings"
	"time"

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
	ID      int          `json:"id"`
	Name    string       `json:"name"`
	Aliases []*RoleAlias `json:"aliases"`
	Notes   *string      `json:"notes"`
	Image   *Image       `json:"image"`
	URL     string       `json:"url"`
}

// RoleAlias represents a role's alias.
type RoleAlias struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	URL  string `json:"url"`
}

// Image represents a role's image.
type Image struct {
	ThumbURL    string     `json:"thumb_url"`
	FullURL     string     `json:"full_url"`
	SubmittedBy *string    `json:"submitted_by"`
	SubmittedAt *time.Time `json:"submitted_at"`
}

// GetRole retrieves a role by its ID.
//
// Scraped page: https://vgmdb.net/role/{id}
func (s *RolesService) GetRole(id int) (*Role, error) {
	urlPath := fmt.Sprintf("/role/%d", id)

	var vgmdbError error
	role := new(Role)

	role.ID = id
	role.Aliases = []*RoleAlias{}
	role.URL, _ = url.JoinPath(baseURL, urlPath)

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
		// Parse the role's name
		role.Name = e.ChildText(fmt.Sprintf(`a[href="%s?alias=0"]`, urlPath))

		// Parse the role's aliases
		e.ForEach(fmt.Sprintf(`#leftfloat a[href^="%s?alias="]`, urlPath), func(_ int, e *colly.HTMLElement) {
			roleAlias := &RoleAlias{
				Name: e.Text,
			}
			if aliasURL, err := url.Parse(baseURL + e.Attr("href")); err == nil {
				roleAlias.URL = aliasURL.String()
				if aliasID, err := strconv.Atoi(aliasURL.Query().Get("alias")); err == nil {
					roleAlias.ID = aliasID
				}
			}
			role.Aliases = append(role.Aliases, roleAlias)
		})

		// Parse the role's image
		if thumbURL := e.ChildAttr(`#leftfloat img[src^="https://thumb-media.vgm.io"]`, "src"); thumbURL != "" {
			role.Image = &Image{
				ThumbURL: thumbURL,
				FullURL:  e.ChildAttr(`#leftfloat a[href^="https://media.vgm.io"]`, "href"),
			}
			if caption := e.ChildText("#leftfloat div.highslide-caption"); caption != "" {
				// Extract submission author & timestamp from caption
				if matches := regexp.MustCompile(`Submitted by (.+) on (.+)`).FindStringSubmatch(caption); len(matches) == 3 {
					// Parse the submission author
					submittedBy := matches[1]
					role.Image.SubmittedBy = &submittedBy

					// Parse the submission timestamp
					if submittedAt, err := time.Parse("Jan 2, 2006 03:04 PM", matches[2]); err == nil {
						role.Image.SubmittedAt = &submittedAt
					}
				}
			}
		}

		// Parse the role's notes
		notes := e.ChildText(`#rightfloat > div[style="background-color: #2F364F;"] > div.smallfont`)
		if notes != "No notes available." {
			role.Notes = &notes
		}
	})

	// Visit the page
	err := s.scraper.collector.Visit(role.URL)
	if err != nil {
		return nil, fmt.Errorf("unable to visit role page: %w", err)
	}

	if vgmdbError != nil {
		return nil, vgmdbError
	}
	return role, nil
}
