<!-- web/src/components/guild/GuildClaimedBy.svelte -->
<!-- `Claimed by` + the holder's crest + their character in class colour (spec 2026-11-05 §3.3): the
     game's own Guild Info frame names a guild master this way, never by an account tag. With no
     character on the roster, `Claimed by the guild master`. -->
<script lang="ts">
  import type { ClaimHolder } from '../../lib/guild/api';
  import { guildHomeCopy } from '../../lib/guild/copy';
  import CharacterName from '../ui/CharacterName.svelte';

  let { holder, testid = 'guild-claimed-by' }: { holder: ClaimHolder | null | undefined; testid?: string } =
    $props();
</script>

{#if holder === null || holder === undefined}
  <span data-testid={testid}>{guildHomeCopy.claimedByMaster}</span>
{:else}
  <span class="inline-flex items-center gap-1.5" data-testid={testid}>
    {guildHomeCopy.claimedBy}
    <CharacterName name={holder.name} cls={holder.class} crestSize={20} />
  </span>
{/if}
