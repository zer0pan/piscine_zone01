package main

import (
	"fmt"
	"os"
	"os/exec"
)

// Function definitions as strings
var quadCheckerSrc = `
package main

import (
    "fmt"
    "os"
    "os/exec"
    "bytes"
)

func main() {
    // Read all input from stdin
    input, err := os.ReadFile("/dev/stdin")
    if err != nil {
        fmt.Println("Error reading input:", err)
        return
    }

    input = bytes.TrimSpace(input)  // Remove leading/trailing spaces and newlines
    inputLines := parseInput(input)
    y := len(inputLines)
    if y == 0 {
        fmt.Println("Not a quad function")
        return
    }
    x := len(inputLines[0])

    // List of valid quad executables
    quads := []string{"./quadA", "./quadB", "./quadC", "./quadD", "./quadE"}
    matchingQuads := []string{}

    // Compare input with outputs from each valid quad function
    for _, quad := range quads {
        quadOutput, err := runQuad(quad, x, y)
        if err != nil {
            // Suppress the "no such file or directory" error
            if os.IsNotExist(err) {
                continue
            }
            // Print other errors
            fmt.Println("Error running", quad, ":", err)
            continue
        }

        quadOutput = bytes.TrimSpace(quadOutput)  // Remove extra spaces/newlines from the quad output

        if string(quadOutput) == string(input) {
            matchingQuads = append(matchingQuads, fmt.Sprintf("[%s] [%d] [%d]", quad[2:], x, y))
        }
    }

    // Output the results or "Not a quad function" if no match
    if len(matchingQuads) > 0 {
        fmt.Println(joinResults(matchingQuads, " || "))
    } else {
        fmt.Println("Not a quad function")
    }
}

// runQuad executes a quad function and captures its output
func runQuad(quad string, x, y int) ([]byte, error) {
    cmd := exec.Command(quad, fmt.Sprintf("%d", x), fmt.Sprintf("%d", y))
    output, err := cmd.Output()

    // If the command doesn't exist, return a specific error
    if err != nil {
        if exitErr, ok := err.(*exec.ExitError); ok && exitErr.Exited() {
            return nil, os.ErrNotExist
        }
        return nil, err
    }
    return output, nil
}

// parseInput converts raw input bytes into a 2D slice of bytes
func parseInput(input []byte) [][]byte {
    var lines [][]byte
    var line []byte
    for _, b := range input {
        if b == '\n' {
            lines = append(lines, line)
            line = []byte{} // reset for new line
        } else {
            line = append(line, b)
        }
    }
    // Add the last line if input doesn't end with newline
    if len(line) > 0 {
        lines = append(lines, line)
    }
    return lines
}

// joinResults manually concatenates slice of strings with a separator
func joinResults(results []string, sep string) string {
    if len(results) == 0 {
        return ""
    }
    result := results[0]
    for _, s := range results[1:] {
        result += sep + s
    }
    return result
}
`

var quadASrc = `
package main
import "fmt"
import "os"
import "strconv"
func main(){
cols,_ := strconv.Atoi(os.Args[1])
rows,_ := strconv.Atoi(os.Args[2])
QuadA(int(cols),int(rows))
}
func QuadA(x, y int) {
	// Handle cases where x or y is non-positive
	if x <= 0 || y <= 0 {
		return
	}

	// Draw the first row (top border)
	if y > 0 {
		fmt.Print("o")
		for i := 1; i < x-1; i++ {
			fmt.Print("-")
		}
		if x > 1 {
			fmt.Println("o")
		} else {
			fmt.Println()
		}
	}

	// Draw the middle rows (vertical borders with spaces in between)
	for j := 1; j < y-1; j++ {
		fmt.Print("|")
		for i := 1; i < x-1; i++ {
			fmt.Print(" ")
		}
		if x > 1 {
			fmt.Println("|")
		} else {
			fmt.Println()
		}
	}

	// Draw the last row (bottom border)
	if y > 1 {
		fmt.Print("o")
		for i := 1; i < x-1; i++ {
			fmt.Print("-")
		}
		if x > 1 {
			fmt.Println("o")
		} else {
			fmt.Println()
		}
	}
}
`

var quadBSrc = `
package main
import "fmt"
import "os"
import "strconv"
func main(){
cols,_ := strconv.Atoi(os.Args[1])
rows,_ := strconv.Atoi(os.Args[2])
QuadB(int(cols),int(rows))
}

func QuadB(x, y int) {
	// Handle cases where x or y is non-positive
	if x <= 0 || y <= 0 {
		return
	}

	// Draw the first row (top border)
	if y > 0 {
		fmt.Print("/")
		for i := 1; i < x-1; i++ {
			fmt.Print("*")
		}
		if x > 1 {
			fmt.Println("\\")
		} else {
			fmt.Println()
		}
	}

	// Draw the middle rows (vertical borders with spaces in between)
	for j := 1; j < y-1; j++ {
		fmt.Print("*")
		for i := 1; i < x-1; i++ {
			fmt.Print(" ")
		}
		if x > 1 {
			fmt.Println("*")
		} else {
			fmt.Println()
		}
	}

	// Draw the last row (bottom border)
	if y > 1 {
		fmt.Print("\\")
		for i := 1; i < x-1; i++ {
			fmt.Print("*")
		}
		if x > 1 {
			fmt.Println("/")
		} else {
			fmt.Println()
		}
	}
}
`

