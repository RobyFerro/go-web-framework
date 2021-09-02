# Change Log

All notable changes to the "Go-Web Framework" will be documented in this file.

## [Unreleased]
### Changed
- Abstract CLI command structure
- Improved kernel registry removing useless declaration.
- Implemented HTTP router in base entities

## [v0.7.1-beta] - 2021-09-02
### Fixed
- Issue #17
- Issue #16

### Added
- Support for request validation in routers
- 
## [v0.7.0-beta] - 2021-08-27
### Changed
- Changed .yml routing system in favour of the router module. Now every route/group have to be registered directly in go structure.
- Changed .yml configuration. Now system and custom configuration are located into the new config module.

### Removed
- Removed Mail struct from configuration


## [v0.6.3-beta] - 2021-08-27
### Fixed
- Fixed cookie store issue

## [v0.6.2-beta] - 2021-08-27
### Changed
- Changed middleware architecture. Now every middleware consists in an isolated structure.

## [v0.6.1-beta] - 2021-08-27
### Fix
- Resolved command service container issue

## [v0.6.0-beta] - 2021-08-25
### Changed
- Split command line interfaces and http server.

## [v0.5.9-beta] - 2021-08-23
### Added
- Gzip compression over HTTP responses

### Fixed
- Fixed IOC arguments order issue

## [v0.5.8-beta] - 2021-08-20

### Fixed
- Fixed IOC arguments order issue

## [v0.5.7-beta] - 2021-08-19

### Fixed
- Implemented command ioc container

## [v0.5.6-beta] - 2021-08-19

### Added
- Merged factory and singleton container abstraction

## [v0.5.5-beta] - 2021-08-19

### Fixed
- Controller raw template with right 'kernel' package

## [v0.5.4-beta] - 2021-07-02

### Added
- Insert readme.md
- Shortcut methods to access on global configuration and cookie store objects.

### Changed

- Go-Web `server:run` command runs outside the IOC container.
- Expose`SingletonServices`

## [v0.5.3-beta] - 2021-06-21

### Fix

- Fix DI by encapsulating custom services into a separate service container.

## [v0.5.2-beta] - 2021-06-21

### Added

- Implemented "View" tool
- Implemented "ValidateRequest" tool
- Implemented "DecodeJsonRequest" tool
- Implemented "RunJob" tool

## [v0.5.1-beta] - 2021-06-16

### Fix

- Fix CLI issue

## [v0.5-beta] - 2021-06-16

### Changed

- Changed repository
- Improved structure with an approach og Clean Architecture

### Removed

- Server daemon cli command
- HTTP Load cli command
- Show route cli command

## [v0.4.4-beta] - 2020-08-30

### Fix

- Warning messages creating new service in Windows

## [v0.4.3-beta] - 2020-08-29

### Removed

- Syscall used to change user/group

## [v0.4.2-beta] - 2020-08-28

### Added

- Controllers and Models auto-register function (If created with Alfred).

## [v0.4.1-beta] - 2020-08-27

### Fix

- Generic command improvements

## [v0.4-beta] - 2020-08-27

### Removed

- Removed some CLI commands in favour of new "Alfred" tool

## [v0.3.7-beta] - 2020-08-07

### Fix

- Fixes #1

## [v0.3.6-beta] - 2020-04-14

### Changed

- Change JWT lifetime with a custom duration

`func (c *Auth) NewToken(key string, duration time.Duration) (string, bool)`

## [v0.3.5-beta] - 2020-04-14

### Fix

- Removed routines from kernel initialization.

## [v0.3.4-beta] - 2020-03-26

### Fix

- Error in route resolution when two routes with same path are declared

## [v0.3.3-beta] - 2020-02-24

### Added

- New HTTP load CLI tool

## [v0.3.2-beta] - 2020-02-24

### Fixed

- Wrong package name on new commands

## [v0.3.1-beta] - 2020-02-24

### Fixed

- Wrong CLI command directory

## [v0.3.0-beta] - 2020-02-24

### Changed

- Basic type structure: controllers, services, models, and middleware.

### Removed

- Removed external configuration.

### Fixed

- Fix authentication flow in JWTAuth (other fix)

## [v0.2.2-beta] - 2020-02-19

### Fixed

- Fix authentication flow in JWTAuth (other fix)

## [v0.2.1-beta] - 2020-02-19

### Added

- Insert changelog

### Fixed

- Fix authentication flow in JWTAuth

## [v0.2-beta] - 2020-02-19

### Added

- Insert app.key generate command

## [v0.1-beta] - 2020-02-17

### Added

- First version
