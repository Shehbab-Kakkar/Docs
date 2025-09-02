Sure! Here's a **step-by-step guide** with an **example** to help you:

---

## ✅ Goal

You want to:

1. Create a Go package.
2. Push it to GitHub.
3. Import and use it in another Go project.

---

## 📁 Example Package: `mymath`

Let’s say you want to create a simple math package that provides a function to add two numbers.

---

### Step 1: Create Your Go Package

```bash
mkdir mymath
cd mymath
go mod init github.com/yourusername/mymath
```

#### File: `mymath/mymath.go`

```go
package mymath

// Add returns the sum of two integers.
func Add(a, b int) int {
    return a + b
}
```

#### File: `go.mod`

This will be created automatically when you run `go mod init github.com/yourusername/mymath`.

---

### Step 2: Push the Package to GitHub

1. Create a repo on GitHub: `https://github.com/yourusername/mymath`
2. Push your code:

```bash
git init
git add .
git commit -m "Initial commit"
git branch -M main
git remote add origin https://github.com/yourusername/mymath.git
git push -u origin main
```

---

## ✅ Now the package is public and importable!

---

### Step 3: Use Your Package in Another Project

Now create another Go project to use the package:

```bash
mkdir myapp
cd myapp
go mod init myapp
```

#### File: `main.go`

```go
package main

import (
    "fmt"
    "github.com/yourusername/mymath"
)

func main() {
    result := mymath.Add(5, 3)
    fmt.Println("Result:", result)
}
```

#### Then run:

```bash
go mod tidy
go run main.go
```

> Output:

```
Result: 8
```

---

## ✅ Done!

---

## 📝 Tips:

* Make sure your GitHub repo is **public** (or setup Go private module access if it’s private).
* You can version your package using Git tags like `v1.0.0`.

```bash
git tag v1.0.0
git push origin v1.0.0
```

Then you can import specific versions in other projects.

---

Would you like me to generate the files for you or give you a GitHub-ready zip?
