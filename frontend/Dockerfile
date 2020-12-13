FROM node:15.4.0-alpine3.12
ENV PORT=3000
EXPOSE $PORT

WORKDIR /app

COPY package.json .
COPY yarn.lock .

RUN yarn install
COPY . .

RUN yarn build
CMD ["sh", "-c", "yarn start -p $PORT"]