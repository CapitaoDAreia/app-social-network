FROM node:latest AS build-frontend
WORKDIR /usr/src/app-frontend

COPY . .
RUN npm install && npm run build
CMD ["npm", "run", "preview"]