#!/bin/bash

echo "Get employee by first name - Expect 200"
curl -i -H "Content-Type: application/json" -X GET http://localhost:8080/employee?firstName=Christian

echo "Create Project with given json data - Expect 200"
curl -i -H "Content-Type: application/json" -X POST http://localhost:8080/project -d '{"name": "Go", "client": "Google", "startDate":"2019-10-03"}'

echo "Delete project -- Expect random error type depending on id"
curl -i -H "Content-Type: application/json" -X DELETE http://localhost:8080/project/9
