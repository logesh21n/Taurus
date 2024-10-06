# Taurus

## Overview

`Taurus` is a command-line utility designed to automate the process of gathering subdomain information and performing reconnaissance on specified domains. It integrates multiple tools and techniques to identify subdomains, check their availability, and assess their security posture.

## Features

- **Subdomain Enumeration**: Utilizes various tools like `subfinder`, `interlace`, `sublist3r`, and `amass` to discover subdomains. Also performs `github` recon.

- **Data Collection**: Retrieves subdomain data from multiple sources, including web services and public datasets.
- **Alive Check**: Checks the live status of discovered subdomains using `httpx`.
- **Vulnerability Scanning**: Optionally performs Nmap scans on valid subdomains to identify potential vulnerabilities. This process can take time.
- **Output Management**: Consolidates results into organized output files for easy analysis.

# Installation Instructions for Taurus

## Prerequisites

Before using `Taurus`, ensure you have the following dependencies installed on your system:

1. **Go**: Version 1.15 or later. You can install it from [the official Go website](https://golang.org/dl/).
2. **subfinder**: Install by following the instructions [here](https://github.com/projectdiscovery/subfinder#installation).
3. **interlace**: Install via [its GitHub repository](https://github.com/cgboal/Interlace#installation).
4. **curl**: Typically pre-installed on most systems. You can check if it's installed by running `curl --version`.
5. **httpx**: Install by following the instructions [here](https://github.com/projectdiscovery/httpx#installation).
6. **gau**: Install using the instructions [here](https://github.com/lc/gau#installation).
7. **waybackurls**: Install via [its GitHub repository](https://github.com/tomnomnom/waybackurls#installation).
8. **anew**: Install from [its GitHub repository](https://github.com/tj/anew#installation).

## Installation Steps

### Step 1: Clone the repository:

```bash
git clone https://github.com/yourusername/Taurus.git
cd Taurus
```

### Step 2: Rebuild the Tool

You need to recompile the tool:

```bash
go build -o Taurus taurus.go
```

### Step 3: Replace the Old Binary

Depending on where you placed the binary, you will need to replace it:

  **If You Placed It in `/usr/local/bin`**:
   Move the newly compiled binary to `/usr/local/bin`:

```bash
sudo mv taurus /usr/local/bin/
```

### Step 4: Verify the Update

 **Run the Tool**:
   Test the tool to ensure that your changes work as expected:

```bash
taurus -h
```


Install these tools using their respective installation methods, typically through package managers or direct downloads from their repositories.

## Usage

1. **Prepare a Domains File**: Create a text file named `domains.txt` containing the list of domains you want to analyze, with one domain per line.

2. **Run Taurus**: Execute `Taurus` using the following command:

```bash
go run Taurus.go -d domains.txt // Name domains.txt which contains your targeted domains.
```

3. **Follow the Prompts**: After running the commands, you will be prompted to continue with an optional Nmap scan on the discovered subdomains. Respond with `yes` or `no`.

## Output

- **Subdomains**: The tool generates multiple output files, including:
  - `subdomains1.txt`, `subdomains2.txt`, ..., `subdomains9.txt` - various sources of discovered subdomains.
  - `sortedSubdomains.txt` - a unique, sorted list of all discovered subdomains.
  - `correctedSubdomains.txt` - validated and filtered subdomains ready for further analysis.
  - `aliveSubdomain.txt` - live subdomains after a health check.
  - `nmap` - results of the optional Nmap scan (if executed).

## Example Commands

- To run the tool with a custom file:
  
```bash
go run Taurus.go -d mydomains.txt
```

- To skip the Nmap scan:
  
```bash
Do you want to continue with the Nmap scan? (yes or no):no
```

## Help

If any problem in running the tool, just ping me or use issues tab.

## Contributing

Contributions are welcome! If you have suggestions or improvements, please fork the repository and submit a pull request.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
