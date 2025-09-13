Here’s a clear, authoritative guide to **HTTP status codes** for **POST, PUT, GET, and DELETE** operations in APIs, covering **success**, **failure**, and **timeout** responses. This is especially valuable for DevOps and backend engineers designing or consuming RESTful services.

***

## HTTP Methods and Typical Success Codes

| Method   | Success (200s)                      | Notes                                                                                   |
|----------|-------------------------------------|-----------------------------------------------------------------------------------------|
| **GET**  | **200 OK**                          | The resource was found and returned in the response body[2][10].                       |
| **POST** | **201 Created**                     | A new resource was created as a result of this request[11].                              |
|          | **200 OK**                          | Sometimes used if the result of the POST is not a resource creation[10].                |
| **PUT**  | **200 OK**                          | The resource was updated; a representation may be returned[3][4][7].                 |
|          | **204 No Content**                  | The resource was updated, no response body needed[1][3][4][7].                      |
|          | **201 Created**                     | If the resource was newly created (not already existing) by the PUT[4][7].            |
| **DELETE** | **200 OK**                          | The resource was deleted; a representation may be returned[3][4][5].                 |
|          | **204 No Content**                  | The resource was deleted, no response body needed[3][4][5][7].                      |
|          | **202 Accepted**                    | The deletion is accepted but not yet completed (async operation)[3][4][5].           |

**Note:** While **200 OK** is common for all methods, **201 Created** is the preferred success code for POST when a new resource is created. For updates (PUT), **200 OK** and **204 No Content** are both appropriate depending on whether a response body is returned. For deletes (DELETE), **204 No Content** is the most RESTful, but **200 OK** is also used if a status description is returned.[3][4][5][7]

***

## HTTP Status Codes for Failures

| Method   | Common Failure Codes                | Notes                                                                                   |
|----------|-------------------------------------|-----------------------------------------------------------------------------------------|
| **GET**  | **404 Not Found**                   | The resource does not exist[5].                                                        |
|          | **400 Bad Request**                 | The request syntax is incorrect.                                                        |
|          | **401 Unauthorized**                | Authentication failed.                                                                  |
|          | **403 Forbidden**                   | Authenticated, but not authorized for the resource.                                     |
| **POST** | **400 Bad Request**                 | Malformed request.                                                                      |
|          | **401 Unauthorized**                | Authentication failed.                                                                  |
|          | **403 Forbidden**                   | Forbidden access.                                                                       |
|          | **409 Conflict**                    | Resource state conflicts with the request.                                              |
|          | **422 Unprocessable Entity**        | Validation error in the request data.                                                   |
| **PUT**  | **404 Not Found**                   | The resource to update does not exist (unless PUT is allowed to create new resources).   |
|          | **400 Bad Request**                 | Malformed request.                                                                      |
|          | **401 Unauthorized**                | Authentication failed.                                                                  |
|          | **403 Forbidden**                   | Forbidden access.                                                                       |
|          | **409 Conflict**                    | Resource state conflicts with the request.                                              |
|          | **422 Unprocessable Entity**        | Validation error in the request data.                                                   |
| **DELETE** | **404 Not Found**                   | The resource does not exist (sometimes **204 No Content** is used for idempotency)[4][5][7]. |
|          | **400 Bad Request**                 | Malformed request.                                                                      |
|          | **401 Unauthorized**                | Authentication failed.                                                                  |
|          | **403 Forbidden**                   | Forbidden access.                                                                       |
|          | **409 Conflict**                    | Resource state conflicts with the request.                                              |

***

## Timeout and Gateway Issues

For all methods, if a server cannot complete the request due to a timeout or other server-side issue, these codes are used:

- **408 Request Timeout**: The server timed out waiting for the request (client-side timeout).
- **504 Gateway Timeout**: The server, while acting as a gateway or proxy, did not receive a timely response from the upstream server.
- **502 Bad Gateway**: The server, while acting as a gateway or proxy, received an invalid response from the upstream server.

These codes are not method-specific; they apply whenever the server cannot process the request due to timeout or communication issues with backends.

***

## Summary Table

| Method   | Success                 | Resource Not Found | Bad Request | Auth Fail       | Conflict         | Timeout/Gateway         |
|----------|-------------------------|--------------------|-------------|-----------------|------------------|-------------------------|
| **GET**  | 200 OK                  | 404 Not Found      | 400 Bad Req | 401/403         | 409 Conflict     | 408/502/504              |
| **POST** | 201 Created / 200 OK    | N/A                | 400 Bad Req | 401/403         | 409 Conflict     | 408/502/504              |
| **PUT**  | 200 OK / 204 No Content | 404 Not Found      | 400 Bad Req | 401/403         | 409 Conflict     | 408/502/504              |
| **DELETE** | 200 OK / 204 No Content| 404 Not Found      | 400 Bad Req | 401/403         | 409 Conflict     | 408/502/504              |

***

## Key Takeaways

- **Success**: Use **200**, **201**, **204**, or **202** depending on the operation and response body requirements.
- **Resource Not Found**: **404** for GET, PUT, DELETE; sometimes **204** for DELETE (idempotency).
- **Bad Request**: **400** for malformed requests.
- **Auth Issues**: **401** (Unauthorized) or **403** (Forbidden).
- **Conflict**: **409** if the request cannot be completed due to a conflict with the current state.
- **Timeout**: **408** (client), **502/504** (server/gateway).

These conventions help ensure predictable, standards-compliant API behavior—critical for automation, observability, and integration in DevOps workflows.

[1](https://restfulapi.net/http-status-codes/)
[2](https://developer.mozilla.org/en-US/docs/Web/HTTP/Reference/Status)
[3](https://www.moesif.com/blog/technical/api-design/Which-HTTP-Status-Code-To-Use-For-Every-CRUD-App/)
[4](https://stackoverflow.com/questions/2342579/http-status-code-for-update-and-delete)
[5](https://restfulapi.net/http-methods/)
[6](https://apidog.com/blog/http-methods/)
[7](https://api7.ai/learning-center/api-101/http-methods-in-apis)
[8](https://www.mailersend.com/help/rest-api-response-codes)
[9](https://www.w3schools.com/tags/ref_httpmethods.asp)
[10](https://beeceptor.com/docs/concepts/http-status-codes/)
[11](https://www.digitalocean.com/community/tutorials/process-management-in-linux)
