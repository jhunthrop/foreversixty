<!-- web/src/components/guild/report/ReportPulls.svelte -->
<!-- The report's Pulls panel (web layout spec 2026-11-05 §7.2 "Raid report"): every pull of the
     night as a Row (the boss and its pull number, or the trash pack; the time, the deaths and what
     killed them; the Kill, Wipe or Clear Tag). A pull with deaths opens them: from 1024 in the
     Inspector at the plate's right (the first wipe with deaths shows at once), below 1024 inline
     under the tapped row. The rows' action is the title's button, stretched over the row. -->
<script lang="ts">
  import EmptyState from '../../ui/EmptyState.svelte';
  import { MediaQuery } from 'svelte/reactivity';
  import activeBuild from '../../../data/active-build.json';
  import type { RaidReport } from '../../../lib/guild/attendance-contract';
  import { defaultPullKey, pullRows, type PullRowModel } from '../../../lib/guild/report-night-view';
  import { reportNightCopy as copy } from '../../../lib/guild/report-night-copy';
  import { NO_SPELL_ICONS, loadSpellIcons, type SpellIconOf } from '../../../lib/guild/report-spell-icons';
  import Inspector from '../../ui/Inspector.svelte';
  import Row from '../../ui/Row.svelte';
  import Section from '../../ui/Section.svelte';
  import Tag from '../../ui/Tag.svelte';
  import PullDeaths from './PullDeaths.svelte';

  const INSPECTOR_ID = 'raid-pull-inspector';
  /** The Inspector sits at the plate's side from here (its `side` placement). */
  const SIDE_QUERY = 'min-width: 1024px';

  let {
    report,
    loadIcons = loadSpellIcons,
  }: {
    report: RaidReport;
    /** The spell icon table's load; a test passes its own. */
    loadIcons?: (build: string) => Promise<SpellIconOf>;
  } = $props();

  const rows = $derived(pullRows(report));
  const side = new MediaQuery(SIDE_QUERY, false);
  /** The row the viewer opened; on the side the first wipe stands in until they pick one. */
  let chosen = $state<string | null>(null);
  const openKey = $derived(side.current ? (chosen ?? defaultPullKey(rows)) : chosen);
  const open = $derived(rows.find((row) => row.key === openKey && row.expandable) ?? null);
  let iconOf = $state<SpellIconOf>(NO_SPELL_ICONS);

  const hasCauses = $derived(report.pulls.some((pull) => (pull.death_causes?.length ?? 0) > 0));

  $effect(() => {
    if (!hasCauses) return;
    let live = true;
    loadIcons(activeBuild.build)
      .then((next) => {
        if (live) iconOf = next;
      })
      .catch(() => undefined);
    return () => {
      live = false;
    };
  });

  function toggle(row: PullRowModel): void {
    chosen = side.current || chosen !== row.key ? row.key : null;
  }

  const panelId = (row: PullRowModel): string => (side.current ? INSPECTOR_ID : `raid-pull-${row.key}`);
</script>

{#snippet deaths()}
  {#if open !== null}<PullDeaths pull={open.pull} {iconOf} />{/if}
{/snippet}

{#snippet head()}
  {#if open !== null}
    <p class="text-strong text-[15px] leading-5 font-bold">{open.title}</p>
    <p class="text-muted text-[13px] leading-[18px]">{open.sub}</p>
  {/if}
{/snippet}

<Section
  title={copy.pulls.head}
  meta={copy.pulls.count(rows.length)}
  anchor="report"
  panel
  flush
  testid="raid-report"
>
  <div role="group" aria-label={copy.pulls.caption} data-testid="raid-pulls">
    {#if rows.length === 0}
      <EmptyState inset message={copy.pulls.none} testid="raid-pulls-empty" />
    {:else}
      <div class="flex min-w-0">
        <div class="min-w-0 flex-1">
          {#each rows as row (row.key)}
            {@const isOpen = open?.key === row.key}
            <div data-testid="raid-pull-row" data-result={row.pull.result}>
              <Row title={row.title} sub={row.sub} selected={isOpen}>
                {#snippet titleContent()}
                  {#if row.expandable}
                    <button
                      type="button"
                      class="focus-visible:outline-gold text-left after:absolute after:inset-0 focus-visible:outline-2"
                      aria-expanded={isOpen}
                      aria-controls={isOpen ? panelId(row) : undefined}
                      onclick={() => toggle(row)}
                      data-testid="raid-pull-open">{row.title}</button
                    >
                  {:else}{row.title}{/if}
                {/snippet}
                {#snippet trail()}<Tag tone={row.tag.tone} testid="raid-pull-result">{row.tag.text}</Tag
                  >{/snippet}
              </Row>
              {#if isOpen && !side.current}
                <Inspector
                  id={panelId(row)}
                  selection={row.title}
                  placement="inline"
                  sections={[{ label: copy.figures.deaths, body: deaths }]}
                  testid="raid-pull-inspector"
                />
              {/if}
            </div>
          {/each}
        </div>
        {#if open !== null && side.current}
          <Inspector
            id={INSPECTOR_ID}
            selection={open.title}
            placement="side"
            {head}
            sections={[{ label: copy.figures.deaths, body: deaths }]}
            testid="raid-pull-inspector"
          />
        {/if}
      </div>
    {/if}
  </div>
</Section>
