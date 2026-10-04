# Would I run this? — wow-player — 2026-10-04

Reviewed as a player deciding whether to install, not as a rubric review. Boards:
`addon-{tooltip,overview,talents,gear}.png`. Spec: `design/specs/2026-10-04-addon.md`. Worked
character: Obnoxious Yell, lvl 23 Human Fury Warrior.

## 1. Verdict

**Install and keep.** Leveling 1-60 and at 60 raiding, if the must-fix list in the design spec
(§10, and the blocker in my mock review) actually lands before release — see §3 below, because
one of those bugs is sitting un-fixed on the board I was handed today.

## 2. Three moments it earns its place

1. **I loot a green at level 14 in Westfall.** I hover it in my bag: "Upgrade for chest: +12 by
   our weights," the BiS row shows the real pick with its source line. Pawn gives me the verdict
   *if* I already imported a weight string from somewhere; AtlasLoot gives me a static BiS list
   with no comparison to what's in my hand. Neither does both at once, scored, in the tooltip,
   for a ruleset (Forever) that generic Classic data doesn't know. This is the one thing I'd
   actually miss if I uninstalled it.
2. **I hit level 20 and the rotation toast fires.** "Level 20: Sunder Armor opens your rotation
   now." A brand-new alt gets told what changed without alt-tabbing to a guide. Questie doesn't
   touch rotation; WeakAuras could be built to do this but nobody builds it for a fresh level-20
   character. Real value for a novice, near-zero for me on my 9th warrior — but it costs me
   nothing to have it fire once.
3. **I come out of a dungeon and open the Gear tab.** "11 of 15 planned pieces equipped," Legs
   still blank. AtlasLoot shows me the static list; it has never once told me which slots I've
   already filled versus which I'm still missing, let alone whether something already in my bag
   beats what's planned. This is the addon doing bookkeeping I currently do by eye against a
   Wowhead tab.

## 3. Three reasons I'd turn it off, most likely first

1. **The Overview board I was just shown contradicts itself.** "TALENT POINTS" says `Fury 11`
   two inches from "Your build" saying `7 of 11 points taken`. My real character pane never shows
   a plan next to reality labeled the same way — this is the kind of number I'd screenshot to
   guild chat as "the addon is lying to me," and it's unresolved on the current board, not a past
   bug. One bad first impression here and I uninstall before giving it a second chance.
2. **Frame fatigue.** I already run a WeakAuras next-point tracker, a Details! window, ElvUI, and
   Questie. A fifth movable frame plus a toast plus a minimap button is a lot of new real estate
   for a player who already solved "what do I spend my next point on" years ago. If it doesn't
   let me fold the tracker into my existing WeakAuras group, raiders drop it first.
3. **Early-band emptiness reads as broken, not honest.** Trinket slots and some bands show no
   pick at all. I now know (Forever fact) that's correct — no sub-41 trinket has a physical DPS
   stat and ten are genuinely unsourced — but a player who doesn't know that nuance sees a blank
   row where Wowhead always shows *something*, even a "World drop," and assumes the addon gave up.

## 4. What would make me tell my guild to install it

Nothing generic — Wowhead, Pawn, and AtlasLoot already cover "good item, bad item" well enough
for vanilla rules. What they cannot do is **know Forever's own ruleset**: refuse to suggest a
second weapon for my enhancement shaman, gate mail/plate picks by the real level-40 line, score a
Troll Warlock or Dwarf Shaman build correctly, and tie the BiS pick to this client's own crafted
epics (Torch of Light, Heart of the Mountain). That's the pitch no other addon can make. Paired
with the planner round trip — "your top 5 upgrades from last night's sim run" landing in my
in-game inbox, dated, sourced — that closes a loop nothing else in the genre does: site simulates,
addon tells me what to go get.

## 5. Cut one, add one

- **Cut:** the Guild tab's officer-facing claim/pending-approval state on the default surface —
  it's a different app (a loot council tool) glued onto a personal gear helper; put it behind an
  officer-only toggle instead of showing to every member.
- **Add first:** Wave C's typed inbox upgrade message (`inboxUpgradeLine`) — "your top upgrade:
  <item> for <slot>, +N, from your sim" — because it's the one feature that makes the site's
  simulator show up in game without me doing anything, which is the whole reason to keep this
  addon enabled instead of just bookmarking foreversixty.gg.
