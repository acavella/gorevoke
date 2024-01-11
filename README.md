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

## Installation Instructions

GoRevoke can be deployed as either a containerized image or installed natively on the host. The following instructions outline basic installation and configuration options.

### Docker Deployment

1. On the host machine create the following directories: `${PWD}/appdata/gorevoke/conf` and `${PWD}/appdata/gorevoke/crl`
2. Copy and rename the configuration example `conf/config.yml.example` to `${PWD}/appdata/gorevoke/conf/config.yml`
3. Pull the latest image from Docker Hub using the following example Docker run command:
```Shell
docker run -d \
--name gorevoke \
-p 80:4000 \
-v ${PWD}/appdata/gorevoke/crl:/usr/local/bin/gorevoke/crl/static \
-v ${PWD}/appdata/gorevoke/config:/usr/local/bin/gorevoke/conf \
--restart=unless-stopped \
s0lution/gorevoke:latest
```

> [!IMPORTANT]
> The Docker Run command above exposes the built-in webserver to the host directly on port 80 and is not recommended for production deploys. For a production configuration we recommend placing a webserver or proxy (such as Apache httpd or nginx) in front of GoRevoke to handle public web requests.

### Native Deployment

1. Download the [latest release](https://github.com/acavella/gorevoke/releases/latest/) archive for the appropriate platform 
   - Linux (amd64): gorevoke-<version>-linux-amd64.tar.gz
   - Windows (amd64): gorevoke-<version>-windows-amd64.zip
2. Extract the archive to the appropriate application directory
   - Linux: /usr/local/bin
   - Windows: C:\Program Files\
3. Edit the provided example configuration file `conf/config.yml.example` and save it as `conf/config.yml`
4. Create a startup file to handle starting the application.

## Container Performance
![Docker Container Performance](assets/docker-stats.png)

## Security Vulnerabilities

I welcome welcome all responsible disclosures. Please do not open an ISSUE to report a security problem. Please use the private reporting system to report security related issues responsibly: https://github.com/acavella/gorevoke/security/advisories/new

## Contributing

Contributions are essential to the success of open-source projects. In other words, we need your help to keep GoRevoke great!

What is a contribution? All the following are highly valuable:

1. **Let us know of the best-practices you believe should be standardized**
   Netdata should out-of-the-box detect as many infrastructure issues as possible. By sharing your knowledge and experiences, you help us build a monitoring solution that has baked into it all the best-practices about infrastructure monitoring.

2. **Let us know if Netdata is not perfect for your use case**
   We aim to support as many use cases as possible and your feedback can be invaluable. Open a GitHub issue, or start a GitHub discussion about it, to discuss how you want to use Netdata and what you need.

   Although we can't implement everything imaginable, we try to prioritize development on use-cases that are common to our community, are in the same direction we want Netdata to evolve and are aligned with our roadmap.

3. **Support other community members**
   Join our community on GitHub, Discord and Reddit. Generally, Netdata is relatively easy to set up and configure, but still people may need a little push in the right direction to use it effectively. Supporting other members is a great contribution by itself!

4. **Add or improve integrations you need**
   Integrations tend to be easier and simpler to develop. If you would like to contribute your code to Netdata, we suggest that you start with the integrations you need, which Netdata does not currently support.

General information about contributions:

Check our Security Policy.
Found a bug? Open a GitHub issue.
Read our Contributing Guide, which contains all the information you need to contribute to Netdata, such as improving our documentation, engaging in the community, and developing new features. We've made it as frictionless as possible, but if you need help, just ping us on our community forums!

## License

Distributed under the MIT License. See `LICENSE` for more information.

## Contact

Tony Cavella - tony@cavella.com

Project Link: [https://github.com/acavella/gorevoke](https://github.com/acavella/gorevoke)

<!-- ACKNOWLEDGEMENTS -->
## Acknowledgements

