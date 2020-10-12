Image cache

For run this project you must to have a mongo db running in the port: 27017 Please feel free to create an .env file and edit the properties that you want.


Please for run the project use: 

````
go build agileengine/imagecache
./imagecache
````

Feel free to run the project with the docker compose, 
in that case, please use the .env.example.docker as template


in case that you want to use the docker compose 
please run this commands

``
mv .env.example.docker .env
docker-compose up --build
``
