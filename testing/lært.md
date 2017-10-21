# Hva har jeg lært?
For å lage en ferdig compilert fil av go i den mappen man er i så må man bruke følgende kommando
```
go build main.go
```

Hvis man lager / importerer noen custom funkskjoner så må følgende kriterier oppfylles:

[ ] De filene som man vil bruke må ligge i en egen mappe

[ ] Funksjonene i filene må starte med stor forbokstav

[ ] Man må bruke go install på pakke

[ ] Pakke navnet må være noe annet en main

[ ] Når man skal bruke funksjonene i main så må man bruke 

[ ] Pakkenavn før man kaller funksjon som vist under: 
```go
func main() {
	d := sum.Sum(4, 5)
	fmt.Println(d)
}
```
