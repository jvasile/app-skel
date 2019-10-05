# App Skel

This is a skeleton repository that should get you a web app skeleton
repo pretty quickly.  It supports a react front end, a golang+gin or
python+flask backend.  We could add more options if we like.

You get a docker container that includes a postgres DB.  We could add
more if we like.

I have a private branch with some ansible stuff to support all this.
Not ready to push yet.

Python+flask is implemented in a separate repo.  I'll add it here
soonish.

## Setup

You'll want to base your repo on this one.  That is, your new repo
should always be a bunch of commits on top of all the commits in this
repo.  That way, when this repo updates, you can pull those changes
in.  You can also offer changes back.

    git clone git@github.com:jvasile/app-skel.git NEWREPO
    cd NEWREPO
    git remote set-url origin [new repo]
    git push
    git remote add app-skel git@github.com:jvasile/app-skel.git
    git fetch app-skel
    git branch --track app-skel app-skel/master
    
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
    
Finally, just run the app directly (again, replace the password with your own)::

    cd app
    export postgres_password=partytimenow
    npm start OR npm run watch
    
Then, point your browser at localhost:9898 or localhost:9897 if you're
using the watch functionality.

## Keeping The Skeleton Up To Date

### Pulling skeleton changes

    git co app-skel
    git fetch
    git rebase app-skel
    git co master
    git rebase app-skel
    [ Resolve conflicts ]
    
### Pushing your skeleton changes upstream

    git push app-skel HEAD:master
    
