// web/src/lib/guild/home-card.ts
// The homepage hero's guild card (spec 2026-09-28) reduced to a pure function: given the
// account, the member-gated guild home (or null when it hasn't loaded/failed/refused), and
// the guild's progression figure (or null when that fetch failed), decide what the card
// says. No fetching, no Svelte, no DOM -- HomeGuildCard.svelte owns all of that and calls
// this once it has data, the same "pure view, dumb component" split the rest of this
// codebase already uses for its API-shaped rule modules (lib/home/timeline.ts, for one).
//
// States are checked in exactly the order the spec lists them, first match wins: a
// pending claim outranks an unclaimed one, an unclaimed guild outranks a roster waiting
// for approval, and none of the three ever reach a plain member -- each is gated on the
// viewer being an officer or the leader, so a member always lands on 'steady' regardless
// of the guild's real claim state. That is deliberate: "never show claim state to a plain
// member on this card" (spec).
import type { Me } from '../account/api';
import { guildClaimHref, guildHref, guildSettingsHref } from '../characters';
import { SETUP_NAV_ITEM } from '../nav';
import type { GuildHome } from './api';
import { homeGuildCardCopy } from './copy';

export type HomeGuildCardState =
  'no-guild' | 'claim-pending' | 'claim-unclaimed' | 'waiting-approval' | 'steady' | 'home-failed';

export interface HomeGuildCardAction {
  label: string;
  href: string;
}

export interface HomeGuildCardView {
  state: HomeGuildCardState;
  /** The one body line below the lead (guild name or "No guild yet."). Empty when the
   *  state has none to show (home-failed). */
  line: string;
  /** The one action, or null for a state with none -- there is none today, but the shape
   *  stays optional rather than every caller checking `.label !== ''`. */
  action: HomeGuildCardAction | null;
}

/** `MeGuild.rank` is a plain, possibly-absent string (the API's own shape), not the
 *  narrower `GuildRank` union `GuildRosterRow` carries -- this is the one place that
 *  narrowing happens for this card, same rule `Guild.svelte`'s own `canManage` uses. */
function isOfficer(rank: string | undefined): boolean {
  return rank === 'officer' || rank === 'leader';
}

export function homeGuildCardView(
  me: Me | null,
  home: GuildHome | null,
  progression: { killed: number; total: number } | null,
): HomeGuildCardView {
  const guild = me !== null && me.guilds.length > 0 ? me.guilds[0] : null;

  if (guild === null) {
    return {
      state: 'no-guild',
      line: homeGuildCardCopy.noGuildBody,
      action: { label: SETUP_NAV_ITEM.label, href: SETUP_NAV_ITEM.href },
    };
  }

  const href = guildHref(guild.region, guild.ruleset, guild.name);
  const officer = isOfficer(guild.rank);

  if (home === null) {
    return {
      state: 'home-failed',
      line: '',
      action: { label: homeGuildCardCopy.viewGuild, href },
    };
  }

  if (officer && home.claim.state === 'pending') {
    return {
      state: 'claim-pending',
      line: homeGuildCardCopy.claimPendingLine,
      action: {
        label: homeGuildCardCopy.confirmClaim,
        href: guildClaimHref(guild.region, guild.ruleset, guild.name),
      },
    };
  }

  if (officer && home.claim.state === 'unclaimed') {
    return {
      state: 'claim-unclaimed',
      line: homeGuildCardCopy.unclaimedLine(guild.name),
      action: {
        label: homeGuildCardCopy.claimThisGuild,
        href: guildClaimHref(guild.region, guild.ruleset, guild.name),
      },
    };
  }

  const waiting = home.roster.filter((row) => !row.verified).length;
  if (officer && waiting > 0) {
    return {
      state: 'waiting-approval',
      line: homeGuildCardCopy.waitingLine(waiting),
      action: { label: homeGuildCardCopy.reviewRoster, href },
    };
  }

  const loggedRecently = home.roster.filter((row) => row.logged_recently).length;
  const reportCount = home.reports.length;
  const line =
    progression === null
      ? homeGuildCardCopy.statsLineNoBosses(loggedRecently, reportCount)
      : homeGuildCardCopy.statsLine(loggedRecently, reportCount, progression.killed, progression.total);

  return {
    state: 'steady',
    line,
    action: officer
      ? {
          label: homeGuildCardCopy.guildSettings,
          href: guildSettingsHref(guild.region, guild.ruleset, guild.name),
        }
      : { label: homeGuildCardCopy.viewGuild, href },
  };
}
