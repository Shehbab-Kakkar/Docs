Let's analyze your original program and the `if` conditions to clarify what is being returned.

### Your code:
```go
func main() {
	fmt.Println("Web service")
	res, err := http.Get("https://jsonplaceholder.typicode.com/todos/1")
	if err != nil {
		fmt.Println("Error getting Get response", err)
		return
	}
	defer res.Body.Close()
	fmt.Printf("Type of response: %T\n", res)
	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println("Error ")
		return
	}
	fmt.Println("response", string(data))
}
```

---

## There are two `if` statements using `return`:

### 1. First `if`:
```go
if err != nil {
    fmt.Println("Error getting Get response", err)
    return
}
```
**What happens here?**
- If there is an error making the HTTP GET request (e.g., no internet, invalid URL), the code prints an error message and then executes `return`.
- This immediately stops execution of the `main()` function. Nothing else runs after this point in `main`.

### 2. Second `if`:
```go
if err != nil {
    fmt.Println("Error ")
    return
}
```
**What happens here?**
- After a successful HTTP GET, the code tries to read the response body.
- If there is an error while reading the response body (e.g., network error, invalid body), the code prints `"Error "` and again does `return`.
- This also immediately stops the execution of the `main()` function.

## **What is being returned?**
- In both cases, `return` is used in the `main()` function, which has a return type of `void` (no return value).
- So, **nothing is returned** (no value is sent back), but the function exits—meaning the program stops running at that point.

## **Summary Table**

| Condition                     | What prints                     | What returns          | What happens next                |
|-------------------------------|----------------------------------|----------------------|----------------------------------|
| `if err != nil` after GET     | Error message with details       | nothing (`void`)     | main() exits, program stops      |
| `if err != nil` after ReadAll | "Error "                        | nothing (`void`)     | main() exits, program stops      |

**In Go, `return` in main means exit the program at that point. No value is returned. Execution stops.**
