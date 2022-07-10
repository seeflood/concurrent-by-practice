package io.github.seeflood.concurrent.primitives.job.sequential;

import java.util.concurrent.Semaphore;

public class Printer {
    int threadCnt;
    int max;
    int v = 0;

    public Printer(int threadCnt, int max) {
        this.threadCnt = threadCnt;
        this.max = max;
    }

    public void run() throws InterruptedException {
        Semaphore[] semaphores = new Semaphore[threadCnt];
        for (int i = 0; i < threadCnt; i++) {
            semaphores[i] = new Semaphore(0);
        }

        // create and start threads
        Thread[] threads = new Thread[threadCnt];
        for (int i = 0; i < threadCnt; i++) {
            int idx = i;
            threads[i] = new Thread(new Runnable() {
                @Override
                public void run() {
                    try {
                        while (true) {
                            int prev = (idx + threadCnt - 1) % threadCnt;
                            // wait for previous node
                            semaphores[prev].acquire();
                            // quit if v>max
                            if (v > max) {
                                semaphores[idx].release();
                                return;
                            }
                            System.out.println("thread " + idx + ": " + v);
                            v++;
                            semaphores[idx].release();
                        }
                    } catch (InterruptedException e) {
                        e.printStackTrace();
                    }
                }
            });
            threads[i].start();
        }
        System.out.println("Start...");
        //    produce
        semaphores[threadCnt - 1].release();
        //    wait
        for (int i = 0; i < threadCnt; i++) {
            threads[i].join();
        }
    }

    public static void main(String[] args) throws InterruptedException {
        new Printer(10, 100).run();
        System.out.println("All printed!");
    }

}
