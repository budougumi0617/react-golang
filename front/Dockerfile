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
RUN yarn install --force --silent
RUN yarn post-install
RUN mkdir -p node_modules/node-sass/vendor/linux-x64-57
RUN curl -L https://github.com/sass/node-sass/releases/download/v4.8.3/linux-x64-57_binding.node -o node_modules/node-sass/vendor/linux-x64-57/binding.node
# RUN yarn global add react-scripts@1.1.1

# start app
CMD ["yarn", "run", "start"]
