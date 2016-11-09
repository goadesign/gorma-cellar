#! /bin/bash

set -e

goagen --design=github.com/goadesign/gorma-cellar/design app
goagen --design=github.com/goadesign/gorma-cellar/design client
goagen --design=github.com/goadesign/gorma-cellar/design swagger
goagen --design=github.com/goadesign/gorma-cellar/design js
goagen --design=github.com/goadesign/gorma-cellar/design schema
goagen --design=github.com/goadesign/gorma-cellar/design gen --pkg-path=github.com/goadesign/gorma

echo "You can use 'go run gen/main.go' for vendor goagen instead"
