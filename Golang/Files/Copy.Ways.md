```go
//Read and Write in the streaming fashion

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	sourceFile, err := os.Open("testing/example3.txt")
	if err != nil {
		panic(err)
	}
	defer sourceFile.Close()

	destFile, err := os.Create("testing/example2.txt")
	if err != nil {
		panic(err)
	}
	defer destFile.Close()

	reader := bufio.NewReader(sourceFile)
	writer := bufio.NewWriter(destFile)

	for {
		b, err := reader.ReadByte()
		if err != nil {
			if err.Error() != "EOF" {
				panic(err)
			}
			break
		}
		e := writer.WriteByte(b)
		if e != nil {
			panic(e)
		}
	}
	writer.Flush()
	fmt.Println("Written to new file successfully")
}


```
---
This Go program **reads a file byte by byte** using a streaming approach and writes its contents to another file. It's an efficient way of copying data without loading the entire file into memory, which is especially useful for **large files**.

---

## ✅ **Purpose of the Program**

To **copy the contents of one file to another** using **streaming I/O**, meaning it reads and writes small parts of the file (bytes) one at a time.

---

## 🗂️ **Example Files**

### Source File: `testing/example3.txt`

**Before running the program:**

```
Hello World!
This is a test file.
Streaming is cool.
```

### Destination File: `testing/example2.txt`

**Before running the program:**

```
(empty or does not exist)
```

**After running the program:**

```
Hello World!
This is a test file.
Streaming is cool.
```

---

## 🧠 **How the Program Works: Step-by-Step**

### 1. **Opening the Source File**

```go
sourceFile, err := os.Open("testing/example3.txt")
```

* `os.Open`: Opens the source file in read-only mode.
* `err`: Checks for errors like "file not found".
* `defer sourceFile.Close()`: Ensures the file is closed when `main()` exits.

---

### 2. **Creating the Destination File**

```go
destFile, err := os.Create("testing/example2.txt")
```

* `os.Create`: Creates or truncates the destination file. If it exists, it will be **emptied** first.
* `defer destFile.Close()`: Closes the destination file when done.

---

### 3. **Creating Buffered Reader & Writer**

```go
reader := bufio.NewReader(sourceFile)
writer := bufio.NewWriter(destFile)
```

* `bufio.NewReader`: Wraps the source file with a buffered reader for efficient reading.
* `bufio.NewWriter`: Wraps the destination file with a buffered writer to reduce system calls (better performance).

---

### 4. **Streaming Data Byte-by-Byte**

```go
for {
    b, err := reader.ReadByte()
    if err != nil {
        if err.Error() != "EOF" {
            panic(err)
        }
        break
    }
    e := writer.WriteByte(b)
    if e != nil {
        panic(e)
    }
}
```

* `reader.ReadByte()`:

  * Reads **1 byte at a time** from the source.
  * Returns `EOF` error when it reaches the end.
* `writer.WriteByte(b)`:

  * Writes that byte to the destination buffer.
* `break`: Loop stops when the file ends (`EOF`).
* `panic(err)`: Crashes the program with a stack trace if an unexpected error happens.

---

### 5. **Flushing the Writer**

```go
writer.Flush()
```

* `Flush()` writes any buffered data to the destination file. **Critical** — without this, the data might not be saved.

---

### 6. **Success Message**

```go
fmt.Println("Written to new file successfully")
```

---

## 📊 **Data Flow Overview**

```plaintext
testing/example3.txt
        |
        v
[Buffered Reader (ReadByte)]
        |
        v
[Buffered Writer (WriteByte)]
        |
        v
testing/example2.txt
```

---

## 🧪 **Why Streaming Instead of Reading All at Once?**

* ✅ Better for **large files**.
* ✅ Reduces memory usage.
* ✅ More **responsive** and **scalable**.
* ❌ Slower for very small files (due to overhead), but the difference is minimal.

---

## 🛠️ **Used Functions & Packages**

| Function/Method      | Description                                                   |
| -------------------- | ------------------------------------------------------------- |
| `os.Open`            | Opens a file for reading                                      |
| `os.Create`          | Creates a new file for writing (overwrites if exists)         |
| `bufio.NewReader`    | Creates a buffered reader for efficient reading               |
| `bufio.NewWriter`    | Creates a buffered writer for efficient writing               |
| `reader.ReadByte()`  | Reads one byte at a time from the buffer                      |
| `writer.WriteByte()` | Writes one byte at a time to the buffer                       |
| `writer.Flush()`     | Flushes the buffer to the actual file (writes it all at once) |
| `defer file.Close()` | Closes the file when `main()` finishes execution              |
| `panic(err)`         | Stops execution if a critical error occurs                    |

