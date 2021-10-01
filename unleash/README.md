**Run from Docker**

 1. Create a network by running 
 ```bash
 docker network create unleash
  ```
 2. Start a postgres database:
```bash
    docker run -e POSTGRES_PASSWORD=some_password \
    -e POSTGRES_USER=unleash_user -e POSTGRES_DB=unleash \
    --network unleash --name postgres postgres
   ```

 3. Start Unleash via docker:
```bash
    docker run -p 4242:4242 \
    -e DATABASE_HOST=postgres -e DATABASE_NAME=unleash \
    -e DATABASE_USERNAME=unleash_user -e DATABASE_PASSWORD=some_password \
    -e DATABASE_SSL=false \
    --network unleash unleashorg/unleash-server
   ```
The first time Unleash starts it will create a default user which you can use to sign-in to your Unleash instance and add more users with:

 - username: admin
 - password: unleash4all

 4. After login go to Advanced > Api Access . Click on Add new Api Key and then copy secret and replace it with Authorization Header in unleash.go
 