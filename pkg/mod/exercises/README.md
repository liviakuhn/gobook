# The Go Programming Language

This repository provides the exercise programs
for the book, "The Go Programming Language"; see http://www.gopl.io.

You can install and run the programs with the following commands:

    # Option 1: RUN
	$ go run <x>.go                         # compile source code, link it with libraries,
                                            # and run temporary binary
    # Option 2: INSTALL
	$ export GOPATH=$HOME/gobook            # choose workspace directory
	$ go install <x>.go                     # install into $GOPATH/bin
	$ $GOPATH/bin/<x>                       # run from binary

    # Option 3: BUILD
	$ go build <x>.go                       # install into current directory
	$ ./<x>                                 # run from binary
