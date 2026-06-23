@echo off
:: Ustawienie kodowania UTF-8, aby polskie znaki wyświetlały się poprawnie w konsoli
chcp 65001 >nul

echo ========================================================
echo 🚀 Uruchamianie projektu Aeroklub e-Chronometraz...
echo ========================================================
echo.

:: 1. Uruchamianie bazy danych PostgreSQL
echo 📦 Startowanie bazy danych PostgreSQL (Docker na porcie 5434)...
docker-compose up -d

:: Dajemy bazie 3 sekundy na pełne uruchomienie
timeout /t 3 /nobreak >nul

:: 2. Uruchamianie Backend (Go) w nowym oknie
echo 🐹 Uruchamianie backendu (Go)...
cd backend
:: "start" otwiera nowe okno z unikalnym tytułem, abyśmy mogli je później łatwo zamknąć
start "Aeroklub_Backend" cmd /c "go mod tidy && go run cmd/api/main.go"
cd ..

:: 3. Uruchamianie Frontend (Vue) w nowym oknie
echo 🌐 Uruchamianie frontendu (Vue 3 + Vite)...
cd frontend
start "Aeroklub_Frontend" cmd /c "npm install && npm run dev"
cd ..

echo.
echo ========================================================
echo ✅ Wszystkie serwery zostaly pomyslnie uruchomione!
echo 👉 Frontend (Panel): http://localhost:5173
echo 👉 Backend (API):    http://localhost:8080
echo ========================================================
echo 🛑 Aby ZATRZYMAC projekt i wylaczyc serwery, wroc do tego okna i...
pause

:: --- SEKCJA ZAMYKANIA PROCESÓW (uruchomi się po wciśnięciu klawisza) ---
echo.
echo 🛑 Zatrzymywanie procesow...

:: Zabijanie procesów po tytułach okien
taskkill /FI "WINDOWTITLE eq Aeroklub_Backend*" /T /F >nul 2>&1
taskkill /FI "WINDOWTITLE eq Aeroklub_Frontend*" /T /F >nul 2>&1

echo 🐳 Zatrzymywanie bazy danych Docker...
docker-compose stop

echo 👋 Do widzenia!
timeout /t 2 >nul