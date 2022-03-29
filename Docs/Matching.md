What happens if you remove the go-command from the Seek call in the main function?
    Hypothesis: The program won't break, but the messages will always be sent to the same people as the program will always be executed in the same order
    Testing: The hypothesis holds up.
What happens if you switch the declaration wg := new(sync.WaitGroup) to var wg sync.WaitGroup and the parameter wg *sync.WaitGroup to wg sync.WaitGroup?
    Hypothesis: The program should not be able to run as Seek() need to reference the specific waitgroup and not a general one
    Testing: The program won't run, as expected :D
What happens if you remove the buffer on the channel match?
    Hypothesis: The program should crash as mainroutine waits until the waitgroup is finished and, as the channel does not have a buffer, the message is expected to be read instantly, creating a deadlock
What happens if you remove the default-case from the case-statement in the main function?
    Hypothesis: The program should run as intended in this case as an empty default probably does the same thing as not including its