package mingati.luis.producerprojeto;

import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;

@SpringBootApplication
public class ProducerProjetoApplication {

    public static void main(String[] args) {

        SpringApplication.run(ProducerProjetoApplication.class, args);
        Producer.runProducer();
    }

}
