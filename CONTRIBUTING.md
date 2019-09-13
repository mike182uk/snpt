# Contributing

Contributions are **welcome!**

Contributions can be made via a Pull Request on [GitHub](https://github.com/mike182uk/snpt).

## Reporting an Issue

Please report issues via the issue tracker on [GitHub](https://github.com/mike182uk/snpt). For security-related issues, please email the maintainer directly.

## Pull Requests

- **Lint & format changes** - Make sure you run `make lint` &  `make fmt` before committing your code.

- **Add tests where appropriate** - Make sure new features or bug fixes are covered by a test.

- **Document any change in behaviour** - Make sure the README and any other relevant documentation are kept up-to-date.

- **Create topic branches** - i.e `feature/some-awesome-feature`.

- **One pull request per feature**

- **Send coherent history** - Make sure each individual commit in your pull request is meaningful. If you had to make multiple intermediate commits while developing, please squash them before submitting.

## Install project dependencies

```bash
make install
```

## Running the tests

```bash
make test
```

## Building the project

```bash
make build
```

## Compile the protocol buffers

```bash
make proto
```

⚠️ You will need to manually install the [Protobuf Runtime](https://github.com/protocolbuffers/protobuf#protobuf-runtime-installation) before you can compile the protocol buffers.

## Notes

- Package layout is based on https://www.ardanlabs.com/blog/2017/02/package-oriented-design.html
