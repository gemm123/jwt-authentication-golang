# jwt-authentication-golang
Practice implementation JWT Authentication in Golang REST APIs

## Dependencies
    go get -u github.com/gin-gonic/gin
    go get gorm.io/gorm
    go get gorm.io/driver/postgres 

## Example
### Register User
POST /api/user/register

    HTTP/1.1 201 Created
    Content-Type: application/json; charset=utf-8
    Date: Thu, 09 Jun 2022 03:41:13 GMT
    Content-Length: 58
    Connection: close
    
    {
      "email": "gemaakbr@go.com",
      "userId": 2,
      "username": "gema23"
    }

### Generate Token
POST /api/token

    HTTP/1.1 200 OK
    Content-Type: application/json; charset=utf-8
    Date: Thu, 09 Jun 2022 03:44:47 GMT
    Content-Length: 179
    Connection: close
    
    {
      "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6ImdlbWEyMyIsImVtYWlsIjoiZ2VtYWFrYnJAZ28uY29tIiwiZXhwIjoxNjU0NzQ5ODg3fQ.hNVHM7FSXFRTwkV1SG0SFunyfO7lkRghNucKBsdr2lM"
    }

### Check Secured 
GET /api/secured/ping

    HTTP/1.1 200 OK
    Content-Type: application/json; charset=utf-8
    Date: Thu, 09 Jun 2022 03:47:28 GMT
    Content-Length: 18
    Connection: close
    
    {
      "message": "pong"
    }
