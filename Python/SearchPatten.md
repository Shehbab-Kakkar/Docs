Here are a few simple ways to **find the pattern `"bb"` in the string `"aabbaba"` using Python**.

---

### ✅ Method 1: Using `in`

```python
s = "aabbaba"

if "bb" in s:
    print("Pattern 'bb' found")
else:
    print("Pattern 'bb' not found")
```

---

### ✅ Method 2: Using `find()` (gives index)

```python
s = "aabbaba"

index = s.find("bb")

if index != -1:
    print("Pattern 'bb' found at index:", index)
else:
    print("Pattern 'bb' not found")
```

---

### ✅ Method 3: Using Regular Expressions (recommended for pattern matching)

```python
import re

s = "aabbaba"

match = re.search("bb", s)

if match:
    print("Pattern 'bb' found at index:", match.start())
else:
    print("Pattern 'bb' not found")
```

---

### ✅ Method 4: Count occurrences

```python
s = "aabbaba"
count = s.count("bb")

print("Number of 'bb' patterns:", count)
```

📌 **Output for this string:**

```
Pattern 'bb' found at index: 2
```

