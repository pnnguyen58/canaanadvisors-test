FROM node:18-alpine

WORKDIR /app

COPY . .

RUN npm ci --omit=dev

RUN npm run build

USER node

CMD ["npm", "run", "start:prod"]