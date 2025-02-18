## Setup

### Install `pnpm` (or `npm`)

### Install `go`

ensure `$GOPATH/bin` is in `$PATH` by adding to `~/.zshrc` or `~/.bashrc`

```sh
PATH=$PATH:$GOPATH/bin
```

#### Install `just`

https://github.com/casey/just?tab=readme-ov-file#installation

## Start App

### Run containers

```sh
just start-infra
```

### Start backend

```sh
cd server
just run-server
```

### Start frontend

```sh
cd client-web
pnpm install
pnpm dev
```