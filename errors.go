package fasthttpunit

import "errors"

var (
	ErrUnsupportedFileType = errors.New(`unsupported file type`)
	ErrNotFoundApi         = errors.New("not found api")
)
