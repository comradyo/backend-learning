FROM nginx:1.27.4-alpine
# Контекст будет на папке выше (описывается в compose.yaml)
#COPY .nginx/nginx.conf /etc/nginx/conf.d/default.conf
# COPY data/static/index.html /usr/share/nginx/html/index.html
COPY data data

EXPOSE 80
CMD ["nginx", "-g", "daemon off;"]
