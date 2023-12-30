package fetch

import (
	"context"

	"github.com/devops-works/binenv/internal/mapping"
)

// Fetcher should implement fetching a release from a version
// and return a path where the release has been downloaded
type Fetcher interface {
	Fetch(ctx context.Context, dist, version string, arch, os string, mapper mapping.Mapper) (string, error)
}

// Fetch contains fetch configuration
type Fetch struct {
	Type string `yaml:"type"`
	URL  string `yaml:"url"`
	OS   string `yaml:"os"`
	Arch string `yaml:"arch"`
}

// Factory returns instances that comply to Fecther interface
func (r Fetch) Factory() Fetcher {
	switch r.Type {
	case "download":
		return Download{
			url:  r.URL,
			os:   r.OS,
			arch: r.Arch,
		}
	default:
		return Download{
			url:  r.URL,
			os:   r.OS,
			arch: r.Arch,
		}
	}
}
