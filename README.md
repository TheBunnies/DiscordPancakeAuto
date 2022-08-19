# SELF BOTS ARE PROHIBITED ON DISCORD
Please note that you can get banned from discord for using this software since self-bots are against Discord TOS.
This software was made for educational purposes ONLY!

# About 
Pancake bot automation for discord. Fishing, trivia, work, selling items.
Bot handles everything itself, you don't have to buy a fishing rod everytime it breaks or be worried about burnt panakces after executing the **.work** command.

## Cron Jobs
.work -> every 5 minutes 6 seconds

.fish -> every 1 minute 4 seconds

.trivia hard -> every 10 minutes 1 second

**Selling Items** -> every 8 minutes 2 seconds

## Resolving multiple account conflicts
You may run multiple instances of the bot
However if you launch every instance on the same time you might run into issues, bots can conflict with each other and the only solution to that is to run them with a small delay of approximately 3 seconds each. 

## Configuration of your .env file
You must supply your **user token** and **channel id** (a channel the bot will type in) to the file (both parameters are required)

## Running bot on your Windows PC
1. Install GO runtime from the [official website](https://go.dev/).
2. Open **.env** and pass down your user token.
3. Run `go run main.go` or build `go build .`


## Linux and Docker support
1. Install [Docker](https://www.docker.com/) on your main OS.
2. Build and run an image `docker-compose up -d --build`
