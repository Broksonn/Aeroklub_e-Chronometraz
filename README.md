## 👨‍💻 My Role: Frontend Developer
In this project, I was responsible for the design, implementation, and API integration of the entire client-side (frontend) application layer, focusing on building a responsive, secure, and intuitive interface.

### 🚀 Key Achievements:
* **Frontend Architecture (Vue.js 3):** Developed the application's core dynamic views (Pilot Panel and Admin Panel) from scratch using the Composition API.
* **State Management (Pinia):** Implemented global application state for authentication and asynchronous fetching of aircraft fleet data directly from the Go/REST API backend.
* **Security & Authentication:** Configured Vue Router Navigation Guards integrated with backend JWT tokens to strictly separate user roles (Admin vs. Pilot) and prevent unauthorized access.
* **External API Integration:** Developed a dynamic weather widget fetching 7-day airport forecasts asynchronously via the Open-Meteo API.
* **Reservation Interface:** Created an interactive date selection calendar and a dynamic system for calculating flight minute rates for specific aircraft.

### 🛠️ Tech Stack & Tools:
* **Framework:** Vue.js 3
* **State Management:** Pinia
* **Routing:** Vue Router
* **Languages:** TypeScript / JavaScript, HTML5, CSS3
* **Communication:** Fetch API (REST integration with Go/PostgreSQL backend)

# ✈️ Aeroklub e-Chronometraż

Kompleksowy system do zarządzania operacjami lotniczymi, flotą aeroklubu oraz rezerwacjami statków powietrznych. Projekt stworzony w architekturze Full-Stack.

## 🛠 Wykorzystane technologie

* **Backend:** Go (Golang), Gin Web Framework, GORM, PostgreSQL, JWT, bcrypt
* **Frontend:** Vue 3 (Composition API), Vite, TypeScript, Pinia, Vue Router, V-Calendar
* **Infrastruktura:** Docker, Docker Compose

---

## 📋 Wymagania wstępne (Prerequisites)

Aby uruchomić ten projekt na swoim komputerze, upewnij się, że masz zainstalowane następujące narzędzia:

