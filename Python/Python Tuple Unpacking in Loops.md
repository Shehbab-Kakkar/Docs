# 📦 Python Tuple Unpacking in Loops — Stock Example

## 📝 Description

This README explains **tuple unpacking** in Python using a simple example:  
Printing the brand name and updated price (with a 10% increase) for a list of stocks.

---

## 🧩 The Example Code

```python
# List of (brand, price) tuples
stocks = [("Apple", 200), ("Samsung", 300), ("Nokia", 400)]

# Loop with tuple unpacking
for brand, price in stocks:
    print(brand, price + (0.1 * price))
```

---

## 🔍 Step-by-Step Explanation

### 1. The Data Structure

```python
stocks = [("Apple", 200), ("Samsung", 300), ("Nokia", 400)]
```
- This is a list of tuples.
- Each tuple contains two elements: the brand (string) and the price (integer).

### 2. Tuple Unpacking in the Loop

```python
for brand, price in stocks:
    # body of the loop
```
- In each iteration, Python takes one tuple from the list (e.g., `("Apple", 200)`).
- It **unpacks** the tuple into the variables `brand` and `price`:
  - `brand` gets `"Apple"`, `price` gets `200`.
- This is called **tuple unpacking**.

### 3. Calculating the Updated Price

```python
print(brand, price + (0.1 * price))
```
- Adds 10% to the original price.
- Prints the brand and the new price.

---

## 🖨️ Sample Output

```
Apple 220.0
Samsung 330.0
Nokia 440.0
```

---

## 📦 What is Tuple Unpacking?

- **Tuple unpacking** allows you to assign each item in a tuple to a variable in a single statement.
- Very useful in `for` loops when iterating over a list of tuples.

**Example:**
```python
person = ("Alice", 30)
name, age = person   # name='Alice', age=30
```

---

## 🏷️ When to Use

- When you have sequences (like tuples/lists) with a fixed size and want to extract elements directly.
- Common in looping over lists of tuples, e.g., from databases or APIs.

---

## 📚 References

- [Python Docs: Sequence unpacking](https://docs.python.org/3/tutorial/datastructures.html#tuples-and-sequences)
- [Python Docs: for Statements](https://docs.python.org/3/tutorial/controlflow.html#for-statements)

---
