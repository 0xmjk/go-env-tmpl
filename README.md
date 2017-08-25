[![Go Report Card](https://goreportcard.com/badge/github.com/0xmjk/go-env-tmpl)](https://goreportcard.com/report/github.com/0xmjk/go-env-tmpl)
[![MIT licensed](https://img.shields.io/badge/license-MIT-blue.svg)](https://github.com/0xmjk/go-env-tmpl/blob/master/LICENSE)
[![Travis CI report](https://travis-ci.org/0xmjk/go-env-tmpl.svg?branch=master)](https://travis-ci.org/0xmjk/go-env-tmpl)

Use environmental variables in a Go text template
===

For example

```sh
export FOO_VARIABLE="Hello"
./go-env-tmpl -prefix FOO <<EOF
{{ .FOO_VARIABLE }}
{{ .FOO_NIL | default "World" }}
EOF
```

will result in:

```
Hello
World
```
