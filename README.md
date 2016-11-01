# Go AWS HTTP Auth Proxy
A simple http aws auth reverse proxy for connecting to services such as elasticsearch that require v4 signing. Written in Go.

* export AWS_ACCESS_KEY_ID and AWS_SECRET_ACCESS_KEY

```
Usage of ./aws-http-auth-proxy:
  -endpoint_url string
    	the remote endpoint to proxy to (default "localhost")
  -local_port string
    	the localhost port to listen on (default "8080")
  -scheme string
    	the http protocol scheme to use (default "https")
```
