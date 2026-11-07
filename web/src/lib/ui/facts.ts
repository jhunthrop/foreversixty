// web/src/lib/ui/facts.ts
// The props of the two list primitives of the sim round (sim spec 2026-11-06 §4.5, ruling S8):
// Facts (a key/value list) and IconList (a wrapping icon + name list).
import type { Snippet } from 'svelte';

/** One Facts pair. */
export interface Fact {
  id: string;
  label: string;
  /** The value as text; `content` replaces it with markup (a control in the More settings sheet). */
  value?: string;
  content?: Snippet;
  /** Literal text a player types (a slash command): drawn as `Code` in place of the value. */
  code?: string;
  /** A small muted line under the value ("2:24 to 3:36"). */
  sub?: string;
  /** The value itself is quiet (an engine build string). */
  muted?: boolean;
  testid?: string;
}

/** One IconList entry. */
export interface IconListItem {
  id: string;
  name: string;
  /** The icon's URL; null keeps the socket. */
  icon: string | null;
  /** No effect on this spec: muted, the icon at 62%. */
  dim?: boolean;
  /** A game-tooltip key (`spell:<id>`). */
  tip?: string;
}
