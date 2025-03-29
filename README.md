# Smart Review - Analisi delle Recensioni e Risposte Automatiche

**Smart Review** è un'applicazione avanzata progettata per analizzare le recensioni degli utenti e generare risposte automatiche basate sull'intelligenza artificiale. Utilizzando una combinazione di tecnologie moderne come **Vite.js** per il frontend, **Golang** per il backend e **Ollama** per l'integrazione AI, **Smart Review** offre una soluzione completa per il miglioramento delle interazioni con i clienti e per il monitoraggio delle opinioni.

### Caratteristiche principali:
- **Analisi delle recensioni**: Elabora e analizza il contenuto delle recensioni, identificando i sentimenti (positivi, negativi, neutri) e le principali tematiche.
- **Generazione di risposte automatiche**: Utilizza modelli di intelligenza artificiale per rispondere automaticamente alle recensioni, migliorando l'efficienza e l'engagement con i clienti.
- **Integrazione con Ollama**: Grazie a Ollama, l'AI è in grado di fornire risposte più naturali e personalizzate, migliorando la qualità dell'interazione.
- **Gestione facile tramite Docker Compose**: Il progetto è completamente containerizzato, facilitando la configurazione e l'esecuzione in ambienti di sviluppo e produzione.

Con **Smart Review**, le aziende possono ottimizzare il processo di gestione delle recensioni, rispondendo in modo rapido e accurato, e ottenere approfondimenti utili per migliorare i loro prodotti e servizi.

## Prerequisiti

Assicurati di avere i seguenti strumenti installati:

- [Docker](https://www.docker.com/get-started)
- [Docker Compose](https://docs.docker.com/compose/)

## Configurazione

Clona questo repository sul tuo sistema:

```bash
git clone https://github.com/devmarcocrotti/smart-review
cd smart-review
```

## Esecuzione del progetto con Docker Compose

Avvia i container

```bash
docker compose up -d
```

## Accesso all'applicazione

Avvia il browser: http://localhost:5174/