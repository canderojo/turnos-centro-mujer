export default function Flor({ className, color = "var(--color-primary)" }) {
  return (
    <svg className={className} viewBox="0 0 60 60" fill="none" aria-hidden="true">
      <g fill={color}>
        <ellipse cx="30" cy="16" rx="8" ry="11" />
        <ellipse cx="30" cy="44" rx="8" ry="11" />
        <ellipse cx="16" cy="30" rx="11" ry="8" />
        <ellipse cx="44" cy="30" rx="11" ry="8" />
      </g>
      <circle cx="30" cy="30" r="7" fill="var(--color-accent)" />
    </svg>
  );
}
