#!/bin/bash
export RES=`curl --silent -X POST -H 'Content-Type: application/json' -d '{"url":"https://example.org"}' http://localhost:8080/post`
