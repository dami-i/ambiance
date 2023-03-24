package ambiance

import (
	"errors"
	"os"
	"strings"
)

// Prepares the environment with variables according to a .env file.
func PrepareEnv() error {
	return errors.New("")
}

// Checks if all environment variables are properly set, according to the sample file.
func precheck() error {
	return errors.New("")
}

func readFile(filePath string) (string, error) {
	return "", nil
}

func mapEnvVars(fileContents string) (map[string]string, error) {
	envMap := make(map[string]string)
	
	file, err := os.Open(fileContents)
	if err != nil {
		return nil, err
	}

	contents, err := os.ReadFile(file.Name())
	if err != nil {
		return nil, err
	}

	lines := strings.Split(string(contents), eol())
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if strings.Split(line, "")[0] == "#" || line == "" {
			continue
		}
		parts := strings.Split(line, "=")
		key := strings.TrimSpace(parts[0])
		value := strings.TrimSpace(parts[1])
		envMap[key] = value
	}

	return envMap, file.Close()
}

func eol() string {
	eol := "\n"
	if string(os.PathSeparator) != "/" {
		eol = "\r\n"
	}
	return eol
}