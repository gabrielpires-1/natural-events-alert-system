package mingati.luis.producerprojeto;

public class Message {
    private String time;
    private String location;
    private String topic;
    private int value;

    public Message() {
    }

    public Message(String time, String location, String topic, int value) {
        this.time = time;
        this.location = location;
        this.topic = topic;
        this.value = value;
    }

    public String getTime() {
        return time;
    }

    public void setTime(String time) {
        this.time = time;
    }

    public String getLocation() {
        return location;
    }

    public void setLocation(String location) {
        this.location = location;
    }

    public String getTopic() {
        return topic;
    }

    public void setTopic(String topic) {
        this.topic = topic;
    }

    public int getValue() {
        return value;
    }

    public void setValue(int value) {
        this.value = value;
    }
}
