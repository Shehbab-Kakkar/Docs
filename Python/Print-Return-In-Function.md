# Why Does Python Print `None` After My Function?

This document explains a common beginner pitfall in Python:  
**Why do I see `None` printed after calling my function?**

---

## 📝 The Example

```python
def add(name='None'):
    print(f'Name of the person is {name}')

result = add("Jony")
print(result)
```

**Output:**
```
Name of the person is Jony
None
```

---

## ❓ Why Does `None` Appear?

- In Python, **every function returns a value**.
- If you don't write a `return` statement, the function returns `None` by default.
- When you assign the result of a function that only prints (does not return), like this:
  ```python
  result = add("Jony")
  ```
  `result` is set to `None`.

- When you print `result`:
  ```python
  print(result)
  ```
  Python prints `None`.

---

## ✅ How to Avoid Printing `None`

### Option 1: Just Call the Function Without Assigning

If your function only prints and doesn't need to return anything,
**don't assign its result to a variable**:

```python
def add(name='None'):
    print(f'Name of the person is {name}')

add("Jony")  # Direct call, no assignment
```

**Output:**
```
Name of the person is Jony
```

---

### Option 2: Return the Value Instead of Printing

If you want to use the result later, **return** the string instead of printing it:

```python
def add(name='None'):
    return f'Name of the person is {name}'

result = add("Jony")
print(result)
```

**Output:**
```
Name of the person is Jony
```

---

### Option 3: Both Print and Return

If you want to both print and return the value:

```python
def add(name='None'):
    output = f'Name of the person is {name}'
    print(output)
    return output

result = add("Jony")
# Only print(result) if you need to see the returned value again
```

---

## ⚠️ Common Mistake

If you do this:

```python
def add(name='None'):
    print(f'Name of the person is {name}')

result = add("Jony")   # Calls add(), prints message, returns None
print(result)          # Prints None
```

You will get:
```
Name of the person is Jony
None
```

Because `add()` does not return anything (so returns `None`), and `print(result)` prints that `None`.

---

## 🏆 Summary Table

| Pattern             | Output                       | Reason                                |
|---------------------|-----------------------------|---------------------------------------|
| `print(add(...))`   | message + None              | add() prints message, returns None    |
| `add(...)`          | message                     | add() prints message, nothing else    |
| `return` then print | message                     | add() returns message, printed later  |

---

## 📚 References

- [Python Docs: Defining Functions](https://docs.python.org/3/tutorial/controlflow.html#defining-functions)
- [Python Docs: None](https://docs.python.org/3/library/constants.html#None)

---


