To import a function from another `.py` file in Python, you can use the `import` statement—just like how you import modules from the standard library. Here's how it works step by step.

---

## ✅ Basic Example

Suppose you have the following two files in the **same directory**:

### `math_utils.py` (your module file)

```python
def add(a, b):
    return a + b
```

### `main.py` (your main script)

```python
from math_utils import add

result = add(5, 3)
print(result)  # Output: 8
```

---

## 📦 If the files are in different folders

Let's say you have this structure:

```
project/
├── main.py
└── utils/
    └── math_utils.py
```

You can import like this in `main.py`:

```python
from utils.math_utils import add
```

Make sure there's an `__init__.py` file in the `utils/` folder (can be empty) to mark it as a package in older Python versions.

---

## 🧪 Alternate import styles

### 1. Import the whole module

```python
import math_utils

result = math_utils.add(5, 3)
```

### 2. Import with alias

```python
import math_utils as mu

result = mu.add(5, 3)
```

### 3. Import multiple functions

```python
from math_utils import add, subtract, multiply
```

---

## ⚠️ Notes

* Make sure the `.py` file you're importing **is in the Python path** or the **same directory**, or use a proper package structure.
* File and function names must not conflict with Python's built-in module names (like `random.py`, `os.py`, etc.).
* Avoid circular imports (where `file_a.py` imports from `file_b.py`, and `file_b.py` imports back from `file_a.py`).

---

