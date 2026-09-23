<!-- web/src/components/character/CharacterPortrait.svelte -->
<!-- The one place that draws a character's avatar/class-icon/letter-square (spec 2026-09-23
     §2.1), replacing four copies of the same markup: `CharacterList.svelte`,
     `HomeAccountPanel.svelte` (hero at 40px and chips at 24px), `Character.svelte`'s public
     header. Renders, in priority order: the avatar (`avatar_url`) -> the class icon
     (`classIconUrl`) over the class-coloured letter square (`classSquare`, so a blank/failed
     icon load still shows the letter) -> the letter square alone. -->
<script lang="ts">
  import { classSquare, classIconUrl } from '../../lib/account/character-descriptor';

  export interface PortraitCharacter {
    name: string;
    class?: string;
    avatar_url?: string;
  }

  /** 28 / 36 / 44px, spec 2026-09-23 §2.1 -- one named map, no magic numbers at the call site. */
  const SIZE_CLASS = {
    sm: { box: 'h-7 w-7', letter: 'text-[12px]' },
    md: { box: 'h-9 w-9', letter: 'text-[15px]' },
    lg: { box: 'h-11 w-11', letter: 'text-[18px]' },
  } as const;

  let {
    character,
    size,
    testid,
  }: { character: PortraitCharacter; size: 'sm' | 'md' | 'lg'; testid: string } = $props();

  const box = $derived(SIZE_CLASS[size].box);
  const letterClass = $derived(SIZE_CLASS[size].letter);
  const square = $derived(classSquare(character));
  const classIcon = $derived(classIconUrl(character));
</script>

{#if character.avatar_url !== undefined}
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
