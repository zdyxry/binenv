package install

import (
	"archive/tar"
	"compress/gzip"
	"io"
	"log"
	"os"

	"github.com/devops-works/binenv/internal/mapping"
	"github.com/devops-works/binenv/internal/tpl"
)

// Tgz handles zip files
type Tgz struct {
	filters []string
}

// Install files from tgz file
func (t Tgz) Install(src, dst, version string, targetArch, targetOS string, mapper mapping.Mapper) error {
	noMatches := ErrNoMatch

	f, err := os.Open(src)
	if err != nil {
		return err
	}
	defer f.Close()

	gzf, err := gzip.NewReader(f)
	if err != nil {
		return err
	}

	tarReader := tar.NewReader(gzf)
	args := tpl.New(version, targetArch, targetOS, mapper)

	for {
		header, err := tarReader.Next()
		if err == io.EOF {
			break
		}

		if err != nil {
			return err
		}

		switch header.Typeflag {
		case tar.TypeReg: // regular file
			ok, err := args.MatchFilters(header.Name, t.filters)
			if err != nil {
				return err
			}
			if !ok {
				continue
			}

			out, err := os.OpenFile(dst, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0755)
			if err != nil {
				return err
			}
			defer out.Close()
			if _, err := io.Copy(out, tarReader); err != nil {
				log.Fatal(err)
			}
			noMatches = nil
		}
	}

	return noMatches
}
