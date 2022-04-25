FROM node:16-alpine

WORKDIR /app

COPY package*.json ./

COPY . .

EXPOSE 8080

RUN npm install

ENV IP_SERVER=localhost

CMD ["npm", "run", "start"]