# Spell text scan, 2026-11-10

Lane `lane/spell-text-vars`. Every description the pipeline ships for the active build
1.60.1.70790: talent ranks (`talents/*.json`), spell, racial, enchant and consumable
tooltips (`spelltips/*.json`), item use/equip/proc text (`items/*.json` `effect_text`) and
set bonuses (`sets.json`). `Data.lua` carries no descriptions; the apps read the same
JSON. Reader and defect classes: `data/pipeline/spelltext_lint.py`
(`uv run python -m pipeline spelltext-check --build <build>`).

Every expected string in the tests was read off Wowhead's Forever tooltip
(`nether.wowhead.com/tooltip/spell/<id>?dataEnv=16`) and the build's own client rows.

## The reported bug: Booming Voice

"reduces the Rage cost by of your Shouts" is the client's own text, not a dropped
value. Spell 12321's `Description_lang` in 1.60.1.70790's hotfix push reads
`...reduces the Rage cost by of your Shouts and Howls by $s2%.`, and Wowhead's Forever
tooltip prints rank 5 as "Increases the area of effect of your Shouts by 50% and reduces
the Rage cost by of your Shouts and Howls by 25%." Both values render. The site prints
what the game prints; the lint lists the string as a client typo
(`CLIENT_VERBATIM`).

## Counts

Occurrences (unique strings in brackets). Before: the committed files. After: a local
`normalize` + `spelltips` of 1.60.1.70790 with this lane's renderer (the local run ships
10,742 descriptions against the committed 10,670; it lacks some of the nightly's item
inputs, so the nightly's own `spelltext-check` is the final word).

| Defect class | Before | After |
|---|---:|---:|
| unrendered `$` token | 538 (197) | 0 |
| stray `{` / `}` | 157 (129) | 0 |
| unrendered `@` token | 4 (4) | 0 |
| client colour code (`\|cFF..\|r`) | 3 (3) | 0 |
| double space mid-sentence | 46 (36) | 0 |
| empty substitution ("by of", "for .") | 0 | 0 |
| negative zero | 0 | 0 |
| "0 sec" / "0 yd" | 0 | 0 |
| "0%" (a value lost) | 5 (5) | 0 |
| **total** | **753** | **0** |

Empty substitutions counted 0 before only because the one real case, Earth Shield's
`$<heal>`, still showed its raw token. The `0 sec` class appeared during the lane
(integer rounding of `${$m1/-1000}` on three S03 set bonuses) and was fixed before the
end.

## Leftover token forms before, and the fix for each

| Form (occurrences) | Example | Now |
|---|---|---|
| `$s2`, `$s3` on an effect the spell lacks (277) | Unflinching Morale 1216997 "by $s2%" | the spell's first effect, as Wowhead prints it: "by 5%" |
| `${..}` holding `$d`, `$b1`, a variable or a stat (157) | Rip 1079 `${6*($m1+$b1)}` | worked out: "57 damage" |
| `$<name>` description variables (100+) | Frostbolt 116 `$<frostdamage>` | SpellDescriptionVariables, now fetched: "19 to 21 Frost damage" |
| `$b1`, `$b` points per combo point (28) | Expose Armor 8647 `${$b*-1}` | "90 Armor" |
| `$o1` with no tick period (21) | Food: "Restores $o1 health" | duration / 5.2 sec, matching Wowhead on six foods and drinks |
| reader stats `$AP $SPI $rap $SPFR $SPS $bh $mcrit` (40+) | Dark Sacrifice 1277324 `${400+$SPI}` | "(440 + Spirit) Mana" |
| `$ec1 $ec2 $ecim` (20) | Enchant 744 "armor by $ec1" | SpellItemEnchantment: "armor by 20" |
| chained, negated, level and spaced conditions (10) | `$?s2060[..]?s5185[..][..]` | the last branch a level-60 reader with no talents sees |
| `$L<one>:<many>;` after the first (11) | Nature's Grasp 16689 "$Lstrikes:strike;" | "strikes" |
| `$e`, `$f`, `$v`, `$z`, `$i`, `$*10;F1` (25) | Mana Shield, Thornspike, Mind Soothe, Astral Recall, Execute | amplitude, chain amplitude, max target level, "Hearthstone Location", target cap, multiplier |
| `$@spelltooltip`, `$@spelldesc` of a missing spell (4) | Feral Charge 1238122; Libram of Banishment 211472 | name + text; nothing (Wowhead prints the missing SoD rune as nothing) |
| `$<id>d`, `$c`, `$m1/$t1` the client has no row for (5) | Berserking `$26635d` | `curated/spelltext-gaps.json`, each value from Wowhead with its link |
| "a 0% chance" on a curve talent's spell (5) | Improved Hamstring 12289 | the talent's top rank from `talents/<class>.json`: "15% chance" |
| `\|CFFFF2020..\|R` (3) | Feral Charge 1238122 | the words kept, the colour dropped |

Also fixed on the way: 1.60 effects now carry their Variance and per-level gain
(Revenge 6572: Wowhead's 20.3077 to 23.6923), `$M` reads the top of a roll instead of the
bottom, a value that rounds to nothing keeps its digits (Wowhead's "0.0200 sec"), and a
fractional talent curve value survives (Gore Drinker 0.5%).

## Accepted, with reasons (`CLIENT_VERBATIM`)

- Booming Voice 12321 "by of": the client's typo.
- Earth Shield 974 "heal the shielded target for .": no variable defines `$<heal>`; Wowhead
  prints the same.
- Earthfury Resolve 2P (457540) `$16280s1` and Immoveable Wrath 4P (468067) `$12970s1`,
  `$12970n`: they name Flurry spells no 1.60.1.70790 table has, and Wowhead's Forever
  tooltips print "1%" and "100 + (5 * Combo Points)% for the next 0 swings", which no
  Flurry reads. They need a sourced value in `curated/spelltext-gaps.json`.

## Follow-ups outside this lane

- Three S03 set bonuses store milliseconds that read as fractions of a second (1226378
  "Arcane Power by 0.02 sec", 1226479, 1226501); Wowhead prints the same "0.0200 sec". The
  client data looks wrong, not the renderer.
- Mana Burn 2691's MaxLevel is 70, so a level-60 reader reads 268 to 272 where Wowhead
  (reading level 70) prints 308 to 312.

## Gates

- `data.yml` runs `spelltext-check` after `spelltips`: any defect in the re-rendered build
  fails the run and lists the strings.
- `tests/test_spelltext_contract.py` holds the active build to zero defects and at least
  5,000 descriptions once its manifest records `spelltext_renderer` 2. The committed
  1.60.1.70790 files predate it, so that test skips until the nightly re-renders them.

## Regeneration needed

`data.yml` for 1.60.1.70790 (fetch picks up SpellDescriptionVariables and
SpellXDescriptionVariables, then normalize and spelltips re-render every description and
stamp the manifest), then `bis.yml` so the BiS and addon outputs rebase on it.
