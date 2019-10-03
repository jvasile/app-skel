# Twitter Explorer

Mess around wtih Twitter APIs.

You'll want to set postgres_password in your environment before
running.  This should get you going:

     docker-compose build
     docker-compose up

Then, point your browser at localhost:9898

## Dev

In dev, you probably want to run the app locally and not in a
container.  That way you get the nice watchify functionality.

First, make sure npm has installed your dependencies:

    cd app; npm install; cd ..

Then, bring up a postgres instance (replace the password with your own):

    export postgres_password=partytimenow
    docker-compose up db
    
Finally, just run the app directly:

    cd app
    npm start
    
Then, point your browser at localhost:9898
