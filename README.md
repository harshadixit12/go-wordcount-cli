# go-wordcount-cli
A command line word count tool built in go, as a learning exercise

This was picked from https://codingchallenges.fyi/challenges/challenge-wc/

## How to use
Run the command `go-wordcount-cli /filePath/fileName.txt`, this should print the number of bytes, chars, lines, and words respectively.  
If you want only some of those, use flags while running the command.

### Supported flags
`-l` or `--lines`: For number of lines  
`-m` or `--chars`: For number of chars  
`-w` or `--words`: For number of words  
`-c` or `--bytes`: For number of bytes  

## Implementation details
The current implementation makes use of cobra module to build a cli. There is only one root command, and it supports 4 flags - lines, words, chars and bytes. The implementation is not optimised for performance, rather ease of building. There is plenty room for performance optimisation in terms of counting words, lines, etc.  
