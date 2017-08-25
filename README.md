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
