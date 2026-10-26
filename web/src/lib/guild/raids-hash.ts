// web/src/lib/guild/raids-hash.ts
// The deep link to one raid night (attendance spec 3.1): `{guildPath}#raids-{report_id}` opens the
// guild page's Raids tab with that night expanded. Tiny on purpose: GuildShell's entry chunk reads it
// to keep the Raids tab for such a hash.

const NIGHT_PREFIX = 'raids-';

/** `{guildPath}#raids-{report_id}`. */
export function nightHref(guildPath: string, reportId: string): string {
  return `${guildPath}#${NIGHT_PREFIX}${encodeURIComponent(reportId)}`;
}

/** The report id a hash names, or null for any other hash. */
export function nightFromHash(hash: string): string | null {
  const fragment = hash.replace(/^#/, '');
  if (!fragment.startsWith(NIGHT_PREFIX) || fragment.length === NIGHT_PREFIX.length) return null;
  try {
    return decodeURIComponent(fragment.slice(NIGHT_PREFIX.length));
  } catch {
    return null;
  }
}

/** The tab id a hash names: a night's hash belongs to the Raids tab. Lower-cased otherwise. */
export function tabIdOfHash(hash: string): string {
  const fragment = hash.replace(/^#/, '');
  return nightFromHash(hash) === null ? fragment.toLowerCase() : 'raids';
}
