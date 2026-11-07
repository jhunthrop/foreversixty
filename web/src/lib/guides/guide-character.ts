// web/src/lib/guides/guide-character.ts
// The character the guide's answer settled on, for the parts of the page that follow it (the
// BiS lane's band and faction, the leveling slider's level; bug B2: a Battle.net character
// counts, not only an addon or pasted one). The answer publishes once per read; a listener that
// subscribes later (the leveling island hydrates when visible) gets the last one at once.
import type { GuideCharacter } from './guide-answer';

export const MAX_LEVEL = 60;

/** The BiS band a level reads (20, 30, 40, 50, 60), as the BiS page bands it. */
export function bandForLevel(level: number | undefined): number {
  if (level === undefined || level >= MAX_LEVEL) return MAX_LEVEL;
  return Math.max(20, Math.floor(level / 10) * 10);
}

type Listener = (character: GuideCharacter | null) => void;

let last: GuideCharacter | null | undefined;
const listeners = new Set<Listener>();

export function publishGuideCharacter(character: GuideCharacter | null): void {
  last = character;
  for (const listener of listeners) listener(character);
}

/** Calls `listener` with every published character, the last one at once; returns the unsubscribe. */
export function onGuideCharacter(listener: Listener): () => void {
  listeners.add(listener);
  if (last !== undefined) listener(last);
  return () => listeners.delete(listener);
}

/** For tests: forget the last character. */
export function resetGuideCharacter(): void {
  last = undefined;
  listeners.clear();
}
