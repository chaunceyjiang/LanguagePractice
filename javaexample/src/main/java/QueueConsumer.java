import com.alibaba.fastjson.JSONObject;
import org.apache.activemq.ActiveMQConnectionFactory;
import org.apache.activemq.command.ActiveMQTextMessage;

import javax.jms.*;

public class amq {
    public static void main(String[] args){
        ConnectionFactory connectionFactory =
                new ActiveMQConnectionFactory("tcp://192.168.2.31:61616");
        // 声明Connection
        Connection connection = null;
        // 声明Session
        Session session = null;
        // 声明Producer
        MessageProducer producer = null;
        MessageConsumer consumer = null;
        try{
            // 从连接工厂中获得连接
            connection = connectionFactory.createConnection();
            // 开启连接
            connection.start();
            /*
             * 从连接中获得会话
             * 参数1:transacted  boolean型
             * 当设置为true时,将忽略参数2,acknowledgment mode被jms服务器设置 SESSION_TRANSACTED。
             * 当一个事务被提交时,消息确认就会自动发生。
             * 当设置为false时,需要设置参数2
             * Session.AUTO_ACKNOWLEDGE为自动确认，当客户成功的从receive方法返回的时候，或者从
             * MessageListener.onMessage方法成功返回的时候，会话自动确认客户收到的消息。
             * Session.CLIENT_ACKNOWLEDGE 为客户端确认。客户端接收到消息后，必须调用javax.jms.Message的
             * acknowledge方法。jms服务器才会删除消息。（默认是批量确认）
             */
            session = connection.createSession(false,Session.AUTO_ACKNOWLEDGE);
            // 创建一个Destination目的地 Queue或者Topic
            Topic topic = session.createTopic("testMessage");
            // 创建一个Producer生产者
            ObjectMessage onj = session.createObjectMessage();
            consumer = session.createConsumer(topic);
            // 创建message
            ActiveMQTextMessage textMessage = new ActiveMQTextMessage();
            textMessage.setText("test");
            // 发送message
            JSONObject obj = JSONObject.parseObject("{\"module\": \"report\", \"menus\": [{\"modelName\": \"report\", \"orderNumber\": 9, \"appId\": \"SIMO\", \"name\": \"\\\\u62a5\\\\u8868\", \"range\": [\"Super_Admin\"], \"id\": \"RPT01\", \"beClass\": true, \"iclass\": \"icon-ux-rpt\", \"parentId\": \"SIMO\", \"url\": \"/visual-static/report\"}]}");
            onj.setObject(obj);
//            producer.send(textMessage);
//            producer.send(onj);
            consumer.setMessageListener(new MessageListener() {
                @Override
                public void onMessage(Message message) {
                    if (message instanceof TextMessage) {
                        try {
                            TextMessage msg = (TextMessage) message;
                            System.out.println("Received:“" + msg.getText()
                                    + "”");
                            // 可以通过此方法反馈消息已收到
                            msg.acknowledge();
                        } catch (Exception e) {
                            e.printStackTrace();
                        }
                    }
                }
            });

        }catch(Exception e){
            e.printStackTrace();
        }finally{
            // 回收资源
//            producer.close();
//            session.close();
//            connection.close();
        }
    }
}
