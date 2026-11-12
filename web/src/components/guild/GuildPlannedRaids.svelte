<!-- web/src/components/guild/GuildPlannedRaids.svelte -->
<!-- The body of the Raids tab's `Planned` Section (raids spec 3.1): the next four rows by start in the
     Planned panel's rows and `Show all {n} planned raids`, which expands in place; the empty line says
     who plans raids (member) or what planning does (officer). The Section (GuildRaids) carries the
     sub-head and its count. -->
<script lang="ts">
  import EmptyState from '../ui/EmptyState.svelte';
  import ShowAll from '../ui/ShowAll.svelte';
  import type { GuildRaidView } from '../../lib/guild/api';
  import { guildRaidHref } from '../../lib/characters';
  import { sortByStart } from '../../lib/guild/planned-row';
  import { raidCopy, showAll } from '../../lib/guild/raid-copy';
  import PlannedRaidRow from './PlannedRaidRow.svelte';

  /** Rows shown before `Show all`. */
  const FIRST_ROWS = 4;

  let {
    planned,
    officer,
    guildPath,
    now,
    viewerTimezone,
  }: {
    planned: readonly GuildRaidView[];
    officer: boolean;
    guildPath: string;
    now: Date;
    viewerTimezone: string;
  } = $props();

  let expanded = $state(false);
  const sorted = $derived(sortByStart(planned));
  const shown = $derived(expanded ? sorted : sorted.slice(0, FIRST_ROWS));
</script>

{#if sorted.length === 0}
  <EmptyState
    message={officer ? raidCopy.emptyPlanned.officer : raidCopy.emptyPlanned.member}
    testid="guild-planned-empty"
  />
{:else}
  <div id="guild-planned-rows" data-testid="guild-planned-rows">
    {#each shown as raid (raid.id)}
      <PlannedRaidRow {raid} href={guildRaidHref(guildPath, raid.id)} {now} {viewerTimezone} />
    {/each}
  </div>
  <ShowAll
    shown={FIRST_ROWS}
    total={sorted.length}
    {expanded}
    collapsible={false}
    onToggle={() => (expanded = true)}
    label={showAll(sorted.length)}
    controls="guild-planned-rows"
    testid="guild-planned-show-all"
  />
{/if}
