FROM node:16-alpine
WORKDIR /
COPY . .
RUN npm install
ENV HOSTNAME_RABBIT 192.168.1.10
EXPOSE 50051
CMD ["npm","start"]