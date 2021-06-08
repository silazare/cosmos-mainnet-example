# Cosmos blockchain node exporter

## Build steps

1) Build binary

```shell
go mod init cosmos_exporter
go get github.com/prometheus/client_golang
go get
go build
```

2) Move binary into `/home/cosmos/go/bin` folder

## Usage example

1) Create a service file from example

```shell
/etc/systemd/system/cosmos-exporter.service
```

2) Install systemd service

```shell
sudo systemctl daemon-reload
sudo systemctl enable cosmos-exporter.service
sudo systemctl start cosmos-exporter.service
sudo systemctl status cosmos-exporter.service
journalctl -e -u cosmos-exporter.service
```

3) Scrape endpoint

```shell
curl -s localhost:9201/metrics | grep -v \# | grep cosmos
```

## Exposed Cosmos metrics

- Number of peers in blockchain network
- Current block number
- Time out of sync between current block time and time now
