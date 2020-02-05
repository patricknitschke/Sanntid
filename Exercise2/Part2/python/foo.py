# Python 3.3.3 and 2.7.6
# python fo.py

from threading import Thread
import threading

# Potentially useful thing:
#   In Python you "import" a global variable, instead of "export"ing it when you declare it
#   (This is probably an effort to make you feel bad about typing the word "global")
i = 0
lock = threading.Lock()

def incrementingFunction():
    global i
    global lock
    lock.acquire()
    # TODO: increment i 1_000_000 times
    for n in range(6):
        i = i+1
        print("i", i)
    lock.release()

def decrementingFunction():
    global i
    global lock

    lock.acquire()
    # TODO: decrement i 1_000_000 times
    for n in range(7):
        i = i-1
        print("d", i)
    lock.release()

def main():
    global i
    global lock

    incrementing = Thread(target = incrementingFunction, args = (),)
    decrementing = Thread(target = decrementingFunction, args = (),)
    
    # TODO: Start both threads

    incrementing.start()
    decrementing.start()

    incrementing.join()
    decrementing.join()
    
    print("The magic number is %d" % (i))


main()
