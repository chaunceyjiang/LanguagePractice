import java.util.concurrent.locks.Lock;
import java.util.concurrent.locks.ReentrantLock;

public class Run {
    public static void main(String[] args) throws InterruptedException {
//        MyThread t = new MyThread("A");
//        t.start();
////        Thread.sleep(1000);
//        System.out.println("运行结束");
        HasSelfProvateNum numref1 = new HasSelfProvateNum();
        HasSelfProvateNum numref2 = new HasSelfProvateNum();
        ThreadA threadA = new ThreadA(numref1);
        threadA.start();
        ThreadB threadB = new ThreadB(numref2);
        threadB.start();
        ThreadService s = new ThreadService();
        s.service1();
        Lock lock = new ReentrantLock();
        lock.lock();

        lock.unlock();
        lock.lockInterruptibly();
        lock.unlock();

    }
}
