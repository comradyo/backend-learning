# For Nginx setup
FROM nginx:alpine

# Set working directory
WORKDIR /app

# COPY .nginx/nginx.conf /app/.nginx
COPY data /app
# Возможно, неправильно
COPY data /usr/share/nginx/html 

# Copy config nginx
COPY .nginx/nginx.conf /etc/nginx/conf.d/default.conf



WORKDIR /usr/share/nginx/html

# Remove default nginx static assets
RUN rm -rf ./*

# Copy static assets from builder stage
# COPY --from=build /app/build .
# Это, кажется, уже наверху сделал

#EXPOSE the port
EXPOSE 80

# Containers run nginx with global directives and daemon off
ENTRYPOINT ["nginx", "-g", "daemon off;"]

# Что-то не работает нифига...
