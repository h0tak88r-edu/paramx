# URL Utility Extractor (paramx)

This command-line tool, `paramx`, is designed to extract URLs from a file based on specified bug types. It supports various bug types including XSS, SQLi, LFI, RCE, IDOR, SSRF, SSTI, redirect and more.

## Features

- Extract URLs based on bug types
- Supports multiple bug types: XSS, SQLi, LFI, RCE, IDOR, SSRF, SSTI, and redirect
- Replace parameter values with custom values
- Add custom templates using the `-tp` flag

## Installation

1. Clone the repository:
    ```sh
    git clone https://github.com/zomasec/paramx.git
    cd paramx
    ```

2. Build the tool:
    ```sh
    go build -o paramx
    ```

## Usage

To use the tool, run the executable with the appropriate flags:

```sh
./paramx -tp <TempletesPath> -l <FileInput> -t <BugType> -rw <ReplaceWith>
```

### Flags

- `-tp` : Directory where YAML configuration files are located. You can use the default templates from [paramx-templetes](https://github.com/zomasec/paramx-templetes) or specify your own.
- `-l`  : Path to a file containing URLs (one per line)
- `-t`  : The type of bug to extract the URLs based on it (xss, sqli, lfi, rce, idor, ssrf, ssti, redirect)
- `-rw` : Replace the parameter value with a custom value

### Examples

1. Extract URLs for XSS bugs using default templates:
    ```sh
    ./paramx -tp ./configs -l urls.txt -t xss
    ```

2. Extract URLs for SQLi bugs and replace parameter values with `' OR '1'='1`:
    ```sh
    ./paramx -tp ./configs -l urls.txt -t sqli -rw "' OR '1'='1"
    ```

3. Use custom templates directory:
    ```sh
    ./paramx -tp /path/to/custom/templates -l urls.txt -t ssrf
    ```

## Templates

The default templates are available at [paramx-templetes](https://github.com/zomasec/paramx-templetes). Users can add their own templates by specifying the `-tp` flag with the path to the custom templates directory.

## Contributing

Contributions are welcome! Please submit a pull request or open an issue to discuss any changes.

## License

This project is licensed under the MIT License.

---

For more information, visit the [repository](https://github.com/zomasec/paramx).