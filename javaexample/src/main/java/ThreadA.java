public class ThreadA extends Thread {
    private HasSelfProvateNum numref;

    public ThreadA(HasSelfProvateNum numref) {
        super();
        this.numref = numref;
    }

    public void run() {
        super.run();
        numref.add("a");
    }
}
