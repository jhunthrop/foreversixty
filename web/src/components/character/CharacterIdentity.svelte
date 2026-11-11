<!-- web/src/components/character/CharacterIdentity.svelte -->
<!-- Portrait + name + descriptor as one block (spec 2026-09-23 §2.2): the account hero band,
     the home hub's hero, and (through CharacterRow) every character-list row draw their name
     and descriptor line through this, instead of four hand-rolled copies.

     `below` (Finding 3, 2026-09-24 landing pass) is one additive slot in the text column,
     after the descriptor line: CharacterRow.svelte renders its build pill and guild line
     through it, so they sit under the name rather than under the whole portrait+name block. -->
<script lang="ts">
  import type { Snippet } from 'svelte';
  import type { MeCharacter } from '../../lib/account/api';
  import { characterDescriptorSegments } from '../../lib/account/character-descriptor';
  import { rulesetLabel } from '../../lib/characters';
  import SupporterChip from '../SupporterChip.svelte';
  import CharacterName from '../ui/CharacterName.svelte';
  import Crest from '../ui/Crest.svelte';

  /** The crest: 28 / 36 / 40 / 44 / 72px, and 64px -> 84px from `lg` for the home hero (home rebuild
   *  spec §3.B.1). */
  const CREST_PX = { sm: 28, md: 36, row: 40, lg: 44, xxl: 72, xl: 84 } as const;
  const XL_CREST_BOX = 'h-16 w-16 lg:h-[84px] lg:w-[84px]';
  /** 14 / 15 / 18px, spec 2026-09-23 §2.2; xl is the home hero's 34px display name (design
   *  system: the home hero h1 exception; spec 2026-10-10-home-signed-in-panel §3.B). */
  const NAME_SIZE = {
    sm: 'text-[14px]',
    md: 'text-[15px]',
    row: 'text-[15px]',
    lg: 'text-[18px]',
    xxl: 'text-[30px]',
    xl: 'text-[34px]',
  } as const;
  /** The xl name is capped at two lines (then an ellipsis) and its box always reserves both,
   *  so a header switch from a one-line to a two-line name never moves what sits below. */
  const XL_NAME_CAP = 'line-clamp-2 leading-[1.1] break-words';
  const XL_NAME_BOX = 'flex min-h-[calc(34px*1.1*2)] flex-col justify-end';

  let {
    character,
    size,
    descriptor,
    href,
    heading = false,
    onNameClick,
    nameTestid,
    descriptorTestid,
    testid = 'character',
    below,
    align = 'center',
    supporter = false,
  }: {
    character: MeCharacter;
    size: 'sm' | 'md' | 'row' | 'lg' | 'xxl' | 'xl';
    descriptor: 'full' | 'realm' | 'none';
    href?: string;
    /** Wraps the name in an `<h1>` -- spec 2026-09-24 Ruling 5: the signed-in home hero's
     *  character name is the page's only heading, the same "name is the h1" rule
     *  Character.svelte's public header already follows. Every other caller omits this. */
    heading?: boolean;
    onNameClick?: (event: MouseEvent) => void;
    nameTestid?: string;
    descriptorTestid?: string;
    testid?: string;
    /** Finding 3, 2026-09-24 landing pass: one additive slot in the text column, after the
     *  descriptor line -- CharacterRow.svelte's own build pill and guild line render
     *  through it, so they sit under the name rather than under the whole portrait+name
     *  block. Every other caller renders nothing extra here, unchanged. */
    below?: Snippet;
    /** The portrait against the text column: centred (every list and chip), or topped against
     *  the name ('start', the character profile: L8.4 fix round 4 item 8), so a taller text
     *  column -- the owner's sync line -- does not move it. */
    align?: 'center' | 'start';
    /** The Supporter chip, inline after the name with the descriptor's gap. Never on the crest. */
    supporter?: boolean;
  } = $props();

  /** After every phrase but the last, on the phrase's own line: a wrapped line ends with
   *  "Level 1 ·" and the next starts clean. The space after it is where a line may break. */
  const SEPARATOR = ' ·';
  const SPACE = ' ';
  // Each phrase stays on one line ("High Order Skyborne Warrior"); the line wraps only between
  // them, never one word per line (2026-11-04, /account at 2000 px).
  const segments = $derived(
    descriptor === 'full'
      ? characterDescriptorSegments(character)
      : descriptor === 'realm'
        ? [rulesetLabel(character.ruleset), character.region.toUpperCase()]
        : [],
  );
  // The heading face only at lg and xl (spec 2026-09-23 §2.2; the face is polish spec
  // 2026-10-28 §4.4's): the account hero band is the one place CharacterIdentity's own name
  // needs it; a row or chip name is never that prominent.
  const nameClass = $derived(
    size === 'lg' || size === 'xxl' || size === 'xl'
      ? `w-fit max-w-full [overflow-wrap:anywhere] ${NAME_SIZE[size]} heading-name${size === 'xl' ? ` ${XL_NAME_CAP}` : ''}`
      : `inline-block max-w-full truncate align-top ${NAME_SIZE[size]} font-semibold`,
  );
  const headingClass = $derived(`m-0 p-0 font-normal${size === 'xl' ? ` ${XL_NAME_BOX}` : ''}`);
</script>

{#snippet nameElement()}
  {#if heading}
    <h1 class={headingClass}>
      <CharacterName
        name={character.name}
        cls={character.class}
        {href}
        onclick={onNameClick}
        weight="inherit"
        class={nameClass}
        testid={nameTestid}
      />
    </h1>
  {:else}
    <CharacterName
      name={character.name}
      cls={character.class}
      {href}
      onclick={onNameClick}
      title={href === undefined ? undefined : character.name}
      weight="inherit"
      class={nameClass}
      testid={nameTestid}
    />
  {/if}
{/snippet}

<div class={`flex min-w-0 gap-3 ${align === 'start' ? 'items-start' : 'items-center'}`}>
  <Crest
    cls={character.class}
    size={CREST_PX[size]}
    class={size === 'xl' ? XL_CREST_BOX : undefined}
    testid={`${testid}-avatar`}
  />
  <div class="flex min-w-0 flex-col gap-0.5">
    {#if supporter}
      <div class="flex min-w-0 flex-wrap items-center gap-x-3 gap-y-1">
        {@render nameElement()}
        <SupporterChip />
      </div>
    {:else}
      {@render nameElement()}
    {/if}
    {#if segments.length > 0}
      <span class="text-muted text-[13px]" data-testid={descriptorTestid}
        >{#each segments as segment, index (index)}{#if index > 0}{SPACE}{/if}<span class="whitespace-nowrap"
            >{segment}{index < segments.length - 1 ? SEPARATOR : ''}</span
          >{/each}</span
      >
    {/if}
    {#if below !== undefined}
      {@render below()}
    {/if}
  </div>
</div>
