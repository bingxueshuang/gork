# gork
`fork` written in go

# Why
Usually launching a gui in terminal blocks until the application is
closed and writes a lot of logs and (unwanted) messages to stdout.

As a workaround we use `&`, i.e.
```bash
$ google-chrome&
```

The problem with the above trick is that some shells do not allow
child processes to continue the execution once the shell is exited.
Thus, quitting from the terminal **hangs up** the application process.

There is another workaround for the above issue in zsh:
```zsh
$ google-chrome&!
```
This effectively **disowns** the job but the stdout is still being
printed. Thus it is tedious to *spawn-and-forget* any gui application.

# What
`gork` simply takes the command as an argument and launches it in
background and deliberately orphans the child process and terminates.
This effectively removes the need for all above workarounds for
*spawn-and-forget*.

`gork` is not obligated to the child process once it is spawned. Thus,
it does not matter if the child process is running in the background
successfully or if it has terminated for some reason. This means that
in case the child process terminates with non-zero exit code, the exit
code of `gork` is still 0 [it has successfully *spawn-and-forget*-ed].
As an example:

```bash
$ gork false
$ echo $? # output: 0
```

# Installation
## Using go
Run:
```bash
$ go install github.com/bingxueshuang/gork@latest
```

# Bugs
Feel free to create an issue in the github repository
