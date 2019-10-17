# logtester

Sample Go files to test loggers.

## Usage

See the help:

```shell
go build && ./logtester help
go build && ./logtester errfmt --help
```

See a few examples:

```shell
./logtester errfmt --testCase errorhttp --formatter syslog

./logtester errfmt --testCase errorhttp --formatter text --extractDetails

./logtester errfmt --testCase errorhttp --formatter text --extractDetails --printStructFieldNames

./logtester errfmt --testCase errorhttp --formatter json --callStackInFields --callStackInHTTPProblem --callStackOnConsole --callStackSkipLast 7
```
