# Hello WASM

Hellow WASM is a venture into the world of WASM using golang.
This repository contains both the server and the UI, both written using go.

The `go-app` golang package makes this possible.

I have overriden the default ResourceProvider so that the static files are served from a Brotli compressed virtual file system.
This is acheived through the `broccoli` golang package.

While you can run this app without SSL, you miss out on some of the benefits of a PWA, like (installablity, proper service workers, etc...).
So the Makefile will generate one for you upon first build that you should add as trusted on your system.

## How to run

` make run `