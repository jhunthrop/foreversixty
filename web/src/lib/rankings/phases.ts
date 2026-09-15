// web/src/lib/rankings/phases.ts
// The content phases rankings are bracketed by. The API derives a row's phase from when
// the fight happened and stores it; the site only needs the list to offer as a filter and
// the label to print. The contract names `launch` and `raids-1`; the list grows with the
// data plan's phase table, and an id the API returns that is not here is shown as-is
// (phaseLabel echoes it) rather than hidden.
export const PHASES = [
  { id: 'launch', label: 'Launch' },
  { id: 'raids-1', label: 'Raids, phase 1' },
] as const;

export type Phase = (typeof PHASES)[number]['id'];

export function phaseLabel(id: string): string {
  return PHASES.find((phase) => phase.id === id)?.label ?? id;
}
