package service

import (
	"crypto/rand"
	"database/sql"
	"encoding/hex"
	"time"

	"github.com/jmoiron/sqlx"

	"github.com/canderojo/turnos-centro-mujer/backend/internal/models"
	"github.com/canderojo/turnos-centro-mujer/backend/internal/repository"
)

// generarCodigo arma el identificador público de un turno: 16 bytes
// aleatorios (128 bits, no se puede adivinar probando) en hexadecimal.
func generarCodigo() (string, error) {
	bytes := make([]byte, 16)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	return hex.EncodeToString(bytes), nil
}

// CrearTurnoInput son los datos que necesitamos para reservar un
// turno: a qué profesional, cuándo, y quién es el paciente (que puede
// ser nuevo o ya existente — se identifica por DNI, sin login).
type CrearTurnoInput struct {
	ProfesionalID   int
	FechaHoraInicio time.Time
	Nombre          string
	DNI             string
	Email           string
	Telefono        string
}

// CrearTurno aplica las reglas de negocio de la reserva (horario de
// atención, superposición, snapshot de precio) antes de delegar la
// escritura a la capa de repository.
func CrearTurno(db *sqlx.DB, input CrearTurnoInput) (*models.Turno, error) {
	profesional, err := repository.ObtenerProfesional(db, input.ProfesionalID)
	if err == sql.ErrNoRows {
		return nil, ErrProfesionalNoExiste
	}
	if err != nil {
		return nil, err
	}

	if input.FechaHoraInicio.Before(time.Now()) {
		return nil, ErrFechaEnElPasado
	}

	fin := input.FechaHoraInicio.Add(time.Duration(profesional.DuracionTurnoMinutos) * time.Minute)
	if !dentroDeHorarioAtencion(input.FechaHoraInicio, fin, *profesional) {
		return nil, ErrFueraDeHorario
	}

	superpuestoProfesional, err := repository.ExisteSuperposicionProfesional(db, input.ProfesionalID, input.FechaHoraInicio, fin)
	if err != nil {
		return nil, err
	}
	if superpuestoProfesional {
		return nil, ErrSuperposicionProfesional
	}

	paciente, err := repository.BuscarOCrearPaciente(db, input.Nombre, input.DNI, input.Email, input.Telefono)
	if err != nil {
		return nil, err
	}

	superpuestoPaciente, err := repository.ExisteSuperposicionPaciente(db, paciente.ID, input.FechaHoraInicio, fin)
	if err != nil {
		return nil, err
	}
	if superpuestoPaciente {
		return nil, ErrSuperposicionPaciente
	}

	codigo, err := generarCodigo()
	if err != nil {
		return nil, err
	}

	turno := models.Turno{
		Codigo:          codigo,
		ProfesionalID:   input.ProfesionalID,
		PacienteID:      paciente.ID,
		FechaHoraInicio: input.FechaHoraInicio,
		FechaHoraFin:    fin,
		Estado:          models.EstadoPendiente,
		// Snapshot: guardamos el precio vigente del profesional al
		// momento de la reserva. Si el profesional cambia su precio
		// después, los turnos ya reservados no se ven afectados.
		Precio: profesional.PrecioConsulta,
	}

	return repository.CrearTurno(db, turno)
}

// dentroDeHorarioAtencion compara solo la hora del día (no la fecha)
// del turno contra el horario de atención del profesional.
func dentroDeHorarioAtencion(inicio, fin time.Time, profesional models.Profesional) bool {
	minutosInicio := minutosDesdeMedianoche(inicio)
	minutosFin := minutosDesdeMedianoche(fin)
	minutosAtencionInicio := minutosDesdeMedianoche(profesional.HoraInicioAtencion.Time)
	minutosAtencionFin := minutosDesdeMedianoche(profesional.HoraFinAtencion.Time)

	return minutosInicio >= minutosAtencionInicio && minutosFin <= minutosAtencionFin
}

func minutosDesdeMedianoche(t time.Time) int {
	return t.Hour()*60 + t.Minute()
}

// ObtenerTurno trae un turno por su código público (no por su ID
// secuencial — regla de negocio 7, ver decisiones.md) aplicando la
// regla de auto-completado (ver autoCompletarSiCorresponde).
func ObtenerTurno(db *sqlx.DB, codigo string) (*models.Turno, error) {
	turno, err := repository.ObtenerTurnoPorCodigo(db, codigo)
	if err == sql.ErrNoRows {
		return nil, ErrTurnoNoExiste
	}
	if err != nil {
		return nil, err
	}
	return autoCompletarSiCorresponde(db, turno)
}

// ListarTurnosDePaciente es como repository.ListarTurnosDePaciente,
// pero aplicando la misma regla de auto-completado a cada turno.
func ListarTurnosDePaciente(db *sqlx.DB, pacienteID int) ([]models.Turno, error) {
	turnos, err := repository.ListarTurnosDePaciente(db, pacienteID)
	if err != nil {
		return nil, err
	}
	for i := range turnos {
		actualizado, err := autoCompletarSiCorresponde(db, &turnos[i])
		if err != nil {
			return nil, err
		}
		turnos[i] = *actualizado
	}
	return turnos, nil
}

// autoCompletarSiCorresponde resuelve la regla de negocio 6: como no
// hay login de profesional/staff que marque a mano que una consulta
// sucedió, un turno "confirmado" cuyo horario ya pasó (y que no fue
// cancelado) se considera completado automáticamente al leerlo.
func autoCompletarSiCorresponde(db *sqlx.DB, turno *models.Turno) (*models.Turno, error) {
	if turno.Estado == models.EstadoConfirmado && time.Now().After(turno.FechaHoraFin) {
		return repository.ActualizarEstadoTurno(db, turno.ID, models.EstadoCompletado)
	}
	return turno, nil
}

// CambiarEstadoTurno aplica la regla de negocio de la máquina de
// estados: solo se permiten las transiciones definidas en
// models.TransicionesPermitidas (por ejemplo, no se puede pasar de
// "pendiente" directo a "completado").
func CambiarEstadoTurno(db *sqlx.DB, codigo string, nuevoEstado string) (*models.Turno, error) {
	turno, err := repository.ObtenerTurnoPorCodigo(db, codigo)
	if err == sql.ErrNoRows {
		return nil, ErrTurnoNoExiste
	}
	if err != nil {
		return nil, err
	}

	permitidos := models.TransicionesPermitidas[turno.Estado]
	esValida := false
	for _, estado := range permitidos {
		if estado == nuevoEstado {
			esValida = true
			break
		}
	}
	if !esValida {
		return nil, ErrTransicionInvalida
	}

	// ActualizarEstadoTurno adentro sí usa el ID numérico interno (la
	// primary key) — nunca se expone, solo se usa acá, después de haber
	// resuelto el turno correcto a través de su código público.
	return repository.ActualizarEstadoTurno(db, turno.ID, nuevoEstado)
}
