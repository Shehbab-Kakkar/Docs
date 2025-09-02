```go
//Read and Write in the straming fashion

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



