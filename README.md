# JavaScript driver for [Babelfish](https://github.com/bblfsh/bblfshd) ![Driver Status](https://img.shields.io/badge/status-beta-dbd25c.svg) [![Build Status](https://travis-ci.org/bblfsh/javascript-driver.svg?branch=master)](https://travis-ci.org/bblfsh/javascript-driver) ![Native Version](https://img.shields.io/badge/javascript%20version-8.9.3--r1-aa93ea.svg) ![Go Version](https://img.shields.io/badge/go%20version-1.9-63afbf.svg)

JavaScript driver for [babelfish](https://github.com/bblfsh/bblfshd).


Development Environment
-----------------------

Requirements:
- `docker`
- [`bblfsh-sdk`](https://github.com/bblfsh/sdk) _(go get -u gopkg.in/bblfsh/sdk.v2/...)_
- UAST converter dependencies _(dep ensure --vendor-only)_

To initialize the build system execute: `bblfsh-sdk update`, at the root of the project. This will generate the `Dockerfile` for this driver.

To execute the tests just execute `bblfsh-sdk test`, this will execute the test over the native and the go components of the driver using Docker.

The build is done executing `bblfsh-sdk build`. To evaluate the result using a docker container, execute:
`bblfsh-sdk build test-driver && docker run -it test-driver`.


License
-------

GPLv3, see [LICENSE](LICENSE)



