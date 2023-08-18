FROM node:18-alpine AS base

COPY . .
RUN npm install 

ENTRYPOINT [ "node", "index.js" ]
