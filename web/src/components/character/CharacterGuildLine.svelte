<!-- web/src/components/character/CharacterGuildLine.svelte -->
<!-- The one guild-line snippet (spec 2026-09-24 §2.1: the hero needs the same line
     CharacterRow.svelte already draws inline for every character list row). Extracted so a
     second hand-rolled copy never appears (identity components' own rule: no private
     copies). -->
<script lang="ts">
  import type { MeCharacterGuild } from '../../lib/account/api';
  import { guildHref, guildRankLabel } from '../../lib/characters';
  import { cameGuildClause } from '../../lib/guild/attendance-copy';
  import { characterListCopy } from '../../lib/account/character-list-copy';

  /**
   * `region`/`ruleset` are the CHARACTER's: a guild lives on its members' region and
   * ruleset, and `MeCharacterGuild` carries neither, so the caller passes them from the
   * character it is drawing. With both, the guild's name is the way to its page -- the
   * homepage hero read "Employed Millennials · Member" with nowhere to click (owner,
   * 2026-09-28: "where do I go to see it?"). Without them the line stays plain text.
   */
  let {
    guild,
    region,
    ruleset,
    testid = 'character-guild',
  }: { guild: MeCharacterGuild; region?: string; ruleset?: string; testid?: string } = $props();

  const href = $derived(
    region !== undefined && ruleset !== undefined ? guildHref(region, ruleset, guild.name) : null,
  );
  /** ` · ` with a no-break space before the dot: a wrapped clause never starts with it. */
  const SEP_TAIL = '\u00a0· ';
</script>

<span class="text-muted text-[13px]" data-testid={`${testid}-line`}>
  {#if href !== null}
    <a {href} class="text-nav" data-testid={`${testid}-link`}>{guild.name}</a>
  {:else}
    {guild.name}
  {/if}
  {#if guild.rank !== undefined}· {guildRankLabel(guild.rank)}{/if}
  {#if guild.verified}
    <span class="text-strong" data-testid={`${testid}-verified`}>{characterListCopy.verified}</span>
  {/if}
  {#if guild.attendance !== undefined}<span class="tabular" data-testid={`${testid}-came`}
      >{SEP_TAIL}<span class="whitespace-nowrap"
        >{cameGuildClause(guild.attendance.present, guild.attendance.nights)}</span
      ></span
    >{/if}
</span>
