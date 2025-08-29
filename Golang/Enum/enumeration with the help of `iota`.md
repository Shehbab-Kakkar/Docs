This Go program demonstrates how to use **enumeration types** with the help of `iota`, which is a Go keyword that simplifies the declaration of sequentially increasing constants.

---

### 🔍 **Overview of the Program**

```go
package main

import "fmt"

// Define a custom type for order statuses
type OrderStatus int

// Define constants using iota
const (
	Received OrderStatus = iota
	confirmed
	Prepared
	Delivered
)

func changeOrderStatus(status OrderStatus) {
	fmt.Println("changing order status to", status)
}

func main() {
	changeOrderStatus(Received)
}
```

---

### 🧠 **What’s Happening Here?**

#### 1. **Custom Type Declaration**

```go
type OrderStatus int
```

This line defines a **new type** called `OrderStatus`, which is based on the built-in `int` type. This helps in creating more **type-safe code**, especially when dealing with specific sets of values like order statuses.

---

#### 2. **Constant Declaration with `iota`**

```go
const (
	Received OrderStatus = iota
	confirmed
	Prepared
	Delivered
)
```

* `iota` is a Go predeclared identifier used in constant declarations.
* It resets to `0` each time a new `const` block starts and increments by 1 for each line.

So here:

| Name      | Value | Explanation           |
| --------- | ----- | --------------------- |
| Received  | 0     | First value, iota = 0 |
| confirmed | 1     | iota increments to 1  |
| Prepared  | 2     | iota increments to 2  |
| Delivered | 3     | iota increments to 3  |

Even though we didn't explicitly assign numbers, **`iota` handles the numbering**.

---

#### 3. **Function `changeOrderStatus`**

```go
func changeOrderStatus(status OrderStatus) {
	fmt.Println("changing order status to", status)
}
```

* Accepts an `OrderStatus` value as a parameter.
* Prints the numeric value of that status.

---

#### 4. **Calling the Function in `main`**

```go
func main() {
	changeOrderStatus(Received)
}
```

* Calls the function with `Received`, which has a value of `0`.
* Output:

```
changing order status to 0
```

---

### 💡 How to Make Output More Readable?

Currently, it just prints the numeric value (`0`, `1`, etc.). To improve this, you could implement the `String()` method for `OrderStatus`:

```go
func (s OrderStatus) String() string {
	switch s {
	case Received:
		return "Received"
	case confirmed:
		return "Confirmed"
	case Prepared:
		return "Prepared"
	case Delivered:
		return "Delivered"
	default:
		return "Unknown"
	}
}
```

Now `fmt.Println("changing order status to", status)` will print:

```
changing order status to Received
```

---

### ✅ Summary

* `iota` helps assign auto-incrementing values to constants.
* Constants like `Received`, `confirmed`, etc., are of custom type `OrderStatus`.
* The function takes this custom type and prints its value.
* You can make it more user-friendly by adding a `String()` method to convert int values to readable strings.

Would you like to see how to use this in a real-world order tracking system?
