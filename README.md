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

```bash
cd frontend
cp .env.example .env   # VITE_API_URL apunta a http://localhost:8080 por defecto

npm install
npm run dev
```

El sitio queda en `http://localhost:5173` (el backend ya tiene ese origen
habilitado en CORS). Requiere el backend corriendo para poder listar
profesionales, ver horarios y reservar turnos.

Pantallas:

- **Profesionales** (`/`) — listado filtrable por especialidad.
- **Detalle de profesional** (`/profesionales/:id`) — elegir fecha, ver
  horarios disponibles y reservar un turno.
- **Mis turnos** (`/mis-turnos`) — buscar turnos propios por DNI o email.
- **Detalle de turno** (`/turnos/:id`) — ver un turno y cambiar su estado
  (confirmar, cancelar, completar) según las transiciones permitidas.

## Documentación de decisiones técnicas

Ver [decisiones.md](./decisiones.md) para el razonamiento detrás de las
elecciones de dominio, stack y arquitectura.
