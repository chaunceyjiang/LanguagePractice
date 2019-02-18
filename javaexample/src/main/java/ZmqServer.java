package zmq.resp;

import org.zeromq.ZMQ;
import org.zeromq.ZMQ.Context;
import org.zeromq.ZMQ.Socket;

public class ZmqServer {
	private transient final Context context = ZMQ.context(1);
	
	private transient Socket socket;
	
	private int HWM = 20000;
	
	private String connect;
	
	public void open() {
		// 采用push/pull模式
		this.socket = context.socket(ZMQ.PULL);
		this.socket.setHWM(HWM);
		this.socket.bind(connect);
	}
	
	public void setConnect(String connect) {
		this.connect = connect;
	}
	
	public void recv() {
		boolean wait = false;
		while(!wait) {
			byte[] bytes = socket.recv();
			String getData = new String(bytes);
			System.out.println(new String(getData));
			System.out.println("=====================");
		}
	}
	
	public static void main(String[] args) {
		ZmqServer server = new ZmqServer();
		server.setConnect("tcp://192.168.2.147:23203");
		server.open();
		server.recv();
	}
}
