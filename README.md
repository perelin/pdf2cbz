# PDF to CBZ Converter

A simple command-line tool to convert PDF comic files to CBZ format.

## Prerequisites

- Go 1.22 or later
- pdftoppm (from poppler-utils)
- zip command-line tool

## Installation

### Local Development Installation

1. Clone the repository:

```bash
git clone https://github.com/perelin/pdf2cbz.git
cd pdf2cbz
```

2. Install the application locally:

```bash
go install ./cmd/pdf2cbz
```

This will install the binary in your `$GOPATH/bin` directory, making it available system-wide.

### Remote Installation

To install directly from the GitHub repository:

```bash
# Install latest version
go install github.com/perelin/pdf2cbz/cmd/pdf2cbz@latest

```

This will download and install the specified version of the application.

## Usage

```bash
pdf2cbz <pdf-file>
```

The tool will:

1. Extract all pages from the PDF as images
2. Create a CBZ archive containing the images
3. Clean up temporary files

The output CBZ file will be created in the same directory as the input PDF file, with the same name but .cbz extension.

## Example

```bash
pdf2cbz mycomic.pdf
# Creates mycomic.cbz in the same directory
```

## Notes

- cmd: pdftoppm -png -f 1 -progress comicfile.pdf ./comictempdir/page
- dev: go build -o pdf2cbz ./cmd/main.go; ./pdf2cbz test.pdf
