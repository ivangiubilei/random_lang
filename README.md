# Random Lang
This script will pull the next programming language to learn 

## Compiling
```
go build rlang.go
```
## Exporting to path
In your terminal configuration, add this:
```
export PATH="<path_to_rlang>:$PATH"
```

## Usage
To generate a language:
```
rlang 
```
Other options are available 
```
rlang --help
```

## Examples
```
rlang --classic=false --functional=true --array=true
```
