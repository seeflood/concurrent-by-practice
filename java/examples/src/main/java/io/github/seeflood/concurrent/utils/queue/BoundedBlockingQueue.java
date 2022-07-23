package io.github.seeflood.concurrent.utils.queue;

import java.util.concurrent.locks.Condition;
import java.util.concurrent.locks.Lock;
import java.util.concurrent.locks.ReadWriteLock;
import java.util.concurrent.locks.ReentrantReadWriteLock;

public class BoundedBlockingQueue {
    int[] arr;
    int   cap;
    int   len;
    final ReadWriteLock rwlock   = new ReentrantReadWriteLock();
    final Lock          rlock    = rwlock.readLock();
    final Lock          wlock    = rwlock.writeLock();
    final Condition     notFull  = wlock.newCondition();
    final Condition     notEmpty = wlock.newCondition();

    public BoundedBlockingQueue(int capacity) {
        arr = new int[capacity];
        cap = capacity;
    }

    public void enqueue(int element) throws InterruptedException {
        wlock.lock();
        try {
            while (this.len == cap) {
                notFull.await();
            }
            arr[len] = element;
            this.len++;
            notEmpty.signalAll();
        } finally {
            wlock.unlock();
        }
    }

    public int dequeue() throws InterruptedException {
        wlock.lock();
        try {
            while (this.len == 0) {
                notEmpty.await();
            }
            int toDel = arr[--len];
            notFull.signalAll();
            return toDel;
        } finally {
            wlock.unlock();
        }
    }

    public int size() {
        rlock.lock();
        try {
            return len;
        } finally {
            rlock.unlock();
        }
    }
}
