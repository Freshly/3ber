# 3ber

![logo](logo.png)

One more than [tuber](github.com/Freshly/tuber).

3ber is a Kubernetes ([GKE](https://cloud.google.com/kubernetes-engine)) release manager and training tool for Freshly's CI/CD pipeline.

The goal of 3ber is to obsolete itself by training developers to use the underlying toolchain.

3ber can always be referenced if your goldfish brain resets at midnight like mine does.

## Prerequisites

Install the gcloud CLI:

https://cloud.google.com/sdk/docs/install

Install kubectl:

https://kubernetes.io/docs/tasks/tools/

## Installation

3ber supports Linux, Mac, and Windows.

Head over to the (release)[https://github.com/Freshly/3ber/releases] page and download the relevant binary. Add executable permissions, if necessary, and move the binary to your PATH. Alternatively, you can use the automated instructions below.

### Mac / Linux

Run the following commands in your terminal:

```
VERSION=0.4.1
curl -sSfL https://github.com/Freshly/3ber/releases/download/${VERSION}/3ber-`uname -s`-`uname -m` -o /tmp/3ber
chmod +x /tmp/3ber
sudo mv /tmp/3ber /usr/local/bin/
```

## Usage

## Building

To cross-compile this project for all target platforms, run:

`just build`

To build the Docker image for the project, run:

`just build_docker`
