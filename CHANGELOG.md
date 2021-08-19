# Change Log

All notable changes to the "Go-Web Framework" will be documented in this file.

## [Unreleased]

## [v0.5-5-beta] - 2021-08-19

### Fixed
- Controller raw template with right 'kernel' package

## [v0.5-4-beta] - 2021-07-02

### Added
- Insert readme.md
- Shortcut methods to access on global configuration and cookie store objects.

### Changed

- Go-Web `server:run` command runs outside the IOC container.
- Expose`SingletonServices`

## [v0.5-3-beta] - 2021-06-21

### Fix

- Fix DI by encapsulating custom services into a separate service container.

## [v0.5-2-beta] - 2021-06-21

### Added

- Implemented "View" tool
- Implemented "ValidateRequest" tool
- Implemented "DecodeJsonRequest" tool
- Implemented "RunJob" tool

## [v0.5-1-beta] - 2021-06-16

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
