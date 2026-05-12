# NoDub WP Explorer

A powerful, lightweight web scraper and cataloger written in Go (Golang). This tool is designed to interface with the **[nodub.com](https://nodub.com/)** library, providing an automated way to browse, track, and organize premium WordPress themes and plugins distributed under the General Public License (GPL).

## Project Overview

The **NoDub WP Explorer** streamlines the process of monitoring the latest WordPress solutions. It focuses on efficiency, performance, and clean data structure, making it an ideal utility for developers who need to manage large portfolios of WP sites or stay updated on verified digital products.

## Key Features

*   **Automated Cataloging:** Quickly extracts product names, categories, and direct links from the NoDub repository.
*   **High Performance:** Built with Go's concurrency model for fast, non-blocking operations.
*   **Structured Output:** Generates clean data formats (JSON/CSV) ready for integration with other automation tools or PBN managers.
*   **GPL Compliance Focus:** Specifically tailored for original products distributed under the General Public License.

## Getting Started

### Prerequisites
*   Go 1.21 or higher
*   Colly framework (`go get github.com/gocolly/colly/v2`)

### Installation
```bash
git clone [https://github.com/tang47687/nodub.git](https://github.com/tang47687/nodub.git)
cd nodub-explorer
go mod tidy
