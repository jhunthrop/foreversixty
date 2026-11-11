<!-- web/src/components/guild/dense/OfficerWeek.svelte -->
<!-- The officer's Overview, dense (design/specs/2026-11-09-guild-dense.md section 5.1): what is
     waiting as one line of buttons, Wednesday against the comp target with its fillers, who needs a
     word, last night's rating risers and fallers; the rail: roles to recruit for, last night's loot,
     the buff cover and who is slipping. Reads the next raid (given), its readiness board, the last
     two logged raids, the attendance grid, the loot history and the open applications. -->
<script lang="ts">
  import type { Snippet } from 'svelte';
  import type {
    GuildRaidView,
    GuildRaidsPage,
    GuildReadinessRow,
    GuildRosterRow,
    NextRaidView,
  } from '../../../lib/guild/api';
  import { fetchGuildReadiness } from '../../../lib/guild/api';
  import {
    guildApplicationsHref,
    guildMemberHref,
    guildRaidHref,
    guildRecruitHref,
  } from '../../../lib/characters';
  import {
    fetchDropCandidates,
    fetchLootNights,
    type LootCouncilCandidate,
    type LootNights,
  } from '../../../lib/guild/loot-council-api';
  import {
    fetchRecentAttendance,
    pingRaid,
    type GuildOfficerActions,
    type GuildRecentAttendance,
  } from '../../../lib/guild/overview-api';
  import { fetchGuildRaid } from '../../../lib/guild/raids-api';
  import { fetchOfficerApplications, type OfficerApplication } from '../../../lib/guild/recruit-api';
  import { nudgeText } from '../../../lib/guild/readiness-view';
  import { counted } from '../../../lib/plural';
  import {
    compGaps,
    needsAWord,
    ratingMovers,
    recruitGaps,
    slipping,
    type Word,
    type WordAction,
  } from '../../../lib/guild/dense/week-board';
  import Button from '../../ui/Button.svelte';
  import Page from '../../ui/Page.svelte';
  import BuffCover from '../overview/BuffCover.svelte';
  import type { GuildTabId } from '../GuildTabs.svelte';
  import { weekdayDayMonth } from '../../../lib/dates';
  import { clockDuration } from '../../../lib/format/duration';
  import LastNightMovers from './LastNightMovers.svelte';
  import NeedsAWord from './NeedsAWord.svelte';
  import OfficerRail from './OfficerRail.svelte';
  import WednesdayPanel from './WednesdayPanel.svelte';

  let {
    guildId,
    next,
    raid,
    roster,
    raidsPage,
    actions,
    guildPath,
    planHref,
    viewerTz,
    now,
    onSelectTab,
    onRaidChanged,
    lead,
  }: {
    guildId: number;
    next: NextRaidView | null;
    raid: GuildRaidView | null;
    roster: readonly GuildRosterRow[];
    raidsPage: GuildRaidsPage | null;
    actions: GuildOfficerActions | null;
    guildPath: string;
    planHref: string;
    viewerTz: string;
    now: Date;
    onSelectTab: (tab: GuildTabId) => void;
    onRaidChanged: () => void;
    /** The pre-raid lead (the countdown, the first week, the activity), heading the main column. */
    lead?: Snippet;
  } = $props();

  const NIGHTS = 5;
  let readiness = $state<readonly GuildReadinessRow[]>([]);
  let grid = $state<GuildRecentAttendance | null>(null);
  let loot = $state<LootNights | null>(null);
  let apps = $state<readonly OfficerApplication[]>([]);
  let last = $state<GuildRaidView | null>(null);
  let before = $state<GuildRaidView | null>(null);

  const logged = $derived((raidsPage?.rows ?? []).flatMap((r) => (r.raid_id == null ? [] : [r.raid_id])));

  $effect(() => {
    void fetchRecentAttendance(guildId, NIGHTS, true)
      .then((g) => (grid = g))
      .catch(() => (grid = null));
    void fetchLootNights(guildId)
      .then((l) => (loot = l))
      .catch(() => (loot = null));
    void fetchOfficerApplications(guildId)
      .then((a) => (apps = a.rows))
      .catch(() => (apps = []));
  });
  $effect(() => {
    if (next === null) return;
    void fetchGuildReadiness(guildId, next.id)
      .then((p) => (readiness = p.rows))
      .catch(() => (readiness = []));
  });
  $effect(() => {
    const [a, b] = logged;
    if (a !== undefined)
      void fetchGuildRaid(guildId, a)
        .then((v) => (last = v))
        .catch(() => (last = null));
    if (b !== undefined)
      void fetchGuildRaid(guildId, b)
        .then((v) => (before = v))
        .catch(() => (before = null));
  });

  const movers = $derived(ratingMovers(last?.report?.parses ?? [], before?.report?.parses ?? [], roster));
  const gaps = $derived(raid === null ? [] : compGaps(raid));
  const tz = $derived(next?.guild_timezone ?? viewerTz);
  const dayOf = (iso: string) => weekdayDayMonth(new Date(iso), tz);
  const words = $derived(
    needsAWord({
      view: raid,
      roster,
      grid,
      readiness,
      movers,
      now,
      lastRaidLabel: last === null ? undefined : dayOf(last.starts_at),
    }),
  );
  const ready = $derived(
    readiness.length === 0
      ? null
      : { ready: readiness.filter((r) => r.failing === 0).length, coming: readiness.length },
  );
  const lastMeta = $derived.by(() => {
    if (last === null || last.report === null) return '';
    const kill = last.report.kill_order[0];
    const parts = [
      kill === undefined || kill.kill_ms === null ? '' : `${kill.name} ${clockDuration(kill.kill_ms)}`,
      counted(last.report.headline.wipes, 'wipe'),
      counted(last.report.headline.deaths, 'death'),
      last.night === null ? '' : counted(last.night.came, 'raider'),
    ];
    return parts.filter((p) => p !== '').join(' · ');
  });
  const lootNight = $derived(loot?.nights[0] ?? null);
  let topOf = $state<ReadonlyMap<number, LootCouncilCandidate>>(new Map());
  $effect(() => {
    const night = lootNight;
    if (night === null) return;
    const open = night.drops.filter((d) => d.state === 'open');
    void Promise.all(
      open.map((d) =>
        fetchDropCandidates(guildId, d.item_id, night.night)
          .then((drop) => [d.item_id, drop.candidates[0]] as const)
          .catch(() => null),
      ),
    ).then((pairs) => {
      topOf = new Map(
        pairs.flatMap((p) => (p === null || p[1] === undefined ? [] : [[p[0], p[1]] as const])),
      );
    });
  });
  const rosterByKey = $derived(new Map(roster.map((r) => [r.character_key, r] as const)));

  function hrefOf(word: Word, action: WordAction): string {
    if (action === 'report' && last !== null) return guildRaidHref(guildPath, last.id);
    const account = rosterByKey.get(word.key)?.account_key;
    return account === undefined ? '#roster' : guildMemberHref(guildPath, account);
  }

  async function ping(): Promise<void> {
    if (raid !== null) await pingRaid(guildId, raid.id).catch(() => undefined);
  }

  function nudgeTextOf(word: Word): string {
    const row = readiness.find((r) => r.character_key === word.key);
    return row === undefined ? '' : nudgeText(row);
  }
