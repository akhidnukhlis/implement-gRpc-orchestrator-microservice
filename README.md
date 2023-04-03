# implement-gRpc-microservice-orchestrator

<!-- ABOUT THE PROJECT -->
## About The Project

The service orchestrator is used as the first layer in the microservice deployment where this layer will be the central setting where a service will be channeled. At this layer, authentication will also be set, file upload settings and connections between other services. All coding logic will be applied to this layer.


### Built With

This section should list any major frameworks that you built your project using. Leave any add-ons/plugins for the acknowledgements section. Here are a few examples.
* [Postgres](https://www.postgresql.org/)
* [Golang](https://golang.com)

<!-- GETTING STARTED -->
## Getting Started

This is how you may give instructions on setting up your project locally. To get a local copy up and running follow these simple example steps.

### Prerequisites

This is how to list things you need to use the software and how to install them.

* terminal
  ```sh
  brew install golang
  ```

### Dependency repositories
- Service Author  : https://github.com/akhidnukhlis/implement-gRpc-microservice-author-service
- Protobank     : https://github.com/akhidnukhlis/implement-gRpc-microservice


### Installation

1. Clone example repo author-service
   ```sh
   git clone https://github.com/akhidnukhlis/implement-gRpc-microservice-author-service
   ```
   
2. Edit `env.example` to `.env` and setup your environment

3. Run command golang
   ```sh
   go mod tidy
   ```
   
   ```sh
   go mod download
   ```
   
4. Run service
   ```sh
   go run main.go
   ```
   
5. Clone example repo orchestrator-service
   ```sh
   git clone https://github.com/akhidnukhlis/implement-orchestrator-gRpc-microservice
   ```
   
6. Edit `env.example` to `.env` and setup your environment

7. Run command golang
   ```sh
   go mod tidy
   ```

   ```sh
   go mod download
   ```

8. Run service
   ```sh
   go run main.go
   ```

<!-- ROADMAP -->
## Roadmap

See the [open issues](https://github.com/akhidnukhlis) for a list of proposed features (and known issues).


<!-- CONTRIBUTING -->
## Contributing

Contributions are what make the open source community such an amazing place to be learn, inspire, and create. Any contributions you make are **greatly appreciated**.

1. Fork the Project
2. Create your Feature Branch (`git checkout -b feature/new-feature`)
3. Commit your Changes (`git commit -m 'Add some new-feature'`)
4. Push to the Branch (`git push origin feature/new-feature`)
5. Open a Pull Request


<!-- CONTACT -->
## Contact

Name - @akhidnukhlis - nukhlis@gmail.com

Project Link: https://github.com/akhidnukhlis

***