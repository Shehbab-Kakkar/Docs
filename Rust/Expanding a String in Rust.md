# Expanding a String in Rust and Assigning the Result

## Original Code (Mutates In-Place, Returns Nothing)

```rust
fn main(){
    let mut s1: String = String::from("Hello I am here");
    add_to_string(&mut s1); // returns ()
    println!("Value of s1 {}", s1);
}

fn add_to_string(string_expand: &mut String){
     string_expand.push_str(", I will expand");
}
```

- `add_to_string` changes `s1` directly and returns nothing (`()`).
- Assigning its result to a variable gives you the unit value `()`:

```rust
let s2 = add_to_string(&mut s1);
println!("s2 = {:?}", s2); // prints "s2 = ()"
```

---

## Option 1: Return a Mutable Reference

Return a mutable reference so you can print or chain:

```rust
fn main() {
    let mut s1 = String::from("Hello I am here");
    let s2 = add_to_string(&mut s1); // s2 is &mut String
    println!("s1 = {}", s1);
    println!("s2 = {}", s2);
}

fn add_to_string(string_expand: &mut String) -> &mut String {
    string_expand.push_str(", I will expand");
    string_expand
}
```

- Both `s1` and `s2` point to the same, expanded string.

---

## Option 2: Return a New Owned String (Clone)

Return a new owned `String` (clone) if you want an independent copy:

```rust
fn main() {
    let mut s1 = String::from("Hello I am here");
    let s2 = add_to_string(&mut s1); // s2 is a new String
    println!("s1 = {}", s1);
    println!("s2 = {}", s2);
}

fn add_to_string(string_expand: &mut String) -> String {
    string_expand.push_str(", I will expand");
    string_expand.clone() // create and return a new String
}
```

- Now `s1` and `s2` are both `String` instances with the expanded value.

---

## Summary Table

| Goal                        | Suggested Return Type | Example                             |
|-----------------------------|----------------------|-------------------------------------|
| Just mutate in-place        | `()`                 | Your original code                  |
| Access the changed string   | `&mut String`        | Return the mutable reference        |
| Create a new owned result   | `String`             | Return `string_expand.clone()`      |

---
```
