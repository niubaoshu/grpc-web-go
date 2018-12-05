FROM node

WORKDIR /helloworld

COPY ./package.json ./package.json
RUN npm install

COPY ./helloworld/helloworld.proto ./helloworld/helloworld.proto
COPY ./server.js ./server.js

CMD node ./server.js
