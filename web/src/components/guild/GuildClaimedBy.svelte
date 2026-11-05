<!-- web/src/components/guild/GuildClaimedBy.svelte -->
<!-- `Claimed by` + the holder's crest + their character in class colour (spec 2026-11-05 §3.3): the
     game's own Guild Info frame names a guild master this way, never by an account tag. With no
     character on the roster, `Claimed by the guild master`. -->
<script lang="ts">
  import type { ClaimHolder } from '../../lib/guild/api';
  import { guildHomeCopy } from '../../lib/guild/copy';
  import { classColorVar } from '../../lib/report/format';
  import ClassCrestRing from '../character/ClassCrestRing.svelte';

  let { holder, testid = 'guild-claimed-by' }: { holder: ClaimHolder | null | undefined; testid?: string } =
    $props();
</script>

{#if holder === null || holder === undefined}
  <span data-testid={testid}>{guildHomeCopy.claimedByMaster}</span>
{:else}
  <span class="inline-flex items-center gap-1.5" data-testid={testid}>
    {guildHomeCopy.claimedBy}
    <ClassCrestRing characterClass={holder.class ?? ''} size={20} />
    <span class="font-semibold" style:color={classColorVar(holder.class ?? undefined)}>{holder.name}</span>
  </span>
{/if}
