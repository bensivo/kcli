# kcli
A kafka client written in go.

## Usage
### Consume messages
```
kcli consume -t topic01
```

Specifying offsets
```
kcli -b localhost:9092 -c -t topic01 -o 20
// Starting at offset 20, get all messages on partition 0

kcli -b localhost:9092 -c -t topic01 -o -2
// Get the last 2 messages on partition 0
```

### Produce messages
```
kcli produce -t topic01
// Type your message and press 'enter' to publish
```

You can also pipe messages in from stdin, separated by newlines
```
echo "My Message" | kcli produce -t topic01
```
