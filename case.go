package fasthttpunit

type Case struct {
	Desc         string            `json:"desc" yaml:"desc"`
	Params       string            `json:"params" yaml:"params"`
	Header       map[string]string `json:"header" yaml:"header"`
	Expected     string            `json:"expected" yaml:"expected"`
	ExpectedType string            `json:"expectedType" yaml:"expectedType"`
}
