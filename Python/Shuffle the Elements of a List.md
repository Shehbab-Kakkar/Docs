# 🎲 Python Program: Shuffle the Elements of a List

This README explains how to shuffle a list in Python using the `shuffle()` function from the `random` module, why you might see `None` in your variable, and how to properly use `shuffle()`.

---

## 📝 The Problem

```python
from random import shuffle

origList = [1, 2, 5, 45, 21, 6, 77, 100]
newList = shuffle(origList)
print(newList)      # None
print(origList)     # [shuffled values]
```

**Output:**
```
None
[5, 2, 100, 1, 77, 21, 45, 6]
```

---

## ❓ Why is `newList` None?

- `shuffle(origList)` **shuffles the list in place** (modifies the original list).
- It **does not return a new list**; instead, it returns `None`.
- So when you do `newList = shuffle(origList)`, `newList` is set to `None`.

---

## ✅ The Correct Way to Shuffle

### Shuffle the List and Print

```python
from random import shuffle

origList = [1, 2, 5, 45, 21, 6, 77, 100]

shuffle(origList)  # Shuffles origList in place

print("Shuffled List:")
print(origList)
```

---

### Keep the Original List Unchanged

If you want to preserve the original list, copy it first:

```python
from random import shuffle

origList = [1, 2, 5, 45, 21, 6, 77, 100]

newList = origList.copy()
shuffle(newList)

print("Original List:")
print(origList)

print("Shuffled Copy:")
print(newList)
```

**Sample Output:**
```
Original List:
[1, 2, 5, 45, 21, 6, 77, 100]
Shuffled Copy:
[21, 77, 5, 6, 45, 2, 1, 100]
```

---

## 🧠 Summary Table

| What You Wrote                 | What It Does                        |
|--------------------------------|-------------------------------------|
| `newList = shuffle(origList)`  | `newList` becomes `None`            |
| `shuffle(origList)`            | Shuffles `origList` in-place        |
| `newList = origList.copy(); shuffle(newList)` | Shuffles a copy, keeps original intact |

---

## 📚 References

- [Python Docs: random.shuffle()](https://docs.python.org/3/library/random.html#random.shuffle)

---
