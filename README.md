# Getting_started_with_Golang
A few hello world and toy Golang examples.

### Go Installation

> cd Downloads

> wget https://go.dev/dl/go1.22.1.linux-amd64.tar.gz

> sudo rm -rf /usr/local/go && sudo tar -C /usr/local -xzf go1.22.1.linux-amd64.tar.gz

> echo \ >> $HOME/.profile

> echo "export PATH=$PATH:/usr/local/go/bin" >> $HOME/.profile

> source $HOME/.profile

>  go version

### Useful commands

**hellog go.mod**

> go mod init example/hello

**Running an app**

> go run .

**Compiling an app**

> go build

**Help**

> go help


**Installing requirements**

> go mod tidy
