FROM node:latest

RUN mkdir -p /resources

#RUN npm install -g @angular/cli

COPY ./ui-ng/ /resources/
COPY entrypoint.sh /

WORKDIR /resources
#RUN npm install
RUN chmod u+x /entrypoint.sh

ENTRYPOINT ["/entrypoint.sh"]
