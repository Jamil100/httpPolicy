# httpPolicy
HTTP Policy processing Service - This service shall expose an endpoint that accepts http-requests that contain a list of rules inside the body and it should return a policy.

A rule inside the request may depend on other rules and require execution in proper order. The service takes care of sorting the rules to create a proper execution order. A policy is an ordered collection of rules where each rule has a head and an optional body, that get evaluated in order of the list.

Sample Request:
```JSON
{
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
                "rule-3"
            ]
        }
    ]
}
```

Expected Output:
```JSON
[
    {
        "id": "rule-1",
        "head": "default allow = false"
    },
    {
        "id": "rule-3",
        "head": "allow",
        "body": "input.user == \"alice\""
    },
    {
        "id": "rule-4",
        "head": "allow",
        "body": "input.user == \"bob\"; method == \"GET\""
    },
    {
        "id": "rule-2",
        "head": "allow",
        "body": "method == \"GET\"; data.roles[\"dev\"][_] == input.user"
    }
]
```
I used Postman to send the requests. Note: The solution creates a text file of the response and places in the passed file path.

Postman Screenshot:
![image](https://user-images.githubusercontent.com/75333239/180608630-2b2189fb-6423-4bd8-98b3-7a3bed8c866e.png)

Endpoint: http://localhost:8080/http-policy?filepath='your-file-path'


The request should contain the body with the http policy rules. Also, it has filePath in the params, which is used to store the output json.txt file.


Sample curl request:
```Linux
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
```
