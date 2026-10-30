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
     regardless of which tab is open rather than being Overview-only content. Fix round 1
     (2026-10-04): the officer strip (claim badge, approvals, invite) moved out of this
     card for the same reason -- it is now part of Guild.svelte's own full-bleed band, so
     it, too, stays visible on every tab rather than Overview-only. -->
<script lang="ts">
  import type { GuildHomeSummary, GuildRaidsPage } from '../../lib/guild/api';
  import {
    fetchGuildLoot,
    fetchGuildProgression,
    fetchGuildRaids,
    fetchGuildReadiness,
  } from '../../lib/guild/api';
  import {
    NO_PUBLIC_REPORTS,
    checksFailingLine,
    encountersDownLine,
    raidNightDay,
    rosterCardParts,
  } from '../../lib/guild/overview-view';
  import { bossIdsOf, raidTallyParts } from '../../lib/guild/raid-view';
  import { sortReadinessWorstFirst } from '../../lib/guild/readiness-view';
  import { counted } from '../../lib/plural';
  import { ACTION_LINK, ACTION_LINK_CHEVRON } from '../../lib/ui/link';
  import Card from '../ui/Card.svelte';
  import Glyph from '../ui/Glyph.svelte';
  import SettingsRow from '../ui/SettingsRow.svelte';
  import FactLine from './FactLine.svelte';
  import { bossPickerLabel, nextUnkilledEncounter } from '../../lib/guild/loot-view';
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
    showCounts = true,
    onSelectTab,
  }: {
    guildId: number | null;
    role: 'public' | 'member' | 'officer' | 'moderator';
    reports: readonly OverviewReport[];
    summary: GuildHomeSummary | undefined;
    /** False for a visitor when the only report is guild-only: the progression card then says
     *  what the public tabs say instead of counting what the visitor cannot open. */
    showCounts?: boolean;
    onSelectTab: (tab: GuildTabId) => void;
  } = $props();

  let readinessSummary = $state<{ name: string; failing: number }[] | null>(null);
  let lootSummary = $state<string | null>(null);
  /** Live defect 2026-10-05: a public visitor has no `home` (401), so `summary` is undefined
   *  and this card read "0 named encounters down" under a band saying 1 boss down. The
   *  progression endpoint is public and already carries the tier figures, so the card reads
   *  it for every role and falls back to `summary` only while it loads. */
  let progressionSummary = $state<{ down: number; names: string[] } | null>(null);
  /** The progression read names the encounters; until it answers the home's own count stands in. */
  const encountersLine = $derived(
    progressionSummary !== null
      ? encountersDownLine(progressionSummary.names)
      : `${counted(summary?.named_encounters_down ?? 0, 'encounter')} down this tier.`,
  );
  /** The raid nights' fights, so a night reads as bosses, kills and trash packs. */
  let raidsPage = $state<GuildRaidsPage | null>(null);
  let bossIds = $state<ReadonlySet<number> | null>(null);

  $effect(() => {
    if (guildId === null) return;
    void fetchGuildProgression(guildId)
      .then((page) => {
        progressionSummary = { down: page.tier.down, names: page.encounters.map((e) => e.name) };
        bossIds = bossIdsOf(page);
      })
      .catch(() => {
        progressionSummary = null;
      });
  });

  $effect(() => {
    if (guildId === null) return;
    void fetchGuildRaids(guildId)
      .then((page) => {
        raidsPage = page;
      })
      .catch(() => {
        raidsPage = null;
      });
  });

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

  /** The night's facts, one piece each ("1 boss (Onyxia)", "1 kill", "2 trash packs", "1 wipe") when its
   *  fights are known, else the two counts the report row carries; none when it carries neither (the
   *  public list). */
  function tallyOf(report: OverviewReport): string[] {
    const night = raidsPage?.rows.find((row) => row.id === report.id);
    if (night !== undefined) return raidTallyParts(night, bossIds);
    const wipes = wipeCount(report);
    return report.kill_count !== undefined && wipes !== undefined
      ? [counted(report.kill_count, 'kill'), counted(wipes, 'wipe')]
      : [];
  }

  const rosterParts = $derived(
    rosterCardParts({
      raiders: summary?.raiders ?? 0,
      waiting_for_approval: summary?.waiting_for_approval ?? 0,
      below_rating_floor: summary?.below_rating_floor ?? 0,
    }),
  );
</script>

{#snippet seeAll(tab: GuildTabId)}
  <button type="button" class={ACTION_LINK} onclick={() => onSelectTab(tab)}>
    See all<Glyph name="chevron" size={16} class={ACTION_LINK_CHEVRON} />
  </button>
{/snippet}

<section class="flex flex-col gap-[22px] lg:gap-6" data-testid="guild-overview-tab">
  {#if role !== 'public'}
    <Card glyph="user" title="Roster" testid="guild-overview-card-roster">
      {#snippet aside()}{@render seeAll('roster')}{/snippet}
      <p data-testid="guild-overview-roster-facts">
        <FactLine
          facts={[
            rosterParts.raiders,
            { text: rosterParts.waiting, class: (summary?.waiting_for_approval ?? 0) > 0 ? 'text-gold' : '' },
            rosterParts.belowFloor,
          ]}
        />
      </p>
    </Card>
  {/if}

  <Card
    glyph="chart"
    title="This week's raid nights"
    body={reports.length === 0 ? 'text' : 'rows'}
    testid="guild-overview-card-raids"
  >
    {#snippet aside()}{@render seeAll('raids')}{/snippet}
    {#if reports.length === 0}
      <p class="text-muted">{role === 'public' ? NO_PUBLIC_REPORTS : 'No reports this week yet.'}</p>
    {:else}
      {#each reports.slice(0, 3) as report (report.id)}
        {@const facts = tallyOf(report)}
        <SettingsRow
          label={report.title === '' ? report.zone : report.title}
          value={raidNightDay(report.created_at)}
          testid="guild-overview-raid-row"
        >
          {#snippet below()}
            {#if facts.length > 0}
              <FactLine {facts} class="text-muted tabular text-[13px]" testid="guild-overview-raid-facts" />
            {/if}
          {/snippet}
        </SettingsRow>
      {/each}
    {/if}
  </Card>

  <Card glyph="trophy" title="Progression" testid="guild-overview-card-progression">
    {#snippet aside()}{@render seeAll('progression')}{/snippet}
    {#if showCounts}
      <p>{encountersLine}</p>
    {:else}
      <p class="text-muted">{NO_PUBLIC_REPORTS}</p>
    {/if}
  </Card>

  {#if role !== 'public'}
    <Card
      glyph="guild"
      title="Readiness"
      body={readinessSummary === null ? 'text' : 'rows'}
      testid="guild-overview-card-readiness"
    >
      {#snippet aside()}{@render seeAll('readiness')}{/snippet}
      {#if readinessSummary === null}
        <p class="text-muted">No readiness data yet.</p>
      {:else}
        {#each readinessSummary as row (row.name)}
          <SettingsRow label={row.name} value={checksFailingLine(row.failing)} />
        {/each}
      {/if}
    </Card>

    <Card glyph="star" title="Loot" testid="guild-overview-card-loot">
      {#snippet aside()}{@render seeAll('loot')}{/snippet}
      <p>{lootSummary ?? 'No bosses killed yet this tier.'}</p>
    </Card>
  {/if}
</section>
