<!-- web/src/components/guild/GuildProgression.svelte -->
<!-- Progression (batch 3, b3-guild-progression; guild control-centre spec §4.D): one framed panel a
     zone, its bosses down over a meter, each boss pulled as one Row (its pulls, wipes, deaths a pull
     and first kill) with its Killed or Progress tag. The zones with no encounter list and the roster
     bests follow on their own panels. Shown to every role. -->
<script lang="ts">
  import { characterHref, splitUnitName, type Region, type Ruleset } from '../../lib/characters';
  import { classTone } from '../../lib/ui/primitives';
  import type { GuildParseRef, GuildProgressionPage } from '../../lib/guild/api';
  import { outputFigure, topOutputLabel } from '../../lib/guild/output-figure';
  import { classColorVar } from '../../lib/report/format';
  import { minuteClock } from '../../lib/guild/minute-clock.svelte';
  import activeBuild from '../../data/active-build.json';
  import { raidZoneIcon } from '../../lib/first-raids';
  import {
    bossLine,
    bossesDownMeta,
    progressionEmpty,
    progressionZones,
  } from '../../lib/guild/progression-view';
  import { rosterBestLabels } from '../../lib/guild/roster-best';
  import { counted } from '../../lib/plural';
  import type { GuildRosterBest } from '../../lib/rankings/api';
  import EmptyState from '../ui/EmptyState.svelte';
  import ItemIcon from '../ui/ItemIcon.svelte';
  import Meter from '../ui/Meter.svelte';
  import Page from '../ui/Page.svelte';
  import Row from '../ui/Row.svelte';
  import Section from '../ui/Section.svelte';
  import Tag from '../ui/Tag.svelte';

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
  const encounterBests = $derived(rosterBest.filter((row) => row.encounter_id > 0 && row.encounter !== ''));

  /** The tank, healer and damage bests that exist, in that order, each with its role. */
  function bestByRole(best: {
    tank?: GuildParseRef;
    healer?: GuildParseRef;
    dps?: GuildParseRef;
  }): { role: string; parse: GuildParseRef }[] {
    return (['tank', 'healer', 'dps'] as const).flatMap((role) => {
      const parse = best[role];
      return parse === undefined ? [] : [{ role, parse }];
    });
  }

  const clock = minuteClock(() => fixedNow);
  const now = $derived(clock.now);
  /** Only the zones and encounters with pulls: a section of nothing is not drawn (week-one spec §3.4). */
  const unnamed = $derived((progression?.unnamed ?? []).filter((zone) => zone.pulls > 0));
  const zones = $derived(progressionZones(progression?.encounters ?? []));
</script>

{#snippet rosterBests()}
  {#if encounterBests.length > 0}
    <Section title="Roster bests" panel testid="guild-roster-best">
      <ul>
        {#each encounterBests as row (`${row.player.key}-${row.encounter_id}-${row.metric}`)}
          {@const labels = rosterBestLabels(row)}
          <li>
            <Row
              title={splitUnitName(row.player.name).name}
              titleTone={classTone(row.player.class)}
              sub={labels.execution === null ? labels.what : `${labels.what} · ${labels.execution}`}
              href={characterHref(region, ruleset, row.player.name)}
              density="slim"
              open
            >
              {#snippet trail()}
                <span class="text-[14px] font-bold tabular-nums" title={labels.hover}>{labels.value}</span>
              {/snippet}
            </Row>
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
    <Page layout="ledger" panels testid="guild-progression-page">
      {#each zones as zone (zone.zone)}
        <Section
          title={zone.zone}
          panel
          frame="framed"
          meta={bossesDownMeta(zone)}
          testid="guild-progression-zone"
        >
          <div class="flex flex-col gap-2">
            <Meter
              value={zone.named === 0 ? 0 : zone.down / zone.named}
              tone="kill"
              label={`${zone.down} of ${zone.named} named bosses down`}
            />
            <ul>
              {#each zone.pulled as encounter (encounter.encounter_id)}
                <li>
                  <Row
                    title={encounter.name}
                    sub={bossLine(encounter, now)}
                    open
                    testid="guild-progression-encounter"
                  >
                    {#snippet lead()}<ItemIcon
                        icon={raidZoneIcon(encounter.zone)}
                        quality={5}
                        build={activeBuild.build}
                        size={36}
                      />{/snippet}
                    {#snippet foot()}
                      {@const best = bestByRole(encounter.best_by_role)}
                      {#if best.length > 0}
                        <!-- Raw output, never called a parse (a percentile to a Classic player). -->
                        <span class="text-[13px]" data-testid="guild-progression-best-parse">
                          {#each best as { role, parse }, index (role)}
                            {#if index > 0}<span aria-hidden="true">&nbsp;·&nbsp;</span>{/if}
                            <span class="text-muted">{topOutputLabel(parse.metric, role)}</span>
                            <span style={`color:${classColorVar(parse.class)};font-weight:700`}
                              >{parse.name}</span
                            >
                            <span class="tabular-nums">{outputFigure(parse.value)}</span>
                          {/each}
                        </span>
                      {/if}
                    {/snippet}
                    {#snippet trail()}
                      {#if encounter.kills > 0}<Tag tone="success">Killed</Tag>{:else}<Tag>Progress</Tag>{/if}
                    {/snippet}
                  </Row>
                </li>
              {/each}
            </ul>
          </div>
        </Section>
      {/each}
      {#if unnamed.length > 0}
        <Section title="Without an encounter list" panel testid="guild-progression-unnamed">
          <ul>
            {#each unnamed as zone (zone.zone)}
              <li>
                <Row title={zone.zone} sub={`${counted(zone.pulls, 'pull')} logged`} density="slim" open />
              </li>
            {/each}
          </ul>
        </Section>
      {/if}
      {@render rosterBests()}
    </Page>
  {/if}
</div>
