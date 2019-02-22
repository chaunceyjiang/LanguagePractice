public class ThreadB extends Thread {
    private HasSelfProvateNum numref;

    public ThreadB(HasSelfProvateNum numref) {
        super();
        this.numref = numref;
    }

    public void run() {
        super.run();
        numref.add("b");
    }
}
