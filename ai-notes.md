Pseudocode:

1. Import the os package
2. Define a function that takes the filename and returns a map of environment variables and their values
3. Open the file using os.Open() function and defer its closure
4. Read the contents of the file using os.ReadFile() function
5. Split the contents of the file by lines using strings.Split() function
6. Iterate through each line, split it by "=" and store the key and value in a map
7. Return the map
8. In the main function, call the previously defined function and iterate through the map returned
9. Use os.Setenv() to set the environment variables with the keys and values from the map

Code:

```go
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func readEnvFile(filename string) map[string]string {
	envMap := make(map[string]string)
	file, err := os.Open(filename)
	defer file.Close()
	if err != nil {
		log.Fatalf("Error opening file: %s", err)
	}
	contents, err := os.ReadFile(file.Name())
	if err != nil {
		log.Fatalf("Error reading file: %s", err)
	}
	lines := strings.Split(string(contents), "\n")
	for _, line := range lines {
		if line == "" {
			continue
		}
		parts := strings.Split(line, "=")
		key := strings.TrimSpace(parts[0])
		value := strings.TrimSpace(parts[1])
		envMap[key] = value
	}
	return envMap
}

func main() {
	envMap := readEnvFile("env.txt")
	for key, value := range envMap {
		os.Setenv(key, value)
		fmt.Printf("%s=%s\n", key, value)
	}
}
```

---

To turn this code into a public go package, you should:

1. Create a new empty go module by executing the command `go mod init <module_name>` in your terminal, inside the directory that contains your code. Replace `<module_name>` with an appropriate name for your new module, for example, `github.com/<username>/env-reader`.
2. Create a new file with the name `env_reader.go`.
3. Copy the code from `main.go` into `env_reader.go`.
4. Change the package name from `main` to an appropriate name for your package, for example, `envreader`.
5. Add documentation comments to each function for better clarity.
6. Use `go fmt` to format the code according to Go standard conventions.
7. Test the package within the same directory or create a separate test directory.
8. Push the package to a version control system like Git.
9. Publish the package on a public Go package hosting platform like https://pkg.go.dev or Github.

Here's what the updated code would look like:

```go
package envreader

import (
	"log"
	"os"
	"strings"
)

// ReadEnvFile reads the environment variables from a file and returns a map of strings.
func ReadEnvFile(filename string) map[string]string {
	envMap := make(map[string]string)
	file, err := os.Open(filename)
	defer file.Close()
	if err != nil {
		log.Fatalf("Error opening file: %s", err)
	}
	contents, err := os.ReadFile(file.Name())
	if err != nil {
		log.Fatalf("Error reading file: %s", err)
	}
	lines := strings.Split(string(contents), "\n")
	for _, line := range lines {
		if line == "" {
			continue
		}
		parts := strings.Split(line, "=")
		key := strings.TrimSpace(parts[0])
		value := strings.TrimSpace(parts[1])
		envMap[key] = value
	}
	return envMap
}

// SetEnvVariables reads the environment variables from a file and sets the corresponding env variables.
func SetEnvVariables(filename string) {
	envMap := ReadEnvFile(filename)
	for key, value := range envMap {
		os.Setenv(key, value)
	}
}
```

You can then test importing this package to another Go application to make sure the package works correctly.

---

Plan:

1. Import the "os" package
2. Call the Getwd() function of the os package which will return the current working directory path
3. Handle errors if any
4. Print the current directory path


Pseudocode:
```go
import "os"

// Call the function getwd to get the current directory path
dirpath, err := os.Getwd()

// Handle the errors if any
if err != nil {
	fmt.Println(err)
}

// Print the current directory path
fmt.Println(dirpath)
```

Code:
```go
package main

import (
	"fmt"
	"os"
)

func main() {
	// Call the function getwd to get the current directory path
	dirpath, err := os.Getwd()

	// Handle the errors if any
	if err != nil {
		fmt.Println(err)
	}

	// Print the current directory path
	fmt.Println(dirpath)
}
```

---

Pseudocode:
```
1. Import the "path/filepath" package.
2. Get the present working directory using os.Getwd().
3. Use filepath.Join() to join the present working directory path and the desired filename.
```

Here's the corresponding Go code:

```go
import (
	"os"
	"path/filepath"
)

func main() {
	// Get the present working directory
	pwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	// Initialize filename variable
	filename := "example.txt"

	// Join the filename to the present working directory
	fileWithPath := filepath.Join(pwd, filename)
}
```

The `fileWithPath` variable will contain the full path to the desired file, combining the present working directory and the filename.
