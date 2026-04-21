### 🤝 Jak pracujemy

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