@echo off
chcp 65001 >nul
echo ===================================================
echo      URUCHAMIANIE PROJEKTU: E-CHRONOMETRAZ
echo ===================================================

echo.
echo [1/3] Startowanie bazy danych (Docker)...
docker-compose up -d

echo.
echo [2/3] Startowanie serwera Backend (Go)...
start "Backend (Go)" cmd /k "cd backend && go run cmd/api/main.go"

echo.
echo [3/3] Startowanie serwera Frontend (Vue)...
start "Frontend (Vue)" cmd /k "cd frontend && npm run dev"

echo.
echo ===================================================
echo Wszystko gotowe! Serwery dzialaja w nowych oknach.
echo.
echo 🌍 Frontend: http://localhost:5173
echo 🗄️  pgAdmin:  http://localhost:5050
echo ===================================================
pause