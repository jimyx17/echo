<a href="https://echo.jimyx17.com"><img height="80" src="https://cdn.jimyx17.com/images/echo-logo.svg"></a>

[![Sourcegraph](https://sourcegraph.com/github.com/jimyx17/echo/-/badge.svg?style=flat-square)](https://sourcegraph.com/github.com/jimyx17/echo?badge)
[![GoDoc](http://img.shields.io/badge/go-documentation-blue.svg?style=flat-square)](http://godoc.org/github.com/jimyx17/echo)
[![Go Report Card](https://goreportcard.com/badge/github.com/jimyx17/echo?style=flat-square)](https://goreportcard.com/report/github.com/jimyx17/echo)
[![Build Status](http://img.shields.io/travis/jimyx17/echo.svg?style=flat-square)](https://travis-ci.org/jimyx17/echo)
[![Codecov](https://img.shields.io/codecov/c/github/jimyx17/echo.svg?style=flat-square)](https://codecov.io/gh/jimyx17/echo)
[![Join the chat at https://gitter.im/jimyx17/echo](https://img.shields.io/badge/gitter-join%20chat-brightgreen.svg?style=flat-square)](https://gitter.im/jimyx17/echo)
[![Forum](https://img.shields.io/badge/community-forum-00afd1.svg?style=flat-square)](https://forum.jimyx17.com)
[![Twitter](https://img.shields.io/badge/twitter-@jimyx17-55acee.svg?style=flat-square)](https://twitter.com/jimyx17)
[![License](http://img.shields.io/badge/license-mit-blue.svg?style=flat-square)](https://raw.githubusercontent.com/jimyx17/echo/master/LICENSE)

## Supported Go versions

As of version 4.0.0, Echo is available as a [Go module](https://github.com/golang/go/wiki/Modules).
Therefore a Go version capable of understanding /vN suffixed imports is required:

- 1.9.7+
- 1.10.3+
- 1.11+

Any of these versions will allow you to import Echo as `github.com/jimyx17/echo/v4` which is the recommended
way of using Echo going forward.

For older versions, please use the latest v3 tag.

## Feature Overview

- Optimized HTTP router which smartly prioritize routes
- Build robust and scalable RESTful APIs
- Group APIs
- Extensible middleware framework
- Define middleware at root, group or route level
- Data binding for JSON, XML and form payload
- Handy functions to send variety of HTTP responses
- Centralized HTTP error handling
- Template rendering with any template engine
- Define your format for the logger
- Highly customizable
- Automatic TLS via Let’s Encrypt
- HTTP/2 support

## Benchmarks

Date: 2018/03/15<br>
Source: https://github.com/vishr/web-framework-benchmark<br>
Lower is better!

<img src="https://i.imgur.com/I32VdMJ.png">

## [Guide](https://echo.jimyx17.com/guide)

### Example

```go
package main

import (
  "net/http"
  "github.com/jimyx17/echo/v4"
  "github.com/jimyx17/echo/v4/middleware"
)

func main() {
  // Echo instance
  e := echo.New()

  // Middleware
  e.Use(middleware.Logger())
  e.Use(middleware.Recover())

  // Routes
  e.GET("/", hello)

  // Start server
  e.Logger.Fatal(e.Start(":1323"))
}

// Handler
func hello(c echo.Context) error {
  return c.String(http.StatusOK, "Hello, World!")
}
```

## Help

- [Forum](https://forum.jimyx17.com)
- [Chat](https://gitter.im/jimyx17/echo)

## Contribute

**Use issues for everything**

- For a small change, just send a PR.
- For bigger changes open an issue for discussion before sending a PR.
- PR should have:
  - Test case
  - Documentation
  - Example (If it makes sense)
- You can also contribute by:
  - Reporting issues
  - Suggesting new features or enhancements
  - Improve/fix documentation

## Credits

- [Vishal Rana](https://github.com/vishr) - Author
- [Nitin Rana](https://github.com/nr17) - Consultant
- [Contributors](https://github.com/jimyx17/echo/graphs/contributors)

## License

[MIT](https://github.com/jimyx17/echo/blob/master/LICENSE)
