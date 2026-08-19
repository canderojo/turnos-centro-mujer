# Decisiones técnicas

Este archivo documenta decisiones no triviales tomadas durante el proyecto,
para poder justificarlas en las defensas orales.

## 2026-08-18 — Elección del dominio y stack (TP2)

**Dominio elegido**: Turnero para un Centro de Salud de la Mujer, con
profesionales de 4 especialidades (dermatología, nutrición, ecografía,
endocrinología) y pacientes que reservan turnos.

**Por qué este dominio y no el sample de la cátedra**:
- Tiene entidades con relaciones claras (Profesional 1→N Turno, Paciente
  1→N Turno) pero sin explotar en complejidad — se puede explicar completo
  en una defensa oral de 15-20 minutos.
- Tiene reglas de negocio no triviales y testeables unitariamente
  (validación de horarios, detección de superposición de turnos, máquina
  de estados de un turno, snapshot de precio), que es justo lo que pide
  la cátedra poder cubrir con tests en el TP5.
- No requiere autenticación para arrancar (identificación por DNI/email),
  lo que simplifica el alcance inicial sin sacrificar reglas de negocio
  reales.

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
hardcodeados. Esto es una decisión pensada para el futuro: cuando se
dockerice el proyecto (trabajo posterior), no debería hacer falta tocar
código, solo cambiar las variables de entorno inyectadas por Docker.

## 2026-08-19 — Esquema de base de datos (TP2)

**Enums de aplicación (`especialidad`, `estado`) como `TEXT` + `CHECK`**,
en vez de tipos `ENUM` nativos de Postgres: agregar un valor nuevo a un
`CHECK` es un `ALTER TABLE` chico y fácil de explicar; modificar un
`ENUM` nativo de Postgres es más raro (hay que usar `ALTER TYPE ... ADD
VALUE`, que además no se puede revertir fácil). Para el tamaño de este
proyecto, `TEXT + CHECK` es más simple sin perder validación a nivel
de base de datos.

**Repository pattern simple (funciones, no interfaces)**: la capa
`internal/repository` son funciones de Go que reciben `*sqlx.DB` y
devuelven structs o errores — sin interfaces ni "repositorio genérico".
Elegido a propósito: es el patrón mínimo que separa el SQL del resto
del código (necesario para poder testear las reglas de negocio del
TP5 sin levantar Postgres), sin la complejidad extra de interfaces que
en un proyecto de este tamaño no se llegan a aprovechar.

## Notas para más adelante (no implementar todavía)

- Cuando se dockerice: el `.env.example` del backend va a necesitar un
  segundo valor documentado para cuando Postgres corra en un contenedor
  separado (el host de conexión cambia de `localhost` a el nombre del
  servicio de Docker Compose, ej. `db`).
- Evaluar en su momento si conviene un `Makefile` o scripts npm/go para
  levantar todo con un solo comando, una vez que exista `docker-compose`.
- Autenticación de pacientes: pendiente de definir con la cátedra: no
  implementar hasta que haya una decisión explícita.
