Certainly! Let's break down your Go program step by step and explain all the concepts involved.

---

## **Imports**

```go
import (
	"encoding/json"
	"fmt"
)
```

- `fmt`: Standard package for formatted I/O (printing to the console).
- `encoding/json`: Standard package for encoding and decoding JSON data.

---

## **Struct Definition**

```go
type Person struct {
	Name    string `json:"name"`
	Age     int    `json:"age"`
	IsAdult bool   `json:"is_adult"`
}
```

- You define a struct `Person` with three fields: `Name` (string), `Age` (int), and `IsAdult` (bool).
- The tags (e.g., `json:"name"`) specify how the field should be named in the JSON output/input.  
  - This means when you encode a `Person` to JSON, the field `Name` will be output as `"name"` in JSON.

---

## **main() Function**

```go
func main() {
	fmt.Println("We are learning JSON")
	person := Person{Name: "John", Age: 34, IsAdult: true}
	fmt.Println("person Data is : ", person)
```

- Prints a message to the console.
- Creates an instance of `Person` named `person` with name "John", age 34, and is adult `true`.
- Prints the `person` struct directly (will show as `{John 34 true}`).

---

### **Encoding (Marshalling) to JSON**

```go
	//convert person into JSON Encoding(Marshalling)
	jsonData, err := json.Marshal(person)
	if err != nil {
		fmt.Println("Error marshalling", err)
		return
	}
	fmt.Println("person Data is :", string(jsonData))
```

- `json.Marshal(person)` encodes the `person` struct into JSON format (returns a byte slice and error).
- Checks for errors. If any error occurs during marshalling, prints the error and exits.
- Converts the resulting byte slice to a string and prints the JSON representation of `person`.

**Sample Output:**
```
person Data is : {"name":"John","age":34,"is_adult":true}
```

---

### **Decoding (Unmarshalling) from JSON**

```go
	//Decoding (Unmarshalling)
	var personData Person
	err = json.Unmarshal(jsonData, &personData)
	if err != nil {
		fmt.Println("Error unmarshalling", err)
	}
	fmt.Println("person Data is : ", personData)
}
```

- Declares a new variable `personData` of type `Person`.
- Uses `json.Unmarshal(jsonData, &personData)` to decode the JSON back into a Go struct.
  - `json.Unmarshal` takes the JSON byte slice and a pointer to the variable that will hold the decoded data.
- Checks for errors during unmarshalling.
- Prints the resulting `personData` struct, which should have the same values as the original `person`.

**Sample Output:**
```
person Data is :  {John 34 true}
```

---

## **Summary and Flow**

1. **Creates a struct value.**
2. **Prints struct.**
3. **Converts struct to JSON (encoding/marshalling).**
4. **Prints JSON.**
5. **Converts JSON back to struct (decoding/unmarshalling).**
6. **Prints recovered struct.**

## **What does it demonstrate?**

- **How to encode a Go struct into JSON** (using `json.Marshal`)
- **How to decode JSON back into a Go struct** (using `json.Unmarshal`)
- **How to use struct field tags for custom JSON field names**
- **Error handling during marshalling/unmarshalling**

## **Real Output Example**

```
We are learning JSON
person Data is :  {John 34 true}
person Data is : {"name":"John","age":34,"is_adult":true}
person Data is :  {John 34 true}
```

## **Key Concepts**

- **Marshalling**: Converting Go struct → JSON.
- **Unmarshalling**: Converting JSON → Go struct.
- **Struct tags**: Control how fields appear in JSON.
- **Error handling**: Always check for errors in encoding/decoding operations.

---

**In summary:**  
This program shows the full round-trip of serializing a Go struct to JSON and then deserializing it back into a struct, with practical output and error handling.
