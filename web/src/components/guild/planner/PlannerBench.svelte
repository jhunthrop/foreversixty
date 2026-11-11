<!-- web/src/components/guild/planner/PlannerBench.svelte -->
<!-- Officers: Bench and swaps (web layout spec 2026-11-05 §7.2). Each benched sign-up is a Row with
     Swap in, which moves them to In through the same officer sign-up write the roster's control
     uses; the comp tally above shows what the swap does to the seats. -->
<script lang="ts">
  import EmptyState from '../../ui/EmptyState.svelte';
  import type { GuildRaidView, GuildSignupState } from '../../../lib/guild/api';
  import { kindCountsRoles, kindOf } from '../../../lib/guild/raid-copy';
  import { benchRows, rowSub } from '../../../lib/guild/raid-planner-view';
  import type { RosterRow } from '../../../lib/guild/raid-roster';
  import { classTone, ROW_ITEM } from '../../../lib/ui/primitives';
  import ClassCrestRing from '../../character/ClassCrestRing.svelte';
  import Button from '../../ui/Button.svelte';
  import Row from '../../ui/Row.svelte';
  import Section from '../../ui/Section.svelte';

  let {
    raid,
    officerTools,
    closed = false,
    busyKey = null,
    onSetState,
  }: {
    raid: GuildRaidView;
    officerTools: boolean;
    closed?: boolean;
    busyKey?: string | null;
    onSetState: (row: RosterRow, state: GuildSignupState) => Promise<boolean>;
  } = $props();

  const rows = $derived(benchRows(raid));
  const roles = $derived(kindCountsRoles(kindOf(raid)));
</script>

<Section
  panel
  title="Bench and swaps"
  meta={rows.length === 0 ? undefined : String(rows.length)}
  testid="raid-bench"
>
  {#if rows.length === 0}
    <EmptyState message="Nobody is on the bench." testid="raid-bench-empty" />
  {:else}
    <ul>
      {#each rows as row (row.character_key)}
        <li class={ROW_ITEM} data-testid="raid-bench-row">
          <Row title={row.name} titleTone={classTone(row.class)} sub={rowSub(row, roles)} open density="slim">
            {#snippet lead()}<ClassCrestRing characterClass={row.class} size={24} />{/snippet}
            {#snippet trail()}
              <Button
                size="sm"
                disabled={!officerTools || closed || busyKey === row.character_key}
                onclick={() => void onSetState(row, 'in')}
                testid={`raid-swap-in-${row.character_key}`}>Swap in</Button
              >
            {/snippet}
          </Row>
        </li>
      {/each}
    </ul>
  {/if}
</Section>
