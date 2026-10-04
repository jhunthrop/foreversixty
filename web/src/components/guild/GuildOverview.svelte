<!-- web/src/components/guild/GuildOverview.svelte -->
<!-- Guild control-centre spec §4.A: round 1's page, trimmed to a summary of every other
     tab, each a "See all" link to the full tab. The Roster/Raids/Progression summaries
     read from the same GuildHome/GuildPage fetch the header already makes (spec §8: no new
     round trip). The Readiness/Loot summary cards need figures the home contract does not
     carry (the three worst-readiness raiders; the current boss and next pick) -- this
     component fetches those two endpoints itself, once, for member/officer viewers only;
     both responses are cached through query.ts, so opening the real Readiness/Loot tab
     afterward costs nothing further. Documented as a deliberate, reported deviation from
     the letter of "no new round trip" where the contract's own `home.summary` has no field
     for either figure (build report, "what was left out and why"). The standing line and
     "before Thursday" sentence are NOT rendered here -- Guild.svelte renders them once,
     directly under the header facts line and above the tab strip itself (spec §4.0/§4.A.1:
     "the tab strip ... sits directly under the role-based line"), so they stay visible
     regardless of which tab is open rather than being Overview-only content. -->
<script lang="ts">
  import type { GuildHomeSummary } from '../../lib/guild/api';
  import { fetchGuildLoot, fetchGuildReadiness } from '../../lib/guild/api';
  import { sortReadinessWorstFirst } from '../../lib/guild/readiness-view';
  import { bossPickerLabel, nextUnkilledEncounter } from '../../lib/guild/loot-view';
  import GuildOfficerStrip from './GuildOfficerStrip.svelte';
  import type { GuildTabId } from './GuildTabs.svelte';

  /** A report row this card can summarize: the full `GuildHomeReport` shape (member/
   *  officer, via `home.reports`) or the public `GuildPage.reports` shape, which carries
   *  no kill/wipe counts -- the card simply omits that clause when they are absent, never
   *  showing a fabricated "0 kills". */
  interface OverviewReport {
    id: string;
    title: string;
    zone: string;
    created_at: string;
    fight_count?: number;
    kill_count?: number;
  }

  let {
    guildId,
    role,
    reports,
    summary,
    claimed,
    claimedByName,
    onSelectTab,
  }: {
    guildId: number | null;
    role: 'public' | 'member' | 'officer' | 'moderator';
    reports: readonly OverviewReport[];
    summary: GuildHomeSummary | undefined;
    /** Only a settled, `claimed` guild shows the officer glance strip here -- an
     *  unclaimed/pending/contested guild's claim controls are the global role-line's own
     *  job (Guild.svelte, unchanged v1 mechanism), not this card's. */
    claimed: boolean;
    claimedByName: string | null;
    onSelectTab: (tab: GuildTabId) => void;
  } = $props();

  let readinessSummary = $state<{ name: string; failing: number }[] | null>(null);
  let lootSummary = $state<string | null>(null);

  $effect(() => {
    if (guildId === null || role === 'public') return;
    void fetchGuildReadiness(guildId)
      .then((page) => {
        readinessSummary = sortReadinessWorstFirst(page.rows)
          .slice(0, 3)
          .map((row) => ({ name: row.name, failing: row.failing }));
      })
      .catch(() => {
        readinessSummary = null;
      });
    void fetchGuildLoot(guildId)
      .then((page) => {
        const next = nextUnkilledEncounter(page.encounters);
        const current = page.encounters.find((e) => e.encounter_id === page.selected);
        const pick = page.items[0]?.name;
        lootSummary = `${current?.name ?? bossPickerLabel(next)}${pick !== undefined ? ` · next pick: ${pick}` : ''}`;
      })
      .catch(() => {
        lootSummary = null;
      });
  });

  function wipeCount(report: OverviewReport): number | undefined {
    return report.fight_count === undefined || report.kill_count === undefined
      ? undefined
      : Math.max(0, report.fight_count - report.kill_count);
  }
</script>

