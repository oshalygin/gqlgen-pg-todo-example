# CHANGELOG

## 1.4.1 (January 05, 2020)

- Update README Dataloader example
  - The example in the readme was a bit outdated
    and its best to just copy the code snippet.
  - Makefile limitation noted.

## 1.4.0 (January 05, 2020)

- Add Todo Mutations and Resolvers(Complete/Delete)
  - Addition of resolvers and mutation schema for todo
    deletion/completion. Ideally the request context carries
    through WHO is initiating the request but building out jwts
    and the whole shebang is not necessary in the scope of this
    example project.

## 1.3.2 (January 05, 2020)

- Update User Dataloader
  - The dataloader was not properly configured to
    account for the fact that the keys array is
    non-sequential which is what the pg query returns.
    Because of this, the actual hydration of a user was
    not valid and would return erroneous data. It is imperative
    that the keys array matches the return value in terms of ids<==>user.

## 1.3.1 (January 02, 2020)

- Update Seed
  - The seed has been updated to try and capture different
    users that belong to a given todo.
  - Small bugfix where new todos were created on every prog
    execution.

## 1.3.0 (January 02, 2020)

- Add Todo Resolvers
  - Addition of resolvers to pull all todos
    or a specific one by id.
  - Sample query provided in README

## 1.2.1 (January 02, 2020)

- Add User Dataloaders
  - The makefile script has been updated to properly
    reference the path of the model.
  - The key type has been redefined as int from string to
    account for the fact that the IDs are auto incrementing
    ids in pg.
  - Middleware created with instructions around extensibility.

## 1.2.0 (December 31, 2019)

- Add User Resolvers
  - Addition of user resolvers that will pull user
    data from a db either by id or a collection.
  - UserCreate has been implemented as well to allow
    creation of new users.

## 1.1.3 (December 31, 2019)

- Add Database Seed
  - Addition of scripts that will seed the database
    with tables and data if that data/tables do
    not already exist against the todos database.

## 1.1.2 (December 31, 2019)

- Add Database Creation Step in Makefile
  - Users can now initialize the the creation of
    the todos database from the makefile if they
    dont have one.

## 1.1.1 (December 31, 2019)

- Update GraphQL Models to Int from ID
  - In the interest of time to not configure and set
    custom types within gqlgen.yaml(will be done later),
    the PK values have been updated to Int from ID which
    GQLGEN defaults to as strings.

## 1.1.0 (December 30, 2019)

- Add Resolvers and Wire Up
  - The actual models themselves have been updated to point
    to the appropriate go model(graphql => go).
  - Resolvers split into their respective entities
  - Basic wireup of the DB with a pass to the root resolver
    which will cascade it down to every subsequent resolver.

## 1.0.3 (December 30, 2019)

- Add Schema Models and Routing
  - This is a first pass at the schema models
    definition and general router layout
    that will coexist in the main call.

## 1.0.2 (December 30, 2019)

- Add initial application structure
  - This is a bare bones application structure
    segregating out the schema, graph, resolvers,
    dataloaders, and other basic skeleton components.

## 1.0.1 (December 30, 2019)

- Add README and Makefile
  - Addition of a README and a Makefile to help users
    get started with the application. This is still a huge
    work in progress but a solid starting point in terms of
    documentation.

## 1.0.0 (December 30, 2019)

- Initial Commit
