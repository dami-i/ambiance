# Ambiance

Use **.env** files in your project.

## Usage

Import it and use it inside `init()` function. You may want to alias it as `amb`.

```go
import (
	"github.com/dami-i/ambiance"
)

func init() {
	// Pass as argument the relative path to the directory where the .env files are located
	ambiance.Config("./")
}

func main() {
	// ...
}
```