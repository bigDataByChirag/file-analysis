Gin Server with Structured Logging
==================================

This project is a Gin-based web server that provides endpoints for file and directory operations. It includes structured logging using slog and environment variable management with godotenv. The server can find the largest file in a directory and calculate the total size of a directory.

Prerequisites
-------------

Before running this application, ensure you have the following installed:

1.  [Go](https://golang.org/doc/install) (version 1.16+)
    
2.  [Git](https://git-scm.com/book/en/v2/Getting-Started-Installing-Git)
    
3.  [Go Modules](https://blog.golang.org/using-go-modules)
    

Installation
------------

1.  git clone https://github.com/yourusername/yourrepository.gitcd yourrepository
    
2.  go mod tidy
    
   

Running the Application
-----------------------

To run the application, follow these steps:

1.  go run main.go
    
2.  The server will start and listen on the port specified in the .env file.
    

Endpoints
---------

### GET /file/

*   **Description:** Retrieves the largest file in the specified directory.
    
*   **Parameters:**
    
    *   name (path parameter): The name of the directory to search in.
        
*   bashCopy codecurl http://localhost:8080/file/example\_directory
    
*   jsonCopy code{ "largest\_file": "example.txt", "size": 1024}
    

### GET /dir/

*   **Description:** Calculates the total size of the specified directory.
    
*   **Parameters:**
    
    *   name (path parameter): The name of the directory to calculate the size of.
        
*   bashCopy codecurl http://localhost:8080/dir/example\_directory
    
*   jsonCopy code{ "total\_size": 20480}
    

Logging
-------

Structured logging is implemented using slog. Logs are output in JSON format and include the source file and line number of the log message. This helps in better tracing and debugging.

Middleware
----------

The application uses a custom Gin middleware for logging requests. The middleware is defined in the mid package and is used to log each incoming request to the server.

License
-------

This project is licensed under the MIT License - see the LICENSE file for details.



Acknowledgements
----------------

*   [Gin](https://github.com/gin-gonic/gin)
    
*   slog
    
*   [godotenv](https://github.com/joho/godotenv)
    
