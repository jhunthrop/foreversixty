<!-- web/src/components/account/CharacterList.svelte -->
<!-- /account's Characters section, extracted from Account.svelte (spec 2026-09-22 §7.2) so
     that file stays clear of the ceiling as the Battle.net import grows its per-character
     detail: realm, level, class colour, and a guild line with a verified pill. A pure render
     of the two props Account.svelte already has in hand -- `me.characters` and
     `me.bnet_imported_at` -- it makes no fetch of its own. -->
<script lang="ts">
  import type { MeCharacter } from '../../lib/account/api';
  import { battlenetStartUrl } from '../../lib/account/api';
  import { characterListCopy } from '../../lib/account/character-list-copy';
  import { characterHref, guildRankLabel, parseCharacterPath, rulesetLabel } from '../../lib/characters';
  import { relativeTime } from '../../lib/dates';
  import { classColorVar } from '../../lib/report/format';
  import { classDisplayName } from '../../lib/sim/spec-label';
  import CharacterHandoffLinks from '../CharacterHandoffLinks.svelte';
  import EmptyState from '../ui/EmptyState.svelte';

  let { characters, bnetImportedAt }: { characters: MeCharacter[]; bnetImportedAt?: string } = $props();

  /** `?next=` on the refresh link (spec §7.2/§7.3): a second Battle.net login re-runs the
   *  import and lands back here with the toast query param. */
  const REFRESH_HREF = battlenetStartUrl('/account?refreshed=1');
</script>

<section class="flex flex-col gap-3" data-testid="account-characters">
  <h2 class="section-title text-[18px]">{characterListCopy.heading}</h2>
  <p class="text-muted text-[13px]">
    {characterListCopy.introLead}
    <a class="text-text underline" href="/addon#paste">{characterListCopy.introPasteLink}</a
    >{characterListCopy.introTail}
  </p>

  {#if bnetImportedAt !== undefined}
    <p class="text-muted text-[13px]" data-testid="bnet-imported">
      {characterListCopy.importedFrom(relativeTime(new Date(bnetImportedAt)))}
      <a href={REFRESH_HREF}>{characterListCopy.refreshFromBattlenet}</a>
    </p>
  {/if}

  {#if characters.length === 0}
    <EmptyState
      message={characterListCopy.empty}
      action={{ label: characterListCopy.refreshFromBattlenet, href: REFRESH_HREF }}
      testid="account-characters-empty"
    />
  {:else}
    <ul class="flex flex-col">
      {#each characters as character (character.key)}
        {@const path = parseCharacterPath(`/character/${character.key}`)}
        <li class="border-line-soft flex min-h-11 flex-wrap items-center gap-3 border-b py-2 text-[14px]">
          <a href={characterHref(character.region, character.ruleset, character.name)}>{character.name}</a>
          <span class="text-muted">
            {#if character.race !== undefined}{character.race}{/if}
            {#if character.class !== undefined}
              <span style:color={classColorVar(character.class)}>{classDisplayName(character.class)}</span>
            {/if}
            {rulesetLabel(character.ruleset)}
            {character.region.toUpperCase()}
            {#if character.realm !== undefined}· {character.realm}{/if}
            {#if character.level !== undefined}· {characterListCopy.levelPrefix(character.level)}{/if}
            {#if character.item_level !== undefined}· {characterListCopy.itemLevelPrefix(
                character.item_level,
              )}{/if}
          </span>
          {#if character.guild !== undefined}
            <span class="text-muted" data-testid="character-guild-line">
              {character.guild.name}
              {#if character.guild.rank !== undefined}· {guildRankLabel(character.guild.rank)}{/if}
              {#if character.guild.verified}
                <span class="text-strong" data-testid="character-guild-verified"
                  >{characterListCopy.verified}</span
                >
              {/if}
            </span>
          {/if}
          {#if path !== null}
            <CharacterHandoffLinks {path} />
          {/if}
        </li>
      {/each}
    </ul>
  {/if}
</section>
