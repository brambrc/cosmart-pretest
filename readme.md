# Library API

This is a simple library API that allows you to fetch books by genre, submit book pick-up schedules, and retrieve book information and schedules. This api made for Cosmart-Pretest 

## Installation
### Debug Mode

1. Clone the repository:

   ```bash
   git clone https://github.com/brambrc/cosmart-pretest.git
   ```
2. Enter the project directory :
    ```bash
    cd cosmart-pretest
    ```
3. Run the program by using this command :
    ```go
    go run main.go
    ```
    Now the program will run at your localhost:8080

### Build Mode
#### Windows :

Just run this code below :

   ```bash
   go build
   ```
Then file `cosmart-pretest.exe` will be generated, after that just run `/.cosmart-pretest.exe` and the server will start on `http:/localhost:8080`

#### Linux :
Just run this code below :

   ```bash
   go build
   ```
Then file `cosmart-pretest` will be generated, after that just run `/.cosmart-pretest` and the server will start on `http:/localhost:8080`



## Usage

* ### GET /books
    This API is for fetch a list of books by genre:

    Api Url:
    ```url
        http://localhost:8080/books
    ```

    Param Query:
    ```query
        genre = love
    ```

    Method :
    ```bash
        GET
    ```

    Curl Example :
    ```bash
        curl -X GET "http://localhost:8080/books?genre=love"
    ```


    Expected Result :
    ```json
    {
        "title": "Wuthering Heights",
        "author": "Emily Brontë",
        "edition_number": 1931
    },
    {
        "title": "The Great Gatsby",
        "author": "F. Scott Fitzgerald",
        "edition_number": 1174
    },
    {
        "title": "Romeo and Juliet",
        "author": "William Shakespeare",
        "edition_number": 971
    }
    ....
    ````

* ### POST /schedules

    This API is for submit a book pick-up schedule :
     Api Url:
    ```url
        http://localhost:8080/schedule
    ```

    Request Body:
    ```json
    {
        "book": {
                    "title": "Book Title", 
                    "author": "Book Author", 
                    "edition_number": 1
                }, 
        "pickup_time": "2023-06-23T10:00:00Z"
    }
    ```

    Method :
    ```bash
        POST
    ```

    Curl Example :
    ```bash
        curl -X POST -H "Content-Type: application/json" -d 
        '{
            "book": {
                "title": "Book Title", 
                "author": "Book Author", 
                "edition_number": 1
                }, 
            "pickup_time": "2023-06-23T10:00:00Z"
        }'

        "http://localhost:8080/schedule"
    ```

    Expected Result:
    ```json
    {
        "message": "Schedule Book Pickup Submitted Successfully!"
    }
    ```


* ### GET /scheduled-books
    This API is for submit a book pick-up schedule :
    
     Api Url:
    ```url
        http://localhost:8080/scheduled-books
    ```
    Method :
    ```bash
        GET
    ```

    Curl Example :
    ```bash
        curl -X GET "http://localhost:8080/scheduled-books"
    ```


    Expected Result :
    ```json
     {
        "book": {
            "title": "The Dialogues of Plato",
            "author": "Plato",
            "edition_number": 247
        },
        "pickup_time": "2023-06-23T10:00:00Z"
    }
    ....
    ````

## Unit Testing
I also prepare several testcase using golang `"github.com/stretchr/testify/assert"` in `library_test.go` file, just simply run :
````bash
go test
````
in the main project folder, and the test case will running.

## Contributing
Contributions are welcome! If you find any issues or have suggestions for improvements, please open an issue or submit a pull request.

## License
This project is licensed under the MIT License.