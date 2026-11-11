<!-- web/src/components/report/GroupBuffs.svelte -->
<!-- The Summary rail's buffs (owner review 2026-11-09): the group's buffs over the pull, compact, with
     the Buffs › link. One row per buff the scope's players held: its uptime on the players
     who had it, the pull's full length counting as 100%. The buffs most of the group held
     come first (flasks, food, the class buffs), then by uptime. The top eight. -->
<script lang="ts">
  import { percent } from '../../lib/format/number';
  import type { SpellIconOf } from '../../lib/guild/report-spell-icons';
  import type { Summary } from '../../lib/report/types';
  import Meter from '../ui/Meter.svelte';
  import Section from '../ui/Section.svelte';
  import StatRow from '../ui/StatRow.svelte';
  import StatTable from '../ui/StatTable.svelte';
  import SpellIcon from './SpellIcon.svelte';

  let {
    summary,
    durationMs,
    players,
    spellIcon,
    tabHref,
  }: {
    summary: Summary;
    durationMs: number;
    /** The report's players; the pull's roster is counted, never a passer-by. */
    players: ReadonlySet<string>;
    spellIcon: SpellIconOf;
    tabHref: string;
  } = $props();

  const ROWS = 8;
  const COLUMNS = [{ id: 'uptime', label: 'Uptime' }] as const;

  interface BuffRow {
    spellId: number;
    name: string;
    school?: number;
    holders: number;
    uptimeMs: number;
  }

  /** The pull's roster in the scope: the players the uptime is counted over. */
  const scope = $derived(
    new Set(summary.roster.map((row) => row.guid).filter((guid) => players.size === 0 || players.has(guid))),
  );
  const rows = $derived.by(() => {
    // eslint-disable-next-line svelte/prefer-svelte-reactivity
    const byName = new Map<string, BuffRow & { guids: Set<string> }>();
    for (const track of summary.auras) {
      if (track.type !== 'BUFF' || !scope.has(track.target_guid) || track.name.startsWith('Spell #'))
        continue;
      const found = byName.get(track.name);
      if (found === undefined)
        byName.set(track.name, {
          spellId: track.spell_id,
          name: track.name,
          school: track.school,
          holders: 0,
          uptimeMs: track.uptime_ms,
          guids: new Set([track.target_guid]),
        });
      else {
        found.uptimeMs += track.uptime_ms;
        found.guids.add(track.target_guid);
      }
    }
    return (
      [...byName.values()]
        .map((row) => ({
          ...row,
          holders: row.guids.size,
          share: Math.min(1, row.uptimeMs / (Math.max(durationMs, 1) * row.guids.size)),
        }))
        // A buff the log saw only as held from before the pull carries no time: no figure to show.
        .filter((row) => row.uptimeMs > 0)
        .sort((a, b) => b.holders - a.holders || b.share - a.share || a.name.localeCompare(b.name))
        .slice(0, ROWS)
    );
  });
</script>

{#if rows.length > 0}
  <Section
    title="Buffs"
    action={{ label: 'Buffs', href: tabHref, arrow: true }}
    panel
    flush
    testid="panel-buffs"
  >
    <div class="group-buffs">
      <StatTable columns={COLUMNS} nameHead="Buff" label="Buffs">
        {#each rows as row (row.name)}
          <StatRow figures={[{ value: percent(row.share * 100) }]}>
            {#snippet icon()}<SpellIcon url={spellIcon(row.spellId, row.name, row.school)} />{/snippet}
            {#snippet name()}<b class="text-strong">{row.name}</b>{/snippet}
            {#snippet meter()}<Meter value={row.share} tone="gold" />{/snippet}
          </StatRow>
        {/each}
      </StatTable>
    </div>
  </Section>
{/if}

<style>
  /* In the rail: a short bar in the Buffs tab's gold. A long name wraps onto a second line in its
     own column, whole, never under the bar and never cut ("Power Word: Fortitude"). */
  .group-buffs :global(.ui-st) {
    --st-bar: minmax(0, 32px);
  }
  .group-buffs :global(.ui-st-name),
  .group-buffs :global(.ui-st-name > *) {
    min-width: 0;
    white-space: normal;
    overflow-wrap: anywhere;
  }
</style>
