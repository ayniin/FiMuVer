FROM node:20-alpine

WORKDIR /app

# Kopiere package.json und package-lock.json
COPY package*.json ./

# Installiere Dependencies (für dev und build)
RUN npm ci

# Kopiere Source Code
COPY . .

# Expose Port
EXPOSE 5173

# Starte dev Server
CMD ["npm", "run", "dev", "--", "--host"]

# Starte die Anwendung
CMD ["serve", "-s", "dist", "-l", "5173"]
