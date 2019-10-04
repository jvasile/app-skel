# App Skel

This is a skeleton repository that should get you a web app skeleton
repo pretty quickly.  It supports a react front end, a golang+gin or
python+flask backend.  We could add more options if we like.

You get a docker container that includes a postgres DB.  We could add
more if we like.

I have a private branch with some ansible stuff to support all this.
Not ready to push yet.

Python+flask is implemented in a separate repo.  I'll add it hear
soonish.

## Setup

    git clone git@github.com:jvasile/app-skel.git
    git set-url origin [new repo]
    git remote add app-skel git@github.com:jvasile/app-skel.git
    git branch --track app-skel app-skel/master
    git co master
    
    Edit the template file.
    ./generate_skeleton

[Note we haven't implemented the template file or generate_skeleton yet.]

## Docker Deploy

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
