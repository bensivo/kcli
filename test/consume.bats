TOPIC=`echo $RANDOM | md5sum | head -c 20`

load '/opt/homebrew/lib/bats-support/load.bash'
load '/opt/homebrew/lib/bats-assert/load.bash'

function setup() {
    ./kcli cluster add local -b localhost:9092
    ./kcli create $TOPIC -p 2 -r 1

    sleep 3

    echo "p0m1" | ./kcli produce $TOPIC -p 0
    echo "p0m2" | ./kcli produce $TOPIC -p 0
    echo "p0m3" | ./kcli produce $TOPIC -p 0

    echo "p1m1" | ./kcli produce $TOPIC -p 1
    echo "p1m2" | ./kcli produce $TOPIC -p 1
    echo "p1m3" | ./kcli produce $TOPIC -p 1

}

@test "Consume - should require topic" {
    run ./kcli consume

    assert_output --partial "Error: accepts 1 arg(s), received 0"
}

@test "Consume - Read all messages" {
    run ./kcli consume $TOPIC -e
    assert_output --partial 'p0m1'
    assert_output --partial 'p0m2'
    assert_output --partial 'p0m3'
}

@test "Consume - Read last message" {
    run ./kcli consume $TOPIC -o -1 -e
    assert_output --partial 'p0m3'
    refute_output --partial 'p0m2'
}

@test "Consume - Read from separate partition" {
    run ./kcli consume $TOPIC -p 1 -e
    assert_output --partial 'p1m1'
    assert_output --partial 'p1m2'
    assert_output --partial 'p1m3'
    refute_output --partial 'p0m1'
}
