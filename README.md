# doubanbookapi
Just as the name, the API implemention for "Douban Book".

You can query book information via api which provided by Douban.com 
and save it to local MongoDB.

## Content
* Requirements
* Settings

## Requirements
1. MongoDB
2. Internet

## Settings

### Site Settings
Website listening on:

    Addr: 0.0.0.0
    Port: 8080
    
If you want to change, just modify the file below: 

```
$RPJ_ROOT/configs/AppConfig.go
```

### DB Settings
 
You can change MongoDB service here:

```$go
$RPJ_ROOT/configs/MongoDBConfig.go
```