var quadCSrc = `
package main
import "fmt"
import "os"
import "strconv"
func main(){
cols,_ := strconv.Atoi(os.Args[1])
rows,_ := strconv.Atoi(os.Args[2])
QuadC(int(cols),int(rows))
}
func QuadC(x, y int) {
	// Handle cases where x or y is non-positive
	if x <= 0 || y <= 0 {
		return
	}

	// Draw the first row (top border)
	if y > 0 {
		fmt.Print("A")
		for i := 1; i < x-1; i++ {
			fmt.Print("B")
		}
		if x > 1 {
			fmt.Println("A")
		} else {
			fmt.Println()
		}
	}

	// Draw the middle rows (vertical borders with spaces in between)
	for j := 1; j < y-1; j++ {
		fmt.Print("B")
		for i := 1; i < x-1; i++ {
			fmt.Print(" ")
		}
		if x > 1 {
			fmt.Println("B")
		} else {
			fmt.Println()
		}
	}

	// Draw the last row (bottom border)
	if y > 1 {
		fmt.Print("C")
		for i := 1; i < x-1; i++ {
			fmt.Print("B")
		}
		if x > 1 {
			fmt.Println("C")
		} else {
			fmt.Println()
		}
	}
}
`

var quadDSrc = `
package main
import "fmt"
import "os"
import "strconv"
func main(){
cols,_ := strconv.Atoi(os.Args[1])
rows,_ := strconv.Atoi(os.Args[2])
QuadD(int(cols),int(rows))
}
func QuadD(x, y int) {
	// Handle cases where x or y is non-positive
	if x <= 0 || y <= 0 {
		return
	}

	// Draw the first row (top border)
	if y > 0 {
		fmt.Print("A")
		for i := 1; i < x-1; i++ {
			fmt.Print("B")
		}
		if x > 1 {
			fmt.Println("C")
		} else {
			fmt.Println()
		}
	}

	// Draw the middle rows (vertical borders with spaces in between)
	for j := 1; j < y-1; j++ {
		fmt.Print("B")
		for i := 1; i < x-1; i++ {
			fmt.Print(" ")
		}
		if x > 1 {
			fmt.Println("B")
		} else {
			fmt.Println()
		}
	}

	// Draw the last row (bottom border)
	if y > 1 {
		fmt.Print("A")
		for i := 1; i < x-1; i++ {
			fmt.Print("B")
		}
		if x > 1 {
			fmt.Println("C")
		} else {
			fmt.Println()
		}
	}
}
`

var quadESrc = `
package main
import "fmt"
import "os"
import "strconv"
func main(){
cols,_ := strconv.Atoi(os.Args[1])
rows,_ := strconv.Atoi(os.Args[2])
QuadE(int(cols),int(rows))
}
func QuadE(x, y int) {
	// Handle cases where x or y is non-positive
	if x <= 0 || y <= 0 {
		return
	}
	// Draw the first row (top border)
	if y > 0 {
		fmt.Print("A")
		for i := 1; i < x-1; i++ {
			fmt.Print("B")
		}
		if x > 1 {
			fmt.Println("C")
		} else {
			fmt.Println()
		}
	}
	// Draw the middle rows (vertical borders with spaces in between)
	for j := 1; j < y-1; j++ {
		fmt.Print("B")
		for i := 1; i < x-1; i++ {
			fmt.Print(" ")
		}
		if x > 1 {
			fmt.Println("B")
		} else {
			fmt.Println()
		}
	}
	// Draw the last row (bottom border)
	if y > 1 {
		fmt.Print("C")
		for i := 1; i < x-1; i++ {
			fmt.Print("B")
		}
		if x > 1 {
			fmt.Println("A")
		} else {
			fmt.Println()
		}
	}
}
`

// Map function source to corresponding executable name
var functions = map[string]string{
	"quadchecker": quadCheckerSrc,
	"quadA":       quadASrc,
	"quadB":       quadBSrc,
	"quadC":       quadCSrc,
	"quadD":       quadDSrc,
	"quadE":       quadESrc,
}

func cleanup() {
	if _, err := os.Stat("go.mod"); !os.IsNotExist(err) {
		os.Remove("go.mod") // Remove go.mod if it exists
	}
	fmt.Println("Cleaned up temporary files and go.mod")
}

func main() {
	// Iterate through each function, write to file, and build executable
	for name, src := range functions {
		// Create a temporary file for the function source code
		tmpFile, err := os.CreateTemp("", name+"*.go")
		if err != nil {
			fmt.Printf("Failed to create temp file for %s: %v\n", name, err)
			continue
		}
		defer os.Remove(tmpFile.Name()) // Remove the file later
		// Write the function source code into the file
		_, err = tmpFile.WriteString(src)
		if err != nil {
			fmt.Printf("Failed to write to temp file for %s: %v\n", name, err)
			continue
		}
		tmpFile.Close() // Close the file after writing
		// Compile the temporary Go file into an executable
		buildCmd := exec.Command("go", "build", "-o", name, tmpFile.Name())
		buildCmd.Stdout = os.Stdout
		buildCmd.Stderr = os.Stderr
		if err := buildCmd.Run(); err != nil {
			fmt.Printf("Failed to build executable for %s: %v\n", name, err)
			continue
		}
		fmt.Printf("Successfully built %s executable\n", name)
		// Set the executable permissions
		if err := os.Chmod(name, 0o755); err != nil {
			fmt.Printf("Failed to set permissions for %s: %v\n", name, err)
		}
	}

	// Cleanup after execution
	cleanup()
}
