<!-- web/src/components/guild/GuildProgression.svelte -->
<!-- Progression (guild dense spec, h-guild-tabs): each boss pulled on its own framed panel (its
     figures and every night it was pulled, EncounterProgress); the rail holds the tier (each zone's
     bosses down over a meter, the zones with no encounter list and their pulls) and the roster bests
     as one ranked table. Shown to every role. -->
<script lang="ts">
  import { characterHref, splitUnitName, type Region, type Ruleset } from '../../lib/characters';
  import type { GuildProgressionPage } from '../../lib/guild/api';
  import { classColorVar } from '../../lib/report/format';
  import { minuteClock } from '../../lib/ui/minute-clock.svelte';
  import { bossesDownMeta, progressionEmpty, progressionZones } from '../../lib/guild/progression-view';
  import { rosterBestLabels } from '../../lib/guild/roster-best';
  import { counted } from '../../lib/plural';
  import type { GuildRosterBest } from '../../lib/rankings/api';
  import Crest from '../ui/Crest.svelte';
  import EmptyState from '../ui/EmptyState.svelte';
  import Meter from '../ui/Meter.svelte';
  import Page from '../ui/Page.svelte';
  import Row from '../ui/Row.svelte';
  import Section from '../ui/Section.svelte';
  import StatRow from '../ui/StatRow.svelte';
  import StatTable from '../ui/StatTable.svelte';
  import { LINK } from '../../lib/ui/primitives';
  import EncounterProgress from './progression/EncounterProgress.svelte';
  import BestByRole from './progression/BestByRole.svelte';
  import { SEP } from '../../lib/guild/raid-copy';
  import { specClassName } from '../../lib/guild/spec-name';

  /** `rosterBest` is v1's own "Roster bests" feature (`GET .../<guild>`'s existing,
   *  working `roster_best` rows) -- carried forward unchanged here (tenet 10) rather than
   *  dropped while the new control-centre `/progression` endpoint is still deploying. */
  let {
    progression,
    rosterBest = [],
    region,
    ruleset,
    fixedNow,
  }: {
    progression: GuildProgressionPage | null;
    rosterBest?: GuildRosterBest[];
    region: Region;
    ruleset: Ruleset;
    /** A fixed clock for tests; the page runs on the minute tick. */
    fixedNow?: Date;
  } = $props();

  /** One row per raider per encounter: the API also sends each raider's best with no encounter
   *  behind it (id 0, no name), which reads as a second row of dashes and is not a boss. */
  const encounterBests = $derived(
    rosterBest
      .filter((row) => row.encounter_id > 0 && row.encounter !== '')
      .toSorted((a, b) => a.metric.localeCompare(b.metric) || b.value - a.value),
  );

  /** The top of the list until the officer asks for every raider. */
  const BESTS_SHOWN = 10;
  let allBests = $state(false);
  const shownBests = $derived(allBests ? encounterBests : encounterBests.slice(0, BESTS_SHOWN));

  const clock = minuteClock(() => fixedNow);
  const now = $derived(clock.now);
  /** The tier's zones without an encounter list that the guild has pulled in. */
  const unnamed = $derived((progression?.unnamed ?? []).filter((zone) => zone.pulls > 0));
  /** One boss: the rows need not repeat its name. */
  const oneBoss = $derived(new Set(encounterBests.map((row) => row.encounter_id)).size === 1);
  const zones = $derived(progressionZones(progression?.encounters ?? []));
</script>

{#snippet rosterBests()}
  {#if encounterBests.length > 0}
    <Section title="Roster bests" panel flush testid="guild-roster-best">
      <StatTable columns={[{ id: 'value', label: 'Best' }]} nameHead="Raider" bar={false} dense>
        {#each shownBests as row (`${row.player.key}-${row.encounter_id}-${row.metric}`)}
          {@const labels = rosterBestLabels(row)}
          <StatRow bar={false} figures={[{ value: labels.value }]}>
            {#snippet icon()}<Crest cls={row.player.class} size={28} />{/snippet}
            {#snippet name()}<a
                href={characterHref(region, ruleset, row.player.name)}
                style:color={classColorVar(row.player.class)}
                class="underline-offset-4 hover:underline">{splitUnitName(row.player.name).name}</a
              >{/snippet}
            {#snippet sub()}<span title={labels.hover}
                >{oneBoss ? specClassName(row.player.spec, row.player.class) : labels.what}</span
              >{#if labels.execution !== null}{SEP}<span title={labels.executionHover}
                  >{labels.execution}</span
                >{/if}{/snippet}
          </StatRow>
        {/each}
      </StatTable>
      {#if !allBests && encounterBests.length > BESTS_SHOWN}
        <button
          type="button"
          class="{LINK} min-h-11 px-5 text-[14px]"
          onclick={() => (allBests = true)}
          data-testid="guild-roster-best-all">All {encounterBests.length}</button
        >
      {/if}
    </Section>
  {/if}
{/snippet}

{#snippet tier()}
  {#if progression !== null}
    <Section
      title={progression.tier.name}
      meta={counted(progression.tier.pulls_total, 'pull')}
      panel
      testid="guild-progression-tier"
    >
      <ul>
        {#each zones as zone (zone.zone)}
          <li class="flex flex-col gap-2 pb-3" data-testid="guild-progression-zone">
            <span class="flex items-baseline justify-between gap-3">
              <span class="text-strong text-[14px] font-bold">{zone.zone}</span>
              <span class="text-muted text-[13px]">{bossesDownMeta(zone)}</span>
            </span>
            <Meter
              value={zone.named === 0 ? 0 : zone.down / zone.named}
              tone="kill"
              label={`${zone.down} of ${zone.named} named bosses down`}
            />
          </li>
        {/each}
        {#each unnamed as zone (zone.zone)}
          <li data-testid="guild-progression-unnamed">
            <Row
              title={zone.zone}
              sub={`no encounter list · ${counted(zone.pulls, 'pull')}${zone.nights > 0 ? ` over ${counted(zone.nights, 'night')}` : ''}`}
              density="slim"
              open
            />
          </li>
        {/each}
      </ul>
    </Section>
  {/if}
{/snippet}

<div data-testid="guild-progression-tab">
  {#if progression === null || progression.tier.pulls_total === 0}
    <!-- Zero pulls: the tab's one block is the shared empty state under its heading (guild round 2). -->
    <Page layout="ledger" panels>
      <Section title="Progression" panel testid="guild-progression-empty-card">
        {#if progression === null}
          <EmptyState message="No pulls recorded yet." testid="guild-progression-empty" />
        {:else}
          {@const empty = progressionEmpty(now)}
          <EmptyState message={empty.message} action={empty.action} testid="guild-progression-empty" />
        {/if}
      </Section>
      <!-- The public read's roster bests stand without the progression endpoint (tenet 10): a
           guild with no tier pulls yet still shows who is best where. -->
      {@render rosterBests()}
    </Page>
  {:else}
    <Page layout="split" panels testid="guild-progression-page">
      {#each zones as zone (zone.zone)}
        {#each zone.pulled as encounter (encounter.encounter_id)}
          <EncounterProgress {encounter} {now} />
        {/each}
      {/each}
      {@render rosterBests()}
      {#snippet rail()}
        {@render tier()}
        {#each zones as zone (zone.zone)}
          {#each zone.pulled as encounter (encounter.encounter_id)}<BestByRole {encounter} />{/each}
        {/each}
      {/snippet}
    </Page>
  {/if}
</div>
