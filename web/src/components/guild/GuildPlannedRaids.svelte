<!-- web/src/components/guild/GuildPlannedRaids.svelte -->
<!-- The body of the Raids tab's `Planned` Section (raids spec 3.1): the next four rows by start in the
     tab's one Plate and `Show all {n} planned raids`, which expands in place; the empty line says
     who plans raids (member) or what planning does (officer). The Section (GuildRaids) carries the
     sub-head and its count. -->
<script lang="ts">
  import Plate from '../ui/Plate.svelte';
  import type { GuildRaidView } from '../../lib/guild/api';
  import { guildRaidHref } from '../../lib/characters';
  import { sortByStart } from '../../lib/guild/planned-row';
  import { raidCopy, showAll } from '../../lib/guild/raid-copy';
  import { LINK } from '../../lib/ui/primitives';
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
  <p class="text-muted text-[14px] leading-[22px]" data-testid="guild-planned-empty">
    {officer ? raidCopy.emptyPlanned.officer : raidCopy.emptyPlanned.member}
  </p>
{:else}
  <Plate testid="guild-planned-rows">
    {#each shown as raid (raid.id)}
      <PlannedRaidRow {raid} href={guildRaidHref(guildPath, raid.id)} {now} {viewerTimezone} />
    {/each}
  </Plate>
  {#if !expanded && sorted.length > FIRST_ROWS}
    <button
      type="button"
      class="{LINK} mt-2 h-11 text-[14px]"
      onclick={() => (expanded = true)}
      data-testid="guild-planned-show-all">{showAll(sorted.length)}</button
    >
  {/if}
{/if}
