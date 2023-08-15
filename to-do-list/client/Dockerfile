FROM node:20 AS frontend-builder

WORKDIR /app/frontend

# coping the front-end source code to /app/frontend folder
COPY client/ .

# Install the dependencies and run build the front-end
RUN npm install
RUN npm run build

# Create a new lightweight image from the official NGINX image
FROM nginx:alpine

COPY --from=frontend-builder app/frontend/build /usr/share/nginx/html

EXPOSE 80

# Start NGINX to serve the frontend
CMD ["nginx", "-g", "daemon off;"]