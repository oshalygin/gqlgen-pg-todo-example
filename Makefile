home				        = 	$(shell home)
software_version	  =	  $(shell cat VERSION)
version_array		    =	  $(subst ., ,$(software_version))
major				        = 	$(word 1,${version_array})
minor				        = 	$(word 2,${version_array})
patch				        = 	$(word 3,${version_array})
pwd 				        = 	$(shell pwd)

patch:
	- @echo "BUMPING PATCH"
	- @echo "Current Version: $(software_version)"
	- $(eval patch=$(shell echo $$(($(patch)+1))))
	- @echo "New Version: $(major).$(minor).$(patch)"
	- @printf $(major).$(minor).$(patch) > VERSION

minor:
	- @echo "BUMPING MINOR"
	- @echo "Current Version: $(software_version)"
	- $(eval minor=$(shell echo $$(($(minor)+1))))
	- @echo "New Version: $(major).$(minor).0"
	- @printf $(major).$(minor).0 > VERSION

major:
	- @echo "BUMPING MAJOR"
	- @echo "Current Version: $(software_version)"
	- $(eval major=$(shell echo $$(($(major)+1))))
	- @echo "New Version: $(major).0.0"
	- @printf $(major).0.0 > VERSION

# This will kick off the gqlgen codegen
gen:
	- go generate ./...

# Usage here looks like the following
# make dataloader loader=User
# the argument loader is REQUIRED and is CASE SENSITIVE.
dataloader:
	- @if test -z "$(loader)"; then echo "the loader argument is required"; exit 1; fi
	# If you're going to use this script to generate your own loaders, make sure you update the actual path here.
	# Note the project path of `github.com/oshalygin/gqlgen-todo-pg/models`
	# You can read more about dataloaden at github.com/vektah/dataloaden or feel free to raise an issue and I can elaborate further
	# Note the fact that int is designated here for the key type
	- cd graph/generated && go run github.com/vektah/dataloaden $(loader)Loader int *github.com/oshalygin/gqlgen-pg-todo-example/models.$(loader)

build:
	- CGO_ENABLED=0 GOOS=linux go build -a -gcflags='-N -l' -installsuffix cgo -o main .

init:
	- psql -c "CREATE DATABASE todos"

start:
	- go run main.go

lint:
	- go vet
	- golint


.PHONY: build patch minor major start gen dataloader init
