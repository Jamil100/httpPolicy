# httpPolicy
HTTP Policy processing Service - This service shall expose an endpoint that accepts http-requests that contain a list of rules inside the body and it should return a policy.

we have an endpoint name 
http://localhost:8080/http-policy?filepath=/Users/apple


it should contain the body with the http policy rules.

the request also has filePath in the params, which is used to store the output json.txt file 


the curl request is attachet here

curl --location --request POST 'http://localhost:8080/http-policy?filepath=/Users/apple' \
--header 'Content-Type: application/json' \
--data-raw '{
 "rules": [
 {
 "id": "rule-1",
 "head": "default allow = false"
 },
 {
 "id": "rule-2",
 "head": "allow",
 "body": "method == \"GET\"; data.roles[\"dev\"][_] == input.user",
 "requires": [
 "rule-3",
 "rule-4"
 ]
 },
 {
 "id": "rule-3",
 "head": "allow",
 "body": "input.user == \"alice\"",
 "requires": [
 "rule-1"
 ]
 },
 {
 "id": "rule-4",
 "head": "allow",
 "body": "input.user == \"bob\"; method == \"GET\"",
 "requires": [
 "rule-3",
 "rule-1"
 ]

 }
 ]
}'