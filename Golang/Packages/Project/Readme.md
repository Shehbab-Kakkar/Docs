Here’s how you should proceed to upload your Go project to GitHub under your username Shehbab-Kakkar:

## 1. Update Your Import Paths

Change your import paths in all Go files from:
```go
"github.com/myaccount/podcast/..."
```
to:
```go
"github.com/Shehbab-Kakkar/podcast/..."
```

## 2. Project Structure & Files

Here is the complete file structure and content with the correct import paths:

```go name=main.go
package main

import (
	"fmt"

	"github.com/Shehbab-Kakkar/podcast/auth"
	"github.com/Shehbab-Kakkar/podcast/user"
)

func main() {
	auth.LoginWithCredentials("shehbab", "secret")
	session := auth.GetSession()
	fmt.Println("session", session)
	user := user.User{
		Email: "user@gmail.com",
		Name:  "John",
	}
	fmt.Println(user.Email, user.Name)
}
```

```go name=auth/auth.go
package auth

// Public function as function start with Capital letter
func GetSession() string {
	return extractSession()
}

// Private function as function start with small letter
func extractSession() string {
	return "loggenIn"
}
```

```go name=auth/login.go
package auth

import "fmt"

func LoginWithCredentials(username string, password string) {
	fmt.Println("login user using", username, password)
}
```

```go name=user/user.go
package user

type User struct {
	Email string
	Name  string
}
```

## 3. Initialize Your Module

In your project root, run:
```bash
go mod init github.com/Shehbab-Kakkar/podcast
go mod tidy
```

## 4. Test Locally

Run:
```bash
go run main.go
```
You should see:
```
login user using shehbab secret
session loggenIn
user@gmail.com John
```

## 5. Push to GitHub

1. Initialize git if you haven't:
    ```bash
    git init
    git add .
    git commit -m "Initial podcast project"
    ```
2. Create an empty repo on GitHub named podcast under Shehbab-Kakkar.
3. Add the remote and push:
    ```bash
    git remote add origin https://github.com/Shehbab-Kakkar/podcast.git
    git branch -M main
    git push -u origin main
    ```

---

Would you like a sample README.md or a .gitignore for your project as well?
