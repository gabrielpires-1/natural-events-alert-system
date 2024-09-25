package mingati.luis.producerprojeto;

import com.rabbitmq.client.Channel;
import com.rabbitmq.client.Connection;
import com.rabbitmq.client.ConnectionFactory;

import java.time.LocalDateTime;
import java.time.format.DateTimeFormatter;
import java.util.Scanner;

public class Producer {
    private static final String EXCHANGE_NAME = "alert_topic";
    private static final String TEMP_KEY = "temperature";
    private static final String PRESSURE_KEY = "pressure";
    private static final String RAIN_KEY = "rain";
    private static final String SISMIC_KEY = "sismic";

    public static void runProducer() {
        System.out.println("Starting system server...");

        String uri = "amqps://jykccpdt:XdmtO-wpTbv9T_XRlRHtGOO86h51W0Ir@jackal.rmq.cloudamqp.com/jykccpdt";

        try {
            ConnectionFactory factory = new ConnectionFactory();
            factory.setUri(uri);

            Connection connection = factory.newConnection();
            Channel channel = connection.createChannel();

            channel.exchangeDeclare(EXCHANGE_NAME, "topic");

            Scanner scanner = new Scanner(System.in);
            System.out.println("You are the producer. Which metric can you share?");
            System.out.println("1 - temperature");
            System.out.println("2 - pressure");
            System.out.println("3 - rain");
            System.out.println("4 - sismic activity");

            int option = scanner.nextInt();
            scanner.nextLine();

            System.out.println("Input location: ");
            String location = scanner.nextLine().toLowerCase();

            String topicKey;
            switch (option) {
                case 1:
                    topicKey = TEMP_KEY;
                    break;
                case 2:
                    topicKey = PRESSURE_KEY;
                    break;
                case 3:
                    topicKey = RAIN_KEY;
                    break;
                case 4:
                    topicKey = SISMIC_KEY;
                    break;
                default:
                    System.out.println("Invalid option");
                    scanner.close();
                    channel.close();
                    connection.close();
                    return;
            }

            publishMessages(scanner, channel, location, topicKey);

            scanner.close();
            channel.close();
            connection.close();

        } catch (Exception e) {
            e.printStackTrace();
        }
    }

    private static void publishMessages(Scanner scanner, Channel channel, String location, String topicKey) throws Exception {
        String routingKey = "sensor." + location + "." + topicKey;
        DateTimeFormatter formatter = DateTimeFormatter.ofPattern("yyyy-MM-dd HH:mm:ss");

        while (true) {
            int value;
            System.out.println("Input " + topicKey + " value:");
            value = scanner.nextInt();

            String message = "{"
                    + "\"time\":\"" + LocalDateTime.now().format(formatter) + "\","
                    + "\"location\":\"" + location + "\","
                    + "\"topic\":\"" + topicKey + "\","
                    + "\"value\":" + value
                    + "}";

            channel.basicPublish(EXCHANGE_NAME, routingKey, null, message.getBytes("UTF-8"));

            System.out.println(" [x] Sent '" + routingKey + "':'" + message + "'");
        }
    }
}
