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

## 2026-08-19 — Frontend: stack, estética y datos de prueba

**CSS plano con variables (`tokens.css`) en vez de un framework
(Tailwind, styled-components, etc.)**: mismo criterio que "por qué chi
en vez de Gin" del backend — nada de "magia" que haya que explicar por
fuera del CSS estándar en una defensa oral. Los tokens (color,
espaciado, tipografía, radios) viven en un solo archivo y el resto de
los componentes los consume por variable, así que cambiar la paleta
entera es editar un archivo, no perseguir clases sueltas.

**Identidad visual basada en sitios reales, no en "vibes"**: la primera
versión (paleta rosa saturada + gradiente + ilustraciones tipo flor
dibujadas a mano) leía como genérica de IA. Se relevaron sitios reales
de centros de salud de la mujer (Hospital Privado Universitario de
Córdoba, Centro Médico Mujer CABA, cgap.com.ar) inspeccionando sus
estilos computados (color, tipografía, radios, sombras) desde el
navegador, no solo mirándolos. De ahí salió la paleta final: rosa con
cuerpo real pero no saturado (`#ac3673`), tipografía serif (Fraunces)
para títulos + sans (Public Sans) para el resto, fondo crema en vez de
blanco puro o gradiente, radios de 6-14px (ni "pastilla" en todo ni
esquina a rajatabla) y sin ilustraciones decorativas — solo íconos
lineales funcionales.

**Datos de prueba en el frontend (`frontend/src/api/mockData.js`),
activables por `VITE_USE_MOCK_DATA`**: mientras se armaba el diseño no
había necesidad de tener Postgres + el backend Go corriendo en
simultáneo para iterar rápido. El mock replica el contrato real de la
API (mismas funciones, mismas formas de respuesta, incluida la
validación de superposición de horarios) para que activarlo/desactivarlo
no cambie el comportamiento de las pantallas, solo el origen de los
datos. Queda en el repo por si hace falta desarrollar el frontend sin
levantar la base, pero el default es `false` (API real).

**Bug de huso horario en `horarios-disponibles`, resuelto en el
frontend**: el backend arma los horarios de atención en reloj de
Argentina pero los serializa con sufijo `Z` (UTC) sin convertir de
verdad — es una app de un solo huso horario, nunca se manejó timezone
real. El navegador, al parsear esas fechas, sí aplicaba la conversión
UTC→local, corriendo todo -3hs (un profesional que atendía hasta las 18
se mostraba disponible solo hasta las 15). Se corrigió leyendo esas
fechas con los getters UTC (`getUTCHours`, etc.) en vez de los locales,
tratando los dígitos tal cual llegan del backend. Quedó documentado acá
en vez de tocar el backend porque el resto de la lógica de negocio (que
si es intencional) ya asume ese mismo criterio "sin huso horario real"
de punta a punta.

**Regla de negocio 6 — auto-completado de turnos por fecha**: no hay
login de profesional/staff en la app, así que ninguna transición de
"confirmado" a "completado" se podía disparar a mano (y no tenía
sentido que el propio paciente se auto-marcara la consulta como
completada). Se resolvió en el backend (`service.autoCompletarSiCorresponde`):
un turno "confirmado" cuyo `fecha_hora_fin` ya pasó se considera
completado al leerlo, y se persiste en la base en ese momento (no es
un cálculo solo de lectura).

**Regla de negocio 7 — los turnos se identifican por un código
aleatorio, no por su ID de base**: `GET /turnos/{id}` y
`PATCH /turnos/{id}/estado` usaban el ID secuencial de Postgres (1, 2,
3...) directo en la URL. Como no hay login, cualquiera podía probar
números consecutivos y ver o cancelar turnos ajenos. Se agregó una
columna `codigo` (16 bytes aleatorios, generados con `crypto/rand` en
Go) que es el identificador público real; el ID numérico sigue
existiendo para las foreign keys y como número prolijo para mostrar
("Turno #4"), pero ya no sirve para acceder a nada por sí solo. Se
evaluó agregar login completo para esto, pero es una capa de
complejidad mucho mayor (autenticación, sesiones, recuperación de
contraseña) para un problema que un token no adivinable ya resuelve
razonablemente bien, dado que tampoco hay datos más sensibles que
nombre/DNI/email/horario de una consulta.

**Página 404 y pasada de responsive**: se agregó una ruta catch-all en
el frontend para URLs que no matchean ninguna pantalla, y se revisó el
sitio a 320-375px de ancho (encontró y corrigió un overflow horizontal
real en el formulario de "Mis turnos" — al input le faltaba
`min-width: 0` dentro del contenedor flex) y se achicó el padding del
header/hero/páginas en pantallas chicas.

