# SN Update Set Extractor

A ServiceNow utility designed to streamline the process of extracting code from update sets.

## Overview

This utility, written in Go, processes ServiceNow update sets. It extracts and organizes code from XML files into corresponding directories and files, making it easier to use in an IDE.

## How to Use

### Prerequisites

- Go (1.16 or later)

### Building & Running

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

### Input & Output

The utility expects to find update sets in the `data` directory. It will process all XML files found in this directory.

The code extracted from the XML files will be placed in the `dist` directory. The directory structure will be organized based on file types. For example:

```
dist/
├── Script Include/
│   └── sys_script_include_fd12c85cc39f55109cdb161ce001318a.js
└── Widget/
    └── MyWidget/
        ├── client_script.js
        ├── link.js
        ├── option_schema.json
        ├── server_script.js
        ├── style.scss
        └── template.html
```

## License

This project is licensed under the terms of the MIT license. See the [LICENSE](LICENSE) file.
