Ah, thanks for the clarification — you're asking what **return code or status code** (e.g., HTTP status code) is returned by **Boto3** when performing AWS operations like **GET**, or when an operation **fails or succeeds**, such as **update failed**.

### ✅ Here's the direct answer:

---

## 🔹 Boto3 Return Codes / Status Handling

### 1. **Boto3 does not directly expose HTTP status codes** in most responses.

* Instead, successful operations return a **Python dictionary** with the results.
* If the operation **fails**, Boto3 (actually `botocore`, under the hood) raises an **exception**, typically `ClientError`, which **includes** the HTTP status code and error message.

---

## 🔹 Where to Find the Return Code?

When a request fails, you can access the HTTP status code like this:

```python
from botocore.exceptions import ClientError

client = boto3.client('s3')
try:
    client.get_object(Bucket='my-bucket', Key='nonexistent.txt')
except ClientError as e:
    status_code = e.response['ResponseMetadata']['HTTPStatusCode']
    error_code = e.response['Error']['Code']
    error_message = e.response['Error']['Message']
    print(f"Error code: {error_code}")
    print(f"HTTP status code: {status_code}")
    print(f"Message: {error_message}")
```

### Example Output:

```
Error code: NoSuchKey
HTTP status code: 404
Message: The specified key does not exist.
```

---

## 🔹 Common AWS HTTP Status Codes

| HTTP Status Code            | Meaning                       | Example Cause                         |
| --------------------------- | ----------------------------- | ------------------------------------- |
| `200 OK`                    | Success                       | Any successful GET, PUT, DELETE, etc. |
| `400 Bad Request`           | Invalid request               | Bad parameters                        |
| `403 Forbidden`             | Permission denied (IAM issue) | Missing permission                    |
| `404 Not Found`             | Resource not found            | Missing bucket/object                 |
| `500 Internal Server Error` | AWS service error             | Retryable error                       |

---

## 🔹 Operation Status (e.g., `UPDATE_FAILED`)

For services like CloudFormation, Boto3 **returns a status string** in the response (not a status code):

```python
cf = boto3.client('cloudformation')
response = cf.describe_stacks(StackName='my-stack')
print(response['Stacks'][0]['StackStatus'])  # e.g., "UPDATE_FAILED"
```

This is **not** an HTTP status code — it's an **internal AWS status** for the resource lifecycle.

---

## 🔚 Summary

* ✅ Successful requests return data (dict) and **HTTP 200 OK** under the hood.
* ❌ Failed requests raise `ClientError`, where you can get:

  * `e.response['Error']['Code']` → AWS-specific error code
  * `e.response['ResponseMetadata']['HTTPStatusCode']` → HTTP status code
* Some AWS services (like CloudFormation) return **operation status strings** (e.g., `CREATE_FAILED`) inside the response payload.

---

Let me know which AWS service or method you're specifically using, and I can show you the exact return format and how to read the status from Boto3.
