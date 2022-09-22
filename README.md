
# NeOwOfetch

Combination of neofetch and uwufetch written in Go, currently in beta.  

## Requirements  
requires Go 1.19 or above. [Install Go](https://github.com/golang/go)
## Installation

Guide for installing neowofetch on Linux

### Step 1:
Clone the repository and move into the directory

```
git clone https://github.com/exhq/neOwOfetch.git && cd neOwOfetch
```

### Step 2:
Build and install the program using Go
```
go build
go install
```

### Adding to path
Find path for Go programs (default is /home/user/go/)
```
go env GOPATH
```
add gopath/bin to your $path   

example:
```
export PATH=$PATH:/home/user/go/bin
```
