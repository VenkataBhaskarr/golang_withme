## This repo contains all my learning about one of the famous open source programming languages GO

## check the main.go file for further documentation with written code samples

## sequential (vs) concurrent (vs) parrallel processing (tasks == process)

sequentail processing is a step by step process solving

PC1 - task1..........task2...........task3......

concurrent processing is switching between the tasks with a little time quantum such that it appeaars like all the tasks are running at any given point of time

PC1 - task1..task3..task2...task1..task2..task3.

with multi core

PC1 - task1...task2...task1...task3...
PC2 - task2...task3...task2...task1...


parallel proccessing 

this is different because here every core executes only a single task until it is completed 

PC1 - task1...........
PC2 - task2...........


go routines are lightweight threads why simple answer is threads are managed by OS which takes some time for schedulking concurrently and so on the otherhand go routines are fully managed by go run time leaving with lot of optimizations

## Go runtime scheduler
when you run a go program it will create 'n' OS threads
where n = no.of cores*hardware threads(no.of threads that can run on a single core)


now these n OS threads are managed by the OS itself and they are heavier in size and evcery thread has a LRQ local run queue on which the goroutines are orchestrated by the goruntime. remaining goroutines are kept in GRQ global run queue

go routines are light weight and can be created so many when compared to OS threads and they are fast because they are natively managed by the runtime itself

goruntime has a scheduler which will multiplex the goroutines to the OS threads (m:n multiplexing)

## we have seen before in go routines the main go routine is exiting which results in undfortunate exiting of the subsequent go rouintes but we handled the problem with time.Sleep but we can solve it effictibvely

here comes the process wait groups its a synchronization primitive offered by go runtime basically it waits for sync completion of treads

```

var wg sync.waitGroup
wg.Add(int)         ---> it will mark the counter which means no.of go routines
wg.done()           ---> it will decrease the internal count established by the add function
wg.wait()           ---> it will wait until the inner counter established by the add function becomes zero

```

```
package main
import (
    "fmt"
    "sync"
)

func square(n int , wg *sync.WaitGroup){
    fmt.Println(n*n)
    defer wg.Done()
}

func main() {
  var wg sync.WaitGroup
  wg.Add(10)
  for i:=0;i<10;i++{
      go square(i ,&wg)
  }
  wg.Wait()
}
```

## what are channels ?

Intraditional programming when we create threads in order to establish communication and synchronization we should folllow some lpw level programming constructs such aas locks, mutexes to the shared memory, deadlocks and thinds in order to ensure thread safety, wheraes in golang channels are the way which are inbuilt uses synchronization methods such that it will ellows us to focus on code rather than all those low level concepts for synchronization and communication.

because in channels the channel waits until there is recieveing side/accepting side hence remaing go routines can't access that channel
maintaing sync without explaicit locks , mutexes and all those things.

channels block the execution of goroutines until thechannel resolves or else there might be the situations where there can be deadlocks 

channels in golang are both buffered and nonbuffered(default) where in buffered the go routine is not blocked when there are availabel bufferes to store the value and vice versa buffered channels follows the queue order while reading or writing the data

this is the basic declaration of channels

```
ch := make(chan string)
ch <- 5
val := <- ch
close(ch)

```






