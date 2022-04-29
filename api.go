package fasthttpunit

type Api struct {
	Desc     string            `json:"desc" yaml:"desc"`
	Method   string            `json:"method" yaml:"method"`
	Path     string            `json:"path" yaml:"path"`
	Header   map[string]string `json:"header" yaml:"header"`
	CaseList []Case            `json:"caseList" yaml:"caseList"`
}
