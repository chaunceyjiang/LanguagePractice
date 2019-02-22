import java.math.BigInteger;
import java.net.InetAddress;
import java.net.UnknownHostException;
import java.nio.ByteOrder;

public class IPUtil {
    public static String longToIp(long ip) {
        return ((ip >> 24 ) & 0xFF) + "." +

                ((ip >> 16 ) & 0xFF) + "." +

                ((ip >>  8 ) & 0xFF) + "." +

                ( ip        & 0xFF);
    }
    public static Long ipToInt(String addr) {

        String[] addrArray = addr.split("\\.");
        long num = 0;
        for (int i = 0; i < addrArray.length; i++) {

            int power = 3 - i;
            num += ((Integer.parseInt(addrArray[i]) % 256 * Math.pow(256, power)));

        }
        return num;

    }

    public static void main(String[] args) {
        int ip = -1062731496;
        String IP = longToIp(ip);
        System.out.println(IP);
        System.out.println(longToIp(3232235800L));
//        System.out.println(IP);
        System.out.println(ipToInt("192.168.1.24"));
    }
}

