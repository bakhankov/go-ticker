`go-ticker` package extends the functionality of the `Ticker` struct from the `time` package in Go. 

It adds 3 new features:
1. **flag `tickOnInit`**: Allows the ticker to tick immediately upon initialization.
2. **method `TickAfter(d time.Duration)`**: Allows the ticker to tick after a specified duration.
3. **method `Tick()`**: Allows the ticker to tick immediately, regardless of the interval.

## Installation
To install the `go-ticker` package, use the following command:

```bash
go get github.com/bakhankov/go-ticker
```

## Usage
Here's a simple example of how to use the `go-ticker` package:

```go
package main
import (
    "fmt"
    "time"

    "github.com/bakhankov/go-ticker"
)

func main() {
	// tickOnInit is set to true, so the ticker will tick immediately upon creation
	t := timeutils.NewTicker(2*time.Second, true)

	// Start the ticker
	go func() {
		for {
			<-t.C
			fmt.Println("Tick at", time.Now())
		}
	}()

	time.Sleep(1 * time.Second)
	// Tick immediately at call
	t.Tick()

	// Schedule a tick after 3 seconds
	t.TickAfter(3 * time.Second)

	time.Sleep(5 * time.Second)
	t.Stop()
}
```
## Features
- **Immediate Ticking**: The ticker can tick immediately upon initialization if `tickOnInit` is set to true.
- **Manual Ticking**: You can manually trigger a tick using the `Tick()` method, regardless of the ticker's interval.
- **Tick After Duration**: You can trigger a tick after a specified duration using the `TickAfter(d time.Duration)` method.
## License
This package is licensed under the MIT License. See the [LICENSE](LICENSE) file for more details.
## Contributing
Contributions are welcome! If you find a bug or have a feature request, please open an issue on the [GitHub repository](https://github.com/bakhankov/go-ticker). Pull requests are also welcome.
## Documentation
For more detailed documentation, please refer to the [GoDoc](https://pkg.go.dev/github.com/bakhankov/go-ticker) page for this package.
## Issues
If you encounter any issues while using the `go-ticker` package, please check the [issues page](https://github.com/bakhankov/go-ticker/issues) on GitHub. If your issue is not listed, feel free to open a new issue with a detailed description of the problem.
## Author
This package is developed and maintained by [bakhankov](https://github.com/bakhankov). If you have any questions or suggestions, feel free to reach out via GitHub.
## Support
If you find this package useful, consider supporting the development by starring the repository on GitHub. Your support helps in maintaining and improving the package.
## Acknowledgements
This package is inspired by the need for a more flexible and feature-rich ticker implementation in Go. It builds upon the standard `time.Ticker` while adding additional functionality to enhance usability in various scenarios.
