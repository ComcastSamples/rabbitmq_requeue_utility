# requeue

* *Language*: Go
* *Description*: A CLI tool to move messages from one AMQ cluster to another

<img src="doc/requeue.png" width="700">

## Build
```
# go build -o amq_requeue *.go
```

## Examples
* Move messages to queues with *no arguments*:
```
# ./amq_requeue -srcHost cluster1.example.com \
                -srcPort 5672 \
	        -srcUser username \
	        -srcPass password \
	        -srcQueue queue_one \
	        -srcVhost cluster1-broker \
	        -dstHost cluster2.example.com \
	        -dstPort 5672 \
	        -dstUser username \
	        -dstPass password \
	        -dstQueue queue_two \
	        -dstVhost cluster2-broker
```

* Move messages to queues with *arguements*:
```
# ./amq_requeue -srcHost cluster1.example.com \
                -srcPort 5672 \
	        -srcUser username \
	        -srcPass password \
	        -srcQueue queue_one \
	        -srcVhost cluster1-broker \
	        -dstHost cluster2.example.com \
	        -dstPort 5672 \
	        -dstUser username \
	        -dstPass password \
	        -dstQueue queue_two \
	        -dstVhost cluster2-broker \
                -qArgs 'x-message-ttl:3600000:int,x-ha-policy:all:string'
```

### License
Licensed under the Apache 2.0 License
