---

### ✅ Code:

```go
package main

import "fmt"

func main() {
	for x, y := range "Golang" {
		fmt.Println(x, y)
	}
}
```

---

### 🔍 Explanation:

#### 1. `package main`

This is the entry point for any standalone Go program. It defines the main package.

#### 2. `import "fmt"`

Imports the `fmt` package, which provides I/O formatting functions like `Println`.

#### 3. `func main() { ... }`

Defines the `main()` function — the execution starts here.

#### 4. `for x, y := range "Golang" { ... }`

This is a **`range` loop** over a string. Here's what happens:

* `"Golang"` is a string literal.
* When you use `range` on a string in Go, it **iterates over Unicode code points (runes)**.
* For each character:

  * `x` is the **index** (position) in the string (in bytes).
  * `y` is the **rune** (i.e., Unicode code point) at that position.

So, this loop goes over each character in `"Golang"` and prints:

* The index of the character (in bytes).
* The rune value (as an `int`).

---

### 🖨 Output:

The output of this program will be:

```
0 71
1 111
2 108
3 97
4 110
5 103
```

These numbers are the ASCII (Unicode) values of the characters:

| Index | Character | Unicode (int) |
| ----- | --------- | ------------- |
| 0     | G         | 71            |
| 1     | o         | 111           |
| 2     | l         | 108           |
| 3     | a         | 97            |
| 4     | n         | 110           |
| 5     | g         | 103           |

---

### ✅ Bonus Tip – Print as characters:

If you want to see the actual characters instead of Unicode values, change the print line to:

```go
fmt.Println(x, string(y))
```

Then the output will be:

```
0 G
1 o
2 l
3 a
4 n
5 g
```

Let me know if you want a version that handles multi-byte Unicode characters (like emojis or accented letters).
