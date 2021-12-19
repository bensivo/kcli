# kcli
A kafka client written in go.

## Installation
To install kcli, checkout this repo and build it from scratch:
```
git clone https://gitlab.com/bensivo/kcli.git
cd kcli
go install ./cmd/kcli

# Ensure that ~/go/bin is in your PATH, or copy ~/go/bin/kcli to /usr/local/bin/kcli
```

## Usage
### Topics
List topics with:
```
kcli topic list -b <bootstrap server>
```

Create a topic with:
```
kcli topic create -b <bootstrap server> -t <topic name> -p <partitions> -r <replication factor>
```

### Consume messages
```
kcli consume -b <bootstrap server> -t <topic> -p <partition> -o <offset>
```
By default, uses partition 0 and the earliest available offset.

Specifying offsets
```
kcli consume -b <bootstrap server> -t <topic> -o 20
// Starting at offset 20, get all messages on partition 0

kcli consume -b <bootstrap server> -t <topic> -o -2
// Get the last 2 messages on partition 0
```

### Produce messages
```
kcli produce -b <bootstrap server> -t <topic> -p <partition>
// Type your message and press 'enter' to publish
```

You can also pipe messages in from stdin, separated by newlines
```
echo "My Message" | kcli produce -b <bootstrap server> -t <topic>
```

## Development
### Run tests
Install bats and helper libraries with brew
```
brew tap kaos/shell
brew install bats-core
brew install bats-assert
brew install bats-file
```

Start local kafka broker:
```
docker compose up
```

Wait for the broker to start, then run tests with:
```
make build
make test
```
