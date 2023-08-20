
# k8s-build-cli 
![Go](https://img.shields.io/badge/-Go-blue) ![Under Development](https://img.shields.io/badge/-Under%20Development-orange)

`k8s-build-cli` is a robust command-line interface crafted with Go, designed specifically for Kubernetes enthusiasts. By integrating seamlessly with Kubernetes operations, it serves as a bridge to the [Open Build Service (OBS) API](https://openbuildservice.org/), allowing users to automate and enhance their build processes, package software, manage distributions, and ensure consistent builds across multiple platforms. The tool relies on the [Viper](https://github.com/spf13/viper) package for efficient configuration management and the [Cobra](https://github.com/spf13/cobra) package for a streamlined CLI experience.

## Tech Stack

- **Language**: Go (Golang)
- **Key Packages**:
  - [Viper](https://github.com/spf13/viper) - For handling configuration
  - [Cobra](https://github.com/spf13/cobra) - For crafting the CLI
- **API**: [Open Build Service (OBS) API](https://openbuildservice.org/)
  
## Setup Procedure

1. **Clone the Repository**:
   ```bash
   git clone https://github.com/[your-username]/k8s-build-cli.git
   cd k8s-build-cli
   ```

2. **Install Dependencies**:
   With Go Modules enabled:
   ```bash
   go mod download
   ```

3. **Build the CLI**:
   Compile the code and generate the binary:
   ```bash
   go build -o k8s-build-cli
   ```

4. **Run the CLI**:
   ```bash
   ./k8s-build-cli [commands/options]
   ```
   To see a list of available commands and options:
   ```bash
   ./k8s-build-cli --help
   ```

## Contributing

Contributions to `k8s-build-cli` are more than welcome! Whether it's bug reports, code enhancements, or new features, your input is valued and appreciated. Submit an issue or pull request to get started.

## License

This project is licensed under the [MIT License](https:github.com/kunxl.gg/k8s-build-cli/LICENSE)


The tags (badges) for "Go" and "Under Development" have been added at the top. These use the shields.io service, which is a popular way to add badges to READMEs. Adjust placeholders and other content as needed.
