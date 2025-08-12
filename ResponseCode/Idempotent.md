# Idempotent - HTTP Concept Explained

## What Does Idempotent Mean?

An **idempotent** operation is one that **can be performed multiple times without changing the result beyond the initial application**.

In simpler terms:  
**No matter how many times you do it, the outcome stays the same.**

---

## Idempotency in HTTP Methods

| HTTP Method | Idempotent? | Explanation                                   |
|-------------|-------------|-----------------------------------------------|
| **GET**     | Yes         | Fetching data multiple times doesn’t change anything. |
| **PUT**     | Yes         | Updating a resource with the same data multiple times has the same effect as once. |
| **DELETE**  | Yes         | Deleting the same resource multiple times has the same effect as deleting once (resource gone). |
| **POST**    | No          | Typically creates a new resource every time, so repeated calls change the state (not idempotent). |
| **PATCH**   | Usually No  | Partially updates; repeated calls might have cumulative effect. |

---

## Example

Suppose you `PUT /users/123` with updated user data.

- You send the same `PUT` request **3 times**.
- The user’s data will be the same after the first request and unchanged by the next two — so the operation is idempotent.

---

## Why Is Idempotency Important?

- Ensures safety in network retries.
- Avoids unintended side effects.
- Helps build reliable and predictable APIs.

---

*This explanation helps understand why some HTTP methods are idempotent and how that affects API design.*
