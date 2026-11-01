// web/src/lib/character-selector/follow-page-run.ts
// Carries out follow-page.ts's decisions on a guild or character page: reads the account and the
// two pointers, then moves the page the way the BiS follow does (`location.assign`). Loaded on the
// first pointer change (and by the guild landing), not with the page, so the account API stays
// out of the islands' entry chunks.
import { fetchMeOnce, type Me } from '../account/api';
import { parseCurrentCharacter, readCurrentRaw, type CurrentCharacterSwitch } from '../current-character';
import { characterFollowHref, guildFollowHref, pointerCharacterHref } from './follow-page';
import { currentGuildOf } from './session-guild';

/** A failed `/v1/me` reads as signed out: the page stays where it is. */
function sessionOrNull(): Promise<Me | null> {
  return fetchMeOnce().catch(() => null);
}

function place(): { pathname: string; hash: string } {
  return { pathname: window.location.pathname, hash: window.location.hash };
}

/** A choice moves the page as a new history entry; the landing's own redirect replaces itself. */
type Navigate = (href: string) => void;

const assign: Navigate = (href) => window.location.assign(href);

function go(href: string | null, navigate: Navigate): void {
  if (href !== null) navigate(href);
}

/** A guild view after a choice of character (`onCurrentCharacterSwitch`). */
export async function followGuildPage(
  change: CurrentCharacterSwitch,
  navigate: Navigate = assign,
): Promise<void> {
  const me = await sessionOrNull();
  go(
    guildFollowHref({
      ...place(),
      previous: currentGuildOf(me, parseCurrentCharacter(change.previousRaw)),
      next: currentGuildOf(me, parseCurrentCharacter(change.nextRaw)),
    }),
    navigate,
  );
}

/** The guild landing on load: the current character's guild page when it has one. */
export function openCurrentGuild(): Promise<void> {
  const raw = readCurrentRaw();
  return followGuildPage({ previousRaw: raw, nextRaw: raw }, (href) => window.location.replace(href));
}

/** A character page after a choice of character (`onCurrentCharacterSwitch`). */
export async function followCharacterPage(change: CurrentCharacterSwitch): Promise<void> {
  const me = await sessionOrNull();
  go(
    characterFollowHref({
      ...place(),
      previous: pointerCharacterHref(me, parseCurrentCharacter(change.previousRaw)),
      next: pointerCharacterHref(me, parseCurrentCharacter(change.nextRaw)),
    }),
    assign,
  );
}
