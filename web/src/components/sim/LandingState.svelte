<!-- web/src/components/sim/LandingState.svelte -->
<!-- What a signed-in member sees when they open /sim: their characters, one action each,
     and no form at all until they ask for one.
     Each row carries its own build-source pill (buildSourcePill, build-pill.ts) rather than
     a single blanket footnote: the site now has a real Battle.net-backed source alongside
     the addon export, so "where did this gear come from" is a per-character fact, not a
     lane-wide one.

     2026-09-24 landing pass (owner-approved UX review), Finding 1: the action follows the
     character's own build state -- Sim for a character with one, a "Paste export" link in
     its place for a character without, and the row's pill is hidden then too (the action
     already says it). Finding 2: a pick that fails for want of a build (the API's own
     sim-input 404, as opposed to the older "no race recorded" refusal) shows its own alert
     here, between the list and "Sim something else" -- SimView.svelte's `failedKey` names
     which row. Finding 4: rows use the full descriptor, same as the account page's list.
     Finding 5: the scope note that used to repeat here is gone; ScopeNote.astro's own two
     sentences above the island are the only copy of it now. -->
<script lang="ts">
  import type { MeCharacter } from '../../lib/account/api';
  import { hasBuild } from '../../lib/account/build-pill';
  import type { CharacterPath } from '../../lib/characters';
  import { parseCharacterPath } from '../../lib/characters';
  import { landingCopy } from '../../lib/sim/landing-copy';
  import { simCopy } from '../../lib/sim/copy';
  import { BUSY_CLASS } from '../../lib/ui/busy';
  import { armorySimHref } from '../../lib/sim/url';
  import CharacterRow from '../character/CharacterRow.svelte';

  let {
    characters,
    busyKey,
    failedKey = null,
    onpick,
    onother,
  }: {
    characters: MeCharacter[];
    busyKey: string | null;
    /** Finding 2: the key of the character whose pick just failed for want of a build --
     *  set only for that one failure, never for the race refusal, which keeps its own hint
     *  where it already was (SimView.svelte's `sim-landing-message` paragraph). */
    failedKey?: string | null;
    onpick: (path: CharacterPath) => void;
    onother: () => void;
  } = $props();

  const failedCharacter = $derived(
    failedKey === null ? null : (characters.find((character) => character.key === failedKey) ?? null),
  );

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
    return armorySimHref(character.key);
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
        descriptor="full"
        hidePillWhenNoBuild
        href={hrefFor(character)}
        onNameClick={(event) => path !== null && follow(event, path)}
        nameTestid={`sim-character-link-${character.key}`}
        pillTestid={`sim-character-build-${character.key}`}
        testid={`sim-character-${character.key}`}
      >
        {#snippet action()}
          {#if hasBuild(character)}
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
          {:else}
            <!-- Finding 1: a character with no build gets no Sim button -- there is nothing
                 to sim yet -- and this link in its place. -->
            <a
              class="text-nav label ml-auto inline-flex min-h-11 shrink-0 items-center underline md:min-h-9"
              href={landingCopy.pasteExportHref}
              data-testid={`sim-paste-${character.key}`}
            >
              {landingCopy.pasteExport}
            </a>
          {/if}
        {/snippet}
      </CharacterRow>
    {/each}
  </ul>

  {#if failedCharacter !== null}
    <!-- Finding 2: the sim-input 404 -- this character has no build recorded at all, as
         opposed to the "no race recorded" refusal SimView.svelte's own paragraph still
         handles below (never this one). The design system's LoadError shape (role="alert",
         a muted line, no retry action here since the remedy is one of the two links, not a
         re-fetch of the same 404), built by hand rather than through that component: its
         `message` prop is plain text and cannot carry the two links this alert needs. -->
    <div
      class="flex min-h-11 flex-wrap items-center gap-3 text-[14px]"
      role="alert"
      data-testid="sim-landing-build-missing"
    >
      <span class="text-muted">
        {landingCopy.buildMissingLead(failedCharacter.name)}
        <a class="text-text underline" href={landingCopy.pasteExportHref}
          >{landingCopy.buildMissingPasteLink}</a
        >,
        {landingCopy.buildMissingMiddle}
        <a class="text-text underline" href={landingCopy.buildMissingAccountHref}
          >{landingCopy.buildMissingAccountLink}</a
        >.
      </span>
    </div>
  {/if}

  <button
    type="button"
    class="text-nav label min-h-11 self-start underline md:min-h-9"
    onclick={onother}
    data-testid="sim-other-character"
  >
    {simCopy.otherCharacter}
  </button>
</section>
