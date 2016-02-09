# mackerel-check-aws-sns-status

A mackerel(and fitting for Nagios, sensu, ...) plugin for checking SNS status (enabled or cert expired)

## Install


```bash
go get github.com/udzura/mackerel-check-aws-sns-status
```

## Usage

```
Usage:
  mackerel-check-aws-sns-status [OPTIONS]

Application Options:
  -v, --version   Show version (default: false)
  -a, --arn=      Platform application ARN to check
  -w, --warn=     A threshold to warn cert expiration (in days) (default: 30)
  -c, --critical= A threshold to judge critical for cert expiration (in days) (default: 14)

Help Options:
  -h, --help  Show this help message
```

e.g.

```console
$ mackerel-check-aws-sns-status --arn="arn:aws:sns:ap-northeast-1:XXXXXXXX:app/APNS/Dev_iPhone"
Cert is OK: 2017-03-09 07:48:50 +0000 UTC
```

## License

see [`./LICENSE`](./LICENSE)

## Contributing

* Usual GitHub flow(fork, patch and pull request)
