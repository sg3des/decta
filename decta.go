package decta

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"path"
	"strings"

	"github.com/goware/urlx"
	"github.com/imroc/req"
	"github.com/op/go-logging"
)

// Address global variable is address of the Decta API
var Address = "https://gate.decta.com/api/v0.6/"

var log = logging.MustGetLogger("DECTA")

// Decta type wrap all decta methods
type Decta struct {
	url           url.URL
	urlPathPrefix string
	authHeader    req.Header
}

// NewDecta initialize Decta instance,
// parse address from the global variable `decta.Address`
// prepare authorization header for requests
func NewDecta(secretkey string) (*Decta, error) {
	u, err := urlx.ParseWithDefaultScheme(Address, "https")
	if err != nil {
		return nil, fmt.Errorf("failed to parse address: %v", err)
	}

	if secretkey == "" {
		return nil, errors.New("secret key is empty")
	}

	d := &Decta{
		url:           *u,
		urlPathPrefix: u.Path,
		authHeader:    req.Header{"Authorization": "Bearer " + secretkey},
	}

	return d, nil
}

// URL method should return prepapred URL-address with specified URL-path
func (d *Decta) URL(urlpath string) string {
	u := d.url
	u.Path = path.Join(d.urlPathPrefix, urlpath)

	if !strings.HasSuffix(u.Path, "/") {
		u.Path += "/"
	}

	return u.String()
}

//
// REQUESTS
//

//
// response error
//

type responseError struct {
	Status     string `json:"-"`
	StatusCode int    `json:"-"`

	json.RawMessage
}

type messageError struct {
	Message string
	Code    string
}

func (err messageError) String() string {
	if err.Message != "" {
		return err.Message
	}
	if err.Code != "" {
		return err.Code
	}

	return "unknown error"
}

func (err responseError) Error() string {
	var all struct {
		All []messageError `json:"__all__"`
	}
	json.Unmarshal(err.RawMessage, &all)

	if len(all.All) > 0 {
		var errLines []string
		for _, msg := range all.All {
			errLines = append(errLines, msg.String())
		}
		return strings.Join(errLines, "; ")
	}

	if err.RawMessage != nil {
		return string(err.RawMessage)
	}

	return err.Status
}

// response type contains response data in `Results` array
type response struct {
	Next     interface{}     `json:"next"`
	Previous interface{}     `json:"previous"`
	Results  json.RawMessage `json:"results"`
}

// handlerResponse parse error if it occurd, if not
// trying to parse response data to destination variable `v`, if it specified
func (d *Decta) handleResponse(resp *req.Resp, v interface{}) error {
	// log.Debug(resp.Dump())

	if r := resp.Response(); r.StatusCode >= 400 {
		err := &responseError{
			Status:     r.Status,
			StatusCode: r.StatusCode,
		}

		if d.contentTypeMatch(r.Header, contentTypeJSON) {
			if e := resp.ToJSON(err); e != nil {
				log.Warning("failed to decode response error")
			}
		}

		return err
	}

	// if destination not specified, then do nothing with response
	if v == nil {
		return nil
	}

	// try to unmarshal how multiple records, how response type
	var r response
	if err := resp.ToJSON(&r); err != nil {
		return err
	}

	// if `results` exists then unmarshal to destination type
	if r.Results != nil {
		return json.Unmarshal(r.Results, v)
	}

	// if unmarshal json to `response` failed, then try unmarshal directly
	// to specified destination type
	return resp.ToJSON(v)
}

const (
	headerContentType = "Content-Type"
	contentTypeJSON   = "application/json"
)

// contentTypeMatch check match contentType if required
func (d *Decta) contentTypeMatch(headers http.Header, contentType string) bool {
	return strings.Contains(headers.Get(headerContentType), contentType)
}

// GET method request by specified path
func (d *Decta) GET(urlpath string, v interface{}) error {
	resp, err := req.Get(d.URL(urlpath), d.authHeader)
	if err != nil {
		return err
	}

	return d.handleResponse(resp, v)
}

func (d *Decta) POST(urlpath string, data, v interface{}) error {
	resp, err := req.Post(d.URL(urlpath), d.authHeader, req.BodyJSON(data))
	if err != nil {
		return err
	}

	return d.handleResponse(resp, v)
}

func (d *Decta) DELETE(urlpath string, v interface{}) error {
	resp, err := req.Delete(d.URL(urlpath), d.authHeader)
	if err != nil {
		return err
	}

	return d.handleResponse(resp, v)
}
