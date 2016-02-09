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
  -a, --arn=  Platform application ARN to check

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
