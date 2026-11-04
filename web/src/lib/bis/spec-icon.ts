// web/src/lib/bis/spec-icon.ts
// A spec's icon is its `icon` in the generated spec list (`lib/sim/specs.ts`, written by
// `python -m pipeline specs` from the active client's TalentTab and form-spell icons). This
// module only checks that this site build publishes the file.
import { existsSync } from 'node:fs';
import path from 'node:path';
import { fileURLToPath } from 'node:url';

function repoRoot(): string {
  return path.resolve(path.dirname(fileURLToPath(import.meta.url)), '../../../..');
}

/** `icon` when this site build publishes its file (`public/data/<build>/icons/`, which the
 *  sync fills before every build), else `undefined`: a build with a partial icon set (the
 *  fixture) renders the caller's neutral square instead of a broken image. */
export function publishedIcon(build: string, icon: string | undefined): string | undefined {
  if (icon === undefined) return undefined;
  const file = path.join(repoRoot(), 'web/public/data', build, 'icons', `${icon}.webp`);
  return existsSync(file) ? icon : undefined;
}
