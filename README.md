# Kafka Konsumer

## Description

Kafka Konsumer provides an easy implementation of Kafka consumer with a built-in retry/exception
manager ([kafka-cronsumer](https://github.com/Trendyol/kafka-cronsumer)).

## Guide

### Installation

```sh
go get github.com/Trendyol/kafka-konsumer@latest
```

### Examples

You can find a number of ready-to-run examples at [this directory](examples).

After running `docker-compose up` command, you can run any application you want.

#### Without Retry/Exception Manager

```go
func main() {
    consumerCfg := &kafka.ConsumerConfig{
        Reader: kafka.ReaderConfig{
            Brokers: []string{"localhost:29092"},
            Topic:   "standart-topic",
            GroupID: "standart-cg",
        },
        ConsumeFn:    consumeFn,
        RetryEnabled: false,
    }

    consumer, _ := kafka.NewConsumer(consumerCfg)
    defer consumer.Stop()
    
    consumer.Consume()
}

func consumeFn(message kafka.Message) error {
    fmt.Printf("Message From %s with value %s", message.Topic, string(message.Value))
    return nil
}

```

#### With Retry/Exception Option Enabled

```go
func main() {
    consumerCfg := &kafka.ConsumerConfig{
        Reader: kafka.ReaderConfig{
            Brokers: []string{"localhost:29092"},
            Topic:   "standart-topic",
            GroupID: "standart-cg",
        },
        RetryEnabled: true,
        RetryConfiguration: kafka.RetryConfiguration{
            Topic:         "retry-topic",
            StartTimeCron: "*/1 * * * *",
            WorkDuration:  50 * time.Second,
            MaxRetry:      3,
        },
        ConsumeFn: consumeFn,
    }
    
    consumer, _ := kafka.NewConsumer(consumerCfg)
    defer consumer.Stop()
    
    consumer.Consume()
}

func consumeFn(message kafka.Message) error {
    fmt.Printf("Message From %s with value %s", message.Topic, string(message.Value))
    return nil
}
```

## Configurations

| config                             | description                                                                                                                           | default          |
|------------------------------------|---------------------------------------------------------------------------------------------------------------------------------------|------------------|
| `reader`                           | [Describes all segmentio kafka reader configurations](https://pkg.go.dev/github.com/segmentio/kafka-go@v0.4.39#ReaderConfig)          |                  |
| `consumeFn`                        | Kafka consumer function, if retry enabled it, is also used to consume retriable messages                                              |                  |
| `logLevel`                         | Describes log level; valid options are `debug`, `info`, `warn`, and `error`                                                           | info             |                          |
| `concurrency`                      | Number of goroutines used at listeners                                                                                                | runtime.NumCPU() |
| `retryEnabled`                     | Retry/Exception consumer is working or not                                                                                            | false            |
| `retryConfiguration.startTimeCron` | Cron expression when retry consumer ([kafka-cronsumer](https://github.com/Trendyol/kafka-cronsumer#configurations)) starts to work at |                  |
| `retryConfiguration.workDuration`  | Work duration exception consumer actively consuming messages                                                                          |                  |
| `retryConfiguration.topic`         | Retry/Exception topic names                                                                                                           |                  |
| `retryConfiguration.maxRetry`      | Maximum retry value for attempting to retry a message                                                                                 | 3                |
| `tls.rootCAPath`                   | [see doc](https://pkg.go.dev/crypto/tls#Config.RootCAs)                                                                               | ""               |
| `tls.intermediateCAPath`           | Same with rootCA, if you want to specify two rootca you can use it with rootCAPath                                                    | ""               |
| `sasl.authType`                    | `SCRAM` or `PLAIN`                                                                                                                    |                  |
| `sasl.username`                    | SCRAM OR PLAIN username                                                                                                               |                  |
| `sasl.password`                    | SCRAM OR PLAIN password                                                                                                               |                  |
| `logger`                           | If you want to custom logger                                                                                                          | info             |                          |
