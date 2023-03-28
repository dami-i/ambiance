# Ambiance

![v0.1.1](https://img.shields.io/badge/version-v0.1.1-blue)

Use **.env** file in your project for setting up the environment variables.

## Usage

There are two methods available: `.Config()` and `.Check()`.

`.Config()` uses a `.env` file for setting the environment variables.

`.Check()` uses a `.env.sample` (or `.env.example`) file for checking if all variables are set.

Import and use them inside `init()` function as needed.

```go
import (
	"github.com/dami-i/ambiance"
)

func init() {
	// Pass as argument the relative path to the directory where the .env file will be located
	ambiance.Config("./", true)
	ambiance.Check("./", false)
}

func main() {
	// ...
}
```

## API

### `ambiance.Config(relativePath string, useSampleAsTemplate bool)`

Sets up the project's environment variables according to a `.env` file.

If `useSampleAsTemplate` is set to `true`, it will look for a `env.sample` (or `.env.example`) as template.

The program panics if anything goes wrong, like file not found.

### `ambiance.Check(relativePath string, allowEmptyValues bool)`

Verifies if all environment variables declared in the sample file are properly defined in the project.

It will search for files named `.env.sample` or `.env.example` in the given directory, in that order.

> **Note:** if both `.env.sample` and `.env.example` files are found, `.env.sample` takes precedence over the other, which in turn will be ignored.

If `allowEmptyValues` is set to `true`, the environment variables whose value is an empty string are not treated as invalid.
