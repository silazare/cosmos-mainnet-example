# Cosmos blockchain node deploy in Mainnet

Test task for Cosmos blockchain node deploy in GCP.

## Requirements

- Ansible - `v2.9.x` or later
- Terraform - `v0.15.x` or later
- Packer - `v1.6.x` or later

## Infrastructure deploy steps

1) Prepare and configure Service Account in GCP with required permissions for Packer and Terraform.  

2) Build Base Packer image with all pre-requisites:

```shell
cd packer
packer build -var 'project_id=<PROJECT_NAME>' -var 'network=default' ubuntu2004-gcp.json
packer build -var 'project_id=<PROJECT_NAME>' ubuntu2004-gcp.json
```

3) Create Terraform infrastructure:

```shell
cd terraform
make plan
make apply
```

4) SSH to the instance

```shell
ssh cosmos@<NODE_IP> -i ~/.ssh/cosmos
```

5) Check OS hardening with Lynis:

```shell
sudo -i
cd /var/lib/lynis
./lynis audit system
```

## Semi-automatic Blockchain node setup

1) Run provison playbook:

```shell
cd ansible
ansible-playbook -u cosmos -i '<NODE_IP>,' --private-key ~/.ssh/cosmos cosmos-provision.yml
```

2) On the node fetch the latest genesis.json and start the service:

```shell
cd ~/.gaia/config
wget https://github.com/cosmos/mainnet/raw/master/genesis.cosmoshub-4.json.gz
gzip -d genesis.cosmoshub-4.json.gz
mv genesis.cosmoshub-4.json ~/.gaia/config/genesis.json
```

3) Start and check the service:

```shell
sudo systemctl start gaiad
sudo systemctl status gaiad
journalctl -e -u gaiad.service
```

## Manual Blockchain node setup

https://hub.cosmos.network/main/gaia-tutorials/join-mainnet.html

1) Export environment variables:

```shell
export GOPATH=$HOME/go
export PATH=$GOPATH/bin:$PATH
```

2) Create a system user for running the service:

```shell
sudo mkdir -p /opt/go/bin
sudo useradd -m -d /opt/gaiad --system --shell /usr/sbin/nologin gaiad
sudo -u gaiad mkdir -p /opt/gaiad/config
```

3) Install gaiad:

```shell
git clone -b v4.2.1 https://github.com/cosmos/gaia
cd gaia
make install
```

4) Copy the binaries to /opt/go/bin/

```shell
sudo cp $HOME/go/bin/gaia* /opt/go/bin/
sudo chown gaiad:gaiad /opt/go/bin/gaiad
sudo chmod 755 -R /opt/go/bin
```

5) Init gaiad:

```shell
gaiad init $(hostname)
```

6) Fetch latest genesis.json:

```shell
wget https://github.com/cosmos/mainnet/raw/master/genesis.cosmoshub-4.json.gz
gzip -d genesis.cosmoshub-4.json.gz
mv genesis.cosmoshub-4.json $HOME/.gaia/config/genesis.json
```

7) Reset gaiad:

```shell
gaiad unsafe-reset-all
```

8) Add seeds into `config.toml`

9) Create /etc/systemd/system/gaiad.service from template.

```shell
sudo systemctl daemon-reload
sudo systemctl enable gaiad.service
```

10) Start the application:

```shell
sudo systemctl start gaiad
sudo systemctl status gaiad
journalctl -e -u gaiad.service
```
