# Natural Events Alert System (NEAS)

A system utilizing a Publish/Subscribe architecture to provide alerts about natural events. It consists of a Java publisher and a Go system with publishers, subscribers, and a log.

#### Publishers
Publishers are data producers that publish messages containing natural metric data about specific locations. They choose the location and the topic of the measurement. Messages are sent to the exchange `alert_topic` with the following routing key format: `sensor.<location>.<topic>`.  

#### Consumers
Consumers are systems that receive and use messages published by the publishers. They can subscribe to one or more topics. A consumer can subscribe to the `alert_topic` exchange with the following routing keys:

- `sensor.#` (log - receives everything)
- `sensor.*.<topic>`
- `sensor.<location>.<topic>`
- `sensor.<location>.*`

#### Example: Sending a Temperature Value from Recife

The message is a JSON object like the following:

```json
{
    "time": "2024-09-30 14:24:08",
    "location": "recife",
    "topic": "temperature",
    "value": 20
}
``` 

The message above is sent to the exchange `alert_topic` with the routing key `sensor.recife.temperature`. Every consumer which is subscribed to the exchange with the routing key `sensor.#`, `sensor.*.temperature`, `sensor.recife.*`, or `sensor.recife.temperature` will receive the message.


## Tools and programming languages
- [RabbitMQ](https://www.rabbitmq.com/)
- [Java](https://www.java.com/)
- [Go](https://go.dev/)

## Executing by executables

###  1. Executing Go system
- **Linux**
```bash
./executables/app-linux
``` 
- **Linux-arm64**
```bash
./executables/app-linux
``` 

- **macOS**
```bash
./executables/app-macos
``` 

- **Windows**
```bash
./executables/app-windows.exe
``` 

### 2. Executing Java system
Just run:

```bash
java -jar ./executables/Producer-projeto-0.0.1-SNAPSHOT.jar
``` 

## Building your own executables and using your own AMQP Cloud Instance

### Prerequisites
Ensure the following tools are installed:
- RabbitMQ (via cloud or locally)    
- Java Development Kit (JDK) and Maven  
- Go Programming Language

### RabbitMQ
1. Sign up at [CloudAMQP](https://www.cloudamqp.com/) 
2. Login in to your account
3. Create a new instance
4. Access the instance and retrieve the AMQP URL from the AMQP details.
5. Put the URL in the file *.env* in project source with the following: `AMQP_CONNECTION=<your_URL>`
6. Go to Producer.java class and change the code line 21 to: `String uri = <your_URL>`

### Setting up Go
In your terminal, in project source, use the command below to install required libraries: 
```bash
go mod tidy
``` 

To run Go code:

```bash
go run .\system\main
``` 

### Setting up Java
In your terminal, run to compile Java code:
```bash
cd Producer-projeto
./gradlew build
``` 

To run the executable:
```bash
java -jar build/libs/Producer-projeto-0.0.1-SNAPSHOT.jar
```
