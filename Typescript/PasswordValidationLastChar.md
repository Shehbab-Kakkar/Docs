# last index we want to check

Certainly! If you want to perform the same kind of password validation check, but this time on the **last character** of the password string instead of the first, you can use the `.charCodeAt()` function with the index set to the last character.

Here's how you do it:

## Checking the Last Character

You can get the last character's index using `password.length - 1`, and then check its code just like before.

```typescript
var password: string = "1H&@G^%FD"
var lastCharIndex: number = password.length - 1;

if (
  password.charCodeAt(lastCharIndex) >= 65 &&
  password.charCodeAt(lastCharIndex) <= 90
) {
  console.log('password is valid')
} else {
  console.log('password is not valid')
}
```


### How It Works

- `password.length - 1` gives the index of the **last character** in the string.
- `password.charCodeAt(lastCharIndex)` gets the Unicode (ASCII) value of that character.
- The code then checks if this value is between 65 and 90 (which means it's an uppercase English letter: A-Z).
- If the last character is uppercase, it prints "password is valid". Otherwise, it prints "password is not valid".


### Example Table

| Password | Last Char | charCodeAt Value | Output message |
| :-- | :-- | :-- | :-- |
| `TestA` | `A` | 65 | password is valid |
| `XYZ0B` | `B` | 66 | password is valid |
| `hello1` | `1` | 49 | password is not valid |
| `Abc123z` | `z` | 122 | password is not valid |

> **Tip:**
> Use `charCodeAt(password.length - 1)` to work with the last character of any string.

Let me know if you want this same format in markdown or have more questions!
