FROM node:18-alpine AS base

COPY . .
RUN npm install
RUN npm run build

CMD ["npm", "run", "start"]