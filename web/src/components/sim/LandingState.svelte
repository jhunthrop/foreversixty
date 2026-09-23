<!-- web/src/components/sim/LandingState.svelte -->
<!-- What a signed-in member sees when they open /sim: their characters, one button each,
     and no form at all until they ask for one.
     Each row carries its own build-source pill (buildSourcePill, build-pill.ts) rather than
     a single blanket footnote: the site now has a real Battle.net-backed source alongside
     the addon export, so "where did this gear come from" is a per-character fact, not a
     lane-wide one. -->
<script lang="ts">
  import type { MeCharacter } from '../../lib/account/api';
  import type { CharacterPath } from '../../lib/characters';
  import { parseCharacterPath } from '../../lib/characters';
  import { simCopy } from '../../lib/sim/copy';
  import { BUSY_CLASS } from '../../lib/ui/busy';
  import { defaultSimState, simSearch, withSimState } from '../../lib/sim/url';
  import CharacterRow from '../character/CharacterRow.svelte';

  let {
    characters,
    busyKey,
    onpick,
    onother,
  }: {
    characters: MeCharacter[];
    busyKey: string | null;
    onpick: (path: CharacterPath) => void;
    onother: () => void;
  } = $props();

  // `MeCharacter.key` is already `<region>/<ruleset>/<slug>` (`characters.ts`'s own
  // `characterKey` shape), so this reuses that module's validated parser -- the same
  // region/ruleset/slug allow-listing `/character/<key>` and `/guild/<key>` already trust --
  // rather than a hand-rolled split and a type cast past it. `/v1/me` is a trusted
  // first-party response, so null is not expected in practice; it is still handled rather
  // than assumed away, the same "validate at the boundary" rule every other source in this
  // lane follows.
  function pathOf(character: MeCharacter): CharacterPath | null {
    return parseCharacterPath(`/character/${character.key}`);
  }

  /**
   * The row's own link. Armory is a loadable source (`store.svelte.ts`'s `bootstrapSource`
   * resolves `armory` through `fromStoredCharacter`, keyed off this same `?source=armory&
   * ref=<key>` pair -- current-character spec, 2026-09-21), so following this URL on its
   * own now loads the character it names. The href still exists so the row is a real link
   * -- copyable, middle-clickable, opens in a new tab -- exactly as the design calls for,
   * not only so a click handler (`follow`, below) can pick it in place.
   */
  function hrefFor(character: MeCharacter): string {
    return `/sim${simSearch(withSimState(defaultSimState(), { source: 'armory', ref: character.key }))}`;
  }

  /** A link that picks the character in place and still opens in a new tab from a middle
   *  click -- the same pattern MechanicsMode.svelte's own `follow` uses for an in-place
   *  link. Fix round 1, Important #1: while anything is busy (the button's own
   *  `disabled={busyKey !== null}` below, unreachable through an `<a>`), this falls through
   *  to the link's own plain navigation instead of picking in place -- an anchor has no
   *  `disabled`, so without this check the row's link was the one entry point that could
   *  still start a race in-place, busy or not. */
  function follow(event: MouseEvent, path: CharacterPath): void {
    if (busyKey !== null) return;
    if (event.button !== 0 || event.metaKey || event.ctrlKey || event.shiftKey || event.altKey) return;
    event.preventDefault();
    onpick(path);
  }
</script>

<section class="mx-[18px] flex flex-col gap-3 md:mx-0" data-testid="sim-landing">
  <h2 class="section-title text-[15px]">{simCopy.yourCharacters}</h2>

  <ul class="border-line bg-raised rounded-panel flex flex-col border px-3">
    {#each characters as character (character.key)}
      {@const path = pathOf(character)}
      <CharacterRow
        {character}
        href={hrefFor(character)}
        onNameClick={(event) => path !== null && follow(event, path)}
        nameTestid={`sim-character-link-${character.key}`}
        pillTestid={`sim-character-build-${character.key}`}
        testid={`sim-character-${character.key}`}
      >
        {#snippet action()}
          <button
            type="button"
            class={`border-line-warm-strong rounded-control text-strong label ml-auto min-h-11 shrink-0 border px-4 disabled:opacity-50 md:min-h-9 ${busyKey === character.key ? BUSY_CLASS : ''}`}
            disabled={busyKey !== null || path === null}
            aria-busy={busyKey === character.key}
            onclick={() => path !== null && onpick(path)}
            data-testid={`sim-pick-${character.key}`}
          >
            {simCopy.simIt}
          </button>
        {/snippet}
      </CharacterRow>
    {/each}
  </ul>

  <!-- task-2-brief.md: this is what a signed-in member reads first on /sim, before they
       have picked a character -- the same scope sentence the Astro shell already carries
       above the fold, repeated here since a member who scrolled straight to their
       character list may never have read the shell's own copy. -->
  <p class="text-muted text-[12px]" data-testid="sim-landing-scope-note">{simCopy.scopeNote}</p>

  <button
    type="button"
    class="text-nav label min-h-11 self-start underline md:min-h-9"
    onclick={onother}
    data-testid="sim-other-character"
  >
    {simCopy.otherCharacter}
  </button>
</section>
