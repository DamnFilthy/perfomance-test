FROM node:20
WORKDIR /app
ADD *.json ./
RUN npm i pnpm -g
RUN pnpm i
COPY . .
ENV PATH ./node_modules/.bin/:$PATH
CMD pnpm run start