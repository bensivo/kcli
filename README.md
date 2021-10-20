# kcli
A kafka client written in go.

## Usage
### Consume messages
```
kcli consume -t topic01
```

Specifying offsets and partitions
```
kcli -b localhost:9092 -c -t topic01 -o 20 -p 0

You can use negative offsets for 'n from end'

kcli -b localhost:9092 -c -t topic01 -o -2 -p 0
```