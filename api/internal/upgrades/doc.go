// Package upgrades is the one home of a character's worn-vs-best-in-slot comparison: which
// band a character reads against, how one slot's gain is measured, which slots are already
// the pick, and the order upgrades are listed in. The Discord /char panel (discordbot) and
// the apps' home answer (app) both read it, so an item shows the same gain in both.
//
// The gain rule is the site's (web/src/lib/home/upgrades.ts): a measured tie is no upgrade;
// a worn item the pick lists as an alternative gains the ranker's own measured gap; a weapon
// with no listed alternative has no figure (a stat-weight score cannot see weapon damage);
// otherwise the pick's score less the worn item's stat-weight score, in DPS through the
// band's reference DPS per point. Words and caps on what a surface shows stay with the
// surface.
package upgrades
