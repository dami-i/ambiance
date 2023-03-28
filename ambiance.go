package ambiance

import (
	"errors"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

var (
	filenames = map[string]string{
		"main":   ".env",
		"sample": ".env.sample",
	}
	nameOption = ".env.example"
	cwd, _     = os.Getwd()
	basePath   = filepath.Join(cwd) // initial value
	paths      = map[string]string{
		"main":   filepath.Join(cwd, filenames["main"]),   // initial value
		"sample": filepath.Join(cwd, filenames["sample"]), // initial value
	}
)

type envObject map[string]string

// Prepares the environment variables according to a .env file.
//
// dirPath is the path to where the .env and sample files are located relative to the project's root.
func Config(relativeDirPath string, useSampleAsTemplate bool) {
	basePath = filepath.Join(cwd, relativeDirPath)
	if useSampleAsTemplate {
		err := determineSampleFilename(basePath)
		if err != nil {
			panic(err)
		}
	}
	updatePaths(basePath)

	if !fileExist(paths["main"]) {
		panic("unable to find '.env' file")
	}

	mainMap, err := parseEnvFile(paths["main"])
	if err != nil {
		panic(err.Error() + " error parsing .env file")
	}
	if useSampleAsTemplate {
		sampleMap, err := parseEnvFile(paths["sample"])
		if err != nil {
			panic(err.Error() + " error parsing sample file")
		}
		if !keysMatch(mainMap, sampleMap) {
			panic("environment variables from main file differ from sample file")
		}
	}

	err = setEnv(mainMap)
	if err != nil {
		panic("error setting environment variables")
	}
}

// Checks if all environment variables are properly set, according to the sample file.
//
// dirPath is the relative path where sample file is located.
func Check(relativeDirPath string, allowEmptyValues bool) {
	basePath = filepath.Join(cwd, relativeDirPath)
	err := determineSampleFilename(basePath)
	if err != nil {
		panic(err)
	}
	updatePaths(basePath)

	sampleMap, err := parseEnvFile(paths["sample"])
	if err != nil {
		panic(err.Error() + " error parsing sample file")
	}

	for key := range sampleMap {
		if val, isSet := os.LookupEnv(key); !isSet || val == "" {
			if !isSet {
				panic(key + " is not set")
			} else {
				if allowEmptyValues {
					// do nothing
				} else {
					panic(key + " value is empty")
				}
			}
		}
	}
}

func determineSampleFilename(basePath string) error {
	sample := filepath.Join(basePath, filenames["sample"])
	example := filepath.Join(basePath, nameOption)
	if fileExist(sample) {
		// do nothing - filenames["sample"] already has the correct value
	} else if fileExist(example) {
		filenames["sample"] = nameOption
	} else {
		return errors.New("missing sample file")
	}
	return nil
}

func updatePaths(dirPath string) {
	paths["main"] = filepath.Join(dirPath, filenames["main"])
	paths["sample"] = filepath.Join(dirPath, filenames["sample"])
}

func fileExist(filePath string) bool {
	if _, err := os.Stat(filePath); errors.Is(err, os.ErrNotExist) {
		return false
	}
	return true
}

func parseEnvFile(filePath string) (envObject, error) {
	contents, err := readTextFile(filePath)
	if err != nil {
		return nil, err
	}
	varMap, err := mapVars(contents)
	if err != nil {
		return nil, err
	}
	return varMap, nil
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
			return nil, errors.New("invalid file contents")
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
