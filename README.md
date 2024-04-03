# Getting_started_with_Golang
The purpose of this repository is to show a few Golang toy project examples.

## Meta
 * **Author:** Alaf DO NASCIMENTO SANTOS
 * **License**: [Apache V2](LICENSE)
 * **Year:** 2024

### Installing Golang
If you are using a Linux machine, the following commands would be enough: 

> cd Downloads

> wget https://go.dev/dl/go1.22.1.linux-amd64.tar.gz

> sudo rm -rf /usr/local/go && sudo tar -C /usr/local -xzf go1.22.1.linux-amd64.tar.gz

> echo \ >> $HOME/.profile

> echo "export PATH=$PATH:/usr/local/go/bin" >> $HOME/.profile

> source $HOME/.profile

>  go version

For more information, please take a look at the Golang official website: https://go.dev/doc/install


## How to compile and execute the project?
After installing Go, inside a given project directory, you can use the command ***go build***, which compiles the project.
For running a Go programme, you can use the command line ***go run .***.

### Some useful commands

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


## How to Contribute to the Project
- Any implementation that could lead to a more optimised code for the different methods/functions already designed would be a nice improvement for this project. 

So feel free to fork, change, and pull request 😊
