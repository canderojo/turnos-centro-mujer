package service

import (
	"database/sql"
	"time"

	"github.com/jmoiron/sqlx"

	"github.com/canderojo/turnos-centro-mujer/backend/internal/models"
	"github.com/canderojo/turnos-centro-mujer/backend/internal/repository"
)

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

	turno := models.Turno{
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

// CambiarEstadoTurno aplica la regla de negocio de la máquina de
// estados: solo se permiten las transiciones definidas en
// models.TransicionesPermitidas (por ejemplo, no se puede pasar de
// "pendiente" directo a "completado").
func CambiarEstadoTurno(db *sqlx.DB, id int, nuevoEstado string) (*models.Turno, error) {
	turno, err := repository.ObtenerTurno(db, id)
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

	return repository.ActualizarEstadoTurno(db, id, nuevoEstado)
}
