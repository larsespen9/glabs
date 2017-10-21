# Hva har jeg lært?
## Tips og triks 
Når vi vil iterere og foreksempel reversere tekst i golang så kan det være lurt å bruke runes isteden for bytes. Dette er spesielt viktig hvis man vil reversere tekst som ikke har asci tegn, men isteden bruker utf8 det vi har norske, kinesike tegn etc. se slutten av [denne videoen](https://www.youtube.com/watch?v=XCsL89YtqCs) for å se hvordan dette fungerer. 

## Testing
For å lage testfiler i go så må man importere test biblotek. Vanlig praksis er å kjøre testene i en egen fil. 
Koden under viser et eksempel på testing av en Sum funksjon. Både testen og sum filen ligger i samme mappe (sum mappen)
```go
package sum
import "testing"
func TestSum(t *testing.T) {
	var tests = []struct {
		a    int
		b    int
		want int
	}{
		{3, 4, 7},
		{3, 5, 8},
		{3, 7, 10},
		{3, -4, -1},
	}
	for _, c := range tests {
		got := Sum(c.a, c.b)
		if got != c.want {
			t.Errorf("Sum(%d,%d) => %d, want %d", c.a, c.b, got, c.want)
		}
	}
}
```








## Compilerling og build av go code med pakker
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
