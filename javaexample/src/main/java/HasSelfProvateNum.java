public class HasSelfProvateNum {
    private int num = 0;

    synchronized public void add(String username) {
        try {
            if (username.equals("a")) {
                num = 100;
                System.out.println("a set over");
                Thread.sleep(10000);
            } else {
                num = 200;
                System.out.println("b set over!");
            }
            System.out.println(username + "num=" + num);
        } catch (InterruptedException e) {
            e.printStackTrace();
        }
    }
}
