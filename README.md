# ExLogGo

---

Logger to integrate with the Gin library and support a chain of log messages.

### Installation

---
To install ExLogGo package, you need to install Go and set your Go workspace first.

1. You first need Go installed (version 1.19+ is required), then you can use the below Go command to install 
```
go get -u github.com/execaus/exloggo
```
2. Import it in your code:
```
import "github.com/execaus/exloggo"
```

### Quick start

---

```go
package main

import (
  "net/http"
  
  "github.com/gin-gonic/gin"
  "github.com/execaus/exloggo"
)

func main() {
  r := gin.Default()
  err := exloggo.SetParameters(exloggo.Parameters{
    Mode: exloggo.DevelopmentMode, // default development mode
    ServerVersion: "development",
    Directory: "/logs", // default "/logs"
  })
  if err != nil {
	  println(err.Error())
	  return
  }
  exloggo.Inject(r)
  r.GET("/ping", func(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{
      "message": "pong",
    })
  })
  r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
```

### Methods

---
Log output methods have variants with static message and formatted (with the prefix "f").

| Base method       | Description                                          |   Post effect |
|-------------------|:-----------------------------------------------------|--------------:|
| **Info**          | Generates an information log entry                   |          none |
| **Warning**       | Generates an warning log entry                       |          none |
| **Error**         | Generates an error log entry                         |          none |
| **Fatal**         | Generates an fatal log entry                         |  exit program |
| **RequestResult** | Generates an Gin request log                         |          none |


### Request Headers

---

The API supports the HTTP request headers listed below.

| Заголовок         | Тип значения | Описание                                                                                                                                                |
|-------------------|--------------|---------------------------------------------------------------------------------------------------------------------------------------------------------|
| Client-Request-ID | UUID         | Mandatory. A unique call identifier that is useful for logs and network tracing for troubleshooting purposes. The value must be set anew for each call. |

### Answer headers

---

The REST API supports the HTTP response headers listed below.

| Header            | Value type | Description                                                                                                                                                                                      |
|-------------------|------------|--------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| Timestamp         | timestamp  | Mandatory. A timestamp indicating when the request reached the API.                                                                                                                              |
| Client-Request-ID | UUID       | Mandatory. A unique call identifier that is useful for logs and network tracing for troubleshooting purposes. The value must be set anew for each call. All operations must include this header. |
| Request-ID        | UUID       | Mandatory. A unique call identifier that is useful for logs and network tracing for troubleshooting purposes. The value must be set anew for each call. All operations must include this header. |

### Error resource type

---

An error response is a separate JSON object containing all the information about the error. You can use the data it returns instead of or together with the HTTP status code. Below is an example of the full JSON error text.

The table below describes the scheme of the error response.

| Field   | Type   | Description                                                                                                                                                                  |
|---------|--------|------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| code    | string | Always returns. Indicates the type of error that occurred. Does not accept null.                                                                                             |
| message | string | Always returns. Contains a detailed description of the error and additional information for debugging. Does not take null, cannot be empty. Maximum length: 1024 characters. |

### ```Code``` property

---

The ```code``` property contains one of the possible values listed below. Applications must be prepared to handle any of these errors.

| Code             | Description                                                                              |
|------------------|------------------------------------------------------------------------------------------|
| accessDenied     | The calling party does not have permission to perform the action.                        |
| generalException | An unspecified error has occurred.                                                       |
| itemNotFound     | Resource not found.                                                                      |
| invalidRequest   | The request is in the wrong format or is incorrect.                                      |
| resourceModified | The updated resource has changed since the last read.                                    |
| unAuthenticated  | The calling object has not been authenticated.                                           |
| sendEmail        | The basic script was executed, but an error occurred while working with the SMTP server. |
| resourceDeleted  | All actions with the resource are forbidden, because it was deleted.                     |
| notConfirmed     | The data intended to confirm the action did not match the expected values.               |
| timeoutExpired   | The time limit for performing the action has expired.                                    |
| headerException  | Mandatory query headers are missing or incorrect.                                        |

### Status codes

---

| Status code   | Status message            | Description                                                                                    |
|---------------|---------------------------|------------------------------------------------------------------------------------------------|
| 200           | Success                   | The request was successful.                                                                    |
| 204           | No content                | The request was successful and the response does not contain a body.                           |
| 400           | Request error             | The request could not be processed because it was in an invalid format or is invalid.          |
| 401           | Not authorized            | The necessary data for authentication is missing or is not valid for the resource.             |
| 403           | Prohibited                | Access to the requested resource is denied. Perhaps the user does not have enough permissions. |
| 404           | Not found                 | The requested resource does not exist.                                                         |
| 405           | Метод не разрешен         | The HTTP method in the request is not allowed for the resource.                                |
| 410           | The method is not allowed | The requested, modified or deleted resource is removed from the server.                        |
| 412           | A prerequisite is not met | The prerequisite specified in the request does not match the current state of the resource.    |
| 432           | Request Headers           | An error occurred while processing the query headers.                                          |
| 500           | Internal server error     | An internal server error occurred while processing the request.                                |
| 501           | Not implemented           | The requested function is not implemented.                                                     |
| 506           | Email                     | The SMTP server returned an error.                                                             |
| 507           | Not enough storage space  | The maximum storage quota has been reached.                                                    |


### ```Message``` property

---

The ```message``` property at the root contains an error message intended for the developer. Error messages are not localized and should not be shown directly to the user. When handling errors, code should not check ```message``` values, as they can change at any time, and they often contain dynamic information related to failed requests. The ```code``` should only handle error codes returned in the code properties.


### Example output

---

```json
{
  "level":"INFO",
  "message":"server started successfully",
  "point":"server/server.go:24",
  "request_id":"",
  "client_request_id":"",
  "timestamp":"2022-10-08T03:29:32Z",
  "server_version":"v1.0.0",
  "request":null,
  "extends":null
}
```
