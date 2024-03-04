# gosort

**gosort** is a Go package that provides a concurrent implementation of the Quick Sort algorithm for sorting slices of integers, unsigned integers, and floats in ascending order.

## Installation

To use gosort in your Go project, simply import it using the following import statement:

```go
import "github.com/jrjaro18/gosort"
```

Then, install the package by running:

```bash
go get github.com/jrjaro18/gosort
```

## Usage

```go
package main

import (
    "fmt"
    "github.com/jrjaro18/gosort"
)

func main() {
    // Example usage with a slice of integers
    integers := []int{5, 3, 7, 1, 9, 2, 6, 4, 8}
    fmt.Println("Unsorted:", integers)

    // Sorting the slice using gosort
    gosort.Sort(&integers)
    fmt.Println("Sorted:", integers)
}
```

## Features

- **Concurrent Sorting**: Utilizes goroutines to parallelize the sorting process, improving performance for large datasets.
- **Support for Multiple Types**: Can sort slices of integers, unsigned integers, and floats.
- **Quick Sort Algorithm**: Employs the Quick Sort algorithm for efficient sorting.

## Contributing

Contributions are welcome! If you encounter any issues or have suggestions for improvements, please feel free to open an issue or submit a pull request on GitHub.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

---
