github.com/therealplato/add
===========================

Adds newline-separated integers from stdin until EOF then output the total to stdout.

Ignores whitespace-only lines.

Logs non-integer lines to stderr.

install
-------
```sh
go get -u github.com/therealplato/add
```

example
-------
total bytes in `/etc/*` files:
```sh
$ find /etc/ -type f | xargs stat --format '%s' | add
find: ‘/etc/cups/certs’: Permission denied
3125911
```
