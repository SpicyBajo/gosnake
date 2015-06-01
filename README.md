# gosnake
Snake written in Go (built from gotetris)

Dependencies

https://github.com/nsf/termbox-go

# Setup the Go workspace

mkdir -p gosnake/src/github.com/adamturner92
cd gosnake
(cd src/github.com/adamturner92 &&
  git clone https://github.com/adamturner92/gosnake.git)

# Do this each time to work on the code from the top-level gotetris directory.
export GOPATH=`pwd`
export PATH=$PATH:$GOPATH/bin

# Install dependencies.
go get -u github.com/nsf/termbox-go

# Build:

go install github.com/adamturner92/gosnake

#Execute

gosnake

# Credits:
Based off of gotetris
https://github.com/jjinux/gotetris