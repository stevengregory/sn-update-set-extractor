# Changelog

## [1.0.0](https://github.com/stevengregory/sn-update-set-extractor/releases/tag/v1.0.0) (2023-06-15)

### Added

- Initial public release, providing a utility to streamline the process of extracting code from ServiceNow update sets.
- Automatically parses XML files from the `data` directory, and organizes the extracted code into corresponding directories and files in the `dist` directory.
- Supported file types: Business Rule, Client Script, Fix Script, Header | Footer, Script Include, Scripted REST Resource, UI Script, Widget, and UI Action.
- Includes a Makefile for easy build and execution.
- Provides a detailed README.md for usage and project understanding.