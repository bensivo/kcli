# kcli
A kafka client written in go.

------------
THIS IS A WORK IN PROGRESS - Expect bugs, and changing interfaces for the near future

------------

## Installation
To install kcli, checkout this repo and build it from scratch:
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
- -m, --sasl-mechanism 
- -p, --sasl-password
- -u, --sasl-username

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
- Interactive: ```kcli produce <topic> -p <partition>```
    - Listens for input from stdin. Each newline sends a new message
- Non-Interactive: ```echo "my message here" | kcli produce <topic> -p <partition>```

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
