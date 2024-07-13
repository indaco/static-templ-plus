<h1 align="center" style="font-size: 2.5rem;">
  static-templ-plus
</h1>
<p align="center">
  <a href="https://github.com/indaco/static-templ-plus/blob/main/LICENSE" target="_blank">
    <img src="https://img.shields.io/badge/License-GNU%20GPL-blue?style=flat-square&logo=none" alt="license" />
  </a>
  &nbsp;
  <a href="https://goreportcard.com/report/github.com/indaco/static-templ-plus/" target="_blank">
    <img src="https://goreportcard.com/badge/github.com/indaco/static-templ-plus" alt="go report card" />
  </a>
  &nbsp;
  <a href="https://pkg.go.dev/github.com/indaco/static-templ-plus/" target="_blank">
      <img src="https://pkg.go.dev/badge/github.com/indaco/static-templ-plus/.svg" alt="go reference" />
  </a>
</p>

This repository is an extension of the original [static-templ](https://github.com/nokacper24/static-templ) and has been developed to include additional functionalities. The original repository provided a solid foundation, and this project aims to build upon that work by introducing new features and improvements.

## About This Project

This project is based on the original repository but has been given a new name to differentiate it from the original.

Below are the key enhancements added to this project to fulfill specific use cases. For each enhancement, a pull request (PR) has been submitted to the original repository:

### Fixes

1. ~~waiting for templ commands to complete~~ ([#12]) -> **merged**.
2. ~~ensure path is OS compatible~~ ([#16]) -> **merged**.

### Enhancements

1. ~~Prevent deletion of input directory when `-i` and `-o` have the same value~~ ([#2]) -> **merged**.
2. ~~Ensure generated HTML filename matches its corresponding templ file~~ ([#4]) -> **merged**.
3. ~~Enable direct execution of `templ fmt` and `templ generate` from `static-templ-plus`~~ ([#5]) -> **merged**.
4. ~~Add `version` subcommand~~ ([#10]) -> **merged**.
5. bump `templ` to v747 and add a new flag `m` for operational modes [#19].

### Refactor

1. ~~improve the overall quality code measures (`gofmt`, `go vet`, `go lint` and `gocyclo`) reported by goreportcard and golangci-lint~~([#18]) -> **merged**.

### Build & Ci

1. ~~Version Management~~ ([#14]) -> **merged**.

### Chore

1. ~~simplify usage func with explicit argument indexes~~ ([#17]) -> **merged**.

## Compatibility

This project aims to remain compatible with the original repository, ensuring that you can use it as a drop-in replacement if needed. However, please review the changes and test your integration thoroughly.

I actively keep this project updated with the latest changes from the original repository. If and when the PRs are accepted, I will evaluate switching back to the original repository to ensure alignment and maintainability.

By using this version, you can take advantage of the new features while still benefiting from the updates and improvements made to the original project.

## Installation

To avoid confusion with the original repository, the Go module name has been changed.

```bash
go install github.com/indaco/static-templ-plus@latest
```

## Usage

The usage of this module is the same as the original one, with the following enhancements:

### Modes

A new flag `-m` has been added to address two specific use-cases:

- **Bundle Mode** (`bundle`): Generates HTML files in the specified output directory, mirroring the structure of the input directory. This mode is useful for converting a full set of pages. It reflects the original `static-templ` way of working.
- **Inline Mode** (`inline`): Generates HTML files in the same directory as their corresponding `.templ` files. This mode is useful for smaller projects, single-component development or for documenting components.

You can use it as follows:

```bash
Usage of static-templ-plus:
static-templ-plus [flags] [subcommands]

Flags:
  -m  Set the operational mode: bundle or inline. (default "bundle").
  -i  Specify input directory (default "web/pages").
  -o  Specify output directory (default "dist").
  -f  Run templ fmt.
  -g  Run templ generate.
  -d  Keep the generation script after completion for inspection and debugging.

Subcommands:
  version  Display the version information.

Examples:
  # Specify input and output directories
  static-templ-plus -i web/demos -o output

  # Specify input directory, run templ generate and output to default directory
  static-templ-plus -i web/demos -g=true

  # Display the version information
  static-templ-plus version
```

## Assumptions

> The assumptions remain the same. The following information is from the original repository

Templ components that will be turned into html files must be **exported**, and take **no arguments**. If these conditions are not met, the component will be ignored. Your components must be in the *input* directory, their path will be mirrored in the *output* directory.

By default (`mode=bundle`) all files other than `.go` and `.templ` files will be copied to the output directory, preserving the directory structure. This allows you to include any assets and reference them using relative paths.

## Contribution

Contributions are welcome! If you have suggestions for improvements or new features, please submit an issue or a pull request.

Before submitting a pull request, please follow these steps to ensure a smooth and consistent development process:

### Setting Up Git Hooks

We use Git hooks to automate versioning and ensure code quality. After cloning the repository, you must set up the Git hooks by running the following script. This step ensures that the hooks are properly installed and executed when needed.

1. Clone the repository:

    ```bash
    git clone https://github.com/indaco/static-templ-plus.git
    cd static-templ-plus
    ```

2. Run the setup script to install the Git hooks:

    **For Unix-based systems (Linux, macOS):**

    ```bash
    ./setup-hooks.sh
    ```

    **For Windows systems:**

    ```cmd
    setup-hooks.bat
    ```

By running the appropriate setup script, you ensure that the pre-commit hook is properly installed. This hook will automatically update the version number in the `.version` file and stage it for commit.

## Acknowledgements

We would like to acknowledge the creators of the original repository for their excellent work. This project would not have been possible without their contributions.

## License

This project is licensed under the same terms as the original repository. For more details, see the [LICENSE](./LICENSE) file.

<!-- Resources -->
[#2]: https://github.com/nokacper24/static-templ/pull/2
[#4]: https://github.com/nokacper24/static-templ/pull/4
[#5]: https://github.com/nokacper24/static-templ/pull/5
[#10]: https://github.com/nokacper24/static-templ/pull/10
[#12]: https://github.com/nokacper24/static-templ/pull/12
[#14]: https://github.com/nokacper24/static-templ/pull/14
[#16]: https://github.com/nokacper24/static-templ/pull/16
[#17]: https://github.com/nokacper24/static-templ/pull/17
[#18]: https://github.com/nokacper24/static-templ/pull/18
[#19]: https://github.com/nokacper24/static-templ/pull/19
