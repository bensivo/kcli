TOPIC=`echo $RANDOM | md5sum | head -c 20`

load '/opt/homebrew/lib/bats-support/load.bash'
load '/opt/homebrew/lib/bats-assert/load.bash'

function setup() {
    ./kcli topic create -b localhost:9092 -t $TOPIC -p 2 -r 1

    sleep 3

    echo "p0m1" | ./kcli produce -b localhost:9092 -t $TOPIC -p 0
    echo "p0m2" | ./kcli produce -b localhost:9092 -t $TOPIC -p 0
    echo "p0m3" | ./kcli produce -b localhost:9092 -t $TOPIC -p 0

    echo "p1m1" | ./kcli produce -b localhost:9092 -t $TOPIC -p 1
    echo "p1m2" | ./kcli produce -b localhost:9092 -t $TOPIC -p 1
    echo "p1m3" | ./kcli produce -b localhost:9092 -t $TOPIC -p 1

}

@test "Read all messages" {
    run ./kcli consume -b localhost:9092 -t $TOPIC -e
    assert_output --partial 'p0m1'
    assert_output --partial 'p0m2'
    assert_output --partial 'p0m3'
}

@test "Read last message" {
    run ./kcli consume -b localhost:9092 -t $TOPIC -o -1 -e
    assert_output --partial 'p0m3'
    refute_output --partial 'p0m2'
}

@test "Read from separate partition" {
    run ./kcli consume -b localhost:9092 -t $TOPIC -p 1 -e
    assert_output --partial 'p1m1'
    assert_output --partial 'p1m2'
    assert_output --partial 'p1m3'
    refute_output --partial 'p0m1'
}
