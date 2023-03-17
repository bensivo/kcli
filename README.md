# kcli
A kafka client written in go.

## Quick Start - docker
The fastest way to use kcli is through docker

1. Create a kcli configuration directory
    ```
    mkdir ~/.kcli
    ```

2. Run the docker container, mounting your directory
    ```
    docker run -v ~/.kcli:/root/.kcli bensivo/kcli:latest <commands>
    ```

3. (Optional) Alias the docker command

    To make life easier, add this to your ~/.bashrc or ~/.zshrc
    ```
    alias kcli='docker run -v ~/.kcli:/root/.kcli bensivo/kcli:latest'
    ```

    Now you can just run `kcli` and it will execute in docker.

See below for all the commands you can use


## Quick Start - golang
If you have golang installed, you can install the native binary like so:
```
go install gitlab.com/bensivo/kcli@latest
```

kcli will be installed in your go isntallation's 'bin' directory. Make sure that `$HOME/go/bin` is in your PATH.
```
export PATH=$PATH:$HOME/go/bin
```

Then you can run kcli with:
```
kcli <commands>
```


## Commands

### Cluster Setup 
Before you can do anything, you must add a cluster with: ``` kcli cluster add ```

Supported flags include:
- -b, --bootstrap-server
- -m, --sasl-mechanism  (supports 'plain', 'scram-sha-256', 'scram-sha-512')
- -p, --sasl-password
- -u, --sasl-username
- --ssl Enable TLS / SSL
- --ssl-ca Filepath to ca-certficate PEM file
- --ssl-skip-verification Turn off certificate-chain and hostname verification

Example: 
```
kcli cluster add localhost -b localhost:9092 -m plain -u my-user -p my-password
```

You can configure multiple kafka clusters at once. kcli offers a few commands to manage them: 
- View your configured clusters: `kcli cluster list`
- See a cluster configuration: `kcli cluster get <name>`
- Switch your active cluster: `kcli cluster use <name>`
    - Or add the --cluster-name, -c flag to manually specify a cluster: `kcli consume my-topic --cluster-name=localhost`
- Delete a cluster configuration: `kcli cluster remove <name>`
- Update a cluster configuration: `kcli cluster edit <name> ...flags`
    - Use any of the flags from the 'cluster add' commadn

### Topic Operations:
- List all topics and partitions: `kcli topic list`
- Create a new topic: `kcli topic create <topic> -p <partitions> -r <replication factor>`

### Consume messages:
- Read all messages: ```kcli consume <topic>```
- Read from a specific partition: ```kcli consume <topic> -p 1```
    - By default, kcli uses partition 0
- Read starting at the 10th message: ```kcli consume <topic> -o 10```
- Read the last 10 messages: ```kcli consume <topic> -o -10```

### Produce messages:
- Stdin: ```kcli produce <topic>``` or ```cat ./data.json | kcli produce <topic>```
    - Read stdin until EOF and send as a single message. If not using pipes, use CTRL + D to send an EOF on stdin
- File: ```kcli product ./data.json```
    - Sends the entire files as a single message


## Development
### Run tests
We use bats-core for testing, for ease of use. This does mean unit tests may be dependent on your available terminal environment.

Install bats and helper libraries with brew
```
brew tap kaos/shell
brew install bats-core
brew install bats-assert
brew install bats-file
```

Start local kafka broker:
```
docker compose up -d
```

Wait for the broker to start, then run tests with:
```
make build
make test
```

### Build and Publish Docker image
We use docker buildx to publish a multi-platform image for x86 and arm processors.

This command was tested using ubuntu, with the latest version of docker installed.

```
docker buildx build -t bensivo/kcli:latest --platform=linux/arm64,linux/amd64 --push .
docker buildx build -t bensivo/kcli:<version> --platform=linux/arm64,linux/amd64 --push .
```
