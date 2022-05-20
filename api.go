package fasthttpunit

import (
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/jhq0113/fasthttpunit/internal"

	jsoniter "github.com/json-iterator/go"
	"gopkg.in/yaml.v2"
)

func loadApi(filePath string) (*Api, error) {
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	var api Api

	if internal.IsYaml(filePath) {
		err = yaml.Unmarshal(data, &api)
		if err != nil {
			return nil, err
		}

		api.load()

		return &api, nil
	}

	if internal.IsJson(filePath) {
		err = jsoniter.Unmarshal(data, &api)
		if err != nil {
			return nil, err
		}

		api.load()

		return &api, nil
	}

	return nil, ErrUnsupportedFileType
}

type Api struct {
	Desc        string            `json:"desc" yaml:"desc"`
	Host        string            `json:"host" yaml:"host"`
	Method      string            `json:"method" yaml:"method"`
	Path        string            `json:"path" yaml:"path"`
	ContentType string            `json:"contentType" yaml:"contentType"`
	Header      map[string]string `json:"header" yaml:"header"`
	CaseList    []Case            `json:"caseList" yaml:"caseList"`
}

func (a *Api) load() {
	if a.Host == "" {
		a.Host = "127.0.0.1"
	}

	if a.Method == "" {
		a.Method = http.MethodGet
	} else {
		a.Method = strings.ToUpper(a.Method)
	}

	if a.Method == http.MethodPost && a.ContentType == "" {
		a.ContentType = "application/x-www-form-urlencoded"
	}
}
