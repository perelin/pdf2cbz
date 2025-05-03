# PDF to CBZ Converter

A simple command-line tool to convert PDF comic files to CBZ format.

## Prerequisites

- Go 1.22 or later
- pdftoppm (from poppler-utils)
- zip command-line tool

## Installation

1. Clone the repository
2. Build the application:

```bash
go build -o pdf2cbz cmd/main.go
```

## Usage

```bash
./pdf2cbz <pdf-file>
```

The tool will:

1. Extract all pages from the PDF as PNG images
2. Create a CBZ archive containing the images
3. Clean up temporary files

The output CBZ file will be created in the same directory as the input PDF file, with the same name but .cbz extension.

## Example

```bash
./pdf2cbz mycomic.pdf
# Creates mycomic.cbz in the same directory
```

## Error Handling

The tool provides clear error messages including:

- Function name where the error occurred
- Input filename
- Specific error details

## Notes

- cmd: pdftoppm -png -f 1 -progress comicfile.pdf ./comictempdir/page
- dev: go build -o pdf2cbz ./cmd/main.go; ./pdf2cbz test.pdf
