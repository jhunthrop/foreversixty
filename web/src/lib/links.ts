// web/src/lib/links.ts
// src/data/links.json ships with PLACEHOLDER values so the site builds before the Discord
// invite and the GitHub repo exist. A deploy must never carry them.
export const LINKS_FILE = 'web/src/data/links.json';
export const PLACEHOLDER = 'PLACEHOLDER';

/** The keys whose link is still a placeholder. */
export function placeholderKeys(links: Record<string, string>): string[] {
  return Object.entries(links)
    .filter(([, value]) => value.includes(PLACEHOLDER))
    .map(([key]) => key);
}

/** Fails the build, naming the file and the keys, if any link is still a placeholder. */
export function assertLinksAreReal(links: Record<string, string>): void {
  const unset = placeholderKeys(links);
  if (unset.length > 0) {
    throw new Error(
      `${LINKS_FILE} still contains ${PLACEHOLDER} links: ${unset.join(', ')}. Replace them before deploying.`,
    );
  }
}
