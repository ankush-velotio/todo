# todo

[![Go Report Card](https://goreportcard.com/badge/github.com/ankush-velotio/todo?style=flat-square)](https://goreportcard.com/report/github.com/ankush-velotio/todo)

## Project Structure

### `.githooks/`
Contains the git hooks to run. It requires the git hooks installed. [Githooks](https://github.com/rycus86/githooks#installation)

### `/cmd`
Main applications for this project.

### `/internal`
Private applications and library code. Code inside this directory is intended to use
in this application only.

### `/test`
Test apps and test data for this project

### `/web`
Web application specific components: static web assets, server side templates and SPAs.
