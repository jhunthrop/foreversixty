<!-- web/src/components/character/ClassCrestRing.svelte -->
<!-- The one shared Svelte implementation of ClassCrest.astro's ringed circular crest recipe
     (home spec §5/§6's "one crest language on the page", design/DESIGN-SYSTEM.md tenet 7)
     for callers outside an Astro entry point -- a Svelte component tree cannot render an
     Astro component (Vite's Astro integration only compiles `.astro` files reached from
     another `.astro` file), so every Svelte caller needs its own `<img>`/`<style>` copy of
     the recipe. The header's old account chip and character-switch rows (since replaced by
     the character selector) were the first callers here (owner-reported defect 2026-10-01: the chip drew CharacterPortrait's
     square avatar-or-letter shape instead of the approved mock's circular ringed crest).

     HomeSwitchCharacterPanel.svelte and CharacterPortrait.svelte's own `xl` branch each
     already inline an equivalent copy, added before this component existed -- left as they
     are here (different interactive-state scope in each, and out of this lane's reported
     defects), so this is a second Svelte copy of the recipe, not a third: new callers reach
     for this one first.

     Guild control-centre live-fix round (defect 4): a character with no known class (the
     guild readiness/roster boards' own "class empty" case, e.g. a verified roster
     character with no fight data and -- before a separate API fix -- no export reading
     either) must never resolve to a broken `/icons/hd/crests/.webp` request. An empty
     characterClass renders CharacterPortrait.svelte's own neutral ringed disc instead, at
     this component's own arbitrary pixel size rather than that component's three fixed
     ones -- the one guard every caller gets for free, rather than GuildReadiness.svelte and
     GuildRosterTable.svelte each repeating it. -->
<script lang="ts">
  import { classCrestSrc } from '../../lib/class-crest';
  import { classColorVar } from '../../lib/report/format';
  import { classSlugFromName } from '../../lib/report/tree-sizes';

  let { characterClass, size, testid }: { characterClass: string; size: number; testid?: string } = $props();

  const slug = $derived(classSlugFromName(characterClass));
  const color = $derived(classColorVar(characterClass));
</script>

{#if characterClass === ''}
  <span
    class="bg-raised border-line inline-block shrink-0 rounded-full border"
    style={`width: ${size}px; height: ${size}px;`}
    aria-hidden="true"
    data-testid={testid ? `${testid}-fallback` : undefined}
  ></span>
{:else}
  <img
    src={classCrestSrc(slug)}
    alt=""
    width={size}
    height={size}
    loading="lazy"
    decoding="async"
    class="crest-ring bg-raised shrink-0 rounded-full object-cover"
    style={`--c: ${color}; width: ${size}px; height: ${size}px;`}
    data-testid={testid}
  />
{/if}

<style>
  /* ClassCrest.astro's own exact recipe (home spec §6): a reduced-opacity ring at rest that
     brightens to the full class colour on the wrapping interactive element's hover/focus-
     visible. `summary`/`button` here, not `a`, since this component's callers (the account
     chip's `<summary>`, the menu's character-switch `<button>` rows) are never an anchor,
     unlike ClassCrest.astro's own callers. */
  .crest-ring {
    box-shadow: 0 0 0 2px color-mix(in srgb, var(--c) 55%, transparent);
    transition: box-shadow 120ms ease-out;
  }
  :global(summary:hover) > .crest-ring,
  :global(summary:focus-visible) > .crest-ring,
  :global(button:hover) > .crest-ring,
  :global(button:focus-visible) > .crest-ring {
    box-shadow: 0 0 0 2px var(--c);
  }
  :global(summary:focus-visible) > .crest-ring,
  :global(button:focus-visible) > .crest-ring {
    outline: 2px solid var(--c);
    outline-offset: 2px;
  }
</style>
