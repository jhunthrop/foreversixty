<!-- web/src/components/character/HomeSwitchCharacterPanel.svelte -->
<!-- Home rebuild spec §3.B.4: the signed-in hero's right column. Review round 1 item 3:
     every crest on the page is the same circular ringed ClassCrest, so this builds its own
     row markup (name, descriptor, trailing action) around ClassCrest at 36px rather than
     reusing CharacterSwitchList/CharacterRow's own CharacterPortrait (avatar-or-letter-
     square) -- the mock's own row uses a class crest here too, never an avatar photo.

     The "N upgrades" trailing stat (blocked in the pre-contract build on the same worn-gear
     gap §3.B.2/§3.B.3 named) now reads the real comparison, `lib/home/upgrades.ts` over
     `lib/home/upgrades-loader.ts`'s fetch, computed per non-current row once that row's own
     spec/gear/band are all known -- never a fabricated number for a row that cannot be
     computed (no spec, no gear export, or no published BiS band), which simply shows the
     crest/name/descriptor alone, same as before this comparison existed. §3.B.4's own
     ruling that the stat is additional information, not a replacement for the click target,
     means the trailing stat sits inside the SAME clickable element as the rest of the row
     (a `<button>` wrapping crest+name+descriptor+stat) rather than a separate interactive
     child beside a non-interactive label -- the whole row is the switch action, exactly one
     focus stop, with the stat read as part of its own accessible name.

     Mounted by HomeAccountPanel.svelte into index.astro's `home-switch-character-slot`, the
     same dynamic-import-once-signed-in trick that island already uses for its own module,
     so a signed-out page never fetches this component either. -->
<script lang="ts">
  import type { Me, MeCharacter } from '../../lib/account/api';
  import { homePanelCopy } from '../../lib/home-panel-copy';
  import { currentCharacterCopy } from '../../lib/current-character-copy';
  import { classColorVar } from '../../lib/report/format';
  import { classSlugFromName } from '../../lib/report/tree-sizes';
  import { homeHeroLevelRaceClassLine } from '../../lib/account/character-descriptor';
  import { classCrestSrc } from '../../lib/class-crest';
  import { loadBisContextFor } from '../../lib/home/upgrades-loader';
  import { upgradesFor } from '../../lib/home/upgrades';

  let {
    me,
    currentKey,
    onswitch,
  }: { me: Me; currentKey: string | null; onswitch: (character: MeCharacter) => void } = $props();

  /** The hero listed first, matching the mock's own row order -- every other character
   *  follows in the account's own order. */
  const ordered = $derived.by(() => {
    const hero = me.characters.find((c) => c.key === currentKey);
    if (hero === undefined) return me.characters;
    return [hero, ...me.characters.filter((c) => c.key !== currentKey)];
  });

  /** `character.key` -> its own upgrade count, filled in as each row's own band comparison
   *  resolves. Absent (not just loading) for a row that cannot be computed at all (no spec,
   *  no gear export) or whose comparison is still in flight -- the row simply shows no stat
   *  rather than a shimmering placeholder for every character every time this panel mounts. */
  let upgradeCounts = $state<Record<string, number>>({});

  $effect(() => {
    const characters = ordered;
    for (const character of characters) {
      if (character.key === currentKey) continue;
      if (character.build?.gear === undefined) continue;
      void loadBisContextFor(character).then((ctx) => {
        if (ctx === null) return;
        const result = upgradesFor(character, ctx.band, ctx.items);
        upgradeCounts = { ...upgradeCounts, [character.key]: result.upgrades.length };
      });
    }
  });
</script>

<div
  class="bg-raised/85 border-line rounded-panel flex min-h-[168px] min-w-0 flex-col gap-1 border px-[18px] py-[14px] [grid-area:1/1]"
  data-testid="home-switch-character-panel"
