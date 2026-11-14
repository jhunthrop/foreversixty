# Officer and raider review of the guild plan (2026-11-14)

Reviewed: 00-facts.md and 01-plan.md. Read as the raid leader and as a raider. No screenshots yet, so this judges the idea and its data, not a layout.

## 1. Does each feature save work every raid week

Loot plan before the raid. Yes, and it is the strongest of the three. Today I open the BiS doc, the loot sheet and the signup list to guess who wants Thursday's drops. With this I would stop building the pre-raid "who needs what" list by hand and stop answering "is this an upgrade for me" in Discord. The council already ranks in two tiers with attendance over four raids, so this is that ranking moved ahead of the raid. It only replaces my sheet if it also shows what each raider already won (loot this tier). The facts say that exists in the council; the plan must show it per row.

Live pull reports. Partly. Companion upload within seconds is real, and deaths with killing blow already exist. I would stop scrubbing Warcraft Logs after the raid to find who died to what, and stop typing wipe notes into the sheet. I would not stop using Logs for parses or deep timelines. The in-game version (addon explains each pull from combat events) needs new combat log work and no companion. That is a second product. I would not count it as part of this feature until it is scoped on its own.

Prep reminder at T-1h. Weakest as written. I stop chasing people in Discord for flasks only if the list is trustworthy. Off game it can only say gear and resistances from Blizzard data plus bags as old as the last companion upload. A raider who last ran the companion on Sunday has bags that are four days old. The relay design (reload the UI, send over GUILD, dedupe several relays) is a lot of moving parts for a result I get from /ready check and one message. It saves me real work only on the "cannot check" list being short.

## 2. What I would not trust, and what would fix it

- Stale gear. Directory gear is up to 72 h old. A raider who swapped trinkets Tuesday shows the old one Thursday and the plan ranks them for a drop they no longer need. Trust comes from the age shown on every row (the plan has this) and from a rule that a row older than the newest raid night, or from the directory, never sorts above a fresh row without a visible flag.
- Missing enchants. Directory gear has none. A "no enchant" claim from there is false. It must read "unknown", never "missing", exactly as consumables do. Only signed-in members and addon exports can be called missing.
- Linear weights against real sims. The gain is a stat-line score delta times DPS per point. That ignores set bonuses, procs, trinket interactions and breakpoints (hit cap especially). I accept it as a rank, I do not accept it as a number. Show "gain" as a rounded band or a rank, not "+41.3 DPS". Label it as an estimate. A top-ranked raider who is actually hit capped wrong will make me look wrong in front of the raid.
- Tanks in score units. A tank "gain 14 score" next to a rogue "gain 38 DPS" cannot be compared and I cannot rank across them. Rank within role only. Tanks get their own column or their own ordering, and the Discord post must not mix the units in one list.
- Healers in HPS. Same issue. HPS from stat weights for a healer is noisy and the best healer for a drop is often the one with the most mana problems, not the most HPS. Keep the label, keep it per role.
- Consumables unknown. The one real Forever log has no combatant info row, so no real boss pull has proven the reading. I would not accept "missing protection potion" naming anyone until at least one real Forever raid log proves it. Until then it says unknown or not shown.
- Relay reload. A button that reloads my UI mid-raid-prep is fine at T-1h and awful during a pull. Ask for it only outside combat and say what it does on the button. Whether ReloadUI needs a hardware event is unverified, so the spike comes first.
- Privacy of public gear. The guild Loot tab reads gear only from addon exports at gear consent. The plan pulls signed-in and directory gear into a guild view without that consent. A raider who opted out of sharing with the guild would find their gear ranked for loot. I would want this blocked by default until the owner decides, and opt-outs (directory_optouts) honored in the guild view too.

What would make me trust it: age and source on every row, "unknown" instead of "missing" where the source cannot know, role-scoped ranking, a weekly check where a real raid's output matches what the officers decided by hand, and consent that matches what the Loot tab already respects.

## 3. Top five objections, ranked

1. Public gear used without guild consent. Fix: guild loot plan uses only addon exports and signed-in members who gave gear consent; directory gear is shown only if the owner decides it may be, behind a per-guild setting, and opt-outs apply.
2. Mixed units in one ranking (score, HPS, DPS). Fix: rank inside tank, healer and damage groups; never a single list; Discord post grouped by role; gains shown as rank or band with "estimate".
3. Naming people after wipes. Fix: pull reports default to officer-only; the member view shows only their own deaths; the Discord post names mechanics and counts, not players, unless an officer chooses to name.
4. Prep reminder promises more than the data. Fix: ship the member-facing in-game check and the officer "cannot check" list first; drop the relay and the reload button from v1 until the spike answers ReloadUI and bag age is under a day; show "bags last seen 4 days ago" plainly.
5. Loot plan skips the fairness half. Fix: every row shows awards this tier, last-4 attendance and MS or OS, and the officer can drop, swap or lock a row in one click; without that I still keep the sheet.

## 4. Where the guild would have drama

- Names after wipes. A post saying "Karl died to Shadow Bolt 3 times" lands in a channel the whole raid reads. People quit over this. Aggregates in public, names for officers.
- Loot ranking visible to raiders. If every raider sees "you gain 9, he gains 31", every drop turns into an argument about the number, and the weights are an estimate. Officers see the ranking; raiders see their own gain and their own status (for example "likely upgrade" or "bench for now") and the plan as a suggestion. The council still decides.
- A pre-raid Discord post that pre-announces winners reads as "decided". Word it as a recommendation, with the council's name on the final call.
- Someone with a private profile or an opt-out appears in the list with a "no data" badge. Keep that neutral and do not call it "slipping".
- Prep naming. A "cannot check" list with names is a shame list. Show counts to the raid and names to officers.
- Consumable accusation when the log simply lacks the sheet. Never state missing without a known reading. One false accusation costs more trust than the feature earns.

## Value test

Would we pay $15 to $25 for this plan as written? The loot plan alone, with fairness and role-scoped ranking, is the addition that flips us from the sheet. Pull reports are a good second tab once consent and naming are settled. The prep reminder is a nice-to-have until it is trustworthy.

Needs data (not today): hit cap and breakpoint awareness in the gain, real combatant info from a Forever log, bag freshness better than the last companion upload, in-game combat event parsing.

## Verdict

CLOSE. The loot plan is worth building as specced, with role-scoped ranking, fairness columns and consent rules. Pull reports follow with officer-only names. The prep relay should wait.
