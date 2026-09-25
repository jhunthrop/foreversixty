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
  import { characterDescriptor } from '../../lib/account/character-descriptor';
  import { rulesetLabel } from '../../lib/characters';
  import { classColorVar } from '../../lib/report/format';
  import CharacterPortrait from './CharacterPortrait.svelte';

  /** 14 / 15 / 18px, spec 2026-09-23 §2.2. */
  const NAME_SIZE = { sm: 'text-[14px]', md: 'text-[15px]', lg: 'text-[18px]' } as const;

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
  }: {
    character: MeCharacter;
    size: 'sm' | 'md' | 'lg';
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
  } = $props();

  const line = $derived(
    descriptor === 'full'
      ? characterDescriptor(character)
      : descriptor === 'realm'
        ? `${rulesetLabel(character.ruleset)} · ${character.region.toUpperCase()}`
        : '',
  );
  // Display font only at lg (spec 2026-09-23 §2.2): the account hero band is the one place
  // CharacterIdentity's own name needs it; a row or chip name is never that prominent.
  const nameClass = $derived(
    `w-fit ${NAME_SIZE[size]} font-semibold${size === 'lg' ? ' [font-family:var(--font-display)]' : ''}`,
  );
</script>

<div class="flex min-w-0 items-center gap-3">
  <CharacterPortrait {character} {size} {testid} />
  <div class="flex min-w-0 flex-col gap-0.5">
    {#if heading}
      <h1 class="m-0 p-0 font-normal">
        {#if href !== undefined}
          <a
            class={nameClass}
            style:color={classColorVar(character.class)}
            {href}
            onclick={onNameClick}
            data-testid={nameTestid}
          >
            {character.name}
          </a>
        {:else}
          <span class={nameClass} style:color={classColorVar(character.class)} data-testid={nameTestid}>
            {character.name}
          </span>
        {/if}
      </h1>
    {:else if href !== undefined}
      <a
        class={nameClass}
        style:color={classColorVar(character.class)}
        {href}
        onclick={onNameClick}
        data-testid={nameTestid}
      >
        {character.name}
      </a>
    {:else}
      <span class={nameClass} style:color={classColorVar(character.class)} data-testid={nameTestid}>
        {character.name}
      </span>
    {/if}
    {#if line !== ''}
      <span class="text-muted text-[13px]" data-testid={descriptorTestid}>{line}</span>
    {/if}
    {#if below !== undefined}
      {@render below()}
    {/if}
  </div>
</div>
