# types

[![Go Reference](https://pkg.go.dev/badge/github.com/gitamix/types.svg)](https://pkg.go.dev/github.com/gitamix/types)
[![License: MIT](https://img.shields.io/badge/License-MIT-blue.svg)](LICENSE)
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/gitamix/types)
[![Release](https://img.shields.io/github/v/release/gitamix/types?style=flat)](https://github.com/gitamix/types/releases)

A Go library of git domain types — `Branch`, `Name`, `Commit`, `Message`, `Subject`, and more.
It turns raw values from your own workflow into typed structures,
so every gitamix project operates on a shared vocabulary
instead of plain strings.

> [!NOTE]
> The project is not a CLI tool — it is a pure Go library.
> It never runs git commands itself: you obtain branch names and commit
> messages in your own workflow (for example, with `git branch` and `git log`),
> wrap them into these types, and pass them to other tools.

## Features

- **Branch types** — `Branch` and `Name`, with ticket extraction
  from branch names such as `feature/TASK-123`.
- **Commit types** — `Commit`, `Hash` (with a short 7-char form),
  `Message`, `Subject`, `Type`, `Scope`, `Description`, and `Body`
  for conventional-commit-shaped messages like `feat(ui): add new button`.
- **Message parsing** — `ParseMessage` splits a raw message into subject
  and body; `ParseSubject` breaks the subject down into type, scope,
  and description. Both accept a `string` or `[]byte`.
- **Ticket integration** — `ticket.ParseTicket` and the `Ticket(re)`
  accessors extract task-tracker IDs with your own regular expression
  (the first capturing group is the ticket name). `Message.Ticket` reads
  the raw subject, so a leading ticket prefix such as
  `[TASK-123] feat: …` survives parsing.
- **Zero dependencies** — only the Go standard library at runtime
  (`testify` is used for tests only).

## Installation

Run the command in terminal:

```sh
go get github.com/gitamix/types
```

## Documentation

Full API reference is available on [pkg.go.dev](https://pkg.go.dev/github.com/gitamix/types)

## Ecosystem

These types are the shared vocabulary
of the [gitamix](https://github.com/gitamix) tooling
and every gitamix project consumes them instead of raw strings or internal objects.

## Requirements

- [Go](https://go.dev/) 1.23.4 or later.
- No third-party runtime dependencies — the standard library is enough.

The library does not run git commands itself.
Obtain branch names and commit messages in your own workflow,
wrap them into `types` values, and pass them to the tools you build.

## Contributing

Want to contribute?
Read [CONTRIBUTING.md](CONTRIBUTING.md)
for the full workflow, repository requirements, and Pull Request process.

Please open an issue to discuss large or breaking changes before implementing.

## License

This project is licensed under the MIT License — see the [LICENSE](LICENSE) file for details.

## Author / Contact

Maintained by [Ilya Sitnikov](https://github.com/gitamix)
