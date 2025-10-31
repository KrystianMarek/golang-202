# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

### Added
- Initial project structure and module setup
- Go 1.24 features package (`pkg/go124/`)
  - Iterator functions for custom sequences
  - `unique` package demonstrations
  - `runtime.AddCleanup` examples
  - Generic type aliases
- OOP patterns package (`pkg/oop/`)
  - Struct embedding and composition
  - Interface polymorphism
  - Component-based design
  - Dependency injection patterns
- Design Patterns (`pkg/oop/patterns/`)
  - Creational: Singleton, Factory, Builder
  - Structural: Adapter, Decorator
  - Behavioral: Observer, Strategy
- Functional programming package (`pkg/functional/`)
  - Higher-order functions (Map, Filter, Reduce)
  - Function composition and currying
  - Immutable data structures
  - Lazy evaluation pipelines
- Go idioms package (`pkg/idioms/`)
  - Interface patterns and duck typing
  - Error handling with `errors.Is` and `errors.As`
  - Concurrency patterns (goroutines, channels)
  - Zero value semantics
- Example applications
  - CLI examples runner (`cmd/examples/`)
  - Game engine integration example
- Documentation
  - Comprehensive README
  - Contributing guidelines
  - Package documentation with godoc
- CI/CD setup
  - GitHub Actions workflows
  - Linting with golangci-lint
  - Automated testing
- Configuration files
  - `.gitignore`
  - `.golangci.yml`
  - `go.mod`

## [0.1.0] - TBD

### Added
- Initial release
- Core package structure
- Basic examples for all major patterns
- Test coverage >80%
- Full documentation

---

## Version History

### Version Numbering

This project uses semantic versioning:
- **MAJOR**: Incompatible API changes
- **MINOR**: New functionality (backward compatible)
- **PATCH**: Bug fixes (backward compatible)

### Release Process

1. Update CHANGELOG.md
2. Update version in documentation
3. Create git tag: `git tag -a v0.1.0 -m "Release v0.1.0"`
4. Push tag: `git push origin v0.1.0`
5. GitHub Actions will create release

[Unreleased]: https://github.com/KrystianMarek/golang-202/compare/v0.1.0...HEAD
[0.1.0]: https://github.com/KrystianMarek/golang-202/releases/tag/v0.1.0

