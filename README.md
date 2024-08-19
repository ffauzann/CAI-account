# User
This service contains user & account usecases such as authentication, account management, and balance movement.

## Guide
### Prerequisite
- Programming language : `go@1.21` or later
- RDMS : `postgresql@14.11` or later
- Caching : `redis-server@7.2.1` or later
- Secret management (optional) : `vault@1.16.2` or later
- Containerization (optional) : `docker@26.1.1` or later
- Orchestrator (optional) : `minikube@1.31.2` or later

### Service Configuration
#### With local config files :
1. Duplicate file `internal/app/user.config.example.yaml` as `internal/app/user.config.yaml`
2. Duplicate file `internal/app/auth.config.example.yaml` as `internal/app/auth.config.yaml`
3. Update both `user.config.yaml` and `auth.config.yaml` file with your own credentials

#### With remote secret management :
1. Setup these environment variables :
    - `VAULT_ADDR` : `https://your-vault-server-host`
    - `VAULT_TOKEN` : `hvs.YOURVAULTTOKEN`
    - `VAULT_MOUNT_PATH` : `path/to/secret/engine/data/`
2. Create or use an existing **KV v2 engine** and make sure its the same engine as `VAULT_MOUNT_PATH` value
3. Create new secret in selected engine named `auth` and `user`
4. Convert both yaml config files into json format and then fill it with your own credentials

### Installation

#### Using `go run` :
1. Complete [Service Configuration](#service-configuration) above
2. Create new database schema under the same name as `database.sql.schema` value in `config.yaml`
3. Run `go mod vendor`
4. Run `go run main.go`

#### Using Docker :
Assuming you already have your own docker enviroment setup, Follow these steps:
1. Complete [Service Configuration](#service-configuration) above. **NOTE: For local development, this service is not yet support using docker alongside remote secret manager.**
2. Update `docker-compose.yaml` as well if needed
3. Run `docker-compose up`

#### Using Kubernetes :
Assuming you already have your own k8s enviroment setup including `kubectl` and `minikube`, Follow these steps:
1. Complete [Service Configuration](#service-configuration) above. **NOTE: For local development, this service is not yet support using k8s alongside remote secret manager.**
2. Make sure you already logged in to your docker hub or any other registry using `docker login`
3. Run `make docker-build USERNAME="<YOUR_USERNAME>" VERSION="<IMAGE_VERSION>"` to build your image and push it to the registry
4. Update `spec.template.spec.containers[0].image` value to the same value you used on 4th step in `k8s/auth-deployment.yaml` file
5. [Optional] If you change port mapping, make sure you update `k8s/auth-service.yaml` file as well
6. Run `make k8s-deploy`
7. If you use minikube, run `minikube service user --url` to see assigned ports in host or use `kubectl port-forward deployments/user <LOCAL_PORT>:<SERVICE_PORT>` if you want to specify the localhost port.

To clear all data, you only need to run `make k8s-destroy`

## License
[MIT](https://github.com/ffauzann/CAI-account/blob/main/LICENSE) 