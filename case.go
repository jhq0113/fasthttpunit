package fasthttpunit

type Case struct {
	Desc   string            `json:"desc" yaml:"desc"`
	Header map[string]string `json:"header" yaml:"header"`
	Body   string            `json:"body" yaml:"body"`
}
