# SEOnaut
[![Go Report Card](https://goreportcard.com/badge/github.com/stjudewashere/seonaut)](https://goreportcard.com/report/github.com/stjudewashere/seonaut) [![GitHub](https://img.shields.io/github/license/StJudeWasHere/seonaut)](LICENSE) [![GitHub Workflow Status](https://img.shields.io/github/actions/workflow/status/StJudeWasHere/seonaut/test.yml)](https://github.com/StJudeWasHere/seonaut/actions/workflows/test.yml)

SEOnaut is an open-source SEO auditing tool designed to analyze websites for issues that may impact search engine rankings. It performs a comprehensive site scan and generates a report detailing any identified issues, organized by severity and potential impact on SEO.

SEOnaut categorizes issues into three levels of severity: critical, high, and low. The tool can detect various SEO-related problems, such as broken links (to avoid 404 errors), redirect issues (temporary, permanent, or loops), missing or duplicate meta tags, incorrectly ordered headings, and more.

A hosted version of SEOnaut is available at [seonaut.org](https://seonaut.org).

![seonaut](https://github.com/user-attachments/assets/6184b418-bd54-4456-9266-fcfd4ce5726d)

## Technology

<<<<<<< HEAD
SEOnaut is a web based application built with the Go programming language and a ~~MySQL~~ MariaDB database for its data storage. On the frontend side, the user interface is designed with simplicity in mind, using custom CSS and minimal vanilla Javascript. To make the dashboard interactive, the application utilizes Apache ECharts.

While it is possible to configure a custom database and compile SEOnaut manually, using the provided Docker files is recommended. These files simplify the setup process and eliminate the need for manual configuration, allowing for quicker and easier deployment.

<<<<<<< HEAD
### Podman and Podman-Compose

Rather than use Docker.io to provide containerization, this fork uses Podman and Podman-Compose. So, you will need to have those installed.

There are two methods to build the image and run containerized environment:

1. With sensitivity and chronological longevity

```bash
podman build -t seonaut:latest .
# then
podman-compose up -d
```

2. With speed and ease of effort

2. **Clone the Repository**  
   Clone the SEOnaut repository:

Whichever of these at your preference.

Once the process is complete, you can access SEOnaut in your web browser by visiting ```http://localhost:9000```.

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
