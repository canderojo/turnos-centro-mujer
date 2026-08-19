package models

import "time"

// Profesional representa una fila de la tabla "profesionales".
//
// Las etiquetas `db:"..."` le dicen a sqlx a qué columna de la base
// corresponde cada campo cuando arma un Profesional a partir de una
// fila devuelta por una query (esto es lo que reemplaza al "mapeo
// automático" que haría un ORM, pero acá es explícito y vos lo ves).
//
// Las etiquetas `json:"..."` controlan el nombre de cada campo cuando
// este struct se convierte a JSON para la respuesta HTTP.
type Profesional struct {
	ID                   int       `db:"id" json:"id"`
	Nombre               string    `db:"nombre" json:"nombre"`
	Especialidad         string    `db:"especialidad" json:"especialidad"`
	HoraInicioAtencion   time.Time `db:"hora_inicio_atencion" json:"hora_inicio_atencion"`
	HoraFinAtencion      time.Time `db:"hora_fin_atencion" json:"hora_fin_atencion"`
	DuracionTurnoMinutos int       `db:"duracion_turno_minutos" json:"duracion_turno_minutos"`
	PrecioConsulta       float64   `db:"precio_consulta" json:"precio_consulta"`
}

// Especialidades válidas. Usamos constantes en vez de "strings mágicos"
// sueltos por el código, así el compilador nos avisa si escribimos mal
// alguna en otro archivo.
const (
	EspecialidadDermatologia   = "dermatologia"
	EspecialidadNutricion      = "nutricion"
	EspecialidadEcografia      = "ecografia"
	EspecialidadEndocrinologia = "endocrinologia"
)
