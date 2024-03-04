# go-sort

**go-sort** is a Go package that provides a concurrent implementation of the Quick Sort algorithm for sorting slices of integers, unsigned integers, and floats in ascending order.

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

---

# Performance Comparison

## Test Conditions:

- Hardware:
    - Processor: 4.7 GHz Intel Core i7-1165G7
    - Ram: 16 GB LPDDR4X
      
- Go version: 1.21.3

## Quick Sort vs. go-sort (Concurrent Quick Sort)

### 1 Million Numbers:
- Traditional Quick Sort: 104 milliseconds
- go-sort Quick Sort: 77 milliseconds

### 10 Million Numbers:
- Traditional Quick Sort: 1.29 seconds
- go-sort Quick Sort: 327 milliseconds

### 100 Million Numbers:
- Traditional Quick Sort: 13 seconds
- go-sort Quick Sort: 3 seconds

**Analysis:**
- go-sort consistently outperforms traditional Quick Sort, showcasing significant speed enhancements across varying dataset sizes.
- The concurrent nature of gosort's implementation leverages Go's goroutines, resulting in improved efficiency, particularly evident with larger datasets.
- These performance gains make gosort a compelling choice for applications requiring high-speed sorting of large datasets.

---

## Contributing

Contributions are welcome! If you encounter any issues or have suggestions for improvements, please feel free to open an issue or submit a pull request on GitHub.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

---