<section class="flex flex-col gap-5" data-testid="guild-overview-tab">
  {#if role === 'officer' && guildId !== null && claimed}
    <GuildOfficerStrip
      {guildId}
      {claimedByName}
      waitingCount={summary?.waiting_for_approval ?? 0}
      onSelectRoster={() => onSelectTab('roster')}
    />
  {/if}

  {#if role !== 'public'}
    <div class="guild-overview-card" data-testid="guild-overview-card-roster">
      <div class="flex items-center justify-between">
        <h3 class="section-title text-[16px]">Roster</h3>
        <button class="guild-see-all" onclick={() => onSelectTab('roster')}>See all</button>
      </div>
      <p class="text-[14px]">
        <span class="tabular font-mono">{summary?.raiders ?? 0}</span> raiders ·
        <span class="tabular text-gold font-mono">{summary?.waiting_for_approval ?? 0}</span> waiting for
        approval ·
        <span class="tabular font-mono">{summary?.below_rating_floor ?? 0}</span> below the rating floor
      </p>
    </div>
  {/if}

  <div class="guild-overview-card" data-testid="guild-overview-card-raids">
    <div class="flex items-center justify-between">
      <h3 class="section-title text-[16px]">This week's raid nights</h3>
      <button class="guild-see-all" onclick={() => onSelectTab('raids')}>See all</button>
    </div>
    {#if reports.length === 0}
      <p class="text-muted text-[14px]">No reports this week yet.</p>
    {:else}
      {#each reports.slice(0, 3) as report (report.id)}
        <div class="border-line-soft flex items-center gap-3 border-b py-1.5">
          <span class="flex-1 text-[14px] font-semibold"
            >{report.title === '' ? report.zone : report.title}</span
          >
          <span class="text-muted tabular font-mono text-[13px]">{report.created_at.slice(0, 10)}</span>
          {#if report.kill_count !== undefined && wipeCount(report) !== undefined}
            <span class="text-muted tabular font-mono text-[13px]">
              {report.kill_count} kill{report.kill_count === 1 ? '' : 's'} · {wipeCount(report)} wipe{wipeCount(
                report,
              ) === 1
                ? ''
                : 's'}
            </span>
          {/if}
        </div>
      {/each}
    {/if}
  </div>

  <div class="guild-overview-card" data-testid="guild-overview-card-progression">
    <div class="flex items-center justify-between">
      <h3 class="section-title text-[16px]">Progression</h3>
      <button class="guild-see-all" onclick={() => onSelectTab('progression')}>See all</button>
    </div>
    <p class="text-[14px]">
      <span class="tabular font-mono">{summary?.named_encounters_down ?? 0}</span> named encounter{(summary?.named_encounters_down ??
        0) === 1
        ? ''
        : 's'}
      down this tier. Barrow Deeps and Hyjal Summit have no published encounter list yet —
      <span class="tabular font-mono">{summary?.pulls_this_tier ?? 0}</span> pulls logged there regardless.
    </p>
  </div>

  {#if role !== 'public'}
    <div class="guild-overview-card" data-testid="guild-overview-card-readiness">
      <div class="flex items-center justify-between">
        <h3 class="section-title text-[16px]">Readiness</h3>
        <button class="guild-see-all" onclick={() => onSelectTab('readiness')}>See all</button>
      </div>
      {#if readinessSummary === null}
        <p class="text-muted text-[13px]">No readiness data yet.</p>
      {:else}
        {#each readinessSummary as row (row.name)}
          <p class="text-[13px]">
            <span class="font-semibold">{row.name}</span> · {row.failing} checks failing
          </p>
        {/each}
      {/if}
    </div>

    <div class="guild-overview-card" data-testid="guild-overview-card-loot">
      <div class="flex items-center justify-between">
        <h3 class="section-title text-[16px]">Loot</h3>
        <button class="guild-see-all" onclick={() => onSelectTab('loot')}>See all</button>
      </div>
      <p class="text-[14px]">{lootSummary ?? 'No bosses killed yet this tier.'}</p>
    </div>
  {/if}
</section>

<style>
  .guild-overview-card {
    display: flex;
    flex-direction: column;
    gap: 10px;
    background: var(--color-raised);
    border: 1px solid var(--color-border);
    border-radius: 6px;
    padding: 16px 18px;
  }
  .guild-see-all {
    font-size: 12px;
    font-weight: 700;
    letter-spacing: 0.06em;
    text-transform: uppercase;
    color: var(--color-gold);
    background: none;
    border: none;
  }
</style>
