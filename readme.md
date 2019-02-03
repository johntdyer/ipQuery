# esquery-go

IP Address whois tool with AWS service lookup.

## Installation

You'll need Go installed:

```
$ go get -u -v github.com/johntdyer/ipquery
```

## Usage

`ipquery` is very easy to use. It will help you perform whois on an ip address. If the IP is an AWS IP it will return the assigned region and AWS service mapping from [AWS ip api](https://ip-ranges.amazonaws.com/ip-ranges.json).

```
Usage of ./ipquery:
  -ip string
    	IP Address to lookup
```

Example:

```
./ipquery -ip 34.234.155.200
 Name         Amazon.com, Inc., US
 Network      AS14618
 Country      US
 Registry     ARIN
 Range        34.224.0.0/12
 IP           34.234.155.200
 Network      34.224.0.0/12
 Services     AMAZON,EC2
 Region       us-east-1
 AllocatedAt  2016-09-12 00:00:00 +0000 UTC
```

## License

MIT
