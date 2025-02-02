// Package vgmdb implements a scraper for the VGMdb.net website.
package vgmdb

import (
	"log"

	"github.com/gocolly/colly/v2"
)

const (
	// Base URL of the VGMdb.net website. Make sure to NOT end with a slash.
	baseURL = "https://vgmdb.net"

	// Default user agent to use for the scraper.
	userAgent = "go-vgmdb"
)

// Scraper represents a scraper for the VGMdb.net website.
type Scraper struct {
	// Collector is Colly's main entity, it provides the scraper instance for a scraping job.
	collector *colly.Collector

	// Services used for scraping different parts of the VGMdb.net website.
	Roles *RolesService
}

// NewScraper returns a new VGMdb.net scraper.
func NewScraper(options ...ScraperOptionFunc) (*Scraper, error) {
	s := &Scraper{
		collector: colly.NewCollector(
			colly.UserAgent(userAgent),
		),
	}

	// Apply any given scraper options.
	for _, fn := range options {
		if fn == nil {
			continue
		}
		if err := fn(s); err != nil {
			return nil, err
		}
	}

	// Set error handler
	s.collector.OnError(func(r *colly.Response, err error) {
		log.Println("Request URL:", r.Request.URL, "failed with response:", r, "\nError:", err)
	})

	// Create all the public services.
	s.Roles = &RolesService{scraper: s}

	return s, nil
}
