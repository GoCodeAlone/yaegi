# Changelog

All notable changes to this fork will be documented in this file.

This project is a maintained fork of [traefik/yaegi](https://github.com/traefik/yaegi).

## [v0.17.0] - 2026-02-22

### Breaking Changes

- **Module path changed** from `github.com/traefik/yaegi` to `github.com/GoCodeAlone/yaegi`
  - Update your imports accordingly
- **Minimum Go version** raised to Go 1.26

### Added

- **Generic function imports** (upstream PR #1647): Import generic functions from extracted Go code as interpreted code using `//yaegi:add` comment directive. Generic code is wrapped in `GenericFunc` string type for runtime interpretation.
- **Extract speedup** (upstream PR #1642): 11-45x performance improvement for `yaegi extract` by switching from `go/importer` to `golang.org/x/tools/go/packages` with caching support.
- **Extract outer arg** (upstream PR #1638): New `Outer` field on `Extractor` struct to support generating import code outside the standard stdlib path.
- **Extract underscore params** (upstream PR #1708): `yaegi extract` now correctly handles underscore parameter names in interfaces, generating proper wrapper code.
- **Eval/EvalPath panic recovery** (upstream PR #1560): `Eval()` and `EvalPath()` now recover from internal panics and return them as `Panic` errors with call stacks, preventing interpreted code from crashing the host process.

### Fixed

- **Binary channel type alias nil pointer** (GoCodeAlone): Fixed nil pointer dereference when sending to a channel that is a type alias defined in a binary package. The `send` function now uses the `elem()` method which correctly handles both source-defined channels and binary type aliases.
- **Re-import of identical package** (upstream PR #1551): Fixed false "redeclared in this block" errors when the same binary package is imported across multiple source files during multi-pass type analysis.
- **Binary-to-source interface conversion** (upstream PR #1562): Fixed interface type assertion between binary and source interfaces, allowing proper conversion when interfaces are defined in both compiled and interpreted code.
- **Nil type error masking** (GoCodeAlone): Generic type constraint nil type check no longer masks more specific error messages (e.g., "non-constant array bound" errors are now properly reported).
- **Generic type constraint crash** (upstream PR #1647): Fixed nil type crash when using generic type constraint interfaces, returning a descriptive error instead of panicking.
- **isGeneric nil guard** (upstream PR #1647): Added nil safety check in `isGeneric()` to prevent panics when checking function type nodes.

### Changed

- Upgraded `golang.org/x/tools` from v0.22.0 to v0.42.0
- Upgraded `golang.org/x/mod` from v0.18.0 to v0.33.0
- Upgraded `golang.org/x/sync` from v0.7.0 to v0.19.0
- Applied `go fix` modernization (144 fixes across 16 files)
- Applied `go fmt` formatting updates

### Upstream PRs Incorporated

| PR | Title | Author |
|----|-------|--------|
| [#1551](https://github.com/traefik/yaegi/pull/1551) | Fix re-import of identical package | @he11olx |
| [#1560](https://github.com/traefik/yaegi/pull/1560) | Recover in Eval and EvalPath | @Bai-Yingjie |
| [#1562](https://github.com/traefik/yaegi/pull/1562) | Fix bin interface to src interface conversion | @laushunyu |
| [#1638](https://github.com/traefik/yaegi/pull/1638) | Add outer arg for extract command | @ludanfeng |
| [#1642](https://github.com/traefik/yaegi/pull/1642) | Use x/tools/go/packages for extract speedup | @kkoreilly |
| [#1647](https://github.com/traefik/yaegi/pull/1647) | Import generic functions from extracted code | @rcoreilly |
| [#1708](https://github.com/traefik/yaegi/pull/1708) | Extract: support underscore params | @nelsam |

[v0.17.0]: https://github.com/GoCodeAlone/yaegi/releases/tag/v0.17.0
