# playing-go-personal-finance

## Table of contents
-   [Requirement](#requirement)
-   [How to configure on your local machine](#how-to-configure-on-your-local-machine)
-   [How to run migration](#how-to-run-migration)
-   [How to run on your local machine](#how-to-run-on-your-local-machine)
-   [Api Documentation](#api-documentation)


## Requirement

1. Go 1.18 or above.
2. MySQL.


## How to configure on your local machine

1. Clone this repostory to your local.

    ```bash
    $ git clone https://github.com/pembajak/playing-go-personal-finance.git
    ```
2. Change working directory to `playing-go-personal-finance` folder.

    ```bash
    $ cd playing-go-personal-finance
    ```

3. Install dependencies

    ```bash
    $ go get .
    ```

4. Create configuration files.

    ```bash
    $ cp env.toml.example env.toml
    ```   

5. Edit configuration values in `env.toml` according to your setting.


## How to run migration

This migration can do these actions:

1. Migration up

    This command will migrate the database to the most recent version available. Migration files can be seen in this folder `migrations/sql`.

    ```bash
    $ go run main.go migrate:up
    ```

2. Migration down

    This command will undo/rollback database migration.

    ```bash
    $ go run main.go migrate:down
    ```

## How to run on your local machine

1. Running the system.
    ```bash
    $ go run main.go
    ```    

## Api Documentation

[api documentation here](https://kirsb.stoplight.io/docs/personal-finance)