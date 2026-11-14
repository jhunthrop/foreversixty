# Lane L11 (cleanup and release 2.0.48)

L11 changes no surface a player sees beyond L8's; its player-facing words are the 2.0.48 release notes
(`addon/ForeverSixty/CHANGELOG.md`) and the rewritten gallery captions 1, 2, 9, 10 and 11
(`design/listing/curseforge-description.md`), reviewed here.

| Round | wow-player | ai-slop-reviewer |
|---|---|---|
| 1 | FIX: a binding claim, the removed commands, the sources, Rotation's how you played it, tab availability, the instance rule, caption 1's line | FIX: an inventory voice, repeats, qualifiers, captions 9 to 11 |
| 2 | SHIP (both nits checked: the switch's label is `Show the Forever Sixty line`; the 2.0.47 window is the Builds tab already) | FIX: a repeat, Rotation's line, the instance qualifier, caption 2 |
| 3 | | PASS (its optional nit on Rotation's line applied) |

Kept against ai-slop-reviewer: the tab-by-tab "What moved where" (the owner asked for it) and the
instance rule (wow-player; the build never opens Set up in an instance).

Gates: full addon suite and luacheck on the Windows box, 2250 successes, 0 warnings (650bcd842 plus the
notes-only changes after it); a Locale guard (`tests/locale_keys_spec.lua`) and a one-window guard in
`tests/widgets_spec.lua` added.
