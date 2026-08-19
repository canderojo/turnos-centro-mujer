export default function HeroIlustracion({ className }) {
  return (
    <svg className={className} viewBox="0 0 420 340" fill="none" aria-hidden="true">
      <ellipse cx="230" cy="180" rx="180" ry="150" fill="var(--color-accent-light)" />
      <ellipse cx="120" cy="120" rx="90" ry="80" fill="var(--color-secondary-light)" opacity="0.8" />

      {/* corazón central: cuidado y salud */}
      <g transform="translate(150 90)">
        <path
          d="M60 118C20 92 0 66 0 40 0 18 17 2 38 2c11 0 21 6 22 15C61 8 71 2 82 2c21 0 38 16 38 38 0 26-20 52-60 78Z"
          fill="var(--color-primary)"
        />
        <path d="M46 48h28M60 34v28" stroke="#fff" strokeWidth="6" strokeLinecap="round" />
      </g>

      {/* flor grande derecha */}
      <g transform="translate(300 60) scale(0.9)">
        <ellipse cx="30" cy="14" rx="7" ry="10" fill="var(--color-primary)" />
        <ellipse cx="30" cy="40" rx="7" ry="10" fill="var(--color-primary)" />
        <ellipse cx="13" cy="27" rx="10" ry="7" fill="var(--color-primary)" />
        <ellipse cx="47" cy="27" rx="10" ry="7" fill="var(--color-primary)" />
        <circle cx="30" cy="27" r="6" fill="var(--color-accent)" />
      </g>

      {/* flor chica izquierda */}
      <g transform="translate(55 210) scale(0.6)">
        <ellipse cx="30" cy="14" rx="7" ry="10" fill="var(--color-secondary)" />
        <ellipse cx="30" cy="40" rx="7" ry="10" fill="var(--color-secondary)" />
        <ellipse cx="13" cy="27" rx="10" ry="7" fill="var(--color-secondary)" />
        <ellipse cx="47" cy="27" rx="10" ry="7" fill="var(--color-secondary)" />
        <circle cx="30" cy="27" r="6" fill="var(--color-accent)" />
      </g>

      {/* flor mediana abajo centro */}
      <g transform="translate(190 250) scale(0.7)">
        <ellipse cx="30" cy="14" rx="7" ry="10" fill="var(--color-accent)" />
        <ellipse cx="30" cy="40" rx="7" ry="10" fill="var(--color-accent)" />
        <ellipse cx="13" cy="27" rx="10" ry="7" fill="var(--color-accent)" />
        <ellipse cx="47" cy="27" rx="10" ry="7" fill="var(--color-accent)" />
        <circle cx="30" cy="27" r="6" fill="var(--color-primary)" />
      </g>

      {/* destellos */}
      <circle cx="70" cy="70" r="4" fill="var(--color-primary)" />
      <circle cx="350" cy="180" r="5" fill="var(--color-secondary)" />
      <circle cx="290" cy="270" r="4" fill="var(--color-primary)" />
      <circle cx="30" cy="150" r="3.5" fill="var(--color-accent)" />
    </svg>
  );
}
