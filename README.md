# BookManager

## Manual
In order to start the database just run docker-compose.yaml file with command: ```docker-compose up``` </br>
After that you can start the application with command (navigate to the main.go file):
```go run .``` </br></br>


## Step 1: Describe the expected user experience
If you want to use CLI just enter ```go run .``` and the wanted CLI command **without ```bookmanager-cli```** . </br> </br>
The other way is to enter ```go build -o bookmanager-cli``` and then ```./bookmanager-cli``` and the wanted CLI command. 

## CLI Commands
### Description: Create a new book. </br>
Options: </br>
--title: Title of the book (required) </br>
--author: Author of the book </br>
--publishedDate: Published date of the book (format: "YYYY-MM-DDTHH:MM:SSZ") </br>
--edition: Edition of the book </br>
--description: Description of the book </br>
--genre: Genre of the book </br>

Command: ```bookmanager-cli get-all``` </br>
### Description: Get a list of all books. </br>

Options: </br>
--page: Page number </br>
--pageSize: Page size </br>
--author: Author filter </br>
--genre: Genre filter </br>
--startDate: Start date filter </br>
--endDate: End date filter </br>

Command: ```bookmanager-cli update``` </br> 
### Description: Update details of an existing book.

Options: </br>
--id: ID of the book to update (required) </br>
--title: New title of the book </br>
--author: New author of the book </br>
--publishedDate: New published date of the book (format: "YYYY-MM-DDTHH:MM:SSZ") </br>
--edition: New edition of the book </br>
--description: New description of the book </br>
--genre: New genre of the book </br></br>


Command: ``` bookmanager-cli delete ``` </br>
Description: Delete a book.

Options: </br>
--id: ID of the book to delete (required)


### Examples

```
 bookmanager-cli create --title "The Great Gatsby" --author "F. Scott Fitzgerald" --publishedDate "1925-04-10T00:00:00Z" --edition "First Edition" --description "A classic novel about the American Dream" --genre "Fiction"

 bookmanager-cli get-all --page 1 --pageSize 10 --author "F. Scott Fitzgerald" --genre "Fiction" --startDate "1900-01-01" --endDate "2022-01-01"

 bookmanager-cli update --id 1 --title "Updated Title" --author "New Author" --publishedDate "2022-01-01T12:00:00Z" --edition "Second Edition" --description "Updated description" --genre "Non-fiction"

 bookmanager-cli delete --id 1
```

## Step 2: Describe the expected REST API
### Description: Create a new book. </br>
Method: POST </br>
Endpoint: /api/v1/books </br>
Body JSON: </br>
```
{
    "title": "The Great Gatsby",
    "author": "F. Scott Fitzgerald",
    "publishedDate": "1925-04-10T00:00:00Z",
    "edition": "First Edition",
    "description": "A classic novel about the American Dream",
    "genre": "Fiction"
}
```
Responses: </br>
* Status Code: 200 OK
```
{
    "status": "Book created successfully!"
}
```
* Status Code: 400 Bad Request
```
{
    "error": "There was an error while creating the book."
}
```

### Description: Get a list of all books. </br>
Method: GET </br>
Endpoint: /api/v1/books </br>
Query Parameters: 
* page: Page number (pagination) </br>
* pageSize: Page size (pagination) </br>
* author: Author filter </br>
* genre: Genre filter </br>
* startDate: Start date filter </br>
* endDate: End date filter </br>

Responses: </br>
* Status Code: 200 OK
```
{
    "books": [
        {
            "id": 1,
            "title": "The Great Gatsby",
            "author": "F. Scott Fitzgerald",
            "publishedDate": "1925-04-10T00:00:00Z",
            "edition": "First Edition",
            "description": "A classic novel about the American Dream",
            "genre": "Fiction"
        }
    ],
    "page": 1,
    "total": 1
}
```
* Status Code: 400 Bad Request
```
{
    "error": "There was an error while getting the books."
}
```

### Description: Update details of an existing book. </br>
Method: PUT </br>
Endpoint: /api/v1/books/{id} </br>
Body JSON: </br>
```
{
    "title": "The Great Gatsby",
    "author": "F. Scott Fitzgerald",
    "publishedDate": "1925-04-10T00:00:00Z",
    "edition": "First Edition",
    "description": "A classic novel about the American Dream",
    "genre": "Fiction"
}
```
Responses: </br>
* Status Code: 200 OK
```
{
    "status": "Book updated successfully!"
}
```
* Status Code: 400 Bad Request
```
{
    "error": "There was an error while updating the book."
}
OR
{
    "error": "Book not found."
}
```

