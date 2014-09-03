#Async Commander

This is a simple command-line application that allows you to manage input and output to processes running asynchronously.  Each process sends its output to STDOUT (unless it is muted).  This allows you to see processes interacting with one another synchronously.  You can send input to a specific process, see its output as well as those of any other processes it affects.  The goal is to allow greater visibility for service-oriented architectures run locally for development.

I created this application as a way to learn and explore Go.  It may not be entirely practical, but it's definitely fun, since you can look sling commands and run 10 different asynchronous processes like a hacker from the 80's.

Features to implement:

###Features 
 - Spawn different processes asynchronously (done)
 - Send input to a specific process (done)
 - Switch default process to send input to
 - Send signal to a specific process
 - Mute/unmute output from different processes
 - dotfile to configure aliases for commands
 - Color-code/flag output to clearly see which process it comes from
