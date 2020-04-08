# EIRINI-DOT-CF - The Eirini Web Page 

This repository contains the resources for `https://eirini.cf`. 

## Deploy the database

1. Navigate to the `db/` directory
1. Run `./helm-install-db.sh`
1. Export the `MYSQL_ROOT_PASSWORD` to an environment variable:
    ```command
    export MYSQL_ROOT_PASSWORD=$(kubectl get secret --namespace pheed-db pheed-mysql -o jsonpath="{.data.mysql-root-password}" | base64 --decode; echo)
    ```
1. Get the node port of the database and export it as `$DB_PORT`:
    ```command
    export DB_PORT=$(kubectl get service -npheed-db pheed-mysql -ojsonpath='{.spec.ports[?(.name=="mysql")].nodePort}')
    ```
1. Get a Node IP and export it as `$DB_IP`

## API

1. Navigate to the `api/` directory
1. Run `cf push --no-start`
1. Create the Custom Service for the Database
    ```command
    cf cups pheed-db -p "$(./db-cups.json.sh)"
    ```
1. Bind the service with the API CF app:
    ```command
    cf bind-service eirinidotcf-api pheed-db
    ```
1. Start the API:

    ```command
    cf start eirinidotcf-api
    ```

## Frontend

1. Navigate to the `/web` directory
1. Set the `api_url` in `src/config.json` to `http://eirinidotcf.<your-cf-domain>`
1. Build the static content using `yarn`:
    ```command
    yarn install
    yarn run build
    ```
1. CF push the web frontend with `$ cf push`

The frontend app should be avialable now at `http://web.<your-cf-domain>`

If you have a domain registered to your cluster node IP and you want to use it for this app, run the following commands:
```command
cf create-domain <org> <domain>
cf map-route web <domain>
```
