***Comments***

-pre allocation is making the code run much quicker

-with runes you can preallocate the exact memory for the specific character

-with strbuilder you can preallocate strings to the application will run much quicker

-with struct-interface we can create an interface, so we have a rule that certain methods
must have the correct cables to connect with the interface. Which in this makes it able
to reuse a function by just changing another method by in other words struct variable.

-with gorutines we can run code in multiple threads at once, which enables to run code
more efficient and faster. You need to use waitgroups to wait for a routine to finish, and
add this to the result. To make sure we don't lose data, we can add mutex locks. We can
also add read mutex locks so we can output information for each data allocation.

-with channels, channels can hold data, it is thread safe, and it can wait before data is
pretty much ready, this is used in go routines

-with generic types, for example you have a function with an int32 as parameter, 
but you want the same data but an int16, now you have to make another function which is
the exact same except for the type. With generics you can add kind of parameter types which
can save an significant amount of code writing.



