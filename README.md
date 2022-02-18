# Contact Data app

## Pre-requisites

Database is expected to be named contacts. The database directory contains the DDL for the two tables, contacts and phone_numbers. I created the database without a user or password, so I haven't catered for that in the database connection string.

## General Layout

There are 5 directories included:
- contact_web: this is a small react form
- database: contains the table creation SQL statements
- models: this contains the code for the two models that have been created, contact and phone_number
- persistance: this contains the code that saves the models into the postgresql database
- web_service: the web service that takes a contact data json and saves it to the database

models, persistance and web_service are all go code

## Running the code - web_service

In the web_service directory, running `go run main.go` will run the web server on port 8080. Or running `go build` will produce a `web_service` binary to run

Data can be added via curl scripts

```
curl -X POST http://localhost:8080/contacts -H "Content-Type: application/json" -d '{"full_name": "Alex Bell", "email": "alex@bell-labs.com", "phone_numbers":["03 8578 6688", "1800728069"]}' 

contacts=# SELECT * FROM contacts;
 id | full_name |       email
----+-----------+--------------------
 26 | Alex Bell | alex@bell-labs.com
(1 row)

contacts=# SELECT * FROM phone_numbers;
 id | contact_id | phone_number
----+------------+---------------
 37 |         26 | +61385786688
 38 |         26 | +611800728069
(2 rows)

curl -X POST http://localhost:8080/contacts -H "Content-Type: application/json" -d '{"full_name": "fredrik IDESTAM", "phone_numbers": ["+6139888998"]}' 

contacts=# SELECT * FROM contacts;
 id |    full_name    | email
----+-----------------+-------
 28 | fredrik IDESTAM |
(1 row)

contacts=# SELECT * FROM phone_numbers;
 id | contact_id | phone_number
----+------------+--------------
 41 |         28 | +6139888998
(1 row)

curl -X POST http://localhost:8080/contacts -H "Content-Type: application/json" -d '{"full_name": "radia perlman", "email": "rperl001@mit.edu", "phone_numbers": ["(03) 9333 7119", "0488445688", "+61488224568"]}' 

contacts=# SELECT * FROM contacts;
 id |   full_name   |      email
----+---------------+------------------
 29 | radia perlman | rperl001@mit.edu
(1 row)

contacts=# SELECT * FROM phone_numbers;
 id | contact_id | phone_number
----+------------+--------------
 42 |         29 | +61393337119
 43 |         29 | +61488445688
 44 |         29 | +61488224568
(3 rows)
```

## Running the code - contact_web

Contact web is a small form to enable a user to enter that data in a web browser. It can be started in the contact_web directory with the command `npm start`

## Tests

There are some tests in the persistance and models modules that can be run with `go test .`

## General thoughts

I wanted to ensure some separation here, some of which I was able to do and some I wasn't.

Each of the three modules are implemented as go modules, that way if a cli application was required to use the same models and persistance, then it could be written using the same modules.

## Database

The two tables, contacts and phone_numbers, are pretty simple and they 

There is no foreign key constraint described, but I have added an index to ensure selecting phone numbers by contact_id is quick

### Persistance

This module handles the mapping of the columns and ensures that a row is inserted into the appropriate table. The insert statement returns the id of the row just inserted. In moving forward, the handling of the SELECT, UPDATE and DELETE statements would also end up in this module as well.

Currently the connection string is hardcoded with values. I'd have those values coming in via ENV variables.
## Models

The models are broken up in the following way:
- model.go: this defines the struct/field used to define the model's attributes, as well as any functions that would be applied to that individual model. For instance, the PhoneNumber model defines the Format() function that will format the phone number in the E.164 format. The Format() function in this instance does do some light validation, but this is where validation rules would go
- model_manager.go: this class would define the more CRUD operations. It would call any validate type methods on the actual model before saving. These functions directly call the functions in the persistance module

With the model_manager.go functions, I think I need to look at how namespacing work in go lang. I'd much rather have a convention of `contact_manager.Save(contact)` rather than `SaveContact(contact)` which I currently have.

## Web_Service

The webservice module uses the gin library for setting up a web server. There is only one end point, POST /contacts for saving contact data. This file is pretty thin, which is by design as it leverages the patterns for the above to keep it that way.

## Contact_Web

This is a small react app with a form. I'm not a react developer, so this was pretty much cobbled together using online resources

## Other things that are missing

I haven't done the web page that would show the list of contacts with their phone numbers.

In order to implement that, I would:
- Update the peristance/postgresql.go to include `SELECT` statements, so a Get function. This would take the output and return the result
- Update the models/*_manager.go files to take the row data and return an array of the appropriate models
- Update the web_service/main.go file to include a GET /contacts end point
- New web form that would show the result of the GET endpoint
