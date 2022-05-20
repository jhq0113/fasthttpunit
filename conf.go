package fasthttpunit

import (
	"errors"
	"fmt"
	"io/ioutil"
)

type Conf struct {
	ApiList []*Api `json:"apiList" yaml:"apiList"`
	Delay   uint64 `json:"delay" yaml:"delay"`
}

func newConf() *Conf {
	return &Conf{
		Delay: 1,
	}
}

func LoadConf(basePath string) (c *Conf, err error) {
	fileList, err := ioutil.ReadDir(basePath)
	if err != nil {
		return nil, err
	}

	conf := newConf()
	conf.ApiList = make([]*Api, 0, len(fileList))

	for _, f := range fileList {
		if f.IsDir() {
			continue
		}

		api, err := loadApi(basePath + "/" + f.Name())
		if err != nil {
			return nil, errors.New(fmt.Sprintf("file %s format error:%s", f.Name(), err.Error()))
		}
		conf.ApiList = append(conf.ApiList, api)
	}

	if len(conf.ApiList) < 1 {
		return nil, ErrNotFoundApi
	}

	return conf, nil
}
