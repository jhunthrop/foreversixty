# Character lookup typeahead

Owner, 2026-10-15: "we should have typeahead search for the character search". Applies to both
lookup boxes: the home page's full box (`CharacterSearch.astro`) and the header's compact one
(`HeaderSearch.astro`, the bar's panel from 1024 and the phone Menu's row). Reference: the
in-game "Who" window's name filter and Warcraft Logs' character search, which list matches as
you type with class colour and realm.

## 1. The endpoint

`GET /v1/characters/search?q=<text>&region=<region>&ruleset=<ruleset>&limit=<n>`

- `q`: the typed text, trimmed. Under 2 characters (after trim) answers `rows: []`, never 400;
  over 24 characters is cut to 24. Matching is a case-insensitive prefix on the character's
  name (`lower(name) like lower(q) || '%'`), diacritics as stored.
- `region` and `ruleset`: optional filters (the enums the character path uses); absent means
  every region and ruleset.
- `limit`: 1 to 10, default 8.
- Order: an exact name match first, then name ascending, then region `us` before the others
  and ruleset `normal` before the others, so the same name on several realms lists
  predictably.
- Excludes characters whose account is anonymised (the rule the character page applies).
- Body: the API envelope with `rows: [{ key, region, ruleset, name, class, level, spec }]`;
  `class`, `level` and `spec` are null when unknown (the row has `class` from the characters
  table; `level` and `spec` come from the identity source the character page already reads
  when that is one cheap join, else null).
- Public, cached 60 seconds (`httpx.CachePublic`), behind the router's shared read limiter.
  Backed by a migration adding `characters_name_prefix_idx` on `(lower(name) text_pattern_ops)`.

## 2. The box

The input becomes a combobox (`role="combobox"`, `aria-autocomplete="list"`,
`aria-expanded`, `aria-controls` the listbox, `aria-activedescendant` the highlighted row).
Typing two or more characters asks the endpoint after a 150ms pause, with the previous request
aborted; the box's own region and ruleset fields filter when it has them (the home box), the
header box searches every realm.

The listbox opens under the input, absolutely positioned so nothing below moves, `--color-raised`
on `--color-line`, radius 4, the input's width, at most 8 rows of 44px: the class crest
(`classCrestSrc`, 28px, ringed in the class colour), the name in the class colour, then muted
"{Spec} {Class} · {Level}" where known and "{REGION} {Ruleset}" always, so Bow Jackzon on US
PvP and EU Normal read apart. A row whose class is unknown shows the neutral ring and the name
in `--color-strong`.

Keys: ArrowDown and ArrowUp move the highlight (wrapping), Enter opens the highlighted row's
character page, Enter with no highlight submits the typed name as today, Escape closes the
list, Tab closes it and moves on. Pointer: hover highlights, click opens. The list closes on
blur and when the text drops under two characters.

States: while a request is in flight the list stays as it was (no spinner, no flicker); no
match shows one muted row "No character named "{q}" yet" that is not selectable, and Enter
still submits the typed name so an unlogged character can be looked up; a failed request
closes the list quietly and the submit path still works. Every row is a real link
(`a[role="option"]`), so middle-click and keyboard parity hold.

Nothing hydrates on load: the wiring lives in `lib/character-search-dom.ts` beside the form
wiring and runs only once the input gets focus, so the home page's LCP and the header's
budgets do not move.

## 3. Copy (verbatim)

Listbox label: `Characters`. Empty row: `No character named "{q}" yet`. The row's hidden
realm text: `{REGION} {Ruleset}` as `rulesetLabel` writes it.

## 4. Tests

API: handler tests for the length floor, the limit bounds, the order, the filters and the
anonymised exclusion; a migration test if the suite pins migrations. Web: jsdom tests of the
combobox wiring with a fake fetch (debounce, abort, keys, the empty row, the failure path);
e2e in `character-typeahead.spec.ts` routing `**/v1/characters/search*` to a fixture: the home
box and the header box each list the rows with crests, ArrowDown and Enter open the right
page, the empty row appears, Enter on a no-match submits the typed name, no sideways scroll on
the phone; existing `home-character-search.spec.ts` and `handoffs-home.spec.ts` keep passing.
