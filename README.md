# dockerfiles
Dockerfiles that I use. If you are not ervis then you don't care about this repository.


Some commands:

    $ git clone https://github.com/ervis/Dockerfiles
    $ cd Dockerfiles

    $ mkdir -p /srv/docker/{gitlab, jenkins, postgresql}

    $ echo "127.0.1.1 gitlab.local" >> /etc/hosts
    $ echo "127.0.1.1 jenkins.local" >> /etc/hosts
    $ echo "127.0.1.1 mattermost.local" >> /etc/hosts
    $ echo "127.0.1.1 redmine.local" >> /etc/hosts

    $ docker-compose up -d


Visit http://jenkins.local/ to access jenkins.
