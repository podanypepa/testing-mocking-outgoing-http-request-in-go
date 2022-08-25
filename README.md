# testing-mocking-outgoing-http-request-in-go

Mocking of `Client` from `net/http` package.

100% coverage function for fetching data from external API.

## running test

```bash
go test -v -coverprofile cover.out .
go tool cover -html=cover.out -o cover.html
open cover.html
```
