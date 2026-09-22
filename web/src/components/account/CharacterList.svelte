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
  import { SECONDARY_BUTTON_FIXED } from '../../lib/planner/styles';
  import { classColorVar } from '../../lib/report/format';
  import CharacterHandoffLinks from '../CharacterHandoffLinks.svelte';

  let { characters, bnetImportedAt }: { characters: MeCharacter[]; bnetImportedAt?: string } = $props();

  /** `?next=` on the refresh link (spec §7.2/§7.3): a second Battle.net login re-runs the
   *  import and lands back here with the toast query param. */
  const REFRESH_HREF = battlenetStartUrl('/account?refreshed=1');
</script>

<section class="flex flex-col gap-3" data-testid="account-characters">
  <h2 class="section-title text-[18px]">{characterListCopy.heading}</h2>

  {#if bnetImportedAt !== undefined}
    <p class="text-muted text-[13px]" data-testid="bnet-imported">
      {characterListCopy.importedFrom(relativeTime(new Date(bnetImportedAt)))}
      <a href={REFRESH_HREF}>{characterListCopy.refreshFromBattlenet}</a>
    </p>
  {/if}

  {#if characters.length === 0}
    <p class="text-muted text-[14px]">{characterListCopy.empty}</p>
    <div class="flex flex-wrap gap-3">
      <a
        class="{SECONDARY_BUTTON_FIXED} border-line-warm-strong text-strong px-4"
        href={REFRESH_HREF}
        data-testid="characters-refresh"
      >
        {characterListCopy.refreshFromBattlenet}
      </a>
      <a
        class="{SECONDARY_BUTTON_FIXED} border-line-warm text-text px-4"
        href="/addon#paste"
        data-testid="characters-paste"
      >
        {characterListCopy.pasteAnExport}
      </a>
    </div>
  {:else}
    <ul class="flex flex-col">
      {#each characters as character (character.key)}
        {@const path = parseCharacterPath(`/character/${character.key}`)}
        <li class="border-line-soft flex min-h-11 flex-wrap items-center gap-3 border-b py-2 text-[14px]">
          <a href={characterHref(character.region, character.ruleset, character.name)}>{character.name}</a>
          <span class="text-muted">
            {#if character.class !== undefined}
              <span style:color={classColorVar(character.class)}>{character.class}</span>
            {/if}
            {rulesetLabel(character.ruleset)}
            {character.region.toUpperCase()}
            {#if character.realm !== undefined}· {character.realm}{/if}
            {#if character.level !== undefined}· Level {character.level}{/if}
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
