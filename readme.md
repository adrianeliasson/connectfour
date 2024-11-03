# Adrian's Connect Four game

Not really sure what goes here...

I write this game to get more familiar with the Go programming language.
Also to have something fun on my personal github page.

#### Some thoughts and ideas to guide development:
I like to be able to run it in the terminal. I want to become better and faster and more comfortable writing great CLI applications. Whenever I come across a really nice CLI application it is always so nice. They say Go is nice for writing good CLI applications. Bubbletea is one of the most popular packages to use for developing CLI applications, so it is both useful and fun to learn that.

So, writing a very simple old game called Connect Four that I can run in the terminal.

Right now I already implemented most of the things that make connect four connect four.

You can now officially play locally vs your friend on the same laptop.

Next step is to make it so you can play "online". I have some ideas of how to make that happen, but in the spirit of making the most straight forward / naive / dumb solution first, and then refactoring, what I am thinking currently is to just have a backend connecting two players.

I need to find a way for the players to 
- share game state
- could be fun to render the opponents cursor in real time
- I think that's it

I think it makes the most sense for the state to be kept in the server, and just let the CLI be the frontend interface. However i want to allow local play as well as online play so I need to keep game logic inside the client too.

Something I want to learn more is also ProtoBuffers / RPC/gRPC. It could be cool to share a ProtoBuf structure between the server and the client and use gRPC to pass the state around... Could there even be a way to avoid keeping any knowledge of the game on the server?? Just use the server to connect "websockets" together. I think thats the most fun challenge.


Sooo.... connecting websockets. Not sure how it works but lets see.

TODO in terms of online play

- Require multiple frontend views (Main menu allowing selection between offline play and online play (I also want a way to configure the default game mode choice so that you end up in your preferred mode without needing to select it every time (very late feature though))). Having multiple frontend views could allow more things like unlockable in-game stuff like mark skins, titles, achievements etc.
- When selecting online play, write the name of the player to want to versus. This would require you to set your own name when entering online play so that others can play vs you.
- Websocket Server should connect players who wrote eachothers names (how this will work in reality I have no clue). 

Okay after having researched websockets a bit (Reading parts of the RFC, reading implementations of it, watching videos of people going over it and also researching packages in GO that implements websocket for us to use freely), I think I have a good enough grasp of the technology to get going with implementing it myself.

I will be having a websocket server in GCP because that makes it easy and convenient for clients to reach no matter where they are.

What data to pass. I am thinking to allow sending a few message types
- Cursor placement (Just pass an int between 0-6 or 1-7 so each client can update their opponents cursor placement)
- Mark placed (another int between 0-6 or 1-7).
I need to decide where to keep track of whose turn it is. Right now the clients keep track of it, but I don't think it would be hard to modify the client to circumvent 
this and just put a bunch of marks. Probably I should do some arbitration on the server side somehow. Anyway for now I will just keep it simple and let clients keep track of it.
- Server will assign whos turn it is to begin. We could make it so that it is the "challenged" player that will start (meaning the second one to enter the room)


## Run it locally
```{bash}
go run .
```
