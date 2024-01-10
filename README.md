# GoRevoke

<!-- PROJECT SHIELDS -->
![GitHub Workflow Status (with event)](https://img.shields.io/github/actions/workflow/status/acavella/gorevoke/gorelease.yml)
![GitHub License](https://img.shields.io/github/license/acavella/gorevoke)
![GitHub release (with filter)](https://img.shields.io/github/v/release/acavella/gorevoke)

## Overview

GoRevoke is a standalone [Certificate Revocation List](https://en.wikipedia.org/wiki/Certificate_revocation_list) Distrution Point written in [Go](https://go.dev/), designed to be lightweight and fully self-contained. Using a simple configuration, GoRevoke automates downloading and serving of remote CRLs. GoRevoke is based on, [revoke](https://github.com/acavella/revoke), a shell based script providing similar function.

### Key Features

- Cross-platform compatiblity; tested on Linux and Windows
- Native and containerized deployment options
- Retrieve remote CRL data via HTTP or HTTPS
- Validation and confirmation of CRL data
- Built-in webserver alleviates the need for additional servers
- Ability to retrieve and serve an unlimited number of CRL sources
- Support for full and delta CRLs

### Planned Features

- OCSP responder

## Installation

GoRevoke is designed to be deployed and run as a container. Additional instructions are provided if you'd like to deploy from source on baremetal.

### Docker Instructions

1. 
2. Create a application configuration directory and file; this example maps a volume to the `/appdata` directory.
```Yaml
---
default:
  gateway: crls.pki.goog    # ip or fqdn to check used for connectivity checks
  interval: 900             # update interval to check for new crls, in seconds
  port: 4000                # port used by http server

ca:
  id: 
    - x21
    - x11
  uri: 
    - http://crls.pki.goog/gts1c3/zdATt0Ex_Fk.crl
    - http://crl.godaddy.com/gdig2s1-5609.crl
```
2. Use the following command to pull the latest image.
```Shell
docker run -d \
--name gorevoke \
-p 80:4000 \
-v /appdata/gorevoke/config:/usr/local/bin/gorevoke/conf \
--restart=unless-stopped \
s0lution/gorevoke:latest
```

### Baremetal Instructions

```Text
Installation instructions here.
```

## Container Performance
![Docker Container Performance](assets/docker-stats.png)

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

