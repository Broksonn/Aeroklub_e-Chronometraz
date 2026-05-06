# 🚀 Szybki start (Uruchomienie na nowym komputerze)

Jeśli pobrałeś projekt na nowy komputer, musisz wykonać kilka kroków instalacyjnych, ponieważ pliki z zależnościami (np. node_modules) i konfiguracje bazy danych nie są wysyłane na GitHuba.
Wymagania wstępne:

1. Zainstalowany Go (Golang)

2. Zainstalowany Node.js (wraz z npm)

3. Zainstalowany PostgreSQL oraz pgAdmin

## Krok 1: Baza danych

1. Otwórz pgAdmin i stwórz nową, pustą bazę danych o nazwie aeroclub.

2. Nie musisz tworzyć tabel ręcznie – backend używa GORM i automatycznie wygeneruje struktury oraz doda dane testowe przy pierwszym uruchomieniu.

## Krok 2: Konfiguracja środowiska (Backend)

1. Wejdź do folderu backend.

2. Utwórz plik o nazwie __*.env*__ i wklej do niego dane z pliku __*.env.examle*__ i zmień hasło na swoje:

```bash
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=twoj_password_od_postgresa
DB_NAME=aeroclub
```

3. Pobierz wszystkie biblioteki Go (Gin, GORM, JWT itp.):

```bash
cd backend
go mod tidy
```

## Krok 3: Instalacja zależności (Frontend)

Wejdź do folderu frontend i zainstaluj paczki NPM:

```bash
cd frontend
npm install
```

## Krok 4: Uruchomienie projektu 

### Backend + Frontend

W głównym folderze projektu (tam, gdzie znajduje się główny plik package.json) uruchom skrypt startowy, który podniesie oba serwery jednocześnie:

```bash
npm run dev
```

###  Tylko Backend

W przypadku potrzeby podniesienia tylko serwera backendowego należy w folderze backend uruchomić polecenie 

```bash
go run cmd/api/main.go  
```

### W obu przypadkach Frontend będzie dostępny pod adresem: http://localhost:5173

# 🏗️ Architektura i Logika Aplikacji (Full-Stack)
⚙️ Backend (Go / Gin / GORM)

## Model Danych i Baza (PostgreSQL):
Zaktualizowano model rozliczeń na stawki minutowe, co odpowiada realiom lotniczym (odczyt z tachometru).
```Go

    type Price struct {
        ID             uint    `json:"id" gorm:"primaryKey"`
        Type           string  `json:"type"`
        PricePerMinute float64 `json:"pricePerMinute"` // Stawka za minutę lotu
    }
```

## Autoryzacja i Bezpieczeństwo (JWT & Bcrypt):
Hasła są szyfrowane przy pomocy Bcrypt. Po weryfikacji serwer generuje token JWT.
 ```Go

    // Weryfikacja hasła w kontrolerze auth.go
    err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(creds.Password))
    if err != nil {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Nieprawidłowe hasło"})
        return
    }
```

## Zabezpieczone Endpointy (Middleware):
Własne middleware w Gin sprawdzają nagłówek Authorization i dekodują token przed wpuszczeniem użytkownika do zasobów.
 ```Go
    // Fragment routes.go
    protected := api.Group("/")
    protected.Use(middleware.RequireAuth()) 
    {
        protected.GET("/airplanes", controllers.GetAirplanes)
        // Tylko dla admina
        admin := protected.Group("/")
        admin.Use(middleware.RequireAdmin())
        admin.PUT("/pricing", controllers.UpdatePrices)
    }
 ```

## Konfiguracja CORS:
Skonfigurowano restrykcyjną politykę CORS, umożliwiając przeglądarkom bezpieczne łączenie się z API, wyjatkiem jest Firefox.

# 🖥️ Frontend (Vue.js 3 / Pinia / Vue Router)

## Zarządzanie Stanem (Pinia z Fetch API):
Lokalne mocki JSON zastąpiono asynchronicznymi zapytaniami do prawdziwego serwera. Każde żądanie wysyła token JWT.

```TS

    // Fragment stores/auth.ts
    const response = await fetch('http://localhost:8080/api/auth/login', {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({ username, password })
    })
```

## Ochrona Tras (Navigation Guards):
Zabezpieczono dostęp do widoków w Vue Router. Strażnik blokuje niezalogowanych użytkowników i weryfikuje ich role.

```TS

    // Fragment router/index.ts
    router.beforeEach((to, from, next) => {
        const auth = useAuthStore()

        if (to.meta.requiresAuth && !auth.token) {
            next('/login') // Brak tokena -> odrzucenie do logowania
        } else if (to.meta.requiresAdmin && auth.user?.role !== 'admin') {
            next('/pilot') // Brak uprawnień -> przekierowanie
        } else {
            next() // Dostęp przyznany
        }
    })
```

## Dynamiczny Interfejs (Reaktywność):
Komponenty UI (np. App.vue i PilotDashboardView.vue) automatycznie renderują odpowiednie widoki, ceny ({{ plane.pricePerMinute }} zł / min) oraz ukrywają niepotrzebne przyciski w zależności od aktywnego stanu w Pinia.