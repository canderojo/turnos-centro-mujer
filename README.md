# Turnero Centro de Salud de la Mujer

Sistema de reserva de turnos para un centro de salud con profesionales de
4 especialidades: dermatología, nutrición, ecografía y endocrinología.

Proyecto para la materia **Ingeniería del Software 3** (UCC 2026). Cada
trabajo práctico del cuatrimestre agrega una capa nueva sobre esta misma
aplicación (CI, testing, CD, contenedores, IaC, seguridad, observabilidad).

## Stack

- **Backend**: Go + [chi](https://github.com/go-chi/chi) (router) + [sqlx](https://github.com/jmoiron/sqlx) + [pgx](https://github.com/jackc/pgx) (driver de Postgres)
- **Frontend**: React + Vite
- **Base de datos**: PostgreSQL

## Estructura del repo

```
CENTROMUJER/
├── backend/     # API en Go
├── frontend/    # SPA en React + Vite
├── decisiones.md
└── README.md
```

## Cómo levantar el proyecto en local

### Requisitos

- Go 1.22+
- Node.js 18+
- PostgreSQL corriendo en local

### Backend

```bash
cd backend
cp .env.example .env   # completar DATABASE_URL con tu password de Postgres

# Crear la base y cargar el esquema (tablas + datos de ejemplo)
createdb turnos_centro_mujer
psql -d turnos_centro_mujer -f db/init.sql

go run .
```

La API queda escuchando en `http://localhost:8080`. Para confirmar que
levantó bien y que puede hablar con Postgres:

```bash
curl http://localhost:8080/health
```

Endpoints disponibles:

- `GET /health` — estado del servidor y de la conexión a la base.
- `GET /profesionales` (opcional `?especialidad=`) y `GET /profesionales/{id}`.
- `GET /profesionales/{id}/horarios-disponibles?fecha=YYYY-MM-DD` — huecos libres para reservar ese día.
- `POST /turnos` — reserva un turno (crea el paciente si no existe, identificándolo por DNI).
- `GET /turnos/{id}` y `GET /turnos?dni=...` / `?email=...` ("mis turnos").
- `PATCH /turnos/{id}/estado` — confirma, cancela o completa un turno.

### Frontend

Pendiente: `frontend/` todavía no está inicializado (ver
[decisiones.md](./decisiones.md) para el stack elegido — React + Vite).

## Documentación de decisiones técnicas

Ver [decisiones.md](./decisiones.md) para el razonamiento detrás de las
elecciones de dominio, stack y arquitectura.
