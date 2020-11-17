#!/bin/bash
echo building domainfinder...
go build -o domainfinder

echo building synonyms...
cd ../synonyms
go build -o ../domainfinder/lib/synonyms synonyms.go

echo building available...
cd ../available
go build -o ../domainfinder/lib/available available.go

echo building sprinkle...
cd ../sprinkle
go build -o ../domainfinder/lib/sprinkle sprinkle.go

echo building coolify...
cd ../coolify
go build -o ../domainfinder/lib/coolify coolify.go

echo building domainify...
cd ../domainify
go build -o ../domainfinder/lib/domainify domainify.go

echo done

