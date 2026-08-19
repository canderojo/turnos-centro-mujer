# Decisiones técnicas

Este archivo documenta decisiones no triviales tomadas durante el proyecto,
para poder justificarlas en las defensas orales.

## 2026-08-18 — Elección del dominio

**Dominio elegido**: Turnero para un Centro de Salud de la Mujer, con
profesionales de 4 especialidades (dermatología, nutrición, ecografía,
endocrinología) y pacientes que reservan turnos.

**Por qué este dominio**:
- Tiene entidades con relaciones claras (Profesional 1→N Turno, Paciente
  1→N Turno) pero sin explotar en complejidad
- Tiene reglas de negocio no triviales y testeables unitariamente
  (validación de horarios, detección de superposición de turnos, máquina
  de estados de un turno, snapshot de precio).
  

**Por qué Go + chi + sqlx/pgx (backend)**:
- **Go**: lenguaje compilado, tipado, con concurrencia simple — bueno para
  mostrar en vivo que el código es legible y fácil de razonar sobre él.
- **chi** en vez de un framework grande tipo Gin: es un router minimalista,
  muy cercano a la librería estándar `net/http` de Go. Menos "magia",
  más fácil de explicar línea por línea en una defensa oral.
- **sqlx + pgx** en vez de un ORM (como GORM): sqlx es una capa fina sobre
  `database/sql` que permite escribir SQL explícito y mapearlo a structs
  de Go. La decisión es intencional: quiero poder mostrar y explicar el
  SQL real que se ejecuta contra Postgres, no una abstracción que lo
  esconda. pgx es el driver de Postgres más usado y performante en el
  ecosistema Go.

**Por qué React + Vite (frontend)**:
- Es el stack SPA más estándar y con más documentación/soporte para
  resolver dudas como principiante.
- Vite da un dev server rápido y una config mínima comparado con
  Create React App (que además está deprecado).

**Por qué Postgres**:
- Pedido/sugerido por la cátedra como base de datos relacional estándar.
- El dominio tiene relaciones (FKs) e integridad referencial real
  (turno → profesional, turno → paciente), que encaja natural con un
  modelo relacional.

**Estructura del repo**: `backend/` y `frontend/` como carpetas separadas
en un mismo repo (monorepo simple), en vez de dos repos aparte. Facilita
tener un solo historial de commits y un solo lugar para correr todo en
las defensas, sin la complejidad de gestionar submódulos o repos
sincronizados.

**Config por variables de entorno desde el día 1**: tanto el backend
(connection string de Postgres, puerto) como el frontend (`VITE_API_URL`)
leen su configuración de variables de entorno en vez de tener valores
hardcodeados. 

