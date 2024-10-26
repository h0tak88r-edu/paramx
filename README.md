# ParamX README

## Overview

ParamX is a tool designed for extracting interesting subdomains and parameters from URLs. It can be particularly useful for security researchers and penetration testers who are looking for specific types of vulnerabilities such as XSS, SQLi, LFI, RCE, IDOR, SSRF, SSTI, and open redirects.

## Features

- Extracts parameters based on specified bug types.
- Supports custom templates.
- Can update and download YAML configuration templates.
- Processes URLs from files or standard input.
- Custom parameter value replacement.

## Installation

To install ParamX:

```sh
go install github.com/cyinnove/paramx/cmd/paramx@latest

```

## Usage

ParamX is executed via command-line interface (CLI) with several options to customize its behavior. Below are the available flags:

- `-tp` : Directory where YAML configuration files are located.
- `-l` : Path to a file containing URLs (one per line).
- `-tag` : The type of bug to extract the URLs based on it (default: "xss"). Supported values: xss, sqli, lfi, rce, idor, ssrf, ssti, redirect.
- `-rw` : Replace the parameter value with a custom value.
- `-t` : Path to a custom template.
- `-ut` : Update the templates.

### Examples

#### Basic Usage

To extract XSS parameters from a list of URLs provided in a file:

```sh
cat urls.txt | paramx -tag xss
```

#### Using Custom Template

To use a custom template for extraction:

```sh
cat urls.txt | paramx -t /path/to/custom_template.yaml  
```

#### Replacing Parameter Values

To replace the parameter value with a custom value:

```sh
paramx -rw "custom_value" -l urls.txt
```

#### Updating Templates

To update the YAML configuration templates:

```sh
paramx -ut
```

## Contributing

Contributions are welcome! Please fork the repository and submit pull requests.

## License

This project is licensed under the MIT License. See the `LICENSE` file for details.

## Contact

For issues, questions, or suggestions, please open an issue on the [GitHub repository](https://github.com/cyinnove/paramx).

---

### Detailed Example

Here’s a more detailed example of how you might run ParamX with various options:

```sh
paramx -tp /path/to/templates -l urls.txt -tag sqli -rw "injected_value" -t /path/to/custom_template.yaml
```

In this example, ParamX will:

1. Use templates from `/path/to/templates`.
2. Read URLs from `urls.txt`.
3. Extract parameters that are prone to SQL injection.
4. Replace parameter values with `injected_value`.
5. Use a custom template located at `/path/to/custom_template.yaml`.

## Internal Structure

The main package imports necessary modules and handles command-line flag definitions and parsing. The core functionalities include:

1. **Template Handling**:
   - Updating and downloading YAML configuration templates.
2. **URL Reading**:
   - Reading URLs from a file or standard input.
3. **Parameter Grepping**:
   - Extracting parameters based on specified tags.
4. **Logging**:
   - Logging important information and errors.

The `runner` package contains the main logic for parameter extraction, while the `utils` package includes utility functions for reading URLs and handling I/O operations.

---

Thank you for using ParamX! We hope this tool aids you in your security research and penetration testing endeavors. For more information, visit our [GitHub repository](https://github.com/cyinnove/paramx).
