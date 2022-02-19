# kcli
A kafka client written in go.

------------
THIS IS A WORK IN PROGRESS - Expect bugs, and changing interfaces for the near future

------------

## Installation
To install kcli, just use go install:
```
go install gitlab.com/bensivo/kcli@latest
```


Or checkout this repo and build it from scratch:
```
git clone https://gitlab.com/bensivo/kcli.git
cd kcli
go install 

# Ensure that ~/go/bin is in your PATH
```

## Setup

### Configure your cluster
Before you can do any cluster operations, add a cluster with: ``` kcli cluster add ```

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
kcli  cluster add my-local-cluster -b localhost:9092 -t 10 -m plain -u my-user -p my-password
```

You can configure multiple kafka clusters at once. kcli offers a few commands to manage them: 
- View your configured clusters: `kcli cluster list`
- Switch your active cluster: `kcli cluster use <name>`
    - Or add the --cluster-name, -c flag to manually specify a cluster: `kcli consume my-topic --cluster-name=localhost`
- Delete a cluster configuration: `kcli cluster remove <name>`

## Usage
Topic Operations:
- List all topics and partitions: `kcli list`
- Create a new topic: `kcli create <topic> -p <partitions> -r <replication factor>`

Consume messages:
- Read all messages: ```kcli consume <topic>```
- Read from a specific partition: ```kcli consume <topic> -p 1```
    - By default, kcli uses partition 0
- Read starting at the 10th message: ```kcli consume <topic> -o 10```
- Read the last 10 messages: ```kcli consume <topic> -o -10```

Produce messages:
- Stdin: ```kcli produce <topic>``` or ```cat ./data.json | kcli produce <topic>```
    - Read stdin until EOF and send as a single message. If not using pipes, use CTRL + D to send an EOF on stdin
- File: ```kcli product ./data.json```
    - Sends the entire files as a single message

Bash / ZSH autocompletion:
- To configure shell autocompletion, add the following to your zshrc or bashrc
    ```
    source <(kcli completion)
    ```

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
