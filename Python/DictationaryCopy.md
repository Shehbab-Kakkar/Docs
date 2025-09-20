Great question! You're asking why, when you assign a dictionary to another variable and then `pop` an element from that new variable, it also removes the element from the **original dictionary**.

Let's walk through it clearly.

---

### 📌 Example

```python
original_dict = {"a": 1, "b": 2}
new_dict = original_dict
new_dict.pop("a")
print(original_dict)  # Output?
```

**Output:**

```python
{'b': 2}
```

Why did `"a"` disappear from `original_dict`?

---

### 🧠 Reason: Variables Hold References

In Python, when you do:

```python
new_dict = original_dict
```

You're **not creating a new dictionary**. You're just creating a **new reference** (or alias) to the **same dictionary in memory**.

So both `new_dict` and `original_dict` point to the same object.

---

### 🔍 Visual Representation

```
original_dict  --->  {"a": 1, "b": 2}
                      ↑
new_dict       -------|
```

Both variables point to the same dictionary in memory. So any change made using either variable affects the same object.

---

### ✅ How to Create a Copy Instead

If you want to create a separate (independent) dictionary so that modifying it doesn't affect the original, you need to make a **copy**.

#### Option 1: Using `.copy()`

```python
new_dict = original_dict.copy()
```

#### Option 2: Using `dict()`

```python
new_dict = dict(original_dict)
```

#### Option 3: Using dictionary unpacking

```python
new_dict = {**original_dict}
```

Now, popping from `new_dict` won't affect `original_dict`.

---

### 🧪 Example with `.copy()`

```python
original_dict = {"a": 1, "b": 2}
new_dict = original_dict.copy()
new_dict.pop("a")
print(original_dict)  # {'a': 1, 'b': 2}
print(new_dict)       # {'b': 2}
```

---

