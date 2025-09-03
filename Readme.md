# OA1
##
I utilized a heap for the first OA question.
In Java, I used a PriorityQueue for implementing it.
After that, I added each element in the heap, and it will automatically sorted for me.
Then, I merge the first and the second elements, and added the result in the heap again.
But for Go, there is no library for the heap, so I implemented after I googled it.
And the rest of logic is the same with the one I resolved with JAVA


# OA2
##
I utilized a hash map for the second OA question.
In Java, I used a hash map for implementing it.
So, there are three parameters -> JobNums // Gateways // m (gate size)
First, I implemented a hash map for JobNums and Gateways.
So, I can set the total jobs for each gate.
Second, by using for loop, I created and update each gateways jobs.
If a specific gate job is 0, for Java, I used a getOrDefault function, 

and for Go, I used an if statement for filtering it.# IBM_OA

# How to run it
##
<OA1>
- go run OA1.go

<OA2>
- go run OA2.go
