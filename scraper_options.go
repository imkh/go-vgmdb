package vgmdb

import (
	"net/http"

	"github.com/gocolly/colly/v2"
)

// ScraperOptionFunc can be used to customize a new VGMdb scraper.
type ScraperOptionFunc func(*Scraper) error

// WithUserAgent can be used to configure a custom User-Agent.
func WithUserAgent(userAgent string) ScraperOptionFunc {
	return func(s *Scraper) error {
		s.collector.UserAgent = userAgent
		return nil
	}
}

// WithHTTPTransport can be used to configure custom HTTP options.
func WithHTTPTransport(transport *http.Transport) ScraperOptionFunc {
	return func(s *Scraper) error {
		s.collector.WithTransport(transport)
		return nil
	}
}

// WithCookie can be used to configure a custom cookie to send with each request.
func WithCookie(cookie string) ScraperOptionFunc {
	return func(s *Scraper) error {
		s.collector.OnRequest(func(r *colly.Request) {
			r.Headers.Set("Cookie", cookie)
		})
		return nil
	}
}
