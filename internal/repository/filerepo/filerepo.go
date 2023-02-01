package filerepo

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/alsolovyev/dummy-api/pkg/slices"
)

var (
	filesDir          = "data"
	supportedFileExts = []string{"", ".json"}
)

type FileRepo struct{}

func New() *FileRepo {
	return &FileRepo{}
}

func (f *FileRepo) GetFile(p string) (interface{}, error) {
	ext := strings.ToLower(filepath.Ext(p))

	if !slices.Contains(supportedFileExts, ext) {
		return nil, errors.New(fmt.Sprintf("Files with '%s' extension are not supported.", ext))
	}

	b, err := f.ReadFile(p)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("File '%s' does not exist", p))
	}

	if ext == ".json" {
		d, err := f.ParseJson(b)
		if err != nil {
			return nil, errors.New(fmt.Sprintf("An error occurred while parsing '%s': %s", p, err.Error()))
		}

		return d, nil
	}

	return string(b), nil
}

func (*FileRepo) ReadFile(p string) ([]byte, error) {
	wd, err := os.Getwd()

	if err != nil {
		return nil, err
	}

	f := path.Join(wd, filesDir, p)

	if _, err := os.Stat(f); err != nil {
		return nil, err
	}

	b, err := os.ReadFile(f)

	if err != nil {
		return nil, err
	}

	return b, nil
}

func (*FileRepo) ParseJson(b []byte) (interface{}, error) {
	var j interface{}

	if err := json.Unmarshal(b, &j); err != nil {
		return nil, err
	}

	return j, nil
}