1. **[Docker](https://www.docker.com/products/docker-desktop/) oraz Docker Compose** (niezbędne do uruchomienia bazy danych PostgreSQL).
2. **[Go](https://go.dev/dl/)** (wersja 1.20 lub nowsza).
3. **[Node.js](https://nodejs.org/)** (wersja 18 lub nowsza) oraz menedżer pakietów **npm**.
4. **Git** (do sklonowania repozytorium).

---

## 🚀 Szybki start (Zalecane)

Dla wygody programistów przygotowaliśmy automatyczne skrypty startowe, które samodzielnie uruchamiają bazę danych, instalują brakujące pakiety, budują aplikację i uruchamiają serwery.

### 🐧 Dla systemu Linux / macOS

1. Otwórz terminal w głównym folderze projektu.
2. Nadaj skryptowi uprawnienia do wykonywania (tylko za pierwszym razem):
```bash
   chmod +x run.sh

```

3. Uruchom skrypt:
```bash
./run.sh

```


4. Aby zatrzymać wszystkie serwery i bazę danych, po prostu wciśnij `Ctrl+C` w tym samym oknie terminala.

### 🪟 Dla systemu Windows

1. Otwórz folder główny projektu w Eksploratorze plików.
2. Kliknij dwukrotnie plik `run.bat`.
3. Skrypt uruchomi bazę danych w tle i otworzy dwa dodatkowe okna konsoli (jedno dla Frontendu, drugie dla Backendu).
4. Aby zatrzymać aplikację, przejdź do głównego okna skryptu i naciśnij dowolny klawisz. Skrypt automatycznie zamknie procesy Go, Vue i zatrzyma kontener Docker.

---

## ⚙️ Uruchamianie ręczne (Krok po kroku)

Jeśli wolisz kontrolować każdy proces osobno, postępuj zgodnie z poniższymi instrukcjami.

### Krok 1: Uruchomienie Bazy Danych

Projekt korzysta z bazy PostgreSQL uruchamianej w kontenerze na niestandardowym porcie `5434` (aby zapobiec konfliktom z lokalnymi bazami).

```bash
docker-compose up -d

```

*(Baza danych zostanie utworzona z poświadczeniami: użytkownik `aeroklub_user`, hasło `aeroklub_pass`, baza `aeroklub_db`).*

### Krok 2: Uruchomienie Backendu (API)

Przejdź do folderu `backend`, zainstaluj moduły Go i uruchom serwer. Przy pierwszym uruchomieniu system automatycznie utworzy tabele w bazie (AutoMigrate) oraz doda testowych użytkowników i samoloty (Seeder).

```bash
cd backend
go mod tidy
go run cmd/api/main.go

```

*Backend będzie nasłuchiwał na porcie: `http://localhost:8080*`

### Krok 3: Uruchomienie Frontendu

Otwórz nową kartę terminala, przejdź do folderu `frontend`, zainstaluj zależności i uruchom środowisko deweloperskie Vite.

```bash
cd frontend
npm install
npm run dev

```

*Frontend będzie dostępny pod adresem: `http://localhost:5173*`

---

## 🔐 Domyślne konta testowe

Podczas pierwszego uruchomienia na czystej bazie danych, system automatycznie (poprzez plik `seed.go`) generuje następujące konta do testowania aplikacji:

| Rola w systemie | Login | Hasło | Opis uprawnień |
| --- | --- | --- | --- |
| **Super Administrator** | `superadmin` | `admin123` | Pełny dostęp. Zarządzanie flotą, cennikiem, tworzenie kont Pilotów, Adminów oraz innych SuperAdminów. |
| **Administrator** | `admin` | `admin` | Zarządzanie flotą i cennikiem. Może tworzyć **tylko** konta Pilotów. Brak uprawnień do nadawania praw administracyjnych. |
| **Pilot** | `pilot` | `123` | Standardowy dostęp. Przegląd floty i cennika, możliwość rezerwacji lotów z kontrolą nakładania się czasu rezerwacji. |

---

## 📝 Dodatkowe informacje

* Rezerwacje są weryfikowane na żywo – system nie pozwoli na zarezerwowanie tego samego statku przez dwie osoby w nakładających się oknach czasowych.
* Rezerwacji należy dokonywać z minimum 30-minutowym wyprzedzeniem.
* Hasła w bazie danych są bezpiecznie haszowane za pomocą algorytmu `bcrypt`. Dostęp do tras API chroniony jest tokenami `JWT`.
    
#    
# 🤝 Jak pracujemy

## 🌿 Branch

Nie pracuj na `main`.

Twórz branch:

```bash
git checkout -b feature/nazwa-zmiany
backend/ykalesnikau -> backend mój
frontend/mbrokos -> frontend mateusza
```

Teraz bezpiecznie pracujecie na własnej wersji projektu i jak coś zepsujecie nie wpływa to na mastera/maina

## 💾 Commit

```bash
git add .
git commit -m "Opis zmian"
```

Tak bierzecie wszystkie zmienione pliki na swoim branchu i zapisujecie je do gita (git lokalny github zdalny)

## Push
```bash
git push origin nazwa-branchu
```
Po commicie, możecie zpushować, wtedy wysyłacie to na naszego githuba

## Pull Request
Zpushowane branche mogą zostać, ale żeby weszły do maina (finalnej wersji projektu) trzeba
1. Wejść na github
2. Kliknąć "Compare & Pull Request"
3. Wybrać main <- branch
4. Kliknąć "Create Pull Request"
Wtedy ja sobie to oglądam i daje okejke jak bedzie git i ostatecznie trafia do projektu

W pull request fajnie jak napiszecie w skrócie co dodaliście

## Forcowanie

Jeżeli chcecie zmienić ostatni commit żęby nie śmiecić to używajcie

```bash
git branch  \<- upewnijcie sie że jesteście na swoim branchu a nie mainie, chociaż git raczej was zatrzyma bo nie macie pełnego prawa do repo
git commit --amend --no-edit
git push --force-with-lease
```
