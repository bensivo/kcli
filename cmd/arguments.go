package cmd

// Variables for all the options that can be specified in all subcommands
var partition int = 0
var partitions int = 1
var replicationFactor int = 0
var offset int = 0
var exit bool = false
var bootstrapServer string = ""
var connectionTimeout int = 10
var saslMechanism string = ""
var saslUsername string = ""
var saslPassword string = ""
var clusterName string = ""
var sslEnabled bool = false
var sslSkipVerification bool = false
var sslCaCertificatePath string = ""
