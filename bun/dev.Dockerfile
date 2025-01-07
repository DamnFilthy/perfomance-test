FROM oven/bun:1 AS base
WORKDIR /app

COPY *.json ./

RUN bun install

COPY . .

ENV NODE_ENV=production

USER bun
EXPOSE 3007/tcp

ENTRYPOINT [ "bun", "run", "start" ]