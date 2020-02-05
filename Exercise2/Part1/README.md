# Mutex and Channel basics

### What is an atomic operation?
> It is program operations that run completely independent of any other processes.

### What is a semaphore?
> It's a variable that limits the number of consumers that's allowed to access a specific resource. For us, this would mean that semaphores is a method of limiting the number of tasks that is allowed to access our memory, databases, files, etc. 

### What is a mutex?
> A mutex is a binary flag that is created by a task to allow it to use a specific resource. This means that there is only one resource that can only be used by one task at a time, controlled by using a mutex.

### What is the difference between a mutex and a binary semaphore?
> Semaphores are used for signalling (Synchronisation) while a mutex is used for resource sharing. Smaphores should either signal or wait for a specific resource. For example, Task B is waiting for a specific event, event occurs, semaphore signals task B to perform tasks, task B returns to waiting.

### What is a critical section?
> A critical section is a section of code that accesses shared variables and has to be executed as an atomic operation. In a group of cooperating processes, only one process must execute its critical section at a time.

### What is the difference between race conditions and data races?
 > Data races is a subsection of race conditions where a minimum of two processes access the same memory, and at least one of them is performing a writing operation

### List some advantages of using message passing over lock-based synchronization primitives.
> Message passing is easier to get right. It also isn't reliant on locks which may improve performance. It is also easier to scale in size and number of threads.

### List some advantages of using lock-based synchronization primitives over message passing.
> It gives better control to the programmer in what can be accessed and when. It will also have a lower memory footprint as it isn't reliant on creating messages that is sent between threads.
