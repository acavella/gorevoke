# Go Revoke
A fully featured CDP and OCSP responder written in Go.
<!-- PROJECT LOGO -->
<p align="center">
  <a href="https://github.com/altCipher/revoke">
    <img src="images/logo.png" alt="Logo">
  </a>
</p>

<!-- PROJECT SHIELDS -->
<p align="center">
  <a href="https://github.com/altCipher/revoke/"><img src="https://img.shields.io/badge/build-passing-brightgreen.svg" alt="Build Status"></a>
  <img src="https://img.shields.io/github/contributors/altCipher/revoke.svg" alt="Contributors">
  <a href="LICENSE"><img src="https://img.shields.io/github/license/altCipher/revoke.svg" alt="License"></a>
  <a href="https://github.com/revokehq/revoke/releases"><img src="https://img.shields.io/github/release/altCipher/revoke.svg" alt="Latest Stable Version"></a>
  <a href="https://bestpractices.coreinfrastructure.org/projects/2731"><img src="https://bestpractices.coreinfrastructure.org/projects/2731/badge"></a>
</p>

## Overview

Automates the download and hosting of CRL data from a remote Certificate Authority.  Revoke is designed to be executed via chron.  

- Retrieve remote CRL data via HTTP or HTTPS
- Validates remote CRL data
- Serves CRLs via local HTTPD
- Written using BASH to maximize native compatibility and remain lighweight
- Retrieve an unlimited number of CRLs
- Support for full and delta CRLs

## Requirements
- Bash
- Apache HTTP Server 2.4 
- OpenSSL 1.0.2 or later
- Curl 7.29 or later

## Installation

Installation instructions here.

## Security Vulnerabilities

If you discover a security vulnerability within revoke, please send an e-mail to [tony@cavella.com](mailto:tony@cavella.com?Revoke%20Security%20Vulnerability). Security vulnerabilities are taken very seriously and will be addressed with the utmost priority.

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

