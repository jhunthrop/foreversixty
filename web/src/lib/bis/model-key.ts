// web/src/lib/bis/model-key.ts
// The key of one BiS page model (faction × band × preset). Its own module so the browser island
// can read it without page-model.ts, whose builders pull in server-only file readers.
import type { BisPresetId, Faction } from './types';

export function bandModelKey(faction: Faction, band: number, preset: BisPresetId): string {
  return `${faction}-${band}-${preset}`;
}
