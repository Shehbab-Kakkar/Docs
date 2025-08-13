# AWS Lambda Custom Authorizer Example for API Gateway

This project provides a simple AWS Lambda function to use as a **custom authorizer** for API Gateway endpoints.  
The function checks for a specific authorization token and returns an IAM policy that either allows or denies access to the API.

---

## 🚀 How It Works

1. **Client Request:**  
   The client sends an HTTP request to your API Gateway endpoint with an `Authorization` header.

   ```
   Authorization: Bearer mysecrettoken
   ```

2. **API Gateway Invocation:**  
   API Gateway invokes this Lambda function, passing it the authorization token and the ARN of the method being accessed.

3. **Token Validation:**  
   - If the token matches `Bearer mysecrettoken`, the Lambda returns an "Allow" IAM policy.
   - Otherwise, it returns a "Deny" policy.

---

## 🧩 Function Breakdown

### 1. `generate_policy(principal_id, effect, resource)`

- **Purpose:** Generates an IAM policy document that API Gateway understands.
- **Parameters:**
  - `principal_id`: The user/entity making the request (example: `"user"`).
  - `effect`: `"Allow"` or `"Deny"`.
  - `resource`: The API Gateway method ARN (`event['methodArn']`).

**Example Output:**
```json
{
  "principalId": "user",
  "policyDocument": {
    "Version": "2012-10-17",
    "Statement": [{
      "Action": "execute-api:Invoke",
      "Effect": "Allow",
      "Resource": "<methodArn>"
    }]
  }
}
```

### 2. `lambda_handler(event, context)`

- **event['authorizationToken']**: The token sent by the client (from the `Authorization` header).
- **event['methodArn']**: The ARN of the API Gateway method being accessed.
- **Logic**:  
  - If the token matches the expected value, return an "Allow" policy.
  - Otherwise, return a "Deny" policy.

---

## 🛠️ Example Usage

1. **Deploy** the Lambda function and set it as a custom authorizer on your API Gateway method.
2. **Send a request** to your API with the correct header:

   ```
   Authorization: Bearer mysecrettoken
   ```

3. **Result:**
   - If the token matches, access is allowed.
   - If not, access is denied.

---

## ⚠️ Important Notes

- **This is a basic example.**  
  Hardcoding tokens is **not secure for production**.
- In real applications, validate JWTs, OAuth tokens, expiration, issuer, etc.
- `principalId` can be any string; it's used for logging or identifying the requester.

---

## 📝 Example Lambda Code

```python
def generate_policy(principal_id, effect, resource):
    return {
        'principalId': principal_id,
        'policyDocument': {
            'Version': '2012-10-17',
            'Statement': [{
                'Action': 'execute-api:Invoke',
                'Effect': effect,
                'Resource': resource
            }]
        }
    }

def lambda_handler(event, context):
    token = event.get('authorizationToken')
    method_arn = event.get('methodArn')
    if token == "Bearer mysecrettoken":
        return generate_policy("user", "Allow", method_arn)
    else:
        return generate_policy("user", "Deny", method_arn)
```

---

## 🔒 Security Recommendation

For production:
- Use environment variables or AWS Secrets Manager for secrets.
- Validate industry-standard tokens (JWT/OAuth).
- Implement token expiration, issuer, and audience checks.

---

## 📚 References

- [AWS Custom Authorizers for API Gateway](https://docs.aws.amazon.com/apigateway/latest/developerguide/apigateway-use-lambda-authorizer.html)
- [AWS Lambda Developer Guide](https://docs.aws.amazon.com/lambda/latest/dg/welcome.html)

---
