# Web-Development using Go's Echo framework
## Table of Contents
1. [Useful links to get started](#useful_links)
2. [Steps to configure and start the server](#steps_to_run)

This is the first application that I have developed in Golang and Echo.
<a name="useful_links"></a>
#### Some useful links to get you started with Golang and the Echo framework are:
- This one is a very useful link for the absolute beginners in Go lang. This covers all the basic concepts of Go.
    > https://www.tutorialspoint.com/go/index.htm
- This and a few other articles are concerned with using Echo framework in Go with Postgres are available in this link.
    > https://www.restapiexample.com/golang-tutorial/creating-golang-api-echo-framework-postgresql/
- The project consumes REST APIs. An easy tutorial on this can be found at
    > https://www.restapiexample.com/golang-tutorial/consume-restful-apis-using-echo-golang/
- Unlike Django or other similar frameworks, Echo doesn't have a directory structure of its own. Managing large volumes of code can be really dificult if we don't create packages and maintain a proper directory structure. There are hardly any resources on the Internet explaining this in detail with respect to the Echo framework. However, this Youtube playlist covers you pretty well
    > https://www.youtube.com/playlist?list=PLFmONUGpIk0YwlJMZOo21a9Q1juVrk4YY
<a name="steps_to_run"></a>
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