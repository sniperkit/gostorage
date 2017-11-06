# gostorage

[![Release][release-image]][release-url] [![Build Status][build-image]][build-url] [![Coverage][coverage-image]][coverage-url] [![GoDoc][doc-image]][doc-url]

A Go library that provides a generic interface for various object storages.

**UNDER DEVELOPMENT**

## Installation

```bash
go get github.com/devfacet/gostorage
```

## Usage

*TODO*

## Build

```bash
go build .
```

## Test

```bash
./test.sh
```

## Release

```bash
git add CHANGELOG.md # update CHANGELOG.md
./release.sh v1.0.0  # replace "v1.0.0" with new version

git ls-remote --tags # check the new tag
```

## Contributing

- Code contributions must be through pull requests
- Run tests, linting and formatting before a pull request (`test.sh`)
- Pull requests can not be merged without being reviewed
- Use "Issues" for bug reports, feature requests and discussions
- Do not refactor existing code without a discussion
- Do not add a new third party dependency without a discussion
- Use semantic versioning and git tags for versioning

## License

Licensed under The MIT License (MIT)  
For the full copyright and license information, please view the LICENSE.txt file.


[release-url]: https://github.com/devfacet/gostorage/releases/latest
[release-image]: https://img.shields.io/github/release/devfacet/gostorage.svg

[build-url]: https://travis-ci.org/devfacet/gostorage
[build-image]: https://travis-ci.org/devfacet/gostorage.svg?branch=master

[coverage-url]: https://coveralls.io/github/devfacet/gostorage?branch=master
[coverage-image]: https://coveralls.io/repos/devfacet/gostorage/badge.svg?branch=master&service=github

[doc-url]: https://godoc.org/github.com/devfacet/gostorage
[doc-image]: https://godoc.org/github.com/devfacet/gostorage?status.svg
