# home-recipes
Because being a developer doesn't mean you have to eat bad!

## Aim of the project
This little project aims to create a simple **Rest API** that can be a middleware between an *storage*, where the actual recipes will be stored, and a *front-end* that will get its information to print on the user's screen the different recipes.  

It is a simple app that doesn't focus on anything else rather than learning go, the *net/http* and the *gorilla mux* package.  

Also it will be used as a sample app for creating a *docker image* out of it, as well as *CI/CD pipelines* probably using *Jenkins*.  

## Running home-recipes in a Docker container
In order to run this app in a docker container, please follow the steps below:
  - Create an image out of the Dockerfile
    ~~~
    docker build -t home-recipes .
    ~~~
  - Run the container without persistent storage (DEV)
    ~~~
    docker run -d -p 8080:8080 \
           -e "RECIPES_PORT=8080" \
           -e "RECIPES_SAMPLE_DATA=yes" \
           --name home-recipes home-recipes
    ~~~
    - Run the container with persistent storage (PROD)
    ~~~
    docker run -d -p 8080:8080 \
           -e "RECIPES_PORT=8080" \
           -e "RECIPES_FILEPATH=/go/home-recipes/storage/recipes.json" \
           -v /home/mikeletux/storage-test:/go/home-recipes/storage \
           --name home-recipes home-recipes
    ~~~

    **Environment variables**  
      - *RECIPES_PORT*: Port where the RestAPI is going to listen.
      - *RECIPES_SAMPLE_DATA*: Either if you want to load some sample data (yes) or not (empty string).
      - *RECIPES_FILEPATH*: Path of the JSON file that will store the recipes.
  
/Miguel Sama 2021