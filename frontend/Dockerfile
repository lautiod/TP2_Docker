# --- build stage ---
FROM node:18-alpine AS build
WORKDIR /app
COPY frontend/package*.json ./
RUN npm ci
COPY frontend/ .
RUN npm run build   # genera /app/dist

# --- runtime stage ---
FROM node:18-alpine
WORKDIR /app
# vite preview sirve estáticos; si preferís otro server, avísame
RUN npm i -g vite
COPY --from=build /app/dist ./dist
EXPOSE 5173
# Preview usa 4173 por defecto; lo forzamos a 5173 para tu compose
CMD ["vite", "preview", "--host", "0.0.0.0", "--port", "5173"]
