<!-- web/src/components/character/CharacterRow.svelte -->
<!-- The list row every character list renders (spec 2026-09-23 §2.3): CharacterIdentity plus
     an optional build pill, an optional guild line, and a trailing action -- the exact row
     `CharacterList.svelte` and `LandingState.svelte` both drew, unified into one component. -->
<script lang="ts">
  import type { Snippet } from 'svelte';
  import type { MeCharacter } from '../../lib/account/api';
  import { buildSourcePill } from '../../lib/account/build-pill';
  import { characterListCopy } from '../../lib/account/character-list-copy';
  import { guildRankLabel } from '../../lib/characters';
  import CharacterIdentity from './CharacterIdentity.svelte';

  let {
    character,
    descriptor = 'realm',
    guildLine = false,
    pillTestid = 'character-build-pill',
    href,
    onNameClick,
    nameTestid,
    descriptorTestid,
    testid,
    action,
  }: {
    character: MeCharacter;
    descriptor?: 'full' | 'realm' | 'none';
    guildLine?: boolean;
    pillTestid?: string;
    href?: string;
    onNameClick?: (event: MouseEvent) => void;
    nameTestid?: string;
    descriptorTestid?: string;
    testid?: string;
    action?: Snippet;
  } = $props();

  const pill = $derived(buildSourcePill(character.build));
  const guild = $derived(guildLine ? character.guild : undefined);
</script>

<li
  class="border-line-soft flex min-h-11 flex-wrap items-center gap-3 border-b py-3 text-[14px] last:border-b-0"
  data-testid={testid}
>
  <div class="flex min-w-0 flex-1 flex-col gap-0.5">
    <CharacterIdentity
      {character}
      size="md"
      {descriptor}
      {href}
      {onNameClick}
      {nameTestid}
      {descriptorTestid}
    />
    <span
      class={pill.pillClass === null ? 'text-muted text-[12px]' : `pill ${pill.pillClass} w-fit`}
      data-testid={pillTestid}
    >
      {pill.label}
    </span>
    {#if guild !== undefined}
      <span class="text-muted text-[13px]" data-testid="character-guild-line">
        {guild.name}
        {#if guild.rank !== undefined}· {guildRankLabel(guild.rank)}{/if}
        {#if guild.verified}
          <span class="text-strong" data-testid="character-guild-verified">{characterListCopy.verified}</span>
        {/if}
      </span>
    {/if}
  </div>
  {#if action !== undefined}{@render action()}{/if}
</li>
