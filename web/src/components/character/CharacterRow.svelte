<!-- web/src/components/character/CharacterRow.svelte -->
<!-- The list row every character list renders (spec 2026-09-23 §2.3): CharacterIdentity plus
     an optional build pill, an optional guild line, and a trailing action -- the exact row
     `CharacterList.svelte` and `LandingState.svelte` both drew, unified into one component.

     Finding 3 (2026-09-24 landing pass): the pill and guild line render through
     CharacterIdentity's own `below` slot, in its text column, rather than as this row's own
     siblings underneath the whole portrait+name block -- so they sit under the name, not
     under the portrait. -->
<script lang="ts">
  import type { Snippet } from 'svelte';
  import type { MeCharacter } from '../../lib/account/api';
  import { buildSourcePill, hasBuild } from '../../lib/account/build-pill';
  import CharacterGuildLine from './CharacterGuildLine.svelte';
  import CharacterIdentity from './CharacterIdentity.svelte';

  let {
    character,
    descriptor = 'realm',
    guildLine = false,
    pillTestid = 'character-build-pill',
    hidePillWhenNoBuild = false,
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
    /** Finding 1 (2026-09-24 landing pass): the sim landing rows show no pill at all for a
     *  character with no build -- the action beside it (a "Paste export" link, in place of
     *  Sim) already says so. Every other caller (the account page's rows) keeps today's
     *  muted "No build yet" text, unaffected by the default `false`. */
    hidePillWhenNoBuild?: boolean;
    href?: string;
    onNameClick?: (event: MouseEvent) => void;
    nameTestid?: string;
    descriptorTestid?: string;
    testid?: string;
    action?: Snippet;
  } = $props();

  const pill = $derived(buildSourcePill(character.build));
  const showPill = $derived(hasBuild(character) || !hidePillWhenNoBuild);
  const guild = $derived(guildLine ? character.guild : undefined);
</script>

<li
  class="border-line-soft flex min-h-11 flex-wrap items-center gap-3 border-b py-3 text-[14px] last:border-b-0"
  data-testid={testid}
>
  <!-- basis-full on a phone: the identity takes the whole first line and the actions wrap
       under it. As flex-1 alone it shrank to make room for the unwrappable actions and the
       descriptor fell to one word per line. -->
  <div class="min-w-0 basis-full md:flex-1 md:basis-0">
    <CharacterIdentity
      {character}
      size="md"
      {descriptor}
      {href}
      {onNameClick}
      {nameTestid}
      {descriptorTestid}
    >
      {#snippet below()}
        {#if showPill}
          <span
            class={pill.pillClass === null ? 'text-muted text-[12px]' : `pill ${pill.pillClass} w-fit`}
            data-testid={pillTestid}
          >
            {pill.label}
          </span>
        {/if}
        {#if guild !== undefined}
          <CharacterGuildLine {guild} testid="character-guild" />
        {/if}
      {/snippet}
    </CharacterIdentity>
  </div>
  {#if action !== undefined}{@render action()}{/if}
</li>
