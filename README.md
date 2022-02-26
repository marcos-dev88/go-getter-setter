# go-getter-setter

## About:
 * This is an open source project to generates both getters and setters to files that you can define for each language;
 * Today it supports PHP only, but, will support other languages in the future.
   
## Language support:
 > PHP 7

<br>

### Requirements:
```shell script
GNU Make v4.0 or later;
docker v20.10.7 or later;
docker-compose v1.25 or later.
```
### Instructions:
1. Define your own configuration inside the file `genconf.json`;
2. Ex:
 
```shell script
{
    "files": [
        {
            "path": "./testFiles",   
            "visibility": "private", 
            "functions": "all"       
        }
    ]
}
```


    path -> define the path of your files;
    
    visibility -> define visibility of functions, like: private | protected | public;
    
    functions -> define wich kind of function will generate, like: 
        all -> get and set;
        get -> only getters;
        set -> only setters;

## Run application/instalation:
### Run by file:
```shell script
make run-file
```

### Run by cli:
```shell script
make run run-cli path="your/path/here" fn="get (you could use: get|set|all )"
```