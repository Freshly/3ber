# 3ber

![logo](logo.png)

One more than [tuber](https://github.com/Freshly/tuber).

3ber is a Kubernetes ([GKE](https://cloud.google.com/kubernetes-engine)) release manager and training tool for Freshly's CI/CD pipeline.

The goal of 3ber is to obsolete itself by training developers to use the underlying toolchain.

## Prerequisites

Install the gcloud CLI:

https://cloud.google.com/sdk/docs/install

Install kubectl:

https://kubernetes.io/docs/tasks/tools/

Install helm:

https://helm.sh/docs/intro/install/

## Installation

3ber supports Mac, Linux, and Windows.

### Mac / Linux

Run the following commands in your terminal:

```
VERSION=`curl -sI https://github.com/Freshly/3ber/releases/latest|grep location:|tr ' ' '\n'|tail -n1|cut -d'/' -f8|tr -d '\r'`
curl -sSfL https://github.com/Freshly/3ber/releases/download/${VERSION}/3ber-`uname -s`-`uname -m` -o /tmp/3ber
chmod +x /tmp/3ber
sudo mv /tmp/3ber /usr/local/bin/
```

### Windows

Head over to the [release](https://github.com/Freshly/3ber/releases) page and download the Windows binary. Move the binary to your PATH, maybe `C:\Windows\system32`.

### Installing From Source

If you have the `go` build toolchain installed, you can build from main:

```
go install github.com/freshly/3ber
```

You will notice version information is not present:

```
$ 3ber version
version dev, git commit HEAD
```

## Usage

You will want to `3ber auth` first to authenticate with Google Cloud and populate your Kube config file.

The 3ber CLI embeds usage information for all commands and subcommands. If you simply run `3ber`, you will see this usage message:

```
Kubernetes release manager and training tool for Freshly's CI/CD pipeline

Usage:
  3ber [command]

Available Commands:
  argo        manage the argo continuous delivery pipeline
  auth        authenticate to google cloud and populate kubernetes cluster contexts
  completion  Generate the autocompletion script for the specified shell
  context     manage kubernetes cluster contexts
  helm        manage helm charts
  help        Help about any command
  version     print the program version

Flags:
  -h, --help    help for 3ber
  -q, --quiet   disable info logging
  -v, --voice   enable voice synthesizer

Use "3ber [command] --help" for more information about a command.
```

## Building

To cross-compile this project for all target platforms, run:

`just build`

To build the Docker image for the project, run:

`just build_docker`
