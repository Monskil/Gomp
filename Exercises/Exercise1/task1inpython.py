
from threading import Thread

global i
i = 0

def createThread1():
    global i
    for n in range (0,1000000):
         i = i +1

def createThread2():
    global i
    for n in range (0,1000000):
         i = i -1



def main():
    thread_1 = Thread( target = createThread1, args=(),)
    thread_1.start()
    thread_2 = Thread( target = createThread2, args =(),)
    thread_2.start()

    
    thread_1.join()
    thread_2.join()

    print("hello")
    print(i)




main()
