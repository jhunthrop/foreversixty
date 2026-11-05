// web/src/lib/guild/join-view.ts
// The join card's "Joining as" row (spec 2026-11-05 §3.4 row 3), from the preview's `joins_as`.
// Pure, so the three states are pinned without a page.
import { specName } from './spec-name';
import { guildJoinCopy } from './copy';
import type { InviteJoinsAs, InvitePeekResult } from './api';

export interface JoiningRow {
  /** The crest's class; '' draws the component's neutral ring. */
  crestClass: string;
  label: string;
  /** Null: the row's label is plain, not in class colour. */
  labelClass: string | null;
  sub: string;
}

/** "Level 52 Night Elf Rogue", each part only when known. */
function descriptor(character: InviteJoinsAs): string {
  const parts = [
    character.level === null ? '' : guildJoinCopy.level(character.level),
    character.race ?? '',
    character.class === null ? '' : specName(character.class),
  ];
  return parts.filter((part) => part !== '').join(' ');
}

export function joiningRow(joinsAs: InviteJoinsAs | null | undefined, guildName: string): JoiningRow {
  if (joinsAs === null || joinsAs === undefined) {
    return {
      crestClass: '',
      label: guildJoinCopy.joinWithAccount,
      labelClass: null,
      sub: guildJoinCopy.joinWithAccountSub,
    };
  }
  const crestClass = joinsAs.class ?? '';
  if (!joinsAs.on_roster) {
    return {
      crestClass,
      label: guildJoinCopy.noneOnRoster(guildName),
      labelClass: null,
      sub: guildJoinCopy.noneOnRosterSub(joinsAs.name),
    };
  }
  return {
    crestClass,
    label: guildJoinCopy.joinAs(joinsAs.name),
    labelClass: joinsAs.class,
    sub: descriptor(joinsAs),
  };
}

/** The card holds room for the join error line only while a join is in flight or after one failed,
 *  so a card at rest ends at its buttons (guild round 2: round 1 left about 50 px under them). */
export function holdsErrorRoom(busy: boolean, error: string): boolean {
  return busy || error !== '';
}

/** `You are invited to join {Guild}.` shows only while there is an invite to take: not to a member
 *  already in the guild, and not once the join went through (guild loop round 3: the joined card
 *  still said it above `Open the guild`). */
export function showsInvitedLine(alreadyIn: boolean, joined: boolean): boolean {
  return !alreadyIn && !joined;
}

/** The preview after a join: read again, so the card's `N members on Forever Sixty` counts the
 *  joiner (guild loop round 3 found `1 member` under `You joined the guild.`). A failed read keeps
 *  the guild and drops the count rather than show a stale one. */
export async function previewAfterJoin(
  before: InvitePeekResult,
  reread: () => Promise<InvitePeekResult>,
): Promise<InvitePeekResult> {
  try {
    const fresh = await reread();
    return { ...before, members: fresh.members };
  } catch {
    return { ...before, members: undefined };
  }
}
