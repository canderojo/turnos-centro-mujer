import Flor from "./illustrations/Flor";
import { ESPECIALIDADES } from "../api/profesionales";
import "./Footer.css";

export default function Footer() {
  return (
    <footer className="footer">
      <div className="footer-inner">
        <div className="footer-brand">
          <span className="footer-brand-mark">
            <Flor color="#fff" />
          </span>
          <div>
            <p className="footer-brand-nombre">Centro de Salud de la Mujer</p>
            <p className="footer-brand-tagline">Turnos online, sin vueltas.</p>
          </div>
        </div>

        <div className="footer-col">
          <p className="footer-col-title">Especialidades</p>
          <ul>
            {ESPECIALIDADES.map((e) => (
              <li key={e.value}>{e.label}</li>
            ))}
          </ul>
        </div>

        <div className="footer-col">
          <p className="footer-col-title">Atención</p>
          <ul>
            <li>Lunes a viernes</li>
            <li>Turnos por especialidad</li>
            <li>Confirmación por email</li>
          </ul>
        </div>
      </div>
      <p className="footer-nota">
        Proyecto académico — Ingeniería del Software 3, UCC.
      </p>
    </footer>
  );
}
