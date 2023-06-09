TOPIC=`echo $RANDOM | md5sum | head -c 20`

load '/opt/homebrew/lib/bats-support/load.bash'
load '/opt/homebrew/lib/bats-assert/load.bash'

@test "Create Topics - create and list topics" {
    ./kcli cluster add local -b localhost:9092
    ./kcli create $TOPIC -p 3 -r 1
    run ./kcli list
    assert_output --partial "$TOPIC 0"
    assert_output --partial "$TOPIC 1"
    assert_output --partial "$TOPIC 2"
}