>
  <div class="mb-1 flex items-center justify-between">
    <span class="label">{homePanelCopy.switchCharacterLabel}</span>
    <a href="/account#add-character" class="text-nav text-[12px] font-bold tracking-[0.06em] uppercase">
      {homePanelCopy.addOneCharacter}
    </a>
  </div>
  <ul class="flex flex-col" data-testid="current-character-bar-switch-list">
    {#each ordered as character (character.key)}
      {@const isCurrent = character.key === currentKey}
      {@const upgradeCount = upgradeCounts[character.key]}
      <li
        class="crest-row border-line-soft border-b py-2 text-[14px] last:border-b-0"
        data-testid={`current-character-bar-switch-row-${character.key}`}
      >
        {#if isCurrent}
          <div class="flex min-h-11 items-center gap-3">
            {@render rowCrest(character)}
            {@render rowText(character)}
            <span class="text-muted text-[12px]" data-testid="current-character-bar-switch-current">
              {currentCharacterCopy.switchCurrentMarker}
            </span>
          </div>
        {:else}
          <button
            type="button"
            class="flex min-h-11 w-full min-w-0 items-center gap-3 text-left"
            data-testid={`current-character-bar-switch-${character.key}`}
            aria-label={`${currentCharacterCopy.switchAction} to ${character.name}`}
            onclick={() => onswitch(character)}
          >
            {@render rowCrest(character)}
            {@render rowText(character)}
            {#if upgradeCount !== undefined}
              <span
                class="mono text-muted shrink-0 text-[12px]"
                data-testid={`current-character-bar-switch-upgrades-${character.key}`}
              >
                {currentCharacterCopy.switchUpgradesStat(upgradeCount)}
              </span>
            {/if}
          </button>
        {/if}
      </li>
    {/each}
  </ul>
</div>

{#snippet rowCrest(character: MeCharacter)}
  {#if character.class !== undefined}
    <!-- The same circular ringed crest ClassCrest.astro renders (review round 1 item 3:
         every crest on the page reads the same way) -- inlined rather than imported, since
         a Svelte component tree cannot render an Astro component. ux-designer round 2: the
         rest/hover/focus-visible states are ClassCrest's own recipe too (55% ring at rest,
         100% on hover, 2px offset outline on focus), scoped to `.crest-row:hover`/`:focus-
         within` below -- the row's own hover/focus-within is the equivalent interactive
         scope, whether the trailing content is a button (non-current rows) or plain text
         (the current row). -->
    <img
      src={classCrestSrc(classSlugFromName(character.class))}
      alt=""
      width="36"
      height="36"
      loading="lazy"
      decoding="async"
      class="row-crest bg-raised shrink-0 rounded-full object-cover"
      style={`--c: ${classColorVar(character.class)};`}
      data-testid="home-switch-character-crest"
    />
  {:else}
    <span class="bg-raised border-line inline-block h-9 w-9 shrink-0 rounded-full border" aria-hidden="true"
    ></span>
  {/if}
{/snippet}

{#snippet rowText(character: MeCharacter)}
  <span class="flex min-w-0 flex-1 flex-col">
    <span
      class="w-fit truncate [font-family:var(--font-display)] text-[14px] font-semibold"
      style={`color: ${classColorVar(character.class)}`}
    >
      {character.name}
    </span>
    <span class="text-muted truncate text-[12px]">{homeHeroLevelRaceClassLine(character)}</span>
  </span>
{/snippet}

<style>
  /* ux-designer round 2: the identical recipe ClassCrest.astro's own <style> block uses --
     a reduced-opacity ring at rest that brightens to the class colour at full opacity on
     hover/focus-within, plus a 2px offset outline once focus lands inside the row (the
     trailing button, since the row itself has no wrapping anchor to :focus-visible on
     directly). Kept in this component rather than shared with ClassCrest's own <style>,
     since Astro's per-component scoped styles cannot be imported into a Svelte file. */
  .row-crest {
    box-shadow: 0 0 0 2px color-mix(in srgb, var(--c) 55%, transparent);
    transition: box-shadow 120ms ease-out;
  }
  .crest-row:hover .row-crest,
  .crest-row:focus-within .row-crest {
    box-shadow: 0 0 0 2px var(--c);
  }
  .crest-row:focus-within .row-crest {
    outline: 2px solid var(--c);
    outline-offset: 2px;
  }
</style>
