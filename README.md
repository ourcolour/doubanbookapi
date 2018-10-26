# doubanbookapi
Just as the name, the API implemention for "Douban Book".

You can query book information via api which provided by Douban.com 
and save it to local MongoDB.

## Content
* Requirements
* Installzation
* Settings

## Requirements
1. MongoDB
2. Internet

## Installation
1. Make sure you have MongoDB installed and own the administrator role.
2. Create one database which named by "smartlibrarian".
3. Restore the Mongo(Structure & Data) by the json format script below.
```
$PRJ_ROOT/install/smartlibrarian.sl_book_new.json
```

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

## How to use

### Query API

You can query book by *Isbn, Author* and even by douban book *Identifier*(id).

- by isbn

    Method 1:
        
        http://localhost:8080/book/isbn/*{isbn}*
        
    Or use the alias method, it's shortter.
        
        http://localhost:8080/book/*{isbn}*
    
- by author

        http://localhost:8080/book/author/*{author}*
    
- by douban idntifier

        http://localhost:8080/book/id/*{id}*


Examples here:

- by isbn
    
        http://localhost:8080/book/9787556820825

- by ahthor

        http://localhost:8080/book/author/斯坦尼斯

- by douban id

        http://localhost:8080/book/id/26952828
