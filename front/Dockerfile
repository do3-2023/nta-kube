FROM node:lts-alpine AS build

WORKDIR /app
COPY . .
RUN npm install
RUN npm run build

FROM node:lts-alpine AS run
COPY --from=build /app/build ./build
COPY --from=build /app/package.json ./package.json
COPY --from=build /app/node_modules ./node_modules
CMD ["node", "build"]