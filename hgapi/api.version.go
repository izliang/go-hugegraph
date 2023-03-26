package hgapi

import (
	"context"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
)

// ----- API Definition -------------------------------------------------------

// CatSegments provides low-level information about the segments in the shards of an index.
//
// See full documentation at https://www.elastic.co/guide/en/elasticsearch/reference/5.x/cat-segments.html.
//
func newVersionFunc(t Transport) Version {
	return func(o ...func(*VersionRequest)) (*VersionResponse, error) {
		var r = VersionRequest{}
		for _, f := range o {
			f(&r)
		}
		return r.Do(r.ctx, t)
	}
}

type Version func(o ...func(*VersionRequest)) (*VersionResponse, error)

type VersionRequest struct {
	Body io.Reader
	ctx  context.Context
}

type VersionResponse struct {
	StatusCode int           `json:"-"`
	Header     http.Header   `json:"-"`
	Body       io.ReadCloser `json:"-"`
	Versions   struct {
		Version string `json:"version"`
		Core    string `json:"core"`
		Gremlin string `json:"gremlin"`
		API     string `json:"api"`
	} `json:"versions"`
}

func (r VersionRequest) Do(ctx context.Context, transport Transport) (*VersionResponse, error) {

	req, _ := newRequest("GET", "/versions", r.Body)

	if ctx != nil {
		req = req.WithContext(ctx)
	}

	res, err := transport.Perform(req)
	if err != nil {
		return nil, err
	}

	versionResp := &VersionResponse{}
	bytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(bytes, versionResp)
	if err != nil {
		return nil, err
	}
	versionResp.StatusCode = res.StatusCode
	versionResp.Header = res.Header
	versionResp.Body = res.Body
	return versionResp, nil
}
