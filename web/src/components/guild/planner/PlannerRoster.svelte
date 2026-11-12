<!-- web/src/components/guild/planner/PlannerRoster.svelte -->
<!-- The planner's one Plate (web layout spec 2026-11-05 §7.2, §3A.2): every sign-up as a Row, in
     the grouping the control bar chose. Role is the shipped Raid-Helper read (Tanks, Healers,
     Melee, Ranged, then Late, Bench, Out); Status and Class regroup the same rows. Each group is a
     caps label with its count over its rows, never a Section (a Plate holds no headed section).
     After an officer's change moves a row to another group, focus follows it to its new place. -->
<script lang="ts">
  import EmptyState from '../../ui/EmptyState.svelte';
  import { tick } from 'svelte';
  import type { GuildRaidView, GuildSignupState } from '../../../lib/guild/api';
  import { kindCountsRoles, kindOf } from '../../../lib/guild/raid-copy';
  import type { OfficerLine } from '../../../lib/guild/raid-page-types';
  import { plannerGroups, type PlannerGrouping } from '../../../lib/guild/raid-planner-view';
  import type { RosterRow } from '../../../lib/guild/raid-roster';
  import type { PlannerFacts } from '../../../lib/guild/planner-facts';
  import PlannerRow from './PlannerRow.svelte';
  import AttendanceStrip from '../dense/AttendanceStrip.svelte';
  import { PLANNER_FACT_COLUMNS, PLANNER_FACT_HEADS } from './planner-columns';

  let {
    raid,
    grouping,
    officer,
    officerTools,
    gateId,
    closed = false,
    busyKey = null,
    line = null,
    facts = null,
    onSetState,
  }: {
    /** Each sign-up's figures by character key; null draws none. */
    facts?: ReadonlyMap<string, PlannerFacts> | null;
    raid: GuildRaidView;
    grouping: PlannerGrouping;
    officer: boolean;
    officerTools: boolean;
    gateId: string;
    /** A cancelled raid takes no sign-up changes: the controls stay, disabled. */
    closed?: boolean;
    busyKey?: string | null;
    line?: OfficerLine | null;
    onSetState: (row: RosterRow, state: GuildSignupState) => Promise<boolean>;
  } = $props();

  const groups = $derived(plannerGroups(raid, grouping));
  const roles = $derived(kindCountsRoles(kindOf(raid)));
  const empty = $derived(groups.every((group) => group.rows.length === 0));

  async function setState(row: RosterRow, state: GuildSignupState): Promise<boolean> {
    const saved = await onSetState(row, state);
    if (!saved) return false;
    await tick();
    document
      .querySelector<HTMLElement>(`[data-testid="raid-roster-state-${row.character_key}-${state}"]`)
      ?.focus();
    return true;
  }
  /** The last-nights strip has data for someone; without it its column head stays blank. */
  const nightsKnown = $derived(facts !== null && [...facts.values()].some((f) => f.cells !== null));
</script>

<div class={closed ? 'opacity-60' : undefined} data-testid="raid-roster">
  <div role="region" aria-label="Sign-ups">
    {#if empty}
      <EmptyState message="Nobody has signed up yet." testid="raid-roster-empty" />
    {:else}
      {#if facts !== null}
        <div class="hidden items-center justify-between gap-4 pt-1 lg:flex">
          {#if nightsKnown}<AttendanceStrip legend />{:else}<span></span>{/if}
          <span class="{PLANNER_FACT_COLUMNS} text-label text-muted grid font-bold" aria-hidden="true">
            {#each PLANNER_FACT_HEADS as head, i (head)}<span>{i === 0 && !nightsKnown ? '' : head}</span
              >{/each}
          </span>
        </div>
      {/if}
      {#each groups as group (group.id)}
        <div class="pb-1" data-testid={`raid-roster-${group.id}`}>
          <h3
            class="text-label text-muted flex items-baseline gap-2 pt-3 pb-1 font-bold"
            data-testid={`raid-roster-${group.id}-heading`}
          >
            {group.label} <span class="tabular-nums">{group.rows.length}</span>
          </h3>
          {#if group.rows.length === 0}
            <p class="text-muted pb-3 text-[13px]">{group.empty}</p>
          {:else}
            <ul>
              {#each group.rows as row (row.character_key)}
                <PlannerRow
                  {row}
                  {roles}
                  {officer}
                  disabled={!officerTools || closed}
                  describedBy={officerTools ? undefined : gateId}
                  busy={busyKey === row.character_key}
                  refusal={line?.control === `roster:${row.character_key}` ? line.message : null}
                  facts={facts?.get(row.character_key) ?? null}
                  stateNamed={group.id === row.state}
                  onSetState={setState}
                />
              {/each}
            </ul>
          {/if}
        </div>
      {/each}
    {/if}
  </div>
  {#if line?.control === 'roster'}
    <p class="text-wipe mt-2 text-[13px]" role="alert" data-testid="raid-roster-error">{line.message}</p>
  {/if}
</div>
