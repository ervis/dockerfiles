FROM ruby

ENV NODE_URL="https://nodejs.org/dist/v7.9.0/node-v7.9.0-linux-x64.tar.xz"
ENV PATH=$PATH:/opt/node/bin

ENV GIT_REPO https://github.com/atech/postal

RUN mkdir /opt/node && \
	curl -Sso /tmp/node.tar.gz "${NODE_URL}" 2>/dev/null && \
	tar -xvf /tmp/node.tar.gz -C /opt/node --strip=1

RUN gem install bundler && \
	gem install procodile && \
	git clone "${GIT_REPO}" /opt/postal/app && \
	ln -s /opt/postal/app/bin/postal /usr/bin/postal

RUN	postal bundle /opt/postal/app/vendor/bundle && \
	postal initialize-config

CMD ["/bin/sh"]
