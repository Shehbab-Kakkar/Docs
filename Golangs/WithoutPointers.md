package main
import "fmt"
func changeNum(num int){
	num = 5
	fmt.Println("In changeNum", num)
}
func main(){
  num := 1
  changeNum(num)
  fmt.Println("After changeNum in main",num)
}


Absolutely! Let’s break down your Go program line by line and explain what’s happening, especially focusing on how Go handles function arguments and variable scope.

---

```go
package main
import "fmt"
```
- This is a standard way to start a Go program.
- `package main` defines the package as an executable program.
- `import "fmt"` imports the formatting package, used here for printing output.

---

### Function Definition

```go
func changeNum(num int){
	num = 5
	fmt.Println("In changeNum", num)
}
```
- This defines a function named `changeNum` that takes a single argument `num` of type `int`.
- Inside the function, you assign `num = 5`. This **changes the value of the local variable `num`** within the function scope.
- The function prints "In changeNum 5".

---

### main Function

```go
func main(){
  num := 1
  changeNum(num)
  fmt.Println("After changeNum in main", num)
}
```
- `main()` is the entry point for your program.
- `num := 1` declares and initializes `num` with the value `1`.
- `changeNum(num)` calls the function `changeNum`, passing the value of `num` (which is `1`) **by value**.
- After `changeNum` returns, `fmt.Println("After changeNum in main", num)` prints the value of `num` in the `main` function.

---

## Key Concept: Passing by Value

- In Go, when you pass a variable (like `num`) to a function, the function receives a **copy** of the value.
- Changes made to the parameter inside the function **do not affect** the original variable outside the function.

---

## Step-by-step Execution

1. **`main` starts.**
2. `num` is set to `1`.
3. `changeNum(num)` is called:
   - Inside `changeNum`, a new local variable `num` is created (copy of `main`'s `num`).
   - This local `num` is set to `5`.
   - Prints: `In changeNum 5`
   - The function ends, and the local `num` is discarded.
4. Back in `main`, `num` is still `1` (unchanged).
5. Prints: `After changeNum in main 1`

---

## Output

```
In changeNum 5
After changeNum in main 1
```

---

## What does this teach us?

- **Function arguments in Go are passed by value** (by default for basic types like `int`).
- If you want to modify the original variable from within a function, you need to pass a pointer (e.g., `*int`).  
- In this example, only the local copy inside `changeNum` is changed; the original value in `main` stays the same.

---

**Summary:**  
This program demonstrates how passing parameters by value works in Go. Changes to parameters inside functions do not affect the arguments outside those functions unless you use pointers.
