# AWS Lambda Custom Authorizer

This project contains a simple AWS Lambda function designed to serve as a **custom authorizer** for an API Gateway endpoint.

## 🔍 Overview

The Lambda function inspects the `Authorization` header of incoming requests and determines whether to allow or deny access based on a hardcoded bearer token.

## 🧠 How It Works

When a client makes a request to a protected API Gateway endpoint:

1. API Gateway invokes this Lambda authorizer function.
2. The function checks the provided token.
3. Based on the token, it returns an IAM policy to **Allow** or **Deny** the request.

## 📄 Lambda Function Code

```python
import json

def generate_policy(principal_id, effect, resource):
    return {
        "principalId": principal_id,
        "policyDocument": {
            "Version": "2012-10-17",
            "Statement": [{
                "Action": "execute-api:Invoke",
                "Effect": effect,
                "Resource": resource
            }]
        }
    }

def lambda_handler(event, context):
    token = event['authorizationToken']
    method_arn = event['methodArn']

    if token == "Bearer mysecrettoken":
        return generate_policy("user", "Allow", method_arn)
    else:
        return generate_policy("user", "Deny", method_arn)
