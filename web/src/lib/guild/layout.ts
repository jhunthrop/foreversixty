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
  // Heading, claim-state line, billing panel, visibility/threshold controls, invite block.
  settings: { lines: 6, minHeight: 'min-h-[400px]' },
} as const;
