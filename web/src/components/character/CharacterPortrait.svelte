<!-- web/src/components/character/CharacterPortrait.svelte -->
<!-- The one place that draws a character's avatar/class-icon/letter-square (spec 2026-09-23
     §2.1), replacing four copies of the same markup: `CharacterList.svelte`,
     `HomeAccountPanel.svelte` (hero at 40px and chips at 24px), `Character.svelte`'s public
     header. Renders, in priority order: the avatar (`avatar_url`) -> the class icon
     (`classIconUrl`) over the class-coloured letter square (`classSquare`, so a blank/failed
     icon load still shows the letter) -> the letter square alone.

     `xl` is its own recipe (home rebuild spec §3.B.1, build review): the mock's hero crest is
     the same CIRCULAR, ringed `ClassCrest` every class-picker on the page uses (home spec
     §5/§6 "one crest language on the page"), never this component's own square avatar-or-
     letter shape -- the account page's hero band is the one caller that still wants the
     square shape, and it stays on `lg`, never `xl`. `xl` is reserved for the home hero
     alone (no other caller uses it today), so this branch only ever replaces that one
     surface. The ringed-crest `<img>`/`<style>` recipe is inlined rather than imported for
     the same reason `HomeSwitchCharacterPanel.svelte`'s own copy is: `ClassCrest.astro`'s
     scoped style cannot be imported into a Svelte file, and a Svelte component tree cannot
     render an Astro component at all. -->
<script lang="ts">
  import { classSquare, classIconUrl } from '../../lib/account/character-descriptor';
  import { classColorVar } from '../../lib/report/format';
  import { classSlugFromName } from '../../lib/report/tree-sizes';
  import { classCrestSrc } from '../../lib/class-crest';

  export interface PortraitCharacter {
    name: string;
    class?: string;
    avatar_url?: string;
  }

  /** 28 / 36 / 44 / 64-84px, spec 2026-09-23 §2.1 (xl added for the home hero, 2026-09-25;
   *  widened 72px -> 84px by the home rebuild spec 2026-09-30 §3.B.1 to match the mock's
   *  measured hero crest exactly -- one token change, one caller, since this hero is the
   *  only `xl` caller today). Review round 1 item 4: below 1024px (Tailwind's `lg`, the
   *  same breakpoint the hero's own two-column grid collapses at) `xl` steps down to 64px
   *  instead of staying fixed at 84px, so the hero crest never crowds a class-coloured
   *  34px name at phone width. */
  const SIZE_CLASS = {
    sm: { box: 'h-7 w-7', letter: 'text-[12px]' },
    md: { box: 'h-9 w-9', letter: 'text-[15px]' },
    lg: { box: 'h-11 w-11', letter: 'text-[18px]' },
    xl: { box: 'h-16 w-16 lg:h-[84px] lg:w-[84px]', letter: 'text-[22px] lg:text-[28px]' },
  } as const;

  let {
    character,
    size,
    testid,
  }: { character: PortraitCharacter; size: 'sm' | 'md' | 'lg' | 'xl'; testid: string } = $props();

  const box = $derived(SIZE_CLASS[size].box);
  const letterClass = $derived(SIZE_CLASS[size].letter);
  const square = $derived(classSquare(character));
  const classIcon = $derived(classIconUrl(character));
</script>

{#if size === 'xl'}
  {#if character.class !== undefined}
    <img
      src={classCrestSrc(classSlugFromName(character.class))}
      alt=""
      loading="lazy"
      decoding="async"
      class={`${box} xl-crest bg-raised shrink-0 rounded-full object-cover`}
      style={`--c: ${classColorVar(character.class)};`}
      data-testid={`${testid}-avatar`}
    />
  {:else}
    <span
      class={`${box} bg-raised border-line inline-block shrink-0 rounded-full border`}
      aria-hidden="true"
      data-testid={`${testid}-avatar-fallback`}
    ></span>
  {/if}
{:else if character.avatar_url !== undefined}
  <img
    class={`${box} shrink-0 rounded-[3px] object-cover`}
    src={character.avatar_url}
    alt=""
    loading="lazy"
    data-testid={`${testid}-avatar`}
  />
{:else}
  <span
    class={`relative flex ${box} shrink-0 items-center justify-center rounded-[3px] ${letterClass} font-bold`}
    style={`background-color: color-mix(in srgb, ${square.color} 22%, transparent); color: ${square.color}`}
    data-testid={`${testid}-avatar-fallback`}
  >
    {square.letter}
    {#if classIcon !== undefined}
      <img
        class={`absolute inset-0 ${box} rounded-[3px] object-cover`}
        src={classIcon}
        alt=""
        loading="lazy"
        data-testid={`${testid}-class-icon`}
      />
    {/if}
  </span>
{/if}

<style>
  /* ClassCrest.astro's own exact recipe (home spec §6), inlined -- see this component's own
     header note for why. No hover/focus state: the home hero's portrait sits in an `<h1>`,
     never inside an anchor, so a hover ring here would never be reachable anyway. */
  .xl-crest {
    box-shadow: 0 0 0 2px color-mix(in srgb, var(--c) 55%, transparent);
  }
</style>
