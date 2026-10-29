// web/src/lib/data/build-files.ts
// Build-time reads of a data build's files on disk (data/builds/<build>/), for pages that render
// at build and need an icon the browser would otherwise fetch. Node only.
import { existsSync, readFileSync } from 'node:fs';
import path from 'node:path';
import { fileURLToPath } from 'node:url';

export function repoRoot(): string {
  return path.resolve(path.dirname(fileURLToPath(import.meta.url)), '../../../..');
}

let iconCache: { build: string; icons: Record<string, string> } | undefined;

/** The build's spell id to icon stem table; empty when the build carries none. */
export function spellIcons(build: string): Record<string, string> {
  if (iconCache?.build === build) return iconCache.icons;
  const file = path.join(repoRoot(), 'data/builds', build, 'spellicons.json');
  const icons = existsSync(file) ? (JSON.parse(readFileSync(file, 'utf8')) as Record<string, string>) : {};
  iconCache = { build, icons };
  return icons;
}

/** Whether the build ships the icon stem's `.webp`. */
export function buildIconExists(build: string, stem: string): boolean {
  return existsSync(path.join(repoRoot(), 'data/builds', build, 'icons', `${stem}.webp`));
}
