# base image
FROM node:8.9.4

# set working directory
RUN mkdir /usr/src/app
WORKDIR /usr/src/app

# add `/usr/src/app/node_modules/.bin` to $PATH
ENV PATH /usr/src/app/node_modules/.bin:$PATH

# install and cache app dependencies
COPY package.json /usr/src/app/package.json
COPY package-lock.json /usr/src/app/package-lock.json
COPY yarn.lock /usr/src/app/yarn.lock
COPY webpack.config.js /usr/src/app/webpack.config.js
RUN npm install yarn@1.5.1
RUN yarn install --silent

# start app
# $ docker run -it -v ${PWD}:/usr/src/app -v /usr/src/app/node_modules -p 8080:8080 --rm sample-app
# ホストのnode_modulesをマウントするとnode-sassなどが失敗するのでコンテナにData Volumeを作って起動する
CMD ["yarn", "run", "start"]
