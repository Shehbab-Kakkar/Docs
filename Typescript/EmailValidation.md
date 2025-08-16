## Introduction

This project includes a **simple email validation utility** using TypeScript. The code checks if the provided email address contains both the `@` symbol and a dot (`.`)—the minimal markers of a valid email format. This check is performed using the standard `indexOf()` method.

## How It Works

- **Default Behavior:** The script examines the input email string for the presence of `@` and `.` characters.
- **Validation Logic:** If either character is missing (`indexOf()` returns `-1`), the email is considered invalid.
- **Output:** The result is shown using a log message indicating whether the email is valid or not.


## Example

```typescript
var email: string = "test@test.com";
var atpos: number = email.indexOf('@');
var dopos: number = email.indexOf('.');

if (atpos == -1 || dopos == -1) {
  console.log(`Email is invalid ${email}`);
} else {
  console.log('Email is valid');
}
```


## Why `-1`?

The `indexOf()` method in JavaScript and TypeScript returns the **index position** of the character if found, or **`-1` if not found**. This property is leveraged to quickly test for the presence of required characters in the email.

- If `email.indexOf('@')` is `-1`, then `@` is not present.
- If `email.indexOf('.')` is `-1`, then `.` is not present.

Presence of both is required for the email to be considered valid by this utility.

## Usage Examples

| Email Address | Output Message |
| :-- | :-- |
| `test@test.com` | `Email is valid` |
| `testtestcom` | `Email is invalid testtestcom` |

> **Note:**
> This is a **basic validation technique** and does not ensure the email address is fully valid (e.g., it doesn’t check for format errors like multiple `@` signs, misplaced dots, etc.). For more robust validation, consider using regular expressions or dedicated validation libraries.
>
> 
