## Parking lot code exercise 
Parking lot code design in Golang. 

The system is an terminal type application that's able to:
- Initialize a parking lot with a limit of cars
- Park a car using its plate and color and giving him a parking place
- Search for cars with specific color and giving its position on the parking lot
- Remove a car from the parking lot

to run: `go build main.go`

Some explanations
---
we are using the command pattern here and the main.go file is an interface 
between the logic and the file that represents the database, so all the 
logic on our commands are free of side effects, its one input and one output.

the data manipulation is done by our parking-lot file and not by our commands, 
that simply manipulate witch information should be saved.


##### Yes we are using channels.
It seems queer that we have channels and goroutines everywhere, but it makes sense 
if you want to change the input type.

Today we only support terminal input, but if we want to support an API as well, we
just need to separate the switch on the main file using the channel as a communication
way between the business and the output strategy.