</script>

{#if actions !== null && (actions.approvals.count > 0 || (actions.applications?.count ?? 0) > 0 || (actions.drops_to_award?.count ?? 0) > 0)}
  <div
    class="flex flex-wrap items-center gap-x-4 gap-y-1 px-(--page-inset) pb-5"
    data-testid="guild-overview-actions"
  >
    {#if actions.approvals.count > 0}
      <Button tone="ghost" size="sm" onclick={() => onSelectTab('roster')}
        >{actions.approvals.count} waiting for approval</Button
      >
    {/if}
    {#if (actions.applications?.count ?? 0) > 0}
      <Button tone="ghost" size="sm" href={guildApplicationsHref(guildPath)} testid="guild-open-applications"
        >{counted(actions.applications?.count ?? 0, 'application')}</Button
      >
    {/if}
    {#if (actions.drops_to_award?.count ?? 0) > 0}
      <Button tone="ghost" size="sm" onclick={() => onSelectTab('loot')}
        >{actions.drops_to_award?.count} drops to award</Button
      >
    {/if}
  </div>
{/if}

<Page layout="split" panels testid="guild-overview-page">
  {@render lead?.()}
  {#if next !== null && raid !== null}
    <WednesdayPanel {guildId} {next} {raid} {ready} {planHref} {viewerTz} {now} onChanged={onRaidChanged} />
  {/if}
  <NeedsAWord {words} {hrefOf} onPing={() => void ping()} {nudgeTextOf} />
  {#if last !== null && before !== null}
    <LastNightMovers
      meta={lastMeta}
      vsLabel={`vs ${dayOf(before.starts_at).replace(/^\w+ /, '')}`}
      up={movers.up}
      down={movers.down}
      reportHref={guildRaidHref(guildPath, last.id)}
      title={dayOf(last.starts_at)}
    />
  {/if}
  {#snippet rail()}
    <OfficerRail
      gaps={recruitGaps(gaps)}
      applications={apps}
      applicationsHref={guildApplicationsHref(guildPath)}
      recruitHref={guildRecruitHref(guildPath)}
      night={lootNight}
      nightLabel={lootNight === null ? '' : dayOf(lootNight.at)}
      lootHref="#loot"
      slipping={slipping(grid)}
      {topOf}
    />
    {#if raid?.cover && raid.cover.length > 0}<BuffCover cover={raid.cover} />{/if}
  {/snippet}
</Page>
