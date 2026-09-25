<!-- web/src/components/character/CharacterGuildLine.svelte -->
<!-- The one guild-line snippet (spec 2026-09-24 §2.1: the hero needs the same line
     CharacterRow.svelte already draws inline for every character list row). Extracted so a
     second hand-rolled copy never appears (identity components' own rule: no private
     copies). -->
<script lang="ts">
  import type { MeCharacterGuild } from '../../lib/account/api';
  import { guildRankLabel } from '../../lib/characters';
  import { characterListCopy } from '../../lib/account/character-list-copy';

  let { guild, testid = 'character-guild' }: { guild: MeCharacterGuild; testid?: string } = $props();
</script>

<span class="text-muted text-[13px]" data-testid={`${testid}-line`}>
  {guild.name}
  {#if guild.rank !== undefined}· {guildRankLabel(guild.rank)}{/if}
  {#if guild.verified}
    <span class="text-strong" data-testid={`${testid}-verified`}>{characterListCopy.verified}</span>
  {/if}
</span>
