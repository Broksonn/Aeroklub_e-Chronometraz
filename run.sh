#!/bin/bash

# Ustawienie kodowania na UTF-8 dla polskich znaków
export LANG=C.UTF-8

echo "🚀 Uruchamianie projektu Aeroklub e-Chronometraż..."

# Funkcja do czyszczenia procesów przy zamykaniu (po wciśnięciu Ctrl+C)
cleanup() {
    echo ""
    echo "🛑 Zatrzymywanie procesów..."
    
    # Zabijanie procesów w tle
    kill $BACKEND_PID 2>/dev/null
    kill $FRONTEND_PID 2>/dev/null
    
    echo "🐳 Zatrzymywanie bazy danych Docker..."
    # Zatrzymuje kontenery, ale NIE usuwa wolumenu z danymi (aby nie stracić bazy)
    docker-compose stop
    
    echo "👋 Do widzenia!"
    exit 0
}

# Przechwytywanie sygnału zakończenia (Ctrl+C) i wywołanie funkcji cleanup
trap cleanup SIGINT SIGTERM

# 1. Uruchamianie bazy danych PostgreSQL
echo "📦 Startowanie bazy danych PostgreSQL (Docker na porcie 5434)..."
docker-compose up -d

# Dajemy bazie 3 sekundy na pełne uruchomienie przed startem backendu
sleep 3

# 2. Uruchamianie Backend (Go)
echo "🐹 Uruchamianie backendu (Go)..."
cd backend
# Pobieranie ewentualnych brakujących pakietów (np. bcrypt, gin)
go mod tidy
# Uruchamianie w tle (znak '&' na końcu)
go run cmd/api/main.go &
BACKEND_PID=$!
cd ..

# 3. Uruchamianie Frontend (Vue)
echo "🌐 Uruchamianie frontendu (Vue 3 + Vite)..."
cd frontend
# Instalacja zależności (jeśli jeszcze ich nie ma)
npm install
# Uruchamianie w tle
npm run dev &
FRONTEND_PID=$!
cd ..

echo "========================================================"
echo "✅ Wszystkie serwery zostały pomyślnie uruchomione!"
echo "👉 Frontend (Panel): http://localhost:5173"
echo "👉 Backend (API):    http://localhost:8080"
echo "========================================================"
echo "Wciśnij [CTRL+C], aby zatrzymać projekt i bezpiecznie wyłączyć serwery."

# Oczekiwanie w nieskończoność, aby skrypt się nie zamknął
wait