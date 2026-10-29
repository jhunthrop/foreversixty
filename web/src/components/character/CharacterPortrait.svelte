<!-- web/src/components/character/CharacterPortrait.svelte -->
<!-- The one place a character's picture is drawn, and it is always the circular ringed HD
     class crest (design/DESIGN-SYSTEM.md tenet 7, "one crest language"): the owner raised the
     square class icon / letter square three times (home hero, account chip, the logs page's
     current-character bar, 2026-10-04), so the square paths and the Blizzard avatar thumbnail
     are gone. Fixed sizes delegate to ClassCrestRing.svelte; `xl` (the home hero alone) keeps
     its own responsive 64px -> 84px recipe since the ring component takes one pixel size.
     A character with no class on file yet shows a neutral disc of the same size holding a plain
     class-less silhouette (Forever Score fix round 2 item 4), so it reads as a person, not a hole. -->
<script lang="ts">
  import { classColorVar } from '../../lib/report/format';
  import { classSlugFromName } from '../../lib/report/tree-sizes';
  import { classCrestSrc } from '../../lib/class-crest';
  import ClassCrestRing from './ClassCrestRing.svelte';

  export interface PortraitCharacter {
    class?: string;
  }

  /** 28 / 36 / 44px (a list row is 40, the account hero 72), and 64px -> 84px from `lg` for the home hero (home rebuild spec §3.B.1). */
  const SIZE_PX = { sm: 28, md: 36, row: 40, lg: 44, xxl: 72 } as const;
  const XL_BOX = 'h-16 w-16 lg:h-[84px] lg:w-[84px]';

  let {
    character,
    size,
    testid,
  }: { character: PortraitCharacter; size: 'sm' | 'md' | 'row' | 'lg' | 'xxl' | 'xl'; testid: string } =
    $props();
</script>

{#if character.class === undefined}
  <span
    class={`${size === 'xl' ? XL_BOX : ''} bg-raised border-line text-muted inline-flex shrink-0 items-end justify-center overflow-hidden rounded-full border`}
    style={size === 'xl' ? '' : `width: ${SIZE_PX[size]}px; height: ${SIZE_PX[size]}px;`}
    aria-hidden="true"
    data-testid={`${testid}-avatar-fallback`}
    ><svg
      class="h-[78%] w-[78%]"
      viewBox="0 0 24 24"
      fill="currentColor"
      data-testid={`${testid}-avatar-silhouette`}
      ><circle cx="12" cy="9" r="4.5" /><path d="M3.5 24c0-5.2 3.8-8.5 8.5-8.5s8.5 3.3 8.5 8.5z" /></svg
    ></span
  >
{:else if size === 'xl'}
  <img
    src={classCrestSrc(classSlugFromName(character.class))}
    alt=""
    loading="lazy"
    decoding="async"
    class={`${XL_BOX} xl-crest bg-raised shrink-0 rounded-full object-cover`}
    style={`--c: ${classColorVar(character.class)};`}
    data-testid={`${testid}-avatar`}
  />
{:else}
  <ClassCrestRing characterClass={character.class} size={SIZE_PX[size]} testid={`${testid}-avatar`} />
{/if}

<style>
  /* ClassCrest.astro's exact ring (home spec §6); the hero portrait sits in an <h1>, never
     inside an anchor, so it has no hover state. */
  .xl-crest {
    box-shadow: 0 0 0 2px color-mix(in srgb, var(--c) 55%, transparent);
  }
</style>
