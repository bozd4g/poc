#!/bin/bash

set -x

VERSION="1.0.0"

curl -X PUT \
    http://localhost:9292/pacts/provider/Server/consumer/Client/version/${VERSION} \
    -H "Content-Type: application/json" \
    -d @/Users/furkan.bozdag/Documents/personal/poc/testing/client/internal/todo/pacts/client-server.json