// web/src/lib/guild/nights.ts
// "1 of 1 raid night", "5 of 8 raid nights": the one attendance figure on the roster, the Loot tab,
// the character guild line and the bot (api guilds attendanceChip). Lifecycle round 8 (F8-3) made
// the rule; guild loop round 3 made it one noun everywhere ("Came 1 of 1 raids" and "1 of 1 raid"
// read beside each other for one night). The noun is singular when the denominator is 1.
export function nightsLabel(present: number, nights: number): string {
  return `${present} of ${nights} raid night${nights === 1 ? '' : 's'}`;
}
