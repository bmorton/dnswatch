# dnswatch

`dnswatch` is a DNS packet capture utility to output all DNS queries being made on a given interface.

## Usage

```
$ dnswatch --help
Usage of dnswatch:
  -interface="eth0": device interface to listen for DNS packets
```

## Sample output

```
$ sudo dnswatch --interface=en0
Watching for DNS packets on en0...
Thu May 21 01:22:07 2015 - www.yammer.com [29926]
          -> A    134.170.148.18
Thu May 21 01:22:15 2015 - mmm.hm [61774]
          -> A    104.219.250.111
Thu May 21 01:22:23 2015 - api.smoot-apple.com.akadns.net [6709]
          -> A    17.249.105.246
Thu May 21 01:22:23 2015 - clients1.google.com [60990]
          -> A    173.194.33.162
          -> A    173.194.33.163
          -> A    173.194.33.164
          -> A    173.194.33.165
          -> A    173.194.33.166
          -> A    173.194.33.167
          -> A    173.194.33.168
          -> A    173.194.33.169
          -> A    173.194.33.174
          -> A    173.194.33.160
          -> A    173.194.33.161
```

## Installation from Github

```
$ curl -L https://github.com/bmorton/dnswatch/releases/download/v0.1.0/dnswatch-linux-amd64.tar.gz | tar -xz
```

## Release packaging notes

```
$ tar -pcvzf dnswatch-linux-amd64.tar.gz dnswatch
```
