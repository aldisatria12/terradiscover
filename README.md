#How to run:
1. Clone repo
   ```bash
   git clone https://github.com/aldisatria12/terradiscover
   cd terradiscover
   ```
3. Setting up environment for backend
   ```bash
   cp ./backend/api/.env.example ./backend/api/.env
   ```
4. Run docker image
   ```bash
   docker compose up --build
   ```
5. Setting up environment for frontend
   ```bash
   cp ./frontend/terradiscover/.env.example ./frontend/terradiscover/.env
   ```
6. Run frontend
   ```bash
   npm install
   npm run dev
   ```
