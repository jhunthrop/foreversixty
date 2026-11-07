// web/tests/e2e/support/guild-parity/clock.ts
// The guild parity fixtures' own clock and zone (design/mocks/guild_data.py NOW and ZONE): the
// mocks read "Starts in 2d 4h" and "8:00 pm CST" against them, so the build is pinned to the same.
export const GUILD_PARITY_NOW = new Date('2026-12-08T22:00:00Z');
export const GUILD_PARITY_ZONE = 'America/Chicago';
