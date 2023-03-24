package ambiance

import (
	"errors"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

var names = map[string]string{
	"main":    ".env",
	"sample":  ".env.sample",
	"example": ".env.example",
}

type envObject map[string]string

// Prepares the environment variables according to a .env file.
func Config(dirPath string) {

	basePath, err := os.Getwd()
	if err != nil {
		panic("error getting cwd")
	}
	envDir := filepath.Join(basePath, dirPath)

	if exist := filesExist(envDir); !exist {
		panic("unable to find '.env' and/or '.env.sample'/'.env.example' file(s)")
	}

	sampleFilename, err := sampleOrExample(envDir)
	if err != nil {
		panic(err)
	}
	mainContents, err := readTextFile(filepath.Join(envDir, names["main"]))
	if err != nil {
		panic(err)
	}
	sampleContents, err := readTextFile(filepath.Join(envDir, names[sampleFilename]))
	if err != nil {
		panic(err)
	}
	mainMap, err := mapVars(mainContents)
	if err != nil {
		panic(err)
	}
	sampleMap, err := mapVars(sampleContents)
	if err != nil {
		panic(err)
	}

	if !keysMatch(mainMap, sampleMap) {
		panic("environment variables from main file don't match with sample file")
	}

	err = setEnv(mainMap)
	if err != nil {
		panic("error setting environment variables")
	}

}

// Checks if all environment variables are properly set, according to the sample file.
//
// dirPath is the relative path where sample file is located.
func Check(dirPath string) error {
	return errors.New("")
}

func filesExist(basePath string) bool {
	env := filepath.Join(basePath, names["main"])
	sample := filepath.Join(basePath, names["sample"])
	example := filepath.Join(basePath, names["example"])

	if _, err := os.Stat(env); errors.Is(err, os.ErrNotExist) {
		return false
	} else {
		_, err1 := os.Stat(sample)
		_, err2 := os.Stat(example)
		if errors.Is(err1, os.ErrNotExist) && errors.Is(err2, os.ErrNotExist) {
			return false
		}
	}
	return true
}

func sampleOrExample(basePath string) (string, error) {
	sample := filepath.Join(basePath, names["sample"])
	example := filepath.Join(basePath, names["example"])

	if _, err := os.Stat(sample); err == nil {
		return "sample", nil
	} else if _, err := os.Stat(example); err == nil {
		return "example", nil
	} else {
		return "", errors.New("sample file not found")
	}
}

func keysMatch(mainMap, sampleMap envObject) bool {
	for key := range mainMap {
		if _, ok := sampleMap[key]; !ok {
			return false
		}
	}
	return true
}

// Returns a string with the text file contents
func readTextFile(filePath string) (string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	contents, err := os.ReadFile(file.Name())
	if err != nil {
		return "", err
	}

	return string(contents), file.Close()
}

func mapVars(fileContents string) (envObject, error) {
	varMap := make(map[string]string)

	lines := strings.Split(string(fileContents), eol())
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if string(line[0]) == "#" || line == "" {
			continue
		}

		parts := strings.Split(line, "=")
		if len(parts) != 2 {
			return nil, errors.New("invalid env file contents")
		}

		key := strings.TrimSpace(parts[0])
		value := strings.TrimSpace(parts[1])
		varMap[key] = value
	}

	return varMap, nil
}

func setEnv(envMap envObject) error {
	var err error
	for key, value := range envMap {
		err = os.Setenv(key, value)
	}
	if err != nil {
		return err
	} else {
		return nil
	}
}

func eol() string {
	eol := "\n"
	if string(os.PathSeparator) != "/" || runtime.GOOS == "windows" {
		eol = "\r\n"
	}
	return eol
}
