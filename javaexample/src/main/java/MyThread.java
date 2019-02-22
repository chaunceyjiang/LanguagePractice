public class MyThread extends Thread {
//    public void run(){
//        super.run();
//        System.out.println(this.getClass());
//        try {
//            Thread.sleep(1000000);
//        } catch (InterruptedException e) {
//            e.printStackTrace();
//        }
//    }
    private int count = 5;
    public MyThread(String name){
        super();
        this.setName(name);
    }
    public void run(){
        super.run();
        while (count >0){
            count--;
            System.out.println("由　"+MyThread.currentThread().getName()+"　计算，count ="+count);
        }
    }
    public static void main(String[] args){
        MyThread a = new MyThread("A");
        MyThread b = new MyThread("B");
        MyThread c = new MyThread("C");
        a.start();
        b.start();
        c.start();
    }
}
