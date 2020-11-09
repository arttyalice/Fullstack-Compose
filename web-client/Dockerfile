FROM node:10

# Create app directory
WORKDIR /app

COPY . .
RUN yarn
RUN yarn build

ENV HOST 0.0.0.0
EXPOSE 3000

# start command
CMD [ "yarn", "start" ]