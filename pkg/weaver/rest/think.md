# REST Client

Rest clients Need to be able to

- Create Client → HTTP / HTTPS
  - HTTP Basic Auth ⇒ user and password
  - `NewBearerClient(token, base [string](<https://pkg.go.dev/builtin#string>))`
- BaseUrl / HostUrl ⇒ clean duplicates in requests path
- UploadType `[UploadType](<https://pkg.go.dev/github.com/kevinburke/rest/restclient#UploadType>)`
- Create w Client Config
  - https config - certs
  - http config
  - default request options
    - Default Retry
    - ErrorHandler func(\*http.Response, \*handlerFunc) error
- ExecuteRequest - Post | Put | Get | Delete
  - Build Request
- Verbs must keep their characteristics

//

Client Config

⇒ Client Transport sets up security level

- set auth token
- Set Root Certificate FromFile | FromText

`client.SetRootCertificate("/path/to/root/pemFile1.pem")`

`client.SetRootCertificateFromString("-----BEGIN CERTIFICATE-----content-----END CERTIFICATE-----")`

- | FromKeyPair

```
// Parsing public/private key pair from a pair of files. The files must contain PEM encoded data.
cert1, err := tls.LoadX509KeyPair("certs/client.pem", "certs/client.key")
```

- Default retry condition & mechanism for client ? Just retry command→ execute request

```

client.AddRetryCondition(
    // RetryConditionFunc type is for retry condition function
    // input: non-nil Response OR request execution error
    func(r *resty.Response, err error) bool {
        return r.StatusCode() == http.StatusTooManyRequests
    },
)
```

- Client config = TLS or Non
- `&tls.Config{ InsecureSkipVerify: true }`
- `&tls.Config{ RootCAs: roots }`
- Set Client Timeout
- Set Retry After seconds

Request Structure - default

- Header
- timeout
- maxRetries
- retry Func
- retry after seconds
- body | bodybytes ?

⇒ Build Request

- Content Type → default to application/json
- Verb
- Add Header → field types content | encryption |
  - content type
  - user agent
- Set Cookie Data
- Add Data → tyeps | JSON | XML
- check data
- convert data to other type
- verify data
- multipart upload
- SetQueryString
- Add Param
  - paramName + string
  - `SetQueryParams` w map
- setTimeout
- BuildUrl → use path & params | lowercase | normalize

Returned Response

- Status Code
  - 200 - ok
  - 300 - moved permanently
  - 400 - bad request
  - 401 - unauthorized
  - 403 - Forbidden
  - 404 - Not Found
  - 500 - Internal Server Error
  - 502 - Bad Gateway
  - 503 - Service Unavailable
  - 504 - Gateway TimeOut
- Error - Bool
- error
  - status code
  - message
- Retry Count
- On Error triggers retry function

Debug level

chainable methods for building Request

- add component(s)
- verify + build

Request / Response middleware

Debug mode - clean and informative logging presentation

`HTTP/2` and `HTTP/1.1`

- Backoff Retry
- Conditional Retry

Request and Response Middleware

`client.OnBeforeRequest(func(c *resty.Client, req *resty.Request) error`

some commands would require some execution on clientside before request is executed to the server
