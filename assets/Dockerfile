FROM node:16.10.0 as Builder

WORKDIR /app

COPY . .

RUN yarn config set registry https://registry.npm.taobao.org && yarn install && yarn run build

FROM nginx

WORKDIR /usr/share/nginx/html

COPY ./deploy/nginx.conf /etc/nginx/conf.d/default.conf

COPY --from=Builder /app/dist /usr/share/nginx/html/

EXPOSE 80
