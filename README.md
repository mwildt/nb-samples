# nb-samples

## Application
Bei diesem Projekt handelt es sich um ein Sample-Projekt für eine Containr-Umgebung.

Der Container kann in zwei unterschiedlichen Modi [tcp|http] gestartet werden. Es wird jeweils ein TCP-Socket oder ein Http Server gestartet.

** Bauen der Anwendung **
```
go build -o main 
```

**Bauen und Starten des TCP-Socker Server**
```
SERVICE_MODE=tcp SERVICE_HOST=:3300 ./main
```

**Bauen und Starten des HTTP Server**
```
SERVICE_MODE=http SERVICE_HOST=:3200 ECHO_HOST=localhost:3300 ./main
```

Die Umgebungsvariable `ECHO_HOST` git dabei an, unter welcher Adresse der TCP-Echo-Service erreichbar ist.

## Docker
Das Image wird automatisch via Gitlab Action gebaut und nach hub.docker.com gepusht:
https://hub.docker.com/repository/docker/maltewildt/nb-sample


