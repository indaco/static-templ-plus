# static-templ-plus

This repository is an extension of the original [static-templ](https://github.com/nokacper24/static-templ) and has been developed to include additional functionalities. The original repository provided a solid foundation, and this project aims to build upon that work by introducing new features and improvements.

## About This Project

This project is based on the original repository but has been given a new name to differentiate it from the original. Below are the key enhancements added to this project to fulfill specific use cases. For each enhancement, a pull request (PR) has been submitted to the original repository:

- **Enhancement 1**: ~~Prevent deletion of input directory when `-i` and `-o` have the same value ~~ ([#2]) -> **merged**.
- **Enhancement 2**: ~~Ensure generated HTML filename matches its corresponding templ file~~ ([#4]) -> **merged**.
- **Enhancement 3**: ~~Enable direct execution of `templ fmt` and `templ generate` from `static-templ-plus`~~ ([#5]) -> **merged**.
- **Enhancement 4**: Add `version` subcommand ([#10]).

I actively keep this project updated with the latest changes from the original repository. If and when the PRs are accepted, I will evaluate switching back to the original repository to ensure alignment and maintainability.

By using this version, you can take advantage of the new features while still benefiting from the updates and improvements made to the original project.

## Compatibility

This project aims to remain compatible with the original repository, ensuring that you can use it as a drop-in replacement if needed. However, please review the changes and test your integration thoroughly.

## Installation

To avoid confusion with the original repository, the Go module name has been changed.

```bash
go install github.com/indaco/static-templ-plus@latest
```

## Usage

The usage of this module is the same as the original one, with the following enhancements:

- Input and output directories can be the same.
- Generated HTML filename matches its corresponding templ file.
- A new flag `-f` has been added to eliminate the need for users to run `templ fmt` separately.
- A new flag `-g` has been added to eliminate the need for users to run `templ generate` separately.
- A new flag `-d` has been added to keep the generation script after completion for inspection and debugging.

You can use it as follows:

```bash
Usage of static-templ-plus:
static-templ-plus [flags] [subcommands]

Flags:
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

All files other than `.go` and `.templ` files will be copied to the output directory, preserving the directory structure. This allows you to include any assets and reference them using relative paths.

## Contribution

Contributions are welcome! If you have suggestions for improvements or new features, please submit an issue or a pull request.

## Acknowledgements

We would like to acknowledge the creators of the original repository for their excellent work. This project would not have been possible without their contributions.

## License

This project is licensed under the same terms as the original repository. For more details, see the [LICENSE](./LICENSE) file.

<!-- Resources -->
[#2]: https://github.com/nokacper24/static-templ/pull/2
[#4]: https://github.com/nokacper24/static-templ/pull/4
[#5]: https://github.com/nokacper24/static-templ/pull/5
[#10]: https://github.com/nokacper24/static-templ/pull/10
