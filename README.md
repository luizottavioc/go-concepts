# go-concepts
This repository houses a collection of Go projects and code snippets that I've developed for educational purposes.

## Services
<a href="./services/">Folder</a> with services that run complete flows.
- <a href="./services/auto-import-users/main.go">auto-import-users</a>: Service that reads, from time to time, an API that returns random users in a random quantity. After reading user data, it is incremented to a .json file in the ```./files``` folder. With this data saved in the file, after a specific number of increments in the file, it is read, deleted and saved in a MongoDB database (accessible through <a href="./docker-compose.yaml">docker-compose.yaml</a> from the root). Reading and writing to the file is controlled using an auxiliary variable that blocks writing during reading so that deletion does not result in data loss.
  
## MongoDB
<a href="./mongodb/">Folder</a> with scripts for manipulating MongoDB. Is accessible by <a href="./docker-compose.yaml">docker-compose.yaml</a> from the project root
- <a href="./mongodb/hello-world/main.go">hello-world</a>: Database connection test that creates a ```"hello-world"``` collection (if it does not exist), cleans the documents in this collection and adds a single ```{"hello":"world"}```.
- <a href="./mongodb/users/main.go">users</a>: Script that adds users provided from an API every 3 records at each run.

## Goroutines
<a href="./goroutines/">Folder</a> with files containing tests for using goroutines and communication channels.
- <a href="./goroutines/channels.go">channels.go</a>: Execution of two goroutines that share common channels. One goroutine prints numbers and another prints letters. They are executed together and only after the registration of the completion channels is the flow completed.
- <a href="./goroutines/routines.go">routines.go</a>: Basic goroutine execution tests.

## Http
<a href="./http/">Folder</a> with files for consuming external APIs via HTTP.
  - <a href="./http/random-data-api.go">random-data-api.go</a>: Reads data from a random API, which can return data from a user, an address or a appliance.
- <a href="./http/cat-api.go">cat-api.go</a>: Consumption of an API that returns random cat species data.

## OS
<a href="./os/">Folder</a> with scripts for file manipulation.
- <a href="./os/incrementing-json.go">incrementing-json.go</a>: Script that increments json results from a cat API into a common file.
- <a href="./os/incrementing-file.go">incrementing-file.go</a>: Script that increments lines in a text file.