package main

import (
	"flag"
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/streadway/amqp"
)

var debug bool = false

func printDebug(string1, string2 string) {
	if debug {
		if string2 == "" {
			fmt.Printf("   --> DEBUG:   %s\n", yellow(string1))
		} else {
			fmt.Printf("   --> DEBUG:   %s -->  %s\n", yellow(string1), yellow(string2))
		}
	}
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", red(msg), err)
	}
}

func main() {

	debugOption := flag.Bool("debug", false, "Enable debug logging")
	srcHost := flag.String("srcHost", "", "Source AMQ cluster")
	srcVhost := flag.String("srcVhost", "", "Source AMQ VHost")
	dstVhost := flag.String("dstVhost", "", "Destination AMQ VHost")
	dstHost := flag.String("dstHost", "", "Destination AMQ cluster")
	srcPort := flag.String("srcPort", "", "Source Port")
	dstPort := flag.String("dstPort", "", "Destination Port")
	srcUser := flag.String("srcUser", "", "Source Username")
	dstUser := flag.String("dstUser", "", "Destination Username")
	srcPass := flag.String("srcPass", "", "Source Password")
	dstPass := flag.String("dstPass", "", "Destination Password")
	srcQueue := flag.String("srcQueue", "", "Source queue to copy messages from")
	dstQueue := flag.String("dstQueue", "", "Destination queue to copy messages to")
	sArgs := flag.String("sArgs", "", "Comma separated list of SRC queue args in form of 'key:value:type'. Example: 'x-message-ttl:3600000:int,x-ha-policy:all:string'")
	dArgs := flag.String("dArgs", "", "Comma separated list of DST queue args in form of 'key:value:type'. Example: 'x-message-ttl:10800000:int,x-ha-policy:all:string'")
	flag.Parse()

	if *debugOption == true {
		debug = true
	}

	// Parse Queue Arguments
	srcQueueArgs := make(map[string]interface{})
	dstQueueArgs := make(map[string]interface{})

	parseArgs := func(inArgs, inType string) map[string]interface{} {
		tempMap := make(map[string]interface{})

		list := strings.Split(inArgs, ",")
		for _, arg := range list {
			args := strings.Split(arg, ":")
			fName := args[0]
			fValue := args[1]
			vType := args[2]
			if vType == "int" {
				i, _ := strconv.Atoi(fValue)
				tempMap[fName] = i
			} else {
				tempMap[fName] = string(fValue)
			}

			printDebug("Setting queue arguement -->  ", arg)
		}
		fmt.Printf("%s Queue Arguments: %s \n", inType, blue(tempMap))

		return tempMap
	}

	srcQueueArgs = parseArgs(*sArgs, "SRC")
	dstQueueArgs = parseArgs(*dArgs, "DST")

	// SRC Cluster Connection
	srcConnString := fmt.Sprintf("amqp://%s:%s@%s:%s/%%2F%s", *srcUser, *srcPass, *srcHost, *srcPort, *srcVhost)
	srcconn, err := amqp.Dial(srcConnString)
	failOnError(err, "Failed to connect to source RabbitMQ")
	defer srcconn.Close()

	srcch, err := srcconn.Channel()
	failOnError(err, "Failed to open source channel")
	defer srcch.Close()

	err = srcch.Qos(1, 0, false)
	failOnError(err, "Could not configure Qos on source cluster")

	srcq, err := srcch.QueueDeclare(
		*srcQueue,    // name
		true,         // durable
		false,        // delete when unused
		false,        // exclusive
		false,        // no-wait
		srcQueueArgs, // args
	)
	failOnError(err, "Failed to delcare source queue")

	//READ from SRC QUEUE
	msgs, err := srcch.Consume(
		srcq.Name, // name
		"",        // consumer
		true,      // auto-ack
		false,     // exclusive
		false,     // no-local
		false,     // no-wait
		nil,       // args
	)
	failOnError(err, "Failed to register a Consumer on source cluster")

	// DST Cluster Connection
	dstConnString := fmt.Sprintf("amqp://%s:%s@%s:%s/%%2F%s", *dstUser, *dstPass, *dstHost, *dstPort, *dstVhost)
	dstconn, err := amqp.Dial(dstConnString)
	failOnError(err, "Failed to connect to destination RabbitMQ cluster")
	defer dstconn.Close()

	dstch, err := dstconn.Channel()
	failOnError(err, "Failed to open a channel on the destination cluster")
	defer dstch.Close()

	dstq, err := dstch.QueueDeclare(
		*dstQueue,    // name
		true,         // durable
		false,        // delete when unused
		false,        // exclusive
		false,        // no-wait
		dstQueueArgs, // args
	)

	forever := make(chan bool)

	// Publish messages to DST cluster
	go func() {
		for srcmsg := range msgs {
			log.Printf("Recieved a message: %s", srcmsg.Body)

			err = dstch.Publish(
				"",        // exchange
				dstq.Name, // routing key
				false,     // mandatory
				false,     // immediate
				amqp.Publishing{
					ContentType: "text/plain",
					Body:        []byte(srcmsg.Body),
				})

			log.Printf(" [x] Sent %s", srcmsg.Body)
			failOnError(err, "Failed to publish a message on destination cluster")
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit, press CTRL+C")
	<-forever
}
