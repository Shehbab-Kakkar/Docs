Here's a simple Go program that **creates a file** (or truncates it if it already exists), and **writes data to it**:

---

## ✅ Go Program: Create a File and Write Data

```go
package main

import (
	"fmt"
	"os"
)

func main() {
	// Step 1: Create or truncate the file
	file, err := os.Create("example.txt")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close() // Always close the file when done

	// Step 2: Write data to the file
	data := "Hello, this is some sample data written to the file.\n"
	_, err = file.WriteString(data)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}

	fmt.Println("File created and data written successfully.")
}
```

---

## 🔍 What This Code Does:

* `os.Create("example.txt")`:

  * Creates a file named `example.txt` in the current directory.
  * If the file already exists, it will be **truncated** (emptied).
* `file.WriteString(data)`:

  * Writes the string to the file.
* `defer file.Close()`:

  * Ensures the file is properly closed at the end, even if there's an error.

---

## 📂 Output:

It will create a file `example.txt` in your current directory with this content:

```
Hello, this is some sample data written to the file.
```

---

Let me know if you want to:

* Append instead of overwriting
* Create the file in a specific path
* Write binary data or JSON
* Read the file afterward

I can modify the example for those too.
