import { get, post, patch } from "./client";
import {
  mockCrearTurno,
  mockObtenerTurno,
  mockListarTurnosDePaciente,
  mockCambiarEstadoTurno,
} from "./mockData";

const MOCK = import.meta.env.VITE_USE_MOCK_DATA === "true";

export const ESTADOS = {
  pendiente: "Pendiente",
  confirmado: "Confirmado",
  cancelado: "Cancelado",
  completado: "Completado",
};

// Transiciones permitidas, replicadas del backend (models.Turno.TransicionesPermitidas)
// para poder decidir qué acciones mostrar sin ida y vuelta al servidor.
export const TRANSICIONES_PERMITIDAS = {
  pendiente: ["confirmado", "cancelado"],
  confirmado: ["completado", "cancelado"],
  cancelado: [],
  completado: [],
};

export function crearTurno(datos) {
  if (MOCK) return mockCrearTurno(datos);
  return post("/turnos", datos);
}

// Los turnos se buscan por su "codigo" (un token aleatorio), nunca por
// su "id" numérico — el id es secuencial y por lo tanto adivinable
// (ver decisiones.md, regla de negocio 7).
export function obtenerTurno(codigo) {
  if (MOCK) return mockObtenerTurno(codigo);
  return get(`/turnos/${codigo}`);
}

export function listarTurnosDePaciente({ dni, email }) {
  if (MOCK) return mockListarTurnosDePaciente();
  const param = dni ? `dni=${encodeURIComponent(dni)}` : `email=${encodeURIComponent(email)}`;
  return get(`/turnos?${param}`);
}

export function cambiarEstadoTurno(codigo, estado) {
  if (MOCK) return mockCambiarEstadoTurno(codigo, estado);
  return patch(`/turnos/${codigo}/estado`, { estado });
}
