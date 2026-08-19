export default function IlustracionVacia({ className }) {
  return (
    <svg className={className} viewBox="0 0 160 140" fill="none" aria-hidden="true">
      <ellipse cx="80" cy="115" rx="55" ry="10" fill="var(--color-surface-alt)" />
      <g transform="translate(50 20)" opacity="0.55">
        <ellipse cx="30" cy="14" rx="8" ry="12" fill="var(--color-secondary)" />
        <ellipse cx="30" cy="46" rx="8" ry="12" fill="var(--color-secondary)" />
        <ellipse cx="14" cy="30" rx="12" ry="8" fill="var(--color-secondary)" />
        <ellipse cx="46" cy="30" rx="12" ry="8" fill="var(--color-secondary)" />
        <circle cx="30" cy="30" r="7" fill="var(--color-accent)" />
        <path d="M30 58v30" stroke="var(--color-confirmado)" strokeWidth="3" strokeLinecap="round" />
      </g>
      <circle cx="30" cy="40" r="3" fill="var(--color-primary)" opacity="0.5" />
      <circle cx="132" cy="60" r="4" fill="var(--color-secondary)" opacity="0.5" />
    </svg>
  );
}
