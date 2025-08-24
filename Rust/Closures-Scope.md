
```rust
fn main() {
    let x = 10;
    let add = |y| x + x;
    println!("{}", add(23));
}
```

### What’s going on here:

* `x = 10`
* `y = 23` (passed as an argument)
* The closure is defined as:

  ```rust
  |y| x + x
  ```

Here’s the **key point**:

> Even though `y` is accepted as a parameter, it's **not used** inside the closure body.

So Rust compiles the closure as:

```rust
let add = |y| {
    // y is unused
    x + x
};
```

Thus, `x + x = 10 + 10 = 20`, and `y = 23` is **ignored**.

---

## ✅ Why Does Rust Allow This?

Rust **does not force you to use closure parameters**. Just like regular functions:

```rust
fn foo(x: i32) -> i32 {
    5 // `x` is not used, and that's valid
}
```

You declared `y`, but unless you actually write `x + y` or something using `y`, it just doesn't affect the result.

---

## ✅ How to Use `y`?

Simply change the closure to use `y`:

```rust
let add = |y| x + y;
```

Now both `x` and `y` are used:

* `x = 10`
* `y = 23`
* `x + y = 33`

### Output:

```rust
println!("{}", add(23)); // ✅ prints 33
```

---

## 🧩 TL;DR

| Symbol | Value        | Used?                               |
| ------ | ------------ | ----------------------------------- |
| `x`    | 10           | ✅ yes                               |
| `y`    | 23           | ❌ no — not used in the closure body |
| Result | `x + x = 20` | ✅ that's why you got 20             |

---

If you want, I can also show how the compiler can warn you about unused parameters — just let me know!
