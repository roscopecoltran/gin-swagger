# gin-swagger - DRY templates for go-swagger

gin-swagger is a tool assisting in writing golang REST APIs based on a [API
First Principle][api-first]. Given a swagger spec of your REST API
`gin-swagger` can generate all the boring boilerplate stuff for you, so you
only have to implement the core business logic of the API.

It is based on [go-swagger][go-swagger].

## Generate code

```
./gin-swagger -A cluster-registry -f swagger.yaml
```

## Features
* [ ] Validate + bind input to gin ctx.
  * [x] bind body input.
    * [ ] [Nice to have] custom input Models with required fields.
  * [x] bind params input.
  * [x] bind query params input.
  * [ ] consume more than `application/json`
* [ ] Security.
  * [ ] basic
  * [ ] apiKey
  * [ ] OAuth2
    * [x] `password` (Bearer token)
    * [ ] `accessCode`
    * [ ] `application`
    * [ ] `implicit`
  * [ ] Auth chain
  * [x] Custom authorize on individual routes.
* [x] Set custom middleware on each router Pre/post 'main' handler.
  * Use case pre: *custom authorization pre handler*.
  * Use case post: *audit report events post handler*.
* [ ] Ginize generated code.
* [ ] Set and get user info (uid, realm) from gin context.
* [ ] Response helper functions
* [ ] Default metrics (prometheus)

## How it works

gin-swagger generates a [gin][gin] based REST API based on your swagger
spec.

Given the following swagger spec:

```yaml
swagger: '2.0'
info:
  version: "0.0.1"
  title: My API
schemes:
    - "https"
basePath: /
produces:
  - application/json
consumes:
  - application/json
paths:
  '/persons/{name}':
    get:
      summary: get Person
      tags:
        - Person
      operationId: getPerson
      responses:
        200:
          description: Person by name.
          schema:
            '$ref': '#/definitions/Person'
parameters:
  name:
    name: name
    in: path
    type: string
    required: true
    pattern: "^[a-z][a-z0-9]$"
definitions:
  Person:
    type: object
    properties:
      name:
        type: string
```

You can generate the REST API using the following command:

```bash
$ gin-swagger -A my-api -f swagger.yaml
```

This will generate two folders `models` and `restapi`. `models` contains model
structs based on the *definitions* defined in your swagger file, including
model validation, this is generated by [go-swagger][go-swagger]. `restapi` will
contain everything generated by `gin-swagger` and can be overwritten when
updating and regenerating your swagger spec. Inside `restapi/api.go` will be a
generated `Service` interface which is all you have to implement in order to
add logic to your REST api. Given the above swagger spec, the `Service`
interface will look like this:

```go
type Service interface {
    Healthy() bool
    GetPerson(ctx *gin.Context, params *persons.GetPersonParams) *api.Response
}
```

The `Healthy() bool` method should return true if the service is considered
healthy. The return value is used by the default health endpoint `/healthz`
provided by `gin-swagger`.

The `GetPerson()` method should implement the business logic of the
`/persons/{name}` path of your REST API.

A simple service implementation looks like this:

```go
type mySvc struct {
    health bool
}

func (m *mySvc) GetPerson(ctx *gin.Context, params *persons.GetPersonParams) *api.Response {
    return &api.Response{
        Code: http.StatusOK,
        Body: &models.Person{Name: params.Name},
    }
}

func (m *mySvc) Healthy() bool {
    return m.health
}
```

With the service implemented the only thing left to do is plug it into
`gin-swagger` and run it from your `main.go` (or were ever you like):

```go
package main

import ...

var apiConfig restapi.Config

func main() {
    err := apiConfig.Parse()
    if err != nil {
        log.Fatal(err)
    }
    svc := &mySvc{Health: false}
    api := restapi.NewAPI(svc, &apiConfig)
    err = api.RunWithSigHandler()
    if err != nil {
        log.Fatal(err)
    }
}
```

`restapi.Config` is a default config object defined by `gin-swagger` it lets
you configure default options through command like flags when starting the
server. For instance you can tell the server to serve HTTP only with the
`--insecure-http` flag (default is to serve HTTPS).

For a full example see the [example folder](example).

[api-first]: https://zalando.github.io/restful-api-guidelines/
[gin]: https://github.com/gin-gonic/gin
[go-swagger]: https://github.com/go-swagger/go-swagger
