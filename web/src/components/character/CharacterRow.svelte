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
  import { ROW_HOVER } from '../../lib/ui/row';
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
    tags,
    settingsRow = false,
  }: {
    character: MeCharacter;
    descriptor?: 'full' | 'realm' | 'none';
    guildLine?: boolean;
    pillTestid?: string;
    /** Finding 1 (2026-09-24 landing pass): the sim landing rows show no pill at all for a
     *  character with no build -- the action beside it (a "Paste export" link, in place of
     *  Sim) already says so. Every other caller (the account page's rows) keeps today's
     *  muted "No gear yet" text, unaffected by the default `false`. */
    hidePillWhenNoBuild?: boolean;
    href?: string;
    onNameClick?: (event: MouseEvent) => void;
    nameTestid?: string;
    descriptorTestid?: string;
    testid?: string;
    action?: Snippet;
    /** The account row's own pills (Current, Main), shown before the build status. */
    tags?: Snippet;
    /** The account Characters card: a settings row (56 px, 40 px crest, hover ground), not a bordered list item. */
    settingsRow?: boolean;
  } = $props();

  const pill = $derived(buildSourcePill(character.build));
  const showPill = $derived(hasBuild(character) || !hidePillWhenNoBuild);
  const guild = $derived(guildLine ? character.guild : undefined);
</script>

{#snippet statusPill()}
  {#if showPill}
    <span
      class={pill.pillClass === null
        ? 'text-muted text-[12px] whitespace-nowrap'
        : `pill ${pill.pillClass} w-fit whitespace-nowrap`}
      data-testid={pillTestid}
    >
      {pill.label}
    </span>
  {/if}
{/snippet}

{#snippet identity(withPill: boolean)}
  <CharacterIdentity
    {character}
    size={settingsRow ? 'row' : 'md'}
    {descriptor}
    {href}
    {onNameClick}
    {nameTestid}
    {descriptorTestid}
  >
    {#snippet below()}
      {#if withPill}{@render statusPill()}{/if}
      {#if guild !== undefined}
        <CharacterGuildLine
          {guild}
          region={character.region}
          ruleset={character.ruleset}
          testid="character-guild"
        />
      {/if}
    {/snippet}
  </CharacterIdentity>
{/snippet}

{#if settingsRow}
  <!-- The account Characters card (design/specs/2026-11-04-account-characters.md section 3): a
       grid, never one wrapping flex line. The identity owns its column, the build status under
       the descriptor; the tags and the actions sit in their own column at md and up, and on a phone on a second line under the
       name (52 px in: the 40 px crest and its 12 px gap). Nothing in the actions can squeeze
       the name or the descriptor. -->
  <li
    class={`${ROW_HOVER} border-line-soft grid min-h-14 grid-cols-[minmax(0,1fr)] items-center gap-x-[14px] gap-y-2 border-t px-5 py-3 text-[14px] md:grid-cols-[minmax(0,1fr)_auto]`}
    data-testid={testid}
  >
    <div class="min-w-0">{@render identity(true)}</div>
    <div
      class="flex flex-wrap items-center gap-2 pl-[52px] md:justify-end md:pl-0"
      data-testid="character-row-actions"
    >
      {#if tags !== undefined}{@render tags()}{/if}
      {#if action !== undefined}{@render action()}{/if}
    </div>
  </li>
{:else}
  <li
    class="border-line-soft flex min-h-11 flex-wrap items-center gap-3 border-b py-3 text-[14px] last:border-b-0"
    data-testid={testid}
  >
    <!-- basis-full on a phone: the identity takes the whole first line and the actions wrap
         under it. As flex-1 alone it shrank to make room for the unwrappable actions and the
         descriptor fell to one word per line. -->
    <div class="min-w-0 basis-full md:flex-1 md:basis-0">{@render identity(true)}</div>
    {#if action !== undefined}{@render action()}{/if}
  </li>
{/if}
