# SN Update Set Extractor

A ServiceNow utility designed to streamline the process of extracting code from update sets.

## Overview

This utility, written in Go, processes ServiceNow update sets. It extracts and organizes code from XML files into corresponding directories and files, making it easier to use in an IDE.

## Prerequisites

- [Go (1.16 or later)](https://go.dev/doc/install)

## Getting Started

Clone the repository & change directory:

```sh
git clone https://github.com/stevengregory/sn-update-set-extractor
cd sn-update-set-extractor
```

## Load Update Sets

In the root of the project, create a `data` directory.

```sh
mkdir data
```

Load your update sets into this directory. When running the project, the utility will process all XML files it contains.

## Building & Running

To build and run the project, use the following commands:

```sh
make clean  # Remove any previous build and output files
make build  # Compile the Go project
make run    # Execute the compiled binary
```

Run the `make` command to perform all of these steps in one go:

```sh
make
```

## Generated Output

After building & running the project, the extracted code is then structured and saved in the `dist` directory. Each file type gets its own directory for easy navigation. For example, the output directory structure might look like:

```
dist/
├── Business Rule/
│   └── MyBusinessRule.js
├── Script Include/
│   └── MyScriptInclude.js
└── Widget/
    └── MyWidget/
        ├── client_script.js
        ├── link.js
        ├── option_schema.json
        ├── server_script.js
        ├── style.scss
        └── template.html
```

## Supported File Types

Currently, the extraction & organization of the following ServiceNow file types is supported:

- Business Rule
- Client Script
- Fix Script
- Header | Footer
- Script Include
- Scripted REST Resource
- UI Action
- UI Script
- Widget

Efforts to support additional file types are ongoing. If there's a specific file type you'd like to see supported, please create an issue in the GitHub repository.

## Changelog

For a detailed breakdown of all changes made to this project, see the [CHANGELOG](CHANGELOG.md) file.

## License

This project is licensed under the terms of the MIT license. See the [LICENSE](LICENSE) file.
