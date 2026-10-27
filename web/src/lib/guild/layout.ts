// web/src/lib/guild/layout.ts
// The Skeleton size each of the four /guild/* views (Guild, GuildJoin, GuildClaim,
// GuildSettings) reserves while loading, read by GuildStatus.svelte's callers. One module
// rather than four inline literals so the four numbers are visible together and cannot
// silently drift apart -- current-character-layout.ts's own reason.
export const GUILD_LOADING = {
  // Header, this-week's-reports panel, roster: the tallest of the four.
  home: { lines: 8, minHeight: 'min-h-[480px]' },
  // A heading and one paragraph of sign-in or join copy: the shortest.
  join: { lines: 2, minHeight: 'min-h-[140px]' },
  // Heading, rules list, current-claim line, one action button.
  claim: { lines: 5, minHeight: 'min-h-[320px]' },
  // Heading, claim-state line, billing panel, the Discord section (its not-connected 140),
  // visibility/threshold controls, invite block.
  settings: { lines: 6, minHeight: 'min-h-[540px]' },
} as const;

/**
 * The #billing section while the billing read loads (launch pricing spec 4, amended in the
 * persona round 2). A first visit reserves the measured ready height of the common states (a
 * paid plan with its seats, Who pays and Records: about 520 px from 768, 760 px on a phone);
 * a return visit reserves the height this guild's section last rendered at that layout, so a
 * warm load holds CLS under 0.02 whatever state the guild is in.
 */
export const GUILD_BILLING_LOADING = { lines: 6, minHeight: 'min-h-[760px] md:min-h-[520px]' } as const;

/** Read and written by GuildBillingSection; one key per guild and per layout (phone or from 768). */
export function guildBillingReserveKey(guildId: number, phone: boolean): string {
  return `fs.guildBilling.reserve.${guildId}.${phone ? 'phone' : 'wide'}`;
}

/** Below this width the section stacks (Tailwind `md`). */
export const GUILD_BILLING_PHONE_MAX_PX = 767;
