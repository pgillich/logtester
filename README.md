# logtester

Sample Go files to test loggers.

## Usage

See the help:

```shell
go build && ./logtester help
go build && ./logtester errorformatter --help
```

See a few examples:

```shell
./logtester errorformatter --testCase errorhttp --formatter syslog

./logtester errorformatter --testCase errorhttp --formatter text --extractDetails

./logtester errorformatter --testCase errorhttp --formatter text --extractDetails --printStructFieldNames

./logtester errorformatter --testCase errorhttp --formatter json --callStackInFields --callStackInHTTPProblem --callStackOnConsole --callStackSkipLast 7
```
