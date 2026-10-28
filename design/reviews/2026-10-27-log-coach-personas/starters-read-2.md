# Task 24 review (ba97669a1)

pytest: `uv run --project data pytest -o addopts='' tools/coach` -> 8 passed. Go not run (package does not compile yet).

## A) Spec compliance: APPROVE with notes
- F18: 28 files in data/curated/coach; generator keys on specs.json `spec`, FULL_SPECS excluded, test asserts 25 starters. OK.
- F6: `role` is in front matter from specs.json; test + Go TestAllPersonas compare it to specs.json. OK.
- F19: voice.md carries "Never tell a healer to heal more" as a `- ` rule line; tests strip only that line. OK.
- F26: N/A. The generator reads no build, only the catalogue (no dir sorting anywhere).
- Section order, status/reviewed front matter, drills by role (healer set has no idle/adherence), Voice verbatim: all tested.
- Notes:
  - LOW: brief says never write a number outside Drills. Numbers are spelled out ("fifty percent", "five minute cooldown", "eighty mana") rather than removed. No digits, but it is still numeric content; hunter-survival line "It costs eighty mana and does fifty damage; the client states both." Decide if the validator (§10) is digit-only. The reading is allowed by the test, not by the brief.
  - LOW: sources lists 4 files (adds specs.json); brief said three. Harmless.
  - LOW: §8.1 item 2 "Forever-specific rules" (Survival melee, shaman no dual wield) are absent by design (guide/APL only). Honest, but the starter for hunter-survival says nothing of melee beyond whatever the APL notes carry.
  - Go: TestStarterPersonas_Reviewed is red by design until reviews are recorded (brief mandates it). It will fail CI on every branch until then; make sure the lane does not merge before the reads.

## B) Quality: CHANGES REQUESTED (1 HIGH, 2 MED)
Good: typed, small functions, deterministic, and test_generator_is_deterministic_and_committed_files_match proves committed == generated.

### HIGH: usable_sentence picks a later sentence that depends on a dropped one (tools/coach/persona_starter.py, usable_sentence ~L88-95)
It returns the first sentence that survives the filters, not the lead sentence. When the lead sentence is dropped (parens or jargon words), a dependent follow-up becomes the line.
- rogue-combat.md, APL item 5: source "Sinister Strike builds when Backstab cannot be cast (no dagger in the main hand). Held until 60 energy so a dagger build, which would have cast Backstab at that point, never spends the energy here." Output: "Held until sixty energy so a dagger build ... never spends the energy here." The subject (Sinister Strike) and the rule are gone; reads as nonsense or a false rule. Same item 3 (Backstab builder, dagger) is dropped entirely, so a Combat rogue's builder rule is missing.
- hunter-survival: "It costs eighty mana and does fifty damage; the client states both." (dangling "It").
- shaman-elemental: two lines start "two or more targets: ..." (the spell they refer to is lost).
Fix: use only a note's FIRST sentence; if it is unusable, drop the whole note. Add a test for rogue-combat (no line beginning "Held until").

### MED: load-bearing priority lines silently dropped; the file claims to be "the order ... from the top" (know_section text)
Notes that are empty or fail the jargon filter vanish with no marker. Kept/total: priest-discipline 7/25, paladin-holy 5/14, shaman-restoration 8/18, hunter-survival 3/8, druid-feral 5/9. Examples:
- warlock-affliction drops the Wrack line (parens) yet keeps "Immolate is kept up as a fourth DoT" (only Corruption, Doom, Immolate remain: the count no longer adds up), and drops "Life Tap is only the backstop below 10 percent".
- druid-restoration keeps "the major-cooldown pass above only catches it below five hundred mana" and "paced like Healing Touch" / "at the same pace" while the pass and the pacing line are not in the list (dangling references).
- mage-frost loses the "only with the Fingers of Frost talent" caveat (acceptable: second sentence, not load-bearing).
Fix: reword the intro to "the lines of the rotation data that read cleanly, in order", or better, keep a fixed per-spec hand list; at minimum assert in a test that no kept line starts with a pronoun/lowercase and report dropped counts from --all.

### MED: JARGON regex and "(" ")" blanket drop (JARGON, ~L48)
Parentheses are the most common cause of loss (Wrack, Backstab). A parenthetical-stripping step would keep the lead sentence intact. Words like "see", "search", "dropped" are too broad as whole-sentence kills.

### Voice shared across roles: acceptable
voice.md says "For a healer, ..." conditionally and states the rule once; F19 permits it. Templating by role would remove two lines of healer text from dps/tank files but costs a second source of truth, and the test "voice == voice.md" would need to change. Keep shared; optional follow-up.

### LOW
- `load_json -> object` then iterated in spec_row/spec_slugs: fails mypy strict; return `Any` or type the rows.
- `string(p.Status) == "starter"` guesses Task 22's constant: fine, note it; swap to the constant when Task 22 lands.
- Tests share a first-60-char style; fine.

## Sampled against APL (5 files)
rogue-combat, mage-frost, warlock-affliction, shaman-enhancement, druid-restoration (quotes above). shaman-enhancement and mage-frost were faithful; the other three show the defects above.
