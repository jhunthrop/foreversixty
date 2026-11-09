<!-- web/src/components/guild/loot/LootCandidates.svelte -->
<!-- The rail for the picked drop, officers (web layout spec 2026-11-05 §7.2): "Candidates for
     {item}", ranked by the API (BiS-list matches with their gain first, then item level), each a
     LootCandidate with votes and Award. An awarded drop leads with its decision record and keeps
     the candidates read-only. A refused write says so once, under the head. -->
<script lang="ts">
  import type { LootCouncilDrop, MsOs, Vote } from '../../../lib/guild/loot-council-api';
  import { CLOSED_LINE, NO_CANDIDATES } from '../../../lib/guild/loot-council-view';
  import LoadError from '../../ui/LoadError.svelte';
  import Section from '../../ui/Section.svelte';
  import Skeleton from '../../ui/Skeleton.svelte';
  import DecisionRow from './DecisionRow.svelte';
  import LootCandidate from './LootCandidate.svelte';

  let {
    itemName,
    drop,
    status,
    build,
    shipsIcon,
    timeZone,
    canAct,
    describedBy,
    busy,
    confirmingKey,
    notice,
    onVote,
    onAsk,
    onAward,
    onCancel,
    onUndo,
    onRetry,
  }: {
    itemName: string;
    drop: LootCouncilDrop | null;
    status: 'loading' | 'ready' | 'failed';
    build: string;
    shipsIcon: boolean;
    timeZone?: string;
    canAct: boolean;
    describedBy?: string;
    busy: boolean;
    confirmingKey: string | null;
    /** A refused or failed write, said once under the head. */
    notice: string;
    onVote: (characterKey: string, vote: Vote | null) => void;
    onAsk: (characterKey: string) => void;
    onAward: (characterKey: string, msOs: MsOs | null, reason: string | null) => void;
    onCancel: () => void;
    onUndo: () => void;
    onRetry: () => void;
  } = $props();
</script>

<Section
  panel
  frame="framed"
  anchor="guild-loot-candidates"
  title={`Candidates for ${itemName}`}
  testid="guild-loot-candidates"
>
  {#if notice !== ''}<p
      class="text-wipe pb-2 text-[13px] leading-5"
      role="alert"
      data-testid="guild-loot-notice"
    >
      {notice}
    </p>{/if}
  {#if status === 'loading' || (status === 'ready' && drop === null)}
    <Skeleton lines={4} minHeight="min-h-[200px]" testid="guild-loot-candidates-skeleton" />
  {:else if status === 'failed'}
    <LoadError message="The candidates did not load." testid="guild-loot-candidates-error" {onRetry} />
  {:else if drop !== null}
    {#if drop.award !== null}
      <p class="text-muted pb-1 text-[13px]" data-testid="guild-loot-closed">{CLOSED_LINE}</p>
      <DecisionRow
        decision={drop.award}
        {build}
        {shipsIcon}
        {timeZone}
        {onUndo}
        undoDisabled={!canAct || busy}
      />
    {/if}
    {#each drop.candidates as candidate (candidate.character_key)}
      <LootCandidate
        {candidate}
        {itemName}
        gear={drop.kind === 'gear'}
        closed={drop.closed}
        {canAct}
        {describedBy}
        confirming={confirmingKey === candidate.character_key}
        {busy}
        onVote={(vote) => onVote(candidate.character_key, vote)}
        onAsk={() => onAsk(candidate.character_key)}
        onAward={(msOs, reason) => onAward(candidate.character_key, msOs, reason)}
        {onCancel}
      />
    {:else}
      <p class="text-muted text-[14px]" data-testid="guild-loot-no-candidates">{NO_CANDIDATES}</p>
    {/each}
  {/if}
</Section>
