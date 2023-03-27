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
func Config(dirPath string) {
	basePath = filepath.Join(cwd, dirPath)
	determineSampleFilename(basePath)
	updatePaths(basePath)
	if exist := filesExist(paths["main"], paths["sample"]); !exist {
		panic("unable to find '.env' and/or sample file(s)")
	}

	mainMap, sampleMap := parseEnvFiles(paths["main"], paths["sample"])
	if !keysMatch(mainMap, sampleMap) {
		panic("environment variables from main file don't match with the ones from sample file")
	}

	err := setEnv(mainMap)
	if err != nil {
		panic("error setting environment variables")
	}
}

// Checks if all environment variables are properly set, according to the sample file.
//
// dirPath is the relative path where sample file is located.
func Check(dirPath string) error {
	// @TODO
	return errors.New("")
}

func determineSampleFilename(basePath string) {
	sample := filepath.Join(basePath, filenames["sample"])
	example := filepath.Join(basePath, nameOption)

	if _, err := os.Stat(sample); err == nil {
		// do nothing, as sample filename is already correct
	} else if _, err := os.Stat(example); err == nil {
		filenames["sample"] = nameOption
	} else {
		panic(".env.sample or .env.example is missing")
	}
}

func updatePaths(dirPath string) {
	paths["main"] = filepath.Join(dirPath, filenames["main"])
	paths["sample"] = filepath.Join(dirPath, filenames["sample"])
}

func filesExist(mainPath, samplePath string) bool {
	if _, err := os.Stat(mainPath); errors.Is(err, os.ErrNotExist) {
		return false
	} else if _, err1 := os.Stat(samplePath); errors.Is(err1, os.ErrNotExist) {
		return false
	}
	return true
}

func parseEnvFiles(mainPath, samplePath string) (envObject, envObject) {

	mainContents, err := readTextFile(mainPath)
	if err != nil {
		panic(err)
	}
	mainMap, err := mapVars(mainContents)
	if err != nil {
		panic(err)
	}

	sampleContents, err := readTextFile(samplePath)
	if err != nil {
		panic(err)
	}
	sampleMap, err := mapVars(sampleContents)
	if err != nil {
		panic(err)
	}

	return mainMap, sampleMap

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
