
````markdown
# Rust HashMap Example

## Code Explanation

This Rust program demonstrates the use of `HashMap` and some key methods like `get()`, `entry()`, `or_insert()`, and `and_modify()`.

### Code:

```rust
use std::collections::HashMap;

fn main() {
    // Create a mutable HashMap with key type &str and value type i32
    let mut marks: HashMap<&str, i32> = HashMap::new();

    // Insert values into the HashMap
    marks.insert("Ram", 32);
    marks.insert("Shyam", 45);
    marks.insert("Geeta", 38);
    marks.insert("Sita", 50);

    // Print all key-value pairs
    println!("Marks: {:?}", marks);

    // Accessing a value (get returns an Option)
    if let Some(mark) = marks.get("Ram") {
        println!("Ram's marks: {}", mark);
    }

    // Updating the value for a specific key
    marks.insert("Ram", 60);  // This replaces the old value (32)

    // Add or update using entry API
    marks.entry("Laxman").or_insert(40); // If "Laxman" doesn't exist, insert 40
    marks.entry("Sita").and_modify(|mark| *mark += 5); // Increase Sita's marks by 5

    // Final state
    println!("Updated Marks: {:?}", marks);
}
````

### Explanation

#### `if let Some(mark) = marks.get("Ram")`

* The `get()` method retrieves a value from the `HashMap`. It returns an `Option<&i32>`.

  * `Some(&32)` if `"Ram"` exists in the `HashMap`.
  * `None` if `"Ram"` does not exist.
* The `if let` syntax is used to only execute the code when `"Ram"` is found.

**Output:**

```text
Ram's marks: 32
```

#### `marks.entry("Laxman").or_insert(40)`

* The `entry()` method checks if a key exists. If `"Laxman"` does not exist, it inserts `("Laxman", 40)`.
* `or_insert(40)` inserts the value `40` if `"Laxman"` is not present.

**Result:**

* If `"Laxman"` was not in the map, it will now be added with the value `40`.

#### `marks.entry("Sita").and_modify(|mark| *mark += 5)`

* The `and_modify()` method is called if `"Sita"` exists. The closure modifies the value of `Sita` by adding `5`.

  * If `"Sita"` had a value of `50`, it becomes `55`.

### Final `marks` HashMap:

```rust
{
    "Ram": 32,
    "Shyam": 45,
    "Geeta": 38,
    "Sita": 55,
    "Laxman": 40
}
```

### Data Flow:

1. **Before the operations:**

   ```rust
   marks = {
       "Ram": 32,
       "Shyam": 45,
       "Geeta": 38,
       "Sita": 50
   }
   ```

2. **After `if let Some(mark) = marks.get("Ram")`:**

   ```text
   Ram's marks: 32
   ```

3. **After `marks.entry("Laxman").or_insert(40)`:**

   * `"Laxman"` is inserted with the value `40`.

   ```rust
   marks = {
       "Ram": 32,
       "Shyam": 45,
       "Geeta": 38,
       "Sita": 50,
       "Laxman": 40
   }
   ```

4. **After `marks.entry("Sita").and_modify(|mark| *mark += 5)`:**

   * `"Sita"`'s marks are updated from `50` to `55`.

   ```rust
   marks = {
       "Ram": 32,
       "Shyam": 45,
       "Geeta": 38,
       "Sita": 55,
       "Laxman": 40
   }
   ```

---

### Conclusion

This program demonstrates how to efficiently manipulate a `HashMap` in Rust using methods like `get()`, `entry()`, `or_insert()`, and `and_modify()`. These methods provide robust handling of key-value pairs, whether you’re accessing, modifying, or inserting data.

---
