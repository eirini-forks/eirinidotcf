# EIRINI-DOT-CF - The Eirini Web Page 

This repository contains the resources for `https://eirini.cf`. 

## Frontend

1. Set the `api_url` in `web/src/config.json` to `http://api.<your-cf-domain>`
1. Build the static content using `yarn`:

   ```command
   $ pushd web 
   $ yarn run build
   ```

1. Run `cf push`
1. The frontend app should be avialable now at `http://feedelphia.<your-cf-domain>`

   At this stage you cannot interact with the UI at all. We will need to setup the backend first.

1. Return to the root directory of `feedelphia`: `$ popd`


## API

1. Navigate to the `api/` directory: `$ pushd api`
1. Run `cf push`
1. Wait for the app to get deployed
1. Navigate back to the root directory of `eirinidotcf`: `$ popd`

## Deploy the database

1. Navigate to the `db/` directory `$ pushd db`
1. Run `./helm-install-db.sh`
1. Export the `MYSQL_ROOT_PASSWORD` to an environment variable:

   ```command
   $ MYSQL_ROOT_PASSWORD=$(kubectl get secret --namespace feed-db feed-mysql -o jsonpath="{.data.mysql-root-password}" | base64 --decode; echo)
   ```

1. Expose the database deployment using NodePort:

   ```command
   $ kubectl expose deployment feed-db --type=NodePort
   ```

1. Get the node port: `$ kubectl get svc feed-db -o yaml | grep nodePort`
1. Navigate back to the root directory of `feedelphia`: `$ popd`

## Bind the api to the MySql Database

1. Navigate back to the `api`: `$ pushd api`
1. Edit the `db-cups.json` and provide the `MYSQL_ROOT_PASSWORD` value to the `password` property
1. Set the `db_address` to `<node-ip>:<node-port>`
1. Create a User-Provided-Service instance:

   ```command
   $ cf cups feed-db -p db-cups.json
   ```

1. Bind the `api` app to the `feed-db` service instance:

   ```command
   $ cf bind-service api feed-db
   ```

1. Restage the `api` app: `$ cf restage api`

