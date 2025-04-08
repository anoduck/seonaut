# SEOnaut

[![Go Report Card](https://goreportcard.com/badge/github.com/stjudewashere/seonaut)](https://goreportcard.com/report/github.com/stjudewashere/seonaut) [![GitHub](https://img.shields.io/github/license/StJudeWasHere/seonaut)](LICENSE) [![GitHub Workflow Status](https://img.shields.io/github/actions/workflow/status/StJudeWasHere/seonaut/test.yml)](https://github.com/StJudeWasHere/seonaut/actions/workflows/test.yml)

> [!NOTE]
> 
> This is a fork of the [SEOnaut](https://github.com/StJudeWasHere/seonaut) project with a few modifications. It substitutes MYSQL for the more open source friendly MariaDB, since syntactically they are fully compatible. It also substitutes Docker.io with the more open source friendly Podman.

SEOnaut is an open-source SEO auditing tool designed to analyze websites for issues that may impact search engine rankings. It performs a comprehensive site scan and generates a report detailing any identified issues, organized by severity and potential impact on SEO.

SEOnaut categorizes issues into three levels of severity: critical, high, and low. The tool can detect various SEO-related problems, such as broken links (to avoid 404 errors), redirect issues (temporary, permanent, or loops), missing or duplicate meta tags, incorrectly ordered headings, and more.

A hosted version of SEOnaut is available at [seonaut.org](https://seonaut.org).

![seonaut](https://github.com/user-attachments/assets/6184b418-bd54-4456-9266-fcfd4ce5726d)

## Technology

SEOnaut is a web-based application built with the Go programming language and a MariaDB database for data storage. The frontend is designed for simplicity, using custom CSS and minimal vanilla JavaScript. Apache ECharts is used to provide an interactive dashboard experience.

While it is possible to configure a custom database and compile SEOnaut manually, using the provided Docker files is recommended. These files simplify the setup process and eliminate the need for manual configuration, allowing for quicker and easier deployment.

### Quick Start Guide

To get started with SEOnaut, follow these steps to run it using Docker:

1. **Install Podman**  
   Ensure Podman is installed on your system. This can be achieved with most Linux distributions by using the respective package management system. If you are running debian or ubuntu you would use something like this to achieve podman installation `sudo apt install podman`, if you are using alpine linux it would look something like this `sudo apk add podman`. 

2. **Clone the Repository**  
   Clone the SEOnaut repository:
   
   `git clone https://github.com/anoduck/seonaut.git`

3. **Navigate to the Project Directory**  
   Change into the project directory:
   
   `cd seonaut`

4. **Build and Run Docker Containers**  
   Run the following command to build and start the Docker containers:
   
   `docker-compose -f docker-compose.yml -f docker-compose.dev.yml up -d --build`

5. **Access the Application**  
   Once the containers are running, open your browser and visit:
   
   `http://localhost:9000`
   
   SEOnaut is set up to run on port 9000 using unencrypted HTTP by default. For added security, it is recommended to configure HTTPS using a reverse proxy. This will ensure encrypted communication between the client and the server.

For more detailed installation and configuration instructions, refer to the [INSTALL.md](docs/INSTALL.md) file.

#### Caveat

If you run seonaut on a remote machine, you will need to use ssh to generate a tunnel and forward seonaut to your host system. Below is an example of a command that can accomplish this for you. 

```bash
ssh -L 9000:localhost:9000 -N -T "$USERNAME"@"$REMOTE_HOST"
```

The reasons for this are due to podman's network stack, and occur regardless of running podman in root or rootless configuration.

## Contributing

Please see [CONTRIBUTING](docs/CONTRIBUTING.md) for details.

## License

SEOnaut is open-source under the MIT license. See [License File](LICENSE) for more information.
