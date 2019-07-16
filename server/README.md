# Code 

## Package main

#### main.go
Creates a router engine using gin
Calls functions to initialize routes and database
Runs thr router

#### route.go

Creates all the rotes


## Package  model

Contains all opeartions on database :
create,update,show,delete


## Handlers
Contains two packages : HomeController and UrlController

### Package HomeController
Displays home page

### Package UrlController
All the routes other than root are handled by the functions here


## Package ShortUrl

short_url  genearation and hashing of long_url
