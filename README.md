# doubanbookapi
Just as the name, the API implemention for "Douban Book".

You can query book information via api which provided by Douban.com 
and save it to local MongoDB.

Now a days, we finished the feature search engine supported via ElasticSearch.

## Version

* 0.0.3

## Content
* Requirements
* Installation
* Settings

## Requirements
1. Internet connection
2. MongoDB v4.0.1 (Single-node or Cluster)
3. *[Optional]* ElasticSearch v6.4.2 (with IK plugin)

## Installation
1. Make sure you have MongoDB installed and own the administrator role.
2. Create one database which named by "smartlibrarian".
3. Restore the Mongo(Structure & Data) by the json format script below.
    ```
    $PRJ_ROOT/install/mongodb/smartlibrarian.sl_book_new.json
    ```
4. Restore the ElasticSearch Index & Type mapping.
    ```
    $PRJ_ROOT/install/elasticsearch/sl_book_new.book.txt
    ```

## Settings

### Site Settings
Website listening on:

    Addr: 0.0.0.0
    Port: 8080
    
If you want to change, just modify the file below: 

```
<RPJ_ROOT>/configs/AppConfig.go
```

### DB Settings
 
You can change MongoDB service here:

```$go
<RPJ_ROOT>/configs/MongoDBConfig.go
```

### ES Settings
 
You can change ElasticSearch service here:

```$go
<RPJ_ROOT>/configs/EsConfig.go
```

## How to use

### Query API 

> Query Mongo directly and no search engine accelerate.

You can query book by *Isbn, Author* and even by Douban book *Identifier*(id).

- by isbn
	
	+ Method 1:
		
			curl -XGET http://localhost:8080/v1/book/isbn/{ISBN}
        
   + Or use the alias method, it's pretty shorter.
			
			curl -XGET http://localhost:8080/v1/book/{ISBN}    
- by author

        curl -XGET http://localhost:8080/v1/book/author/{AUTHOR}
    
- by douban idntifier

        curl -XGET http://localhost:8080/v1/book/id/{DOUBAN_ID}


Examples here:

- by isbn
    
        curl -XGET http://localhost:8080/v1/book/9787556820825

- by ahthor

        curl -XGET http://localhost:8080/v1/book/author/斯坦尼斯

- by douban id

        curl -XGET http://localhost:8080/v1/book/id/26952828
    
### CIP API 

> Douban book API could not provide CIP information.
> 
> But we can fetch books' CIP information from [opac.calis.edu.cn](CALIS)
> 
> **P.S. Special Thanks to CALIS.**

- Query by `ISBN`

	curl -XGET http://localhost:8080/v1/book/cip/{ISBN}
	
- Batch update local book CIPs
	
	> Query local existed books which do not contain cip field and fetch CIP by ISBN.

	curl -XPOST http://localhost:8080/v1/book/cip

### Search API

> Data accelerated by ElasticSearch.

- by `KEYWORD`

    + Basic query:
    
 		curl -XGET http://localhost:8080/v1/search/?k={KEYWORD}
        
    + Paged query result:
		
		curl -XGET http://localhost:8080/v1/search/?k={KEYWORD}&pageNo=2&pageSize=10
        
    The records in result set would be splited into 10 per page.
    
    Current page is the #*2nd* page.

