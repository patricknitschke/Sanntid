# Reasons for concurrency and parallelism


To complete this exercise you will have to use git. Create one or several commits that adds answers to the following questions and push it to your groups repository to complete the task.

When answering the questions, remember to use all the resources at your disposal. Asking the internet isn't a form of "cheating", it's a way of learning.

 ### What is concurrency? What is parallelism? What's the difference?
 > Concurrency refers to the concurrent(several at once) running, starting and completion of tasks in an overlapping time period. Parallelism refers to the running of two or more tasks at the same time. The difference is that parallelism is the simultaneous starting, running and completion of tasks to complete a specific objective whereas concurrency are (possibly independent) tasks that start and end at different times, but have overlapping time periods when they run.
 
 ### Why have machines become increasingly multicore in the past decade?
 > Since the speed at which processor performance increases has slowed in recent years, creating machines with muliple cores and utilising parallelism is a way of boosting the performance of computers without directly improving the processor. 
 
 ### What kinds of problems motivates the need for concurrent execution?
 (Or phrased differently: What problems do concurrency help in solving?)
 > Concurrent execution allows computers to bypass slow tasks which require waiting, such as I/O tasks, by allowing the computer to concurrently start a new task and run that, while the previous task is waiting.

 ### Does creating concurrent programs make the programmer's life easier? Harder? Maybe both?
 (Come back to this after you have worked on part 4 of this exercise)
 > Concurrent programs are subject to complex bugs, and are difficult to write in a good way. However, they are necessary for a lot of processes. 
 
 ### What are the differences between processes, threads, green threads, and coroutines?
 > A process is a programme in execution, while a thread is a segment of a process. A green thread is a thread that is scheduled by a runtime library or virtual machiene. Coroutines are very similar to threads. However, coroutines are cooperatively multitasked, whereas threads are typically preemtively multitasked. This means instead of being interrupted by the OS like threads, coroutines voluntarily give up their processes periodically to the OS.
 
 ### Which one of these do `pthread_create()` (C/POSIX), `threading.Thread()` (Python), `go` (Go) create?
 > These create threads.
 
 ### How does pythons Global Interpreter Lock (GIL) influence the way a python Thread behaves?
 > The GIL ensures that only one thread can execute at the time. This is necessary because python's memory managing is not thread safe.
 
 ### With this in mind: What is the workaround for the GIL (Hint: it's another module)?
 > You can download packages such as thread which will then allow you to use threads in python.
 
 ### What does `func GOMAXPROCS(n int) int` change? 
 > The function decides the number of CPUs that can run at the same time.
