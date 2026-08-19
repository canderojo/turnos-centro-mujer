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

> Instrucciones completas y detalladas se agregan a medida que el backend
> y el frontend queden armados (ver secciones más abajo a medida que se
> completen los pasos).

### Requisitos

- Go 1.22+
- Node.js 18+
- PostgreSQL corriendo en local

Documentación detallada de cada paso (backend, frontend, base de datos)
se va completando en este README a medida que avanza el proyecto.

## Documentación de decisiones técnicas

Ver [decisiones.md](./decisiones.md) para el razonamiento detrás de las
elecciones de dominio, stack y arquitectura.
