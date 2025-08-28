This Go program demonstrates the use of **interfaces**, **structs**, and **polymorphism** through the calculation of areas for different geometric shapes.

---

## 🔄 **Data Flow & Execution Overview**

### 1. **Importing packages**

```go
import "fmt"
import "math"
```

* `fmt`: For printing output to the console.
* `math`: Provides mathematical constants and functions, like `math.Pi`.

---

### 2. **Interface Definition**

```go
type Shape interface {
	Area() float64
}
```

* `Shape` is an **interface** with a single method:

  * `Area()` which returns a `float64`.

👉 Any type that has a method `Area() float64` **automatically implements** this interface.

---

### 3. **Struct Definitions**

```go
type Rectangle struct {
	width, height float64
}
type Circle struct {
	radious float64
}
```

* `Rectangle` and `Circle` are **concrete types** (structs) representing shapes.
* They have fields needed to calculate their areas:

  * Rectangle: `width`, `height`
  * Circle: `radious` (misspelled, should be `radius` for clarity)

---

### 4. **Implementing the Interface Methods**

Both structs implement the `Area()` method, which makes them satisfy the `Shape` interface.

#### Rectangle:

```go
func (r Rectangle) Area() float64 {
	return r.width * r.height
}
```

* Calculates area: `width × height`

#### Circle:

```go
func (c Circle) Area() float64 {
	return math.Pi * c.radious * c.radious
}
```

* Calculates area: `π × radius²`

---

### 5. **Function Using Interface**

```go
func caculateArea(s Shape) float64 {
	return s.Area()
}
```

* This function **accepts any type that implements the `Shape` interface**.
* It calls the `Area()` method polymorphically:

  * It doesn't care whether it's a Rectangle or Circle.

*Note: There's a typo — `caculateArea` should be `calculateArea`.*

---

### 6. **Main Function Execution**

```go
func main() {
   rect := Rectangle{width: 2, height: 2}
   circle := Circle{radious: 2}
   fmt.Println("Rectangle Area:", caculateArea(rect))
   fmt.Println("Circle Area:", caculateArea(circle))
}
```

#### Step-by-step:

1. A `Rectangle` with width 2 and height 2 is created.
2. A `Circle` with radius 2 is created.
3. Both are passed to `caculateArea`, which treats them as `Shape` and calls their respective `Area()` methods.
4. The results are printed.

---

### ✅ **Sample Output**

```
Rectangle Area: 4
Circle Area: 12.566370614359172
```

---

## 🧠 Key Concepts

| Concept          | Description                                                                   |
| ---------------- | ----------------------------------------------------------------------------- |
| **Interface**    | A contract: any type that implements its methods is considered to fulfill it. |
| **Struct**       | Concrete data types with fields.                                              |
| **Method**       | A function with a receiver type (e.g., `func (r Rectangle) Area()` )          |
| **Polymorphism** | Calling methods on interface types without knowing the concrete type.         |

---

## ✍️ Suggested Improvements

* Fix typos:

  * Rename `caculateArea` → `calculateArea`
  * Rename `radious` → `radius`
* Use more shapes (Triangle, Square) to demonstrate extensibility.

Would you like me to show a version with those improvements?