### Description: Delete a book. </br>
Method: DELETE </br>
Endpoint: /api/v1/books/{id} </br>
Responses: </br>
* Status Code: 200 OK
```
{
    "status": "Book deleted successfully!"
}
```
* Status Code: 400 Bad Request
```
{
    "error": "Book not found!"
}
```

### Description: Create a new book collection. </br>
Method: POST </br>
Endpoint: /api/v1/collections </br>
Body JSON: </br>
```
{
    "name": "My Collection"
}
```
Responses: </br>
* Status Code: 200 OK
```
{
    "status": "Book collection created successfully!"
}
```
* Status Code: 400 Bad Request
```
{
    "error": "There was an error while creating the collection."
}
```

### Description: Get a list of all book collections. </br>
Method: GET </br>
Endpoint: /api/v1/collections </br>
Query Parameters:
* page: Page number (pagination) </br>
* pageSize: Page size (pagination) </br>

Responses: </br>
* Status Code: 200 OK
```
{
    "collections": [
        {
            "id": 1,
            "name": "My Collection"
        }
    ],
    "total": 1,
    "page": 1
}
```

### Description: Get a list of all books in a collection. </br>
Method: GET </br>
Endpoint: /api/v1/collections/{id}/books </br>
Query Parameters:
* page: Page number (pagination) </br>
* pageSize: Page size (pagination) </br>
* author: Author filter </br>
* genre: Genre filter </br>
* startDate: Start date filter </br>
* endDate: End date filter </br>

Responses: </br>
* Status Code: 200 OK
```
{
    "books": [
        {
            "id": 1,
            "title": "The Great Gatsby",
            "author": "F. Scott Fitzgerald",
            "publishedDate": "1925-04-10T00:00:00Z",
            "edition": "First Edition",
            "description": "A classic novel about the American Dream",
            "genre": "Fiction"
        }
    ],
    "page": 1,
    "total": 1
}
```
* Status Code: 400 Bad Request
```
{
    "error": "Book colleciton not found"
}
```

### Description: Add a book to a collection. </br>
Method: POST </br>
Endpoint: /api/v1/collections/{collectionId}/books/{bookId} </br>
Responses: </br>
* Status Code: 200 OK
```
{
    "status": "Book added to collection successfully!"
}
```
* Status Code: 400 Bad Request
```
{
    "error": "Book colleciton not found."
}
OR
{
    "error": "Book not found."
}
```

### Description: Remove a book from a collection. </br>
Method: DELETE </br>
Endpoint: /api/v1/collections/{collectionId}/books/{bookId} </br>
Responses: </br>
* Status Code: 200 OK
```
{
    "status": "Book removed from collection successfully!"
}
```
* Status Code: 400 Bad Request
```
{
    "error": "Book colleciton not found."
}
OR
{
    "error": "Book not found."
}
```

### Description: Delete a book collection. </br>
Method: DELETE </br>
Endpoint: /api/v1/collections/{id} </br>
Responses: </br>
* Status Code: 200 OK
```
{
    "status": "Book collection deleted successfully!"
}
```
* Status Code: 400 Bad Request
```
{
    "error": "Book colleciton not found!"
}
```


## Step 3: Describe the database structure
## Database Structure
### Table 'books'
Columns:
* ID (Primary Key, Auto-increment): ID of the book.
*  Title (String): Title of the book.
*  Author (String): Author of the book.
*  PublishedDate (Timestamp): Published date of the book.
*  Edition (String): Edition of the book.
*  Description (String): Description of the book.
*  Genre (String): Genre of the book.
*  CollectionId (Foreign Key, Nullable): ID of the collection to which the book belongs.

### Table 'book_collections'
Columns:
* ID (Primary Key, Auto-increment): ID of the collection.
* Name (String): Name of the collection.

*Notes*:</br>
The books table has a foreign key relationship with the book_collections table, indicating which collection a book belongs to. </br>

The CollectionId in the books table is nullable, allowing books that are not part of any collection. </br>

Indexes for some fields based on query patterns are added improving query performance when filtering by these fields. </br>

Relation between two tables could also be *ManyToMany* where new table is created to store the relationship between the two tables. </br>

