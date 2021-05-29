# Go HTTP Redirect

[![Build Status][build-status-svg]][build-status-url]
[![Go Report Card][goreport-svg]][goreport-url]
[![License][license-svg]][license-url]
[![Heroku][heroku-svg]][heroku-url]

This is a simple HTTP redirect service which can be run as a stand-alone service. It can be optionally run on Heroku.

## Configuration

Set the following Environment Variables:

1. `REDIRECT_URL`: URL where the requestor should be redirected.
1. `HTTP_STATUS`: the HTTP status code to use. `301` is used by default.
1. `PORT`: port number which the service should run on, e.g. `8080`. Note: Heroku will automatically set this environment variable.

## Installation

### Heroku

```sh
$ heroku create
$ git push heroku master
$ heroku open
```
or

[![Deploy](https://www.herokucdn.com/deploy/button.svg)](https://heroku.com/deploy)

## Credits

Heroku support provided via [`github.com/grokify/goheroku`](https://github.com/grokify/goheroku).

 [build-status-svg]: https://github.com/grokify/go-http-redirect/workflows/go%20build/badge.svg
 [build-status-url]: https://github.com/grokify/go-http-redirect/actions
 [goreport-svg]: https://goreportcard.com/badge/github.com/grokify/go-http-redirect
 [goreport-url]: https://goreportcard.com/report/github.com/grokify/go-http-redirect
 [docs-godoc-svg]: https://pkg.go.dev/badge/github.com/grokify/go-http-redirect
 [docs-godoc-url]: https://pkg.go.dev/github.com/grokify/go-http-redirect
 [license-svg]: https://img.shields.io/badge/license-MIT-blue.svg
 [license-url]: https://github.com/grokify/go-http-redirect/blob/master/LICENSE
 [heroku-svg]: https://img.shields.io/badge/%E2%86%91_deploy-Heroku-7056bf.svg?style=flat
 [heroku-url]: https://heroku.com/deploy
