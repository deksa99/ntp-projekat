# Servis automobila

Aplikacija za podršku rada servisa automobila. Omogućava pronalaženje servisa, biranje usluga i rezervisanje termina, kao i konfiguraciju samog servisa.


# Funkcionalnosti

Sistem koriste sledeći korisnici

 - Neregistrovani korisnik
 - Registrovani korisnik
 - Serviser
 - Administrator
 
 **Neregistrovani korisnik**
 - Prijava
 - Registracija
 - Pregled servisa
 - Pregled usluga servisa
 
 **Registrovani korisnik**
 - Funkcionalnosti *neregistrovanog* korisnika
 - Pronalaženje najbližeg servisa*
 - Pregled/izmena profila
 - Promena lozinke
 - Pregled/izmena svojih vozila
 - Zakazivanje i otkazivanje termina
 - Pregled svojih termina
 - Ostavljanje recenzija
 - Prijava komentara

**Serviser**

- Vršenje zakazanih intervencija
- CRUD usluga servisa
- Pregled izveštaja o poslovanju*

**Administrator**

- Pregled prijavljenih komentara
- Blokiranje korisnika na 3 dana

# Arhitektura
- Servisi - *podložno manjim promenama u toku implementacije*
	- API gateway - Go
	- Korisnički servis - Go
	- Email servis - Go
	- Servis za rezervaciju usluga auto servisa - Go
	- Servis za automobile - Go
	- Servis za izveštaje - Rust
	- Servis za manipulaciju auto servisima - Go
- Frontend
	- Korisnička aplikacija u React-u
- Baza
	- SQL
