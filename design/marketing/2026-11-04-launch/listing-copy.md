# Addon listing copy for launch (CurseForge and Wago)

Built from design/listing/curseforge-description.md and data-addon-description.md, which stay the
source of truth. This file holds the launch-day changes only: a new one-line summary, a new top
section, a refreshed "What's new", and the tracked site link. To ship it, paste these sections
into those two files, re-render with `node design/listing/md2html.mjs`, and paste into each
dashboard. Neither site lets the release workflow set a description.

The tracked links are `/go/curseforge` and `/go/wago`. Both already exist from migration 0066
with target `/`; see [tracked-links.md](tracked-links.md).

## Forever Sixty (the addon)

### Short description (CurseForge summary, Wago tagline)

Current:

> Your next upgrade, your next talent and your Forever Score, on the game's own frames. No setup, no network, nothing in combat.

Launch version (CurseForge allows 200 characters; this is 129):

> Play your class better in Forever. Your next upgrade, your next talent and your Forever Score on the game's own frames. No setup.

Why: the short line is the only text the CurseForge app shows in search, and "Forever" is the
word people search with this week. "No network, nothing in combat" stays in the long description,
where it has room to be read.

### Long description: new top section

Replace everything from `## Play your class better, in game` down to `### Install in two minutes`
with this:

```markdown
## Play your class better, in game

Forever Sixty shows the next thing that makes your character better, right where you would act on
it: the talent window, the item tooltip, the character panel, the raid frame. There is no window
to open and nothing to set up. A new character follows the default build for its spec from level
one. A surface with nothing to say stays out of the way.

- **Leveling gear:** a green arrow on the bag item that beats what you wear, and a tooltip line
  saying by how much.
- **Talents:** your build's points numbered on the talent window, the next one ringed.
- **Your rotation at your level,** on the character panel.
- **Forever Score** on every player tooltip and in chat.
- **Guild and group:** your roster's gear gap and score, and a ready check that names who lacks
  flask, food or a buff.

It is the in-game half of **[foreversixty.gg](https://foreversixty.gg/go/curseforge)**: class
guides, leveling BiS at 20, 30, 40, 50 and 60, a sim, a talent planner and log coaching for World
of Warcraft: Forever. The numbers in game are the site's numbers, from the same sim and the same
stat weights.

### What's new for launch

- **Built for the launch client, 1.60.1.70755.** Item, talent and spell data is regenerated for
  every new client build, usually the same day.
- **No offer you cannot wear.** The minimap line, the bag rows and the counts check level, class
  and your armor and weapon training first. `/fs loot` says why a piece is not for you.
- **Clear upgrade lines.** Hovering the line names the slot and what the piece replaces ("Hands:
  Earthfury Gauntlets replaces Gloves of the Fang, +3.5 DPS").
- **A guild window that makes sense on day one.** With nobody scored yet it says so once and sorts
  by level.
```

For the Wago copy, use the same Markdown with the link changed to
`https://foreversixty.gg/go/wago`.

Keep the rest of the long description as it is (install, the chip, talent marks, tooltips,
commands, "What it never does").

### Footer line

Replace the last line of the description with:

```markdown
**[foreversixty.gg](https://foreversixty.gg/go/curseforge)** · [Get set up](https://foreversixty.gg/setup) · [Discord](https://discord.gg/NZhBtPwvHt) · [Source and changelog](https://github.com/jhunthrop/foreversixty)
```

Use `/go/wago` on Wago. Leave the setup link untracked, so the two store links count once each.

## Forever Sixty Data (the ratings pack)

### Short description

Current:

> The nightly Forever Scores and guild standings from foreversixty.gg, read in game by the Forever Sixty addon. Updated every night.

Launch version (118 characters):

> Forever Scores and guild standings from foreversixty.gg, updated every night. Install it with the Forever Sixty addon.

Why: "Install it with" answers the first question a player has about a data pack: do I need this?

### Long description: one line at the top

Add above `## The Forever Score, in game`:

```markdown
> You need this only with the **Forever Sixty** addon, which installs it for you. On its own it
> does nothing. Scores start to fill in as players upload their first Forever logs on
> [foreversixty.gg](https://foreversixty.gg/go/curseforge).
```

(`/go/wago` on Wago.)

## Owner checks before pasting

- The "What's new" list names client build 1.60.1.70755 (addon 2.0.23). If a newer build ships
  before you paste, change the number.
- Gallery stays as it is: the twelve plates in design/listing/gallery, in the listed order.
