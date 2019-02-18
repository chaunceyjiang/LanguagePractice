import java.util.Scanner;

public class StringDemo1 {
    public static void main(String[] args){
        String s1 = new String();
        System.out.println("s1:"+s1);
        System.out.println("s1.length:"+s1.length());
        System.out.println("--------------------------");



        byte[] bytes = { 97, 98, 99, 100, 101 };
        String s2 = new String(bytes);
        System.out.println("s1:"+s2);
        System.out.println("s1.length:"+s2.length());
        System.out.println("--------------------------");
        char[] chs = { 'a', 'b', 'c', 'd', 'e', '林', '青', '霞' };


        String s4 = new String(chs);
        System.out.println("s4:" + s4);
        System.out.println("s4.length():" + s4.length());
        System.out.println("--------------------------");
        Character ch = new Character('a');
        System.out.println(Character.isUpperCase('A'));
        String s5 = new String(chs,5,1);
        System.out.println("s4:" + s5);
        System.out.println("s4.length():" + s5.length());
        System.out.println("--------------------------");

        String s6 = new String("hello");
        String s7 = "hello";
        System.out.println(s6==s7);
        System.out.println(s6.equals(s7));
        System.out.println(s6.charAt(1));
        System.out.println(s6.substring(1));

    }
}
