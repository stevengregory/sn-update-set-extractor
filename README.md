# SN Update Set Extractor

A ServiceNow utility designed to streamline the process of extracting code from update sets.

## Overview

This utility, written in Go, processes ServiceNow update sets. It extracts and organizes code from XML files into corresponding directories and files, making it easier to use in an IDE.

## Prerequisites

- Go (1.16 or later)

## Building & Running

You can build and run the project using the `make` command. This command executes the steps defined in the Makefile, which are:

1. `clean`: Remove any previous build and output files.
1. `build`: Compile the Go project.
1. `run`: Execute the compiled binary.

To perform all of these steps in one go, use:

```sh
make
```

If you prefer to run these steps individually, you can do so using the following commands:

```sh
make clean  # Remove any previous build and output files
make build  # Compile the Go project
make run    # Execute the compiled binary
```

## Input & Output

Load your ServiceNow update sets into the `data` directory, and the utility will process all XML files it contains.

The extracted code is then structured and saved in the `dist` directory. Each file type gets its own directory for easy navigation. For example, the output directory structure might look like:

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
- UI Script
- Widget

Efforts to support additional file types are ongoing. If there's a specific file type you'd like to see supported, please create an issue in the GitHub repository.

## License

This project is licensed under the terms of the MIT license. See the [LICENSE](LICENSE) file.
