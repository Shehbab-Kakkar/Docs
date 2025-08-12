# HTTP Status Codes: 201 Created vs 202 Accepted

This guide explains the **key differences** between HTTP status codes **201 (Created)** and **202 (Accepted)**, with clear examples and a quick reference table.

---

## ✅ 201 Created

- **Meaning:** The request was successfully processed, and a new resource has been created as a result.
- **Use Case:** Commonly used for `POST` requests that create something (such as a new user or blog post).
- **Server Action:** The new resource is **already created** by the time the server sends the response.
- **Response Includes:**
  - Usually a `Location` header pointing to the new resource.
  - Often includes the resource details in the body.

**Example:**
```
HTTP/1.1 201 Created
Location: /users/123
Content-Type: application/json

{
  "id": 123,
  "name": "Alice"
}
```

---

## 🕒 202 Accepted

- **Meaning:** The request has been accepted for processing, but the processing has **not yet been completed**.
- **Use Case:** Useful when the server needs more time to process the request (such as uploading a large file or handling a background job).
- **Server Action:** The request is acknowledged, but **completion is not guaranteed**.
- **Response Includes:**
  - Optional details about the processing status or how to track it.

**Example:**
```
HTTP/1.1 202 Accepted
Content-Type: application/json

{
  "message": "Your request is being processed."
}
```

---

## 🔁 Summary Table

| Feature              | 201 Created                  | 202 Accepted                     |
|----------------------|-----------------------------|----------------------------------|
| **Resource created?**| Yes, immediately            | Not yet; may happen later        |
| **Used for?**        | Creating a new resource     | Long-running or deferred process |
| **Guarantees action?**| Yes                        | No, only accepted for processing |
| **Common methods**   | POST                        | POST, PUT, DELETE                |

---

## 📝 When to Use

- **201 Created:**  
  Use when your API **immediately creates** a resource and can return its details.

- **202 Accepted:**  
  Use when your API **starts a process** that may finish later (asynchronously), such as background jobs or large uploads.

---

## 📚 References

- [RFC 9110: HTTP Semantics - 9.5.7. 201 Created](https://datatracker.ietf.org/doc/html/rfc9110#name-201-created)
- [RFC 9110: HTTP Semantics - 9.5.8. 202 Accepted](https://datatracker.ietf.org/doc/html/rfc9110#name-202-accepted)

---
