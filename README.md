# Web-Development using Go's Echo framework

## Steps to run the application


- #### Clone this repository locally
    - Run the command:
        > ```git clone 'https://github.com/kadarsh835/Golang-ECHO.git'```
    - Navigate to this git repository
- #### Create and add details to '.env' file
    - Create a file named '.env'
    - Copy the contents from '.env.dev' and paste them here.
    - Fill in the values to the environment variables in '.env'
- #### Configure $GOPATH
    - Make sure you are in the folder containing the src folder.
    - Run the following command:
        > ```export GOPATH=`pwd` ```
- #### Confiure Postgres Server
    - Start the postgres server and fill in the authentication and connectivity details in the .env file.
- #### Run the Server
    - Build the main package
        > ```go install main```
    - Run the Server
        > ```bin/main```
- #### Access the server in your browser
    - This server runs on the port 8080 by default. So, in your browser, visit
        > http://localhost:8080
    - In case you want to change the port, find the port no in the last lines of main.go file and change it, or, configure it in .env and change easily.