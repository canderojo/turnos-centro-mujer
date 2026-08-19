import { NavLink } from "react-router-dom";
import Flor from "./illustrations/Flor";
import "./Header.css";

export default function Header() {
  return (
    <header className="header">
      <div className="header-inner">
        <NavLink to="/" className="header-brand">
          <span className="header-brand-mark">
            <Flor />
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
    </header>
  );
}
