Great question!

In Go, **strings are UTF-8 encoded**, meaning that some characters (like emojis or accented letters) can take **more than 1 byte**.

Let’s see an example using a character that takes more than 1 byte — say, **`é`** or an emoji like **`😊`**.

---

### ✅ Example Program (Multi-byte characters):

```go
package main

import "fmt"

func main() {
	for i, ch := range "Golang😊é" {
		fmt.Printf("Index: %d, Character: %c, Unicode: %U\n", i, ch, ch)
	}
}
```

---

### 🔍 Explanation:

* `"Golang😊é"` contains:

  * ASCII characters: `G`, `o`, `l`, `a`, `n`, `g` → each 1 byte.
  * Emoji: `😊` → takes **4 bytes**.
  * Accented character: `é` → takes **2 bytes**.

---

### 🖨 Output:

```
Index: 0, Character: G, Unicode: U+0047
Index: 1, Character: o, Unicode: U+006F
Index: 2, Character: l, Unicode: U+006C
Index: 3, Character: a, Unicode: U+0061
Index: 4, Character: n, Unicode: U+006E
Index: 5, Character: g, Unicode: U+0067
Index: 6, Character: 😊, Unicode: U+1F60A
Index: 10, Character: é, Unicode: U+00E9
```

---

### 📌 Key Points:

* Indexes `0–5` are each 1 byte apart (ASCII).
* `😊` starts at byte index `6` and takes **4 bytes**, so the next character (`é`) starts at index `10`.
* `é` takes **2 bytes**.

---

### ✅ How to Check Byte Lengths

To check the size of each character in bytes:

```go
package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	s := "Golang😊é"
	for i := 0; i < len(s); {
		r, size := utf8.DecodeRuneInString(s[i:])
		fmt.Printf("Char: %c, Bytes: %d\n", r, size)
		i += size
	}
}
```

---

Let me know if you want a visual breakdown of the actual byte sequences too!
