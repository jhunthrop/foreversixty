// web/src/lib/guides/beta.ts
// Whether the beta's level 30 cap is behind us: launch (the PHASES entry the API files fights
// under) lifts it. The guides prerender, so the nightly build flips the copy after launch.
import { PHASES } from '../rankings/phases';

const LAUNCH_START = PHASES.find((phase) => phase.id === 'launch')!.start;

export function betaOver(now: Date | number): boolean {
  return new Date(now).getTime() >= Date.parse(LAUNCH_START);
}
