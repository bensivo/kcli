load '/opt/homebrew/lib/bats-support/load.bash'
load '/opt/homebrew/lib/bats-assert/load.bash'

@test "create and list topics" {
    ./kcli topic create -b localhost:9092 -t topic01 -p 3 -r 1
    run ./kcli topic list -b localhost:9092
    assert_output --partial 'topic01 0'
    assert_output --partial 'topic01 1'
    assert_output --partial 'topic01 2'
}

@test "basic produce/consume" {
    echo "Test1" | ./kcli produce -b localhost:9092 -t topic01 -p 0
    run ./kcli consume -b localhost:9092 -t topic01 -p 0 -o -1 -e
    assert_output --partial 'Test1'
}

@test "produce/consume new partition" {
    echo "Test2" | ./kcli produce -b localhost:9092 -t topic01 -p 1
    run ./kcli consume -b localhost:9092 -t topic01 -p 1 -o -1 -e
    assert_output --partial 'Test2'
}

@test "produce/consume negative 2 offset" {
    echo "Test3 (1/2)" | ./kcli produce -b localhost:9092 -t topic01 -p 1
    echo "Test3 (2/2)" | ./kcli produce -b localhost:9092 -t topic01 -p 1
    run ./kcli consume -b localhost:9092 -t topic01 -p 1 -o -2 -e
    assert_output --partial 'Test3 (1/2)'
    assert_output --partial 'Test3 (2/2)'
}