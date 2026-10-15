// web/src/lib/character-selector/session-guild-dom.ts
// How the selector island hands the session's guild to the header's Guild panel. Two
// channels, because their scripts load in no fixed order: the island stamps the guild on the
// header element (so a header script that wires later still finds it) and dispatches
// SESSION_GUILD_EVENT (so a header already wired follows a later change).
import { SESSION_GUILD_EVENT } from './config';
import type { SessionGuild } from './session-guild';

const HEADER_SELECTOR = '[data-site-header]';
const STAMP = 'sessionGuild';

export function publishSessionGuild(guild: SessionGuild | null): void {
  const header = document.querySelector<HTMLElement>(HEADER_SELECTOR);
  if (header !== null) header.dataset[STAMP] = JSON.stringify(guild);
  window.dispatchEvent(new CustomEvent<SessionGuild | null>(SESSION_GUILD_EVENT, { detail: guild }));
}

/** The guild the island last published, or null when nothing (or nothing readable) is stamped. */
export function readPublishedSessionGuild(header: HTMLElement): SessionGuild | null {
  const raw = header.dataset[STAMP];
  if (raw === undefined) return null;
  try {
    const parsed: unknown = JSON.parse(raw);
    return isSessionGuild(parsed) ? parsed : null;
  } catch {
    return null;
  }
}

function isSessionGuild(value: unknown): value is SessionGuild {
  return (
    typeof value === 'object' &&
    value !== null &&
    typeof (value as SessionGuild).name === 'string' &&
    typeof (value as SessionGuild).href === 'string'
  );
}