---

## 🧾 Summary

This program efficiently copies a file byte-by-byte using buffered I/O, suitable for both small and large files. The use of streaming avoids memory overload and ensures the program can handle very large files gracefully.

Would you like a version that reads/writes **line-by-line** or **all at once** for comparison?




---

If you want to copy the content of one file to another using **less code** in Go, you can use the built-in `io.Copy()` function — it's simple, efficient, and handles large files internally with buffering.

---

## ✅ **Minimal Version: Copy File in Go**

```go
package main

import (
	"io"
	"os"
	"log"
)

func main() {
	src, err := os.Open("testing/example3.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer src.Close()

	dst, err := os.Create("testing/example2.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer dst.Close()

	_, err = io.Copy(dst, src)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("File copied successfully")
}
```

---

## 🧠 What This Does

| Function/Package     | Purpose                                                 |
| -------------------- | ------------------------------------------------------- |
| `os.Open()`          | Opens the source file                                   |
| `os.Create()`        | Creates or truncates the destination file               |
| `io.Copy(dst, src)`  | Copies from source to destination using internal buffer |
| `defer file.Close()` | Ensures files are closed after use                      |
| `log.Fatal()`        | Exits with error if something goes wrong                |

---

## 🔁 How `io.Copy()` Works Internally

* Uses a buffer (32 KB by default) to stream data from the source reader to the destination writer.
* Efficient for both small and large files.
* It avoids writing manual loops.

---

## 📋 Before & After Example

### `example3.txt` (source):

```
Hello from the source file!
Line 2
Line 3
```

### `example2.txt` (destination):

**Before**: Empty or doesn't exist
**After running**:

```
Hello from the source file!
Line 2
Line 3
```

---

## ✅ When to Use `io.Copy`

* Best for general-purpose file copying
* Good balance of performance and code simplicity

Let me know if you want a version with error handling using `errors.Is()` or a progress bar for large files.


---
Great — here's a version of the `io.Copy()` approach with **detailed and clean error handling** using `errors.Is()` and Go’s idiomatic practices.

---

## ✅ **Go Program: File Copy with Proper Error Handling**

```go
package main

import (
	"errors"
	"fmt"
	"io"
	"os"
)

func main() {
	// Open the source file
	src, err := os.Open("testing/example3.txt")
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			fmt.Println("Source file does not exist.")
		} else {
			fmt.Printf("Error opening source file: %v\n", err)
		}
		return
	}
	defer src.Close()

	// Create the destination file
	dst, err := os.Create("testing/example2.txt")
	if err != nil {
		fmt.Printf("Error creating destination file: %v\n", err)
		return
	}
	defer dst.Close()

	// Copy contents from source to destination
	bytesCopied, err := io.Copy(dst, src)
	if err != nil {
		fmt.Printf("Error copying file: %v\n", err)
		return
	}

	fmt.Printf("Copied %d bytes successfully.\n", bytesCopied)
}
```

---

## 🧠 **Explanation of Error Handling**

| Code                             | Meaning                                       |
| -------------------------------- | --------------------------------------------- |
| `errors.Is(err, os.ErrNotExist)` | Checks if the error is "file not found"       |
| `fmt.Printf(...)`                | Used for formatted output of errors           |
| `return`                         | Exits early on error                          |
| `defer file.Close()`             | Ensures file handles are closed even on error |

---

## 🔄 Example Run Output

### ✅ Success

```
Copied 75 bytes successfully.
```

### ❌ Source File Missing

```
Source file does not exist.
```

### ❌ Permission Issue or Disk Full

```
Error creating destination file: permission denied
```

---

## 🧰 Optional: Check If Destination Already Exists (Prevent Overwriting)

If you want to **prevent overwriting** the destination file:

```go
if _, err := os.Stat("testing/example2.txt"); err == nil {
	fmt.Println("Destination file already exists. Aborting.")
	return
}
```

Put that check before the `os.Create()` call.

---

Let me know if you want:

* A version that shows a **progress bar** for large files.
* A **command-line version** that accepts file names as arguments.

