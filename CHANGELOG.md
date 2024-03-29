# Changelog

## [1.6.1](https://github.com/stevengregory/sn-update-set-extractor/releases/tag/v1.6.1) (2023-08-05)

### Fixed

- Corrected the `FixScript` struct in the XML parser to properly extract the `script` field from XML files. Previously, the parser was set to extract the `description` field.

## [1.6.0](https://github.com/stevengregory/sn-update-set-extractor/releases/tag/v1.6.0) (2023-07-25)

### Added

- Support for the "Widget Angular Provider" file type. The code extractor can now handle files associated with "Widget Angular Provider" types.
- Updated the file operations and helper functions to include handling for "Widget Angular Provider" files.
- README and CHANGELOG updates to include the information about "Widget Angular Provider" support.

## [1.5.0](https://github.com/stevengregory/sn-update-set-extractor/releases/tag/v1.5.0) (2023-06-20)

### Added

- Support for the "Style Sheet" file type. The code extractor can now handle SCSS files associated with "Style Sheet" types.
- README and CHANGELOG updates to include the information about Style Sheet support.

## [1.4.0](https://github.com/stevengregory/sn-update-set-extractor/releases/tag/v1.4.0) (2023-06-19)

### Added

- Support for the "Page" file type. The code extractor can now handle Page configurations associated with "Page" files.
- README and CHANGELOG updates to include the information about Page support.
- Ignoring `/tmp` and `/temp` directories in the `.gitignore` to avoid committing temporary files.

### Refactored

- Removed the unused `supportedFileTypes` function for cleaner codebase.

## [1.3.0](https://github.com/stevengregory/sn-update-set-extractor/releases/tag/v1.3.0) (2023-06-18)

### Added

- Reorganized the file directory structure for better navigation and categorization of scripts.
- Update the README with the new directory structure.

## [1.2.0](https://github.com/stevengregory/sn-update-set-extractor/releases/tag/v1.2.0) (2023-06-17)

### Added

- Support for the "Angular ng-template" file type. The code extractor can now handle HTML templates associated with "Angular ng-template" files.
- README and CHANGELOG updates to include the information about Angular ng-template support.

## [1.1.0](https://github.com/stevengregory/sn-update-set-extractor/releases/tag/v1.1.0) (2023-06-16)

### Added

- Support for the "Theme" file type. The code extractor can now handle SASS stylesheets associated with "Theme" files.
- README and CHANGELOG updates to include the information about UI Action support.

## [1.0.0](https://github.com/stevengregory/sn-update-set-extractor/releases/tag/v1.0.0) (2023-06-15)

### Added

- Initial public release, providing a utility to streamline the process of extracting code from ServiceNow update sets.
- Automatically parses XML files from the `data` directory, and organizes the extracted code into corresponding directories and files in the `dist` directory.
- Supported file types: Business Rule, Client Script, Fix Script, Header | Footer, Script Include, Scripted REST Resource, UI Script, Widget, and UI Action.
- Includes a Makefile for easy build and execution.
- Provides a detailed README.md for usage and project understanding.
