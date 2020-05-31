# Konzept

## Systemstart

Beim starten werden die full-data, die max-change-id und die min-change-id aus
dem Cache gezogen.

* Die Daten werden atomic gezogen
* sind die Daten nicht da wird eine sekunde gewartet un dann wiederholt.


## Hintergrund Dienst

Anschließend wird auf einen Redis-Stream gehört. Der Stream enthält alle
geänderten Daten und die jeweilige Change-id.

Ist die Change id mehr als eins größer als die höhste Change id, dann werden
die Daten aller fehlenden change ids manuell geladen bevor die Daten aus dem
Stream bearbeitet werden. (Für manuelles laden siehe unten.)

Ist die change id kleiner als eins größer als die höhste change id, dann wird das
paket verworfen.

Um eine Change ID oder ein range einer change id manuell zu laden wird das
change-id feld aus redis ausgelesen und die entsprechenden keys geladen. Dies
muss nicht atomic sein.


## Verbindungsaufbau Client

Ein Client gibt bei der Übergabe die change-id an, die er zuletzt gesehen hat. 0
bedeutet, er hat noch keine Daten gesehen.

* Ist die id kleiner als min-change-id - 1 dann muss der Client alle Daten erhalten.
* Ist die id gleich max-change-id, dann muss nichts passieren.
* Wie muss reagiert werden, wenn die id größer als max-change-id ist?

Anderenfalls werden aus redis alle keys größer der übergebenen change-id und
max-change-id (inklusive) gezogen und die daten mit max-change-id an den Client
gesandt.

Wenn der Client alle Daten erhalten muss, dann wird ihm dies so separat
mitgeteilt. Er erhält max-change-id.


## Autoupdate Client

Kommt ein neuer Datensatz über redis, dann erhalten alle verbundenen clients die
Daten inklusive der übergebenen neuen change-id.


## Datenübermittlung

Nachdem entschieden wurde, welche Daten ein client erhält, werden die keys aus
dem cache gezogen. Anschließend werden sie zum restricter gesandt.

* Was ist, wenn keine Daten übrig bleiben? Wird dann nur die Change id
  übermittelt?
* Wie funktioniert das mit dem gelöschten/weg restricteten Daten?
* Sollten nur Daten gesandt werden, die sich geändert haben? Würde das bei den
  gelöschten Daten helfen?


## Projektor

* Wie zieht der Dienst am Anfang die Daten? Gibt es einen stream? einfach nur
  ein key?
* Ist es möglich die Daten im Autoupdate-Dienst zu berechnen?


### Konzept: Die Daten werden im autoupdate-dienst berechnet

Ganz am Anfang und nach jedem Update vom Server wird der Projektor neu
berechnet. Zusätzlich wird immer noch ein hash gebildet. Ändert sich der hash,
wird ein projektor[id] changed getriggert.

Clients können eine projektor-id abonnieren. Ändert sich der Wert, werden die
neuen Daten gesandt.


