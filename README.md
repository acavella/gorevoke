# GoRevoke

<!-- PROJECT SHIELDS -->
![GitHub Actions Workflow Status](https://img.shields.io/github/actions/workflow/status/acavella/gorevoke/gorelease.yml?logo=go)
![GitHub Actions Workflow Status](https://img.shields.io/github/actions/workflow/status/acavella/gorevoke/dockerbuild.yml?logo=docker)
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

1. Copy and rename the configuration example `gorevoke.yml` to a path on your host system, e.g. `$HOME/gorevoke.yml`
2. Pull the latest image from Docker Hub using the following example Docker run command:
```Shell
docker run -d \
--name gorevoke \
-p 80:4000 \
-v /path/to/gorevoke.yml:/etc/gorevoke.yml \
--restart=unless-stopped \
ghcr.io/acavella/gorevoke:latest
```

> [!IMPORTANT]
> The Docker Run command above exposes the built-in webserver to the host directly on port 80 and is not recommended for production deploys. For a production configuration we recommend placing a webserver or proxy (such as Apache httpd or nginx) in front of GoRevoke to handle public web requests.

### Native Deployment

1. Download the [latest release](https://github.com/acavella/gorevoke/releases/latest/) archive for the appropriate platform 
   - Linux (amd64): gorevoke-<version>-linux-amd64.tar.gz
   - Windows (amd64): gorevoke-<version>-windows-amd64.zip
2. Extract the archive to the appropriate application directory
   - Linux: `/usr/local/bin`
   - Windows: `C:\Program Files\`
3. (optional) Edit the provided example configuration file `gorevoke.yml` and save it as `/etc/gorevoke.yml`
4. (optional) Create a system user for GoRevoke: `useradd --system --no-create-home --shell=/sbin/nologin gorevoke`
5. Create a systemd service file `/etc/systemd/service/gorevoke.service`. Example unit files:
```ini
### Using a static-file configuration
[Unit]
Description=GoRevoke CDP Server
After=network-online.target

[Service]
Type=simple
ExecStart=/usr/local/bin/gorevoke
User=gorevoke
Restart=always

[Install]
WantedBy=multi-user.target default.target
```
```ini
### Using configuration via environment vars
[Unit]
Description=GoRevoke CDP Server
After=network-online.target

[Service]
Type=simple
ExecStart=/usr/local/bin/gorevoke
User=gorevoke
Environment=GOREVOKE_DEFAULT_INTERVAL=900
Environment=GOREVOKE_DEFAULT_WEBSERVER=true
Environment=GOREVOKE_DEFAULT_CRLDIR=/var/www/public_html
Environment=GOREVOKE_CRLS='{"x21":"http://crls.pki.goog/gts1c3/zdATt0Ex_Fk.crl", "x11":"http://crl.godaddy.com/gdig2s1-5609.crl"}'
Restart=always

[Install]
WantedBy=multi-user.target default.target
```
6. Set the permissions `sudo chmod 664 /etc/systemd/service/gorevoke.service`
7. Reload the systemd configuration `sudo systemctl daemon-reload`
8. Enable and start the service:
```shell
sudo systemctl enable --now gorevoke.service
```

## Configuration
A list of all available configuration options is available at [gorevoke.yml](gorevoke.yml), with comments provided inline. Configuration can be set via a static file, in which case the following paths are checked:

- `$PWD/gorevoke.yml`
- `$HOME/.gorevoke/gorevoke.yml`
- `/etc/gorevoke.yml`

Optionally, all configuration values can be specified via environment variables, upper-cased and prefixed with `GOREVOKE`. For example, the configuration item `default.interval` can be set via the `GOREVOKE_DEFAULT_INTERVAL` variable. If specifying the list of CRLs as an environment var (`GOREVOKE_CRLS`), the CRLs must be provided as a json dict. See the systemd unit example, above.

## Container Performance
![Docker Container Performance](assets/docker-stats.png)

## Security Vulnerabilities

I welcome welcome all responsible disclosures. Please do not open an ISSUE to report a security problem. Please use the private reporting system to report security related issues responsibly: https://github.com/acavella/gorevoke/security/advisories/new

## Contributing

Contributions are essential to the success of open-source projects. In other words, we need your help to keep GoRevoke great!

What is a contribution? All the following are highly valuable:

1. **Let us know of the best-practices you believe should be standardized**   
   GoRevoke is designed to be compliant with applicable RFCs out-of-the box. By sharing your experiences and knowledge you help us build a solution that takes into account best-practices and user experience.

2. **Let us know if things aren't working right**   
   We aim to provide a perfect application and test it extensively, however, we can't imagine or replicate every deployment scenario possible. If you run into an issue that you think isn't normal, please let us know.

3. **Add or improve features**   
   Have an idea to add or improve functionality, then let us know! We want to make GoRevoke the best total solution it can be.

**General information about contributions:**

Check our [Security Policy](https://github.com/acavella/gorevoke#).   
Found a bug? Open a [GitHub issue](https://github.com/acavella/gorevoke/issues).   
Read our [Contributing Code of Conduct](https://github.com/acavella/gorevoke?tab=coc-ov-file#), which contains all the information you need to contribute to GoRevoke!

## License

Distributed under the MIT License. See `LICENSE` for more information.

## Contact

- Tony Cavella - tony@cavella.com
- Project Link: [https://github.com/acavella/gorevoke](https://github.com/acavella/gorevoke)

## Acknowledgements
- [@Deliveranc3](https://github.com/Deliveranc3) - Containerfile development and additions to config logic
