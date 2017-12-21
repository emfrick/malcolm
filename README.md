# Malcolm

Malcolm is a stupid simple middleware chaining syntax for Go.

## Usage

`import github.com/emfrick/malcolm`

Call the `Then()` function on `malcolm` and pass it middleware of type `func(next http.Handler) http.Handler`. Call `Then()` as many times as necessary to chain middleware. Call `Create()` at the end of the chain to create the middleware chain. Use this chain to wrap your handler of type `func (h http.Handler) http.Handler`.

## Example
Create logging middleware:
```go
func LoggerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("LoggerMiddleware Begin")
		next.ServeHTTP(w, r)
		fmt.Println("LoggerMiddleware End")
	})
}
```

Create authentication middleware:
```go
func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("AuthMiddleware Begin")
		next.ServeHTTP(w, r)
		fmt.Println("AuthMiddleware End")
	})
}
```

Create your route handler
```go
func RouteHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Println("RouteHandler Begin")

    // ... Do All The Things

    fmt.Println("RouteHandler End")
}
```

Wrap your route handler with the middleware
```go
func main() {

    routeHandler := http.HandlerFunc(RouteHandler)

    standardChain := malcom.
                        Then(LoggerMiddleware).
                        Then(AuthMiddleware).
                        Create()

    http.Handle("/route", standardChain(routeHandler))

    http.ListenAndServe(":8080", nil)
}
```

Log output:
```
LoggerMiddleware Begin
AuthMiddleware Begin
RouteHandlerBegin
RouteHandlerEnd
AuthMiddleware End
LoggerMiddleware End
```