#!/bin/bash

curl --silent --location --request POST 'http://localhost:8001/contact-us/' \
--header 'Content-Type: application/json' \
--data-raw '{
  "name": "test",
  "email": "test@example.com",
  "contact_number": "2",
  "user_type": "coach",
  "message": "TEST"
}'
