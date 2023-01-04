// Package vgmdb implements a VGMdb API client.
package vgmdb

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

const (
	defaultBaseURL = "https://vgmdb.info/"
	userAgent      = "go-vgmdb"
)

// A Client manages communication with the VGMdb API.
type Client struct {
	// HTTP client used to communicate with the API.
	client *http.Client

	// Base URL for API requests. Defaults to the public VGMdb API, but can be
	// set to a domain endpoint to use with a self hosted VGMdb server. baseURL
	// should always be specified with a trailing slash.
	baseURL *url.URL

	// User agent used when communicating with the VGMdb API.
	UserAgent string

	// Services used for talking to different parts of the VGMdb API.
	Albums   *AlbumsService
	Products *ProductsService
}

// NewClient returns a new VGMdb API client.
func NewClient(options ...ClientOptionFunc) (*Client, error) {
	client, err := newClient(options...)
	if err != nil {
		return nil, err
	}
	return client, nil
}

func newClient(options ...ClientOptionFunc) (*Client, error) {
	c := &Client{UserAgent: userAgent}

	// Configure the HTTP client.
	c.client = http.DefaultClient

	// Set the default base URL.
	c.setBaseURL(defaultBaseURL)

	// Apply any given client options.
	for _, fn := range options {
		if fn == nil {
			continue
		}
		if err := fn(c); err != nil {
			return nil, err
		}
	}

	// Create all the public services.
	c.Albums = &AlbumsService{client: c}
	c.Products = &ProductsService{client: c}

	return c, nil
}

// BaseURL return a copy of the baseURL.
func (c *Client) BaseURL() *url.URL {
	u := *c.baseURL
	return &u
}

// setBaseURL sets the base URL for API requests to a custom endpoint.
func (c *Client) setBaseURL(urlStr string) error {
	// Make sure the given URL end with a slash
	if !strings.HasSuffix(urlStr, "/") {
		urlStr += "/"
	}

	baseURL, err := url.Parse(urlStr)
	if err != nil {
		return err
	}

	// Update the base URL of the client.
	c.baseURL = baseURL

	return nil
}

// NewRequest creates a new API request. The method expects a relative URL
// path that will be resolved relative to the base URL of the Client.
// Relative URL paths should always be specified without a preceding slash.
// If specified, the value pointed to by body is JSON encoded and included
// as the request body.
func (c *Client) NewRequest(method, path string) (*http.Request, error) {
	u := *c.baseURL
	unescaped, err := url.PathUnescape(path)
	if err != nil {
		return nil, err
	}

	// Set the encoded path data
	u.RawPath = c.baseURL.Path + path
	u.Path = c.baseURL.Path + unescaped

	// Set format query string
	q := u.Query()
	q.Add("format", "json")
	u.RawQuery = q.Encode()

	// Create a request specific headers map.
	reqHeaders := make(http.Header)
	reqHeaders.Set("Accept", "application/json")

	if c.UserAgent != "" {
		reqHeaders.Set("User-Agent", c.UserAgent)
	}

	req, err := http.NewRequest(method, u.String(), nil)
	if err != nil {
		return nil, err
	}

	// Set the request specific headers.
	for k, v := range reqHeaders {
		req.Header[k] = v
	}

	return req, nil
}

// Do sends an API request and returns the API response. The API response is
// JSON decoded and stored in the value pointed to by v, or returned as an
// error if an API error has occurred. If v implements the io.Writer
// interface, the raw response body will be written to v, without attempting to
// first decode it.
func (c *Client) Do(req *http.Request, v interface{}) (*http.Response, error) {
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	err = CheckResponse(resp)
	if err != nil {
		// Even though there was an error, we still return the response
		// in case the caller wants to inspect it further.
		return resp, err
	}

	if v != nil {
		if w, ok := v.(io.Writer); ok {
			_, err = io.Copy(w, resp.Body)
		} else {
			err = json.NewDecoder(resp.Body).Decode(v)
		}
	}

	return resp, err
}

// An ErrorResponse reports one or more errors caused by an API request.
type ErrorResponse struct {
	Body     []byte
	Response *http.Response
	Message  string
}

func (e *ErrorResponse) Error() string {
	path, _ := url.QueryUnescape(e.Response.Request.URL.Path)
	u := fmt.Sprintf("%s://%s%s", e.Response.Request.URL.Scheme, e.Response.Request.URL.Host, path)
	return fmt.Sprintf("%s %s: %d %s", e.Response.Request.Method, u, e.Response.StatusCode, e.Message)
}

// CheckResponse checks the API response for errors, and returns them if present.
func CheckResponse(r *http.Response) error {
	switch r.StatusCode {
	case 200, 201, 202, 204, 304:
		return nil
	}

	errorResponse := &ErrorResponse{Response: r}
	data, err := io.ReadAll(r.Body)
	if err == nil && data != nil {
		errorResponse.Body = data
		errorResponse.Message = parseError(string(data))
	}

	return errorResponse
}

// Format:
//
// <!DOCTYPE HTML PUBLIC "-//IETF//DTD HTML 2.0//EN">
// <html>
//
//	<head>
//	    <title>Error: 404 Not Found</title>
//	    <style type="text/css">
//	      html {background-color: #eee; font-family: sans;}
//	      body {background-color: #fff; border: 1px solid #ddd;
//	            padding: 15px; margin: 15px;}
//	      pre {background-color: #eee; border: 1px solid #ddd; padding: 5px;}
//	    </style>
//	</head>
//	<body>
//	    <h1>Error: 404 Not Found</h1>
//	    <p>Sorry, the requested URL <tt>'https://vgmdb.info/album/01'</tt>
//	       caused an error:</p>
//	    <pre>Item not found</pre>
//	</body>
//
// </html>
//
// or:
//
// <html>
// <head><title>502 Bad Gateway</title></head>
// <body>
// <center><h1>502 Bad Gateway</h1></center>
// <hr><center>nginx/1.17.6</center>
// </body>
// </html>
func parseError(data string) string {
	errorHtmlTag := "pre"
	errorHtmlTagOpen := fmt.Sprintf("<%s>", errorHtmlTag)
	errorHtmlTagClose := fmt.Sprintf("</%s>", errorHtmlTag)

	i := strings.Index(data, errorHtmlTagOpen)
	j := strings.Index(data, errorHtmlTagClose)

	if i != -1 && j != -1 && i+len(errorHtmlTagOpen) < j {
		return data[i+len(errorHtmlTagOpen) : j]
	}

	return fmt.Sprintf("failed to parse unknown error format: %s", data)
}
