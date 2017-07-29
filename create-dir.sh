mkdir -p /srv/docker/{gitlab, jenkins, postgresql}


echo "127.0.1.1 gitlab.local" >> /etc/hosts
echo "127.0.1.1 jenkins.local" >> /etc/hosts
echo "127.0.1.1 mattermost.local" >> /etc/hosts
