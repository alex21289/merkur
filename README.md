# :full_moon: Merkur

Merkur is a easy to use Go `http.Client` Wrapper.

Origin from: [Federico León](https://github.com/federicoleon/go-httpclient)

## Usage

### Configure the Client

```go
headers := make(http.Header)
headers.Set("Some-Common-Header", "value-for-all-requests")

// Create a new builder:
client := merkur.NewBuilder().

	// You can set global headers to be used in every request made by this client:
	SetHeaders(headers).

	// Configure the timeout for getting a new connection:
	SetConnectionTimeout(2 * time.Second).

	// Configure the timeout for performing the actual HTTP call:
	SetResponseTimeout(3 * time.Second).

	// Configure the User-Agent header that will be used for all of the requests:
	SetUserAgent("Your-User-Agent").

	// Finally, build the client and start using it
	Build()
```

### Simple Get Request

```go
client := merkur.NewBuilder().Build()

response, err := client.Get("https://api.github.com")
if err != nil {
  panic(err)
}

// Get Body as String
bodyString := response.String()
// Get Body as Bytes
bodyBytes := response.Bytes()
// Unmarshal Json to Struct
var foo myStruct
err := response.UnmarshalJson(&foo)
if err != nil {
  panic(err)
}

```

### Get Request with Params

```go
params := merkur.NewParams()
params.Add("tag", "foo")
params.Add("tag", "bar")
params.Add("name", "foobar")

// Get Request to https://api.example.com/search?tag=foo&tag=bar&name=foobar
response, err := client.GetQuery("https://api.example.com/search", params)

```

### Post Request FormData

```go
	payload := formdata.NewFormData()
	payload.Set("grant_type", "password")
	payload.Set("username", "foo")
	payload.Set("password", "bar")

	headers := make(http.Header)
	headers.Set("Content-Type", mmime.ContentTypeXFormUrlencoded)

	response, err := client.Post("http://localhost:9191/api", payload, headers)
	if err != nil {
		log.Fatal(err)
	}
```

### Post Request JSON

```go

headers = make(http.Header)
headers.Set(mmime.HeaderContentType, mmime.ContentTypeJson)

postRequest := User{
  FirstName: "John",
  LastName: "Doe",
  Email: "john.doe@example.com",
}

response, err := client.Post("https://api.example.com/user/create", postRequest, headers)
if err != nil {
	log.Fatal(err)
}
```
