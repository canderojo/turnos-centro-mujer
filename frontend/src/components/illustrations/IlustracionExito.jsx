export default function IlustracionExito({ className }) {
  return (
    <svg className={className} viewBox="0 0 160 120" fill="none" aria-hidden="true">
      <circle cx="80" cy="60" r="46" fill="var(--color-accent-light)" />
      <g transform="translate(56 36)">
        <path
          d="M24 47C8 37 0 25 0 15 0 7 6.5 0 15 0c4.5 0 8.5 2.3 9 6C24.5 2.3 28.5 0 33 0c8.5 0 15 7 15 15 0 10-8 22-24 32Z"
          fill="var(--color-primary)"
        />
      </g>
      <path
        d="M62 60l12 12 24-26"
        stroke="#fff"
        strokeWidth="5"
        strokeLinecap="round"
        strokeLinejoin="round"
      />
      <circle cx="24" cy="30" r="4" fill="var(--color-secondary)" />
      <circle cx="140" cy="40" r="5" fill="var(--color-secondary)" />
      <circle cx="130" cy="90" r="3.5" fill="var(--color-primary)" />
      <circle cx="20" cy="90" r="3.5" fill="var(--color-accent)" />
    </svg>
  );
}
