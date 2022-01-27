TOPIC=`echo $RANDOM | md5sum | head -c 20`

load '/opt/homebrew/lib/bats-support/load.bash'
load '/opt/homebrew/lib/bats-assert/load.bash'

function setup() {
    ./kcli cluster add local -b localhost:9092
    ./kcli create $TOPIC -p 2 -r 1

    sleep 3
}

@test "Produce - should require a topic" {
    run ./kcli produce
    assert_output --partial 'Error: accepts 1 arg(s), received 0'
}

@test "Produce - should accept message from pipe" {
    MSG=`echo $RANDOM | md5sum | head -c 20`
    echo "$MSG" | ./kcli produce $TOPIC
    run ./kcli consume $TOPIC -o -1 -e
    assert_output --partial $MSG
}

@test "Produce - should accept messages interactively" {
    MSG=`echo $RANDOM | md5sum | head -c 20`
    expect -c "
    spawn ./kcli produce $TOPIC
    expect \"connected\"
    sleep 1
    send \"$MSG\n\"
    sleep 1
    send \x03
    "
    run ./kcli consume $TOPIC -o -1 -e
    assert_output --partial $MSG
}