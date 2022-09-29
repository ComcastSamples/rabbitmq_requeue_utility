package main

import (
	"fmt"
)

func printUsage() {

	fmt.Printf("%s: \n\n", blue("Usage"))

	fmt.Printf("  %s: \n", magenta("Options"))
	fmt.Printf("\t%s\n\n", blue("-debug: Enable debug messages"))
	fmt.Printf("\t%s\n\n", blue("-sArgs: Comma separated list of SRC queue args in form of 'key:value:type'. Example: 'x-message-ttl:3600000:int,x-ha-policy:all:string'"))
	fmt.Printf("\t%s\n\n", blue("-dArgs: Comma separated list of DST queue args in form of 'key:value:type'. Example: 'x-message-ttl:3600000:int,x-ha-policy:all:string'"))
	fmt.Printf("\t%s\n\n", blue("-srcTls: Use TLS/AMQPS connection for Source Host"))
	fmt.Printf("\t%s\n\n", blue("-dstTls: Use TLS/AMQPS connection for Destination Host"))
	fmt.Printf("\t%s\n\n", blue("-verifyTls: Verify TLS certificate"))

	fmt.Printf("  %s: \n", magenta("Example"))

	fmt.Printf("\t%s\n", cyan("amq_requeue.osx -srcHost <fqdn>"))
	fmt.Printf("\t\t%s\n", cyan("-srcPort 5671 \\"))
	fmt.Printf("\t\t%s\n", cyan("-srcTls \\"))
	fmt.Printf("\t\t%s\n", cyan("-srcUser username \\"))
	fmt.Printf("\t\t%s\n", cyan("-srcPass ******** \\"))
	fmt.Printf("\t\t%s\n", cyan("-srcQueue queue_name \\"))
	fmt.Printf("\t\t%s\n", cyan("-srcVhost broker_name \\"))
	fmt.Printf("\t\t%s\n", cyan("-srcHost <fqdn> \\"))
	fmt.Printf("\t\t%s\n", cyan("-srcPort 5671 \\"))
	fmt.Printf("\t\t%s\n", cyan("-dstTls \\"))
	fmt.Printf("\t\t%s\n", cyan("-dstUser username \\"))
	fmt.Printf("\t\t%s\n", cyan("-dstPass ******** \\"))
	fmt.Printf("\t\t%s\n", cyan("-dstQueue queue_name \\"))
	fmt.Printf("\t\t%s\n", cyan("-dstVhost broker_name \\"))
	fmt.Printf("\t\t%s\n", cyan("-sArgs 'x-message-ttl:10800000:int,x-ha-policy:ha-all:string' \\"))
	fmt.Printf("\t\t%s\n", cyan("-dArgs 'x-message-ttl:10800000:int,x-ha-policy:ha-all:string' \\"))
	fmt.Printf("\t\t%s\n\n", cyan("-verifyTls=true"))
}
