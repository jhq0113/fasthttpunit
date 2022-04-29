package internal

const (
	yml  = `.yml`
	yaml = `.yaml`
	json = `.json`
)

func IsYaml(fileName string) bool {
	if len(fileName) > len(yml) {
		if fileName[len(fileName)-len(yml):] == yml {
			return true
		}

	}

	return false
}

func IsJson(fileName string) bool {
	if len(fileName) > len(json) {
		return fileName[len(fileName)-len(json):] == json
	}
	return false
}
