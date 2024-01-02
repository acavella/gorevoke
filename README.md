# Go Revoke

<!-- PROJECT SHIELDS -->
<p align="center">
  <a href="https://github.com/acavella/gorevoke/"><img src="https://img.shields.io/badge/build-passing-brightgreen.svg" alt="Build Status"></a>
  <img src="https://img.shields.io/github/contributors/acavella/gorevoke.svg" alt="Contributors">
  <a href="LICENSE"><img src="https://img.shields.io/github/license/acavella/gorevoke.svg" alt="License"></a>
  <a href="https://github.com/revokehq/revoke/releases"><img src="https://img.shields.io/github/release/acavella/gorevoke.svg" alt="Latest Stable Version"></a>
  <a href="https://bestpractices.coreinfrastructure.org/projects/2731"><img src="https://bestpractices.coreinfrastructure.org/projects/2731/badge"></a>
</p>

## Overview

Automates the download and hosting of CRL data from remote Certificate Authorities. This is a rewrite of the original shell based CDP, [revoke](https://github.com/acavella/revoke).

- Written in Golang for crossplatform compatibility
- Retrieve remote CRL data via HTTP or HTTPS
- Validates remote CRL data
- Serves CRLs via local webserver
- Retrieves an unlimited number of CRLs
- Support for full and delta CRLs

## Planned Features

- OCSP implementation

## Requirements
- N/A

## Installation

```Installation instructions here.
```
## Security Vulnerabilities

I welcome welcome all responsible disclosures. Please do not open an ISSUE to report a security problem. Please use the private reporting system to report security related issues responsibly: https://github.com/acavella/gorevoke/security/advisories/new

## Contributing

Contributions are what make the open source community such an amazing place to be learn, inspire, and create. Any contributions you make are **greatly appreciated**.

1. Fork the Project
2. Create your Feature Branch (`git checkout -b feature/AmazingFeature`)
3. Commit your Changes (`git commit -m 'Add some AmazingFeature`)
4. Push to the Branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

## License

Distributed under the MIT License. See `LICENSE` for more information.

## Contact

Tony Cavella - tony@cavella.com

Project Link: [https://github.com/acavella/gorevoke](https://github.com/acavella/gorevoke)

<!-- ACKNOWLEDGEMENTS -->
## Acknowledgements

