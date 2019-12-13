import com.alibaba.fastjson.JSONObject;
import org.apache.activemq.ActiveMQConnectionFactory;
import org.apache.activemq.command.ActiveMQTextMessage;

import javax.jms.*;

public class QueueConsumer {

        private ConnectionFactory connectionFactory;
        private Connection connection;
        private Session session;
        private Destination destination;
        // 注意这里是消息接收（消费）者
        private MessageConsumer messageConsumer;

        public static void main (String[]args){
            QueueConsumer consumer = new QueueConsumer();
            consumer.doReceive();
        }

        public void doReceive () {
            try {
                connectionFactory = new ActiveMQConnectionFactory("failover:tcp://192.168.2.31:61616");
                connection = connectionFactory.createConnection("admin","admin");
                connection.start();
                session = connection.createSession(true, Session.AUTO_ACKNOWLEDGE);
                destination = session.createTopic("testMessage");

                /**
                 * 注意：这里要创建一个消息消费，并指定目的地（即消息源队列）
                 */
                messageConsumer = session.createConsumer(destination);

                // 方式一：监听接收
                receiveByListener();

                // 方式二：阻塞接收
//                 receiveByManual();

                /**
                 * 注意：这里不能再关闭对象了
                 */
                // messageConsumer.close();
            } catch (Exception e) {
                e.printStackTrace();
            } finally {
                /**
                 * 注意：这里不能再关闭Connection了
                 */
                // connection.close();
            }

        }

        /**
         * 通过注册监听器的方式接收消息，属于被动监听
         */
        private void receiveByListener () {
            try {
                messageConsumer.setMessageListener(new MessageListener() {

                    @Override
                    public void onMessage(Message message) {
                        if (message instanceof TextMessage) {
                            try {
                                TextMessage msg = (TextMessage) message;
                                System.out.println("Received:“" + msg.getText()
                                        + "”");
                                // 可以通过此方法反馈消息已收到
                                JSONObject o = JSONObject.parseObject(msg.getText());
                                System.out.println(o.get("1").getClass().toString());
                                msg.acknowledge();
                            } catch (Exception e) {
                                e.printStackTrace();
                            }

                        }

                        else {
                            System.out.println( message.getClass().toString());
                        }

                    }
                });
            } catch (Exception e) {
                e.printStackTrace();
            }
        }

        /**
         * 通过手动去接收消息的方式，属于主动获取
         */
        private void receiveByManual() {
            while (true) {
                try {
                    /**
                     * 通过receive()方法阻塞接收消息，参数为超时时间（单位：毫秒）
                     */
                    TextMessage message = (TextMessage) messageConsumer
                            .receive(60000);
                    if (message != null) {
                        System.out.println("Received:“" + message.getText() + "”");
                    }
                } catch (Exception e) {
                    e.printStackTrace();
                }

            }
        }
    }

