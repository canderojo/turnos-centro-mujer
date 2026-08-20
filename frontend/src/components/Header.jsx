import { NavLink } from "react-router-dom";
import { ESPECIALIDADES } from "../api/profesionales";
import "./Header.css";

export default function Header() {
  return (
    <header className="header">
      <div className="header-inner">
        <NavLink to="/" className="header-brand">
          <span className="header-brand-mark">
            <svg viewBox="0 0 24 24" fill="none" aria-hidden="true">
              <path d="M12 4v16M4 12h16" stroke="#fff" strokeWidth="2.4" strokeLinecap="round" />
            </svg>
          </span>
          Centro de Salud de la Mujer
        </NavLink>
        <nav className="header-nav">
          <NavLink
            to="/"
            end
            className={({ isActive }) => (isActive ? "header-link active" : "header-link")}
          >
            Profesionales
          </NavLink>
          <NavLink
            to="/mis-turnos"
            className={({ isActive }) => (isActive ? "header-link active" : "header-link")}
          >
            Mis turnos
          </NavLink>
        </nav>
      </div>
      <p className="header-ribbon">
        {ESPECIALIDADES.map((e, i) => (
          <span key={e.value}>
            {i > 0 && <span className="header-ribbon-sep">✦</span>}
            {e.label}
          </span>
        ))}
      </p>
    </header>
  );
}
