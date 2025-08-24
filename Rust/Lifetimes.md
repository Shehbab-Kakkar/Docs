Err:

Compiling playground v0.0.1 (/playground)
error[E0597]: `x` does not live long enough
 --> src/main.rs:5:14
  |
4 |         let x = 5;
  |             - binding `x` declared here
5 |          r = &x;
  |              ^^ borrowed value does not live long enough
6 |        //  println!("{}",r);
7 |     }
  |     - `x` dropped here while still borrowed
8 |     println!("Print r : {}",r);
  |                             - borrow later used here

For more information about this error, try `rustc --explain E0597`.
error: could not compile `playground` (bin "playground") due to 1 previous error


You're seeing this error because Rust's borrow checker prevents *dangling references*. In your code:

```rust
fn main() {
    let r;
    {
        let x = 5;
        r = &x;
    }
    println!("Print r : {}", r);
}
```

You're attempting to assign `r` a reference to `x`, which is declared **inside** the inner block. As soon as that block ends, `x` is dropped—yet `r` would still point to it. Rust disallows this to ensure safety:

```
error[E0597]: `x` does not live long enough
```

The **Rust Book** illustrates this exact scenario, highlighting how lifetimes overlap and why it's invalid:

* `x` has a short inner scope (`'b`),
* `r` has a longer outer scope (`'a`),

but `r` cannot reference a value (`x`) that ends before it does.([Massachusetts Institute of Technology][1])

So what's going on conceptually?

* **`x` goes out of scope** at the end of the inner block.
* `r` would then refer to freed memory—creating a dangling reference.
* Rust… *won’t allow that*.

Even if you tried to compile this, Rust catches it at compile time:

```text
borrowed value does not live long enough...
`x` dropped here while still borrowed...
borrow later used here
```

This is **error E0597**, which explicitly flags this scenario.([Rust Documentation][2], [TypeError][3])

---

### How to Fix It

1. **Move `x` to the outer scope** so its lifetime matches or exceeds `r`'s:

   ```rust
   fn main() {
       let x = 5;
       let r = &x;
       println!("Print r : {}", r);
   }
   ```

2. Or restructure the code so that the reference doesn't outlive the value.

---

### In a Nutshell

* **Rust ensures memory safety** by preventing references to values that have already been dropped.
* You can't assign a reference that lives longer than the referred value.
* Error E0597 pops up when the borrowed value dies before usage.
* The fix is always: make the owner live as long as (or longer than) the reference.


Want to explore that next?

[1]: https://web.mit.edu/rust-lang_v1.25/arch/amd64_ubuntu1404/share/doc/rust/html/book/second-edition/ch10-03-lifetime-syntax.html?utm_source=chatgpt.com "Validating References with Lifetimes - The Rust Programming Language"
[2]: https://doc.rust-lang.org/error_codes/E0597.html?utm_source=chatgpt.com "E0597 - Error codes index"
[3]: https://www.typeerror.org/docs/rust/book/ch10-03-lifetime-syntax?utm_source=chatgpt.com "10.3. Validating References with Lifetimes - Rust Documentation - TypeError"
