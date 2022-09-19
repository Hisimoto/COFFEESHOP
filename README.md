### COFFEESHOP

#Task for Go interview

##DB

Postgres.

#Create Database

createdb restapi_dev

#For data
Use the script DB_init.sql

##Endpoint
There is few endpoints that was for "personal development". 
Relative to task is only 1 endpoint(how it was required): 

curl --request POST \
  --url http://localhost:8080/buycoffe \
  --header 'CoffeeType: 1' \
  --header 'ID: 1'

#CoffeeType:
1- espresso
2-americano
3-capucino

#membership
1 - Basic
2 - Coffeelover
3 - Espresso Maniac
