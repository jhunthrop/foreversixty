<!-- web/src/components/guild/GuildProgression.svelte -->
<!-- Guild control-centre spec §4.D, in the flat system (web layout spec 2026-11-05 §3A; GUILD-ALL
     2026-11-07): each named encounter's pulls to kill as an open Section of Rows in the main column,
     the tier's count with its Meter, the zones with no encounter list and the roster bests in the
     rail. No card, no box: hierarchy by type and spacing. Shown to every role. -->
<script lang="ts">
  import { characterHref, splitUnitName, type Region, type Ruleset } from '../../lib/characters';
  import { classColorVar } from '../../lib/report/format';
  import { classTone } from '../../lib/ui/primitives';
  import type { GuildParseRef, GuildProgressionPage } from '../../lib/guild/api';
  import { minuteClock } from '../../lib/guild/minute-clock.svelte';
  import { progressionNightDay } from '../../lib/guild/overview-view';
  import { firstKillLine, progressionEmpty } from '../../lib/guild/progression-view';
  import { outputFigure, topOutputLabel } from '../../lib/guild/output-figure';
  import { rosterBestLabels } from '../../lib/guild/roster-best';
  import { counted } from '../../lib/plural';
  import type { GuildRosterBest } from '../../lib/rankings/api';
  import EmptyState from '../ui/EmptyState.svelte';
  import Figure from '../ui/Figure.svelte';
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
  const encounters = $derived((progression?.encounters ?? []).filter((encounter) => encounter.pulls > 0));
</script>

<div data-testid="guild-progression-tab">
  {#if progression === null || progression.tier.pulls_total === 0}
    <!-- Zero pulls: the tab's one block is the shared empty state under its heading (guild round 2). -->
    <Page layout="ledger">
      <Section title="Progression" testid="guild-progression-empty-card">
        {#if progression === null}
          <EmptyState message="No pulls recorded yet." testid="guild-progression-empty" />
        {:else}
          {@const empty = progressionEmpty(now)}
          <EmptyState message={empty.message} action={empty.action} testid="guild-progression-empty" />
        {/if}
      </Section>
    </Page>
  {:else}
    {@const tier = progression.tier}
    <Page layout="split" testid="guild-progression-page">
      {#each encounters as encounter (encounter.encounter_id)}
        {@const kill = firstKillLine(encounter, now)}
        {@const best = bestByRole(encounter.best_by_role)}
        <Section
          title={encounter.name}
          meta={`${encounter.zone} · pulls to kill`}
          testid="guild-progression-encounter"
        >
          <div class="flex flex-col gap-3">
            {#if kill !== null}
              <p class="text-[14px]" data-testid="guild-progression-kill">{kill}</p>
            {:else if encounter.kills === 0}
              <p class="text-[14px]" data-testid="guild-progression-no-kill">
                {counted(encounter.pulls, 'pull')}, no kill yet.
              </p>
            {/if}
            <ul>
              {#each encounter.pulls_by_night as night (night.report_id)}
                <li>
                  <Row
                    title={progressionNightDay(night.date)}
                    sub={counted(night.pulls, 'pull')}
                    density="slim"
                    open
                  >
                    {#snippet trail()}
                      {#if night.killed}<Tag tone="success">Kill</Tag>{/if}
                    {/snippet}
                  </Row>
                </li>
              {/each}
            </ul>
            <p class="text-muted text-[14px]" data-testid="guild-progression-deaths">
              Deaths per pull: {encounter.deaths_per_pull.toFixed(1)} on average.
            </p>
            {#if best.length > 0}
              <!-- Raw output, never called a parse (a percentile to a Classic player): `Top HPS`
                   per role, the figure written as the Raids tab and the roster write it. -->
              <p class="text-[14px]" data-testid="guild-progression-best-parse">
                {#each best as { role, parse }, index (role)}
                  {#if index > 0}<span class="hidden md:inline" aria-hidden="true">&nbsp;·&nbsp;</span>{/if}
                  <span class="block md:inline">
                    <span class="text-muted">{topOutputLabel(parse.metric, role)}</span>
                    <span style={`color:${classColorVar(parse.class)};font-weight:700`}>{parse.name}</span>
                    <span class="tabular-nums">{outputFigure(parse.value)}</span>
                  </span>
                {/each}
              </p>
            {/if}
          </div>
        </Section>
      {/each}

      {#snippet rail()}
        <Section title="This tier" meta={tier.name} testid="guild-progression-tier">
          <div class="flex flex-col gap-2">
            <Figure
              value={`${tier.down} of ${tier.named_encounters}`}
              caption="Named bosses down"
              size="sm"
              testid="guild-progression-down"
            />
            <Meter
              value={tier.named_encounters === 0 ? 0 : tier.down / tier.named_encounters}
              label={`${tier.down} of ${tier.named_encounters} named bosses down`}
            />
          </div>
        </Section>
        {#if unnamed.length > 0}
          <Section title="Without an encounter list" testid="guild-progression-unnamed">
            <!-- Our to-do, so only when this guild has pulls it applies to (P.1, player review S2-5). -->
            <p class="text-muted pb-2 text-[14px]" data-testid="guild-progression-unnamed-note">
              Barrow Deeps and Hyjal Summit have no encounter list yet. Their pulls count toward the tier
              total but cannot be named or ranked per boss.
            </p>
            <ul>
              {#each unnamed as zone (zone.zone)}
                <li>
                  <Row title={zone.zone} sub={`${counted(zone.pulls, 'pull')} logged`} density="slim" open />
                </li>
              {/each}
            </ul>
          </Section>
        {/if}
        {#if encounterBests.length > 0}
          <Section title="Roster bests" testid="guild-roster-best">
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
                      <span class="text-[14px] font-bold tabular-nums" title={labels.hover}
                        >{labels.value}</span
                      >
                    {/snippet}
                  </Row>
                </li>
              {/each}
            </ul>
          </Section>
        {/if}
      {/snippet}
    </Page>
  {/if}
</div>
