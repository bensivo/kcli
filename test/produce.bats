TOPIC=`echo $RANDOM | md5sum | head -c 20`

load '/opt/homebrew/lib/bats-support/load.bash'
load '/opt/homebrew/lib/bats-assert/load.bash'

function setup() {
    ./kcli topic create -b localhost:9092 -t $TOPIC -p 2 -r 1

    sleep 3
}

@test "Produce to topic from pipe" {
    MSG=`echo $RANDOM | md5sum | head -c 20`
    echo "$MSG" | ./kcli produce -b localhost:9092 -t $TOPIC
    run ./kcli consume -b localhost:9092 -t $TOPIC -o -1 -e
    assert_output --partial $MSG
}

@test "Produce to topic interactively" {
    MSG=`echo $RANDOM | md5sum | head -c 20`
    expect -c "
    spawn ./kcli produce -b localhost:9092 -t $TOPIC
    expect \"connected\"
    sleep 1
    send \"$MSG\n\"
    sleep 1
    send \x03
    "
    run ./kcli consume -b localhost:9092 -t $TOPIC -o -1 -e
    assert_output --partial $MSG
}