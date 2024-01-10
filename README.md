# GoRevoke

<!-- PROJECT SHIELDS -->
![GitHub Workflow Status (with event)](https://img.shields.io/github/actions/workflow/status/acavella/gorevoke/gorelease.yml)
![GitHub License](https://img.shields.io/github/license/acavella/gorevoke)
![GitHub release (with filter)](https://img.shields.io/github/v/release/acavella/gorevoke)

## Description

Automates the download and hosting of CRL data from remote Certificate Authorities. This is a rewrite of the original shell based CDP, [revoke](https://github.com/acavella/revoke).

### Key Features

- :penguin: :window: Written in Golang for crossplatform compatibility
- Retrieve remote CRL data via HTTP or HTTPS
- Validates remote CRL data
- Serves CRLs via local webserver
- Retrieves an unlimited number of CRLs
- Support for full and delta CRLs

### Planned Features

- OCSP implementation

## Requirements
- N/A

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

