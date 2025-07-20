# fast [![Github All Releases](https://img.shields.io/github/downloads/ddo/fast/total.svg?style=flat-square)]()
> Minimal zero-dependency utility for testing your internet download speed from terminal

*Powered by Fast.com - Netflix*

<p align="center"><a href="https://asciinema.org/a/80106"><img src="https://asciinema.org/a/80106.png" width="50%"></a></p>

## Features

- **IPv4-only**: Forces all connections through IPv4
- **Download-only**: Tests only download speed (no upload)
- **Clean output**: No animations, just results
- **Server info**: Shows test server domain and IP
- **Interface binding**: Specify network interface for testing

## Installation

#### Bin

> replace the download link with your os one

> https://github.com/ddo/fast/releases

> below is ubuntu 64 bit example

```sh
curl -L https://github.com/ddo/fast/releases/download/v0.0.4/fast_linux_amd64 -o fast

# or wget
wget https://github.com/ddo/fast/releases/download/v0.0.4/fast_linux_amd64 -O fast

# then chmod
chmod +x fast

# run
./fast
```

#### Docker

> ~10 MB

```sh
docker run --rm ddooo/fast
```

#### Snap

```sh
snap install fast
```

#### Arch Linux (AUR)

```sh
yay -S fast || paru -S fast
```

#### Brew

> *soon*

#### For golang user

> golang user can install from the source code

```sh
go get -u github.com/ddo/fast
```

## Usage

### Basic usage
```sh
$ ./fast
正在连接测试服务器...
测试服务器: xxx.fast.com
服务器 IP: xxx.xxx.xxx.xxx
正在测试下载速度...

下载速度: 125.67 Mbps
```

### Command line options

| Flag | Description |
| ---- | ----------- |
| -k   | Forces output into Kbps |
| -m   | Forces output into Mbps |
| -g   | Forces output into Gbps |
| -i   | Specify network interface to use (e.g., eth0, tun0) |

### Examples

```sh
# Use default network interface
./fast

# Force specific units
./fast -k    # Display in Kbps
./fast -m    # Display in Mbps
./fast -g    # Display in Gbps

# Use specific network interface
./fast -i eth0     # Use eth0 interface
./fast -i tun0     # Use tun0 interface
./fast -i wlan0    # Use wlan0 interface

# Combine options
./fast -i eth0 -m  # Use eth0 interface, display in Mbps
```

### Using specific network interface

When using the `-i` option, the program will:
1. Bind to the specified network interface
2. Show which interface is being used
3. Use only that interface's IPv4 address for the test

Example output:
```sh
$ ./fast -i tun0
使用网络接口: tun0
正在连接测试服务器...
测试服务器: xxx.fast.com
服务器 IP: xxx.xxx.xxx.xxx
正在测试下载速度...

下载速度: 98.45 Mbps
```

## Build

### Docker

```sh
# build alpine binary file from root folder
docker run --rm -v "$PWD":/go/src/fast -w /go/src/fast golang:alpine go build -v -o fast

# build docker image
mv fast build/docker/
cd build/docker/
docker build -t ddooo/fast .
```

### Local Go environment

```sh
# requires Go 1.11 or later
go build -v -o fast
```

### Snap

```sh
cd build/snap/
snapcraft
snapcraft push fast_*.snap
snapcraft release fast <revision> <channel>
```

## Technical Details

- All connections are forced through IPv4 (IPv6 disabled)
- HTTP client uses custom transport with IPv4-only dialer
- When interface is specified, binds to that interface's IPv4 address
- Only tests download speed (as per Fast.com's design)
- No progress bars or animations for cleaner output

## Bug

for bug report just open new issue