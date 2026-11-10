<!-- web/src/components/guild/dense/MemberWeek.svelte -->
<!-- A member's week on the Overview (dense spec 5.2): Your week (the next raid, your own answer, and
     your standing in figures: item level and its place among your spec, last night's rating against
     the night before and your parse, nights and streak, items this tier), then your own line on the
     last two nights. `loot` is the rail's Your loot. Reads the member's own profile and the last two
     logged raids (a member's report carries only their own parses). -->
<script lang="ts">
  import type { Snippet } from 'svelte';
  import activeBuild from '../../../data/active-build.json';
  import type {
    GuildHomeStanding,
    GuildRaidView,
    GuildRaidsPage,
    GuildRosterRow,
    NextRaidView,
  } from '../../../lib/guild/api';
  import { fetchMemberProfile, type MemberProfile } from '../../../lib/guild/recruit-api';
  import { fetchGuildRaid } from '../../../lib/guild/raids-api';
  import { fetchLootNights } from '../../../lib/guild/loot-council-api';
  import { lootIconOf } from '../../../lib/guild/loot-icon';
  import { percentileToken } from '../../../lib/report/format';
  import { guildPerNight, myNight, streak, type MyNight } from '../../../lib/guild/dense/member-week';
  import Figure from '../../ui/Figure.svelte';
  import ItemIcon from '../../ui/ItemIcon.svelte';
  import Row from '../../ui/Row.svelte';
  import Section from '../../ui/Section.svelte';
  import StatRow from '../../ui/StatRow.svelte';
  import StatTable from '../../ui/StatTable.svelte';
  import NextRaidRow from '../overview/NextRaidRow.svelte';

  let {
    guildId,
    next,
    raid,
    myRow,
    roster,
    standing,
    raidsPage,
    viewerTz,
    now,
    you,
    part,
  }: {
    guildId: number;
    next: NextRaidView | null;
    raid: GuildRaidView | null;
    myRow: GuildRosterRow | null;
    roster: readonly GuildRosterRow[];
    standing: GuildHomeStanding | null;
    raidsPage: GuildRaidsPage | null;
    viewerTz: string;
    now: Date;
    you?: Snippet;
    /** Which block to draw: the week panel, your nights, or the rail's loot. */
    part: 'week' | 'nights' | 'loot';
  } = $props();

  const build = activeBuild.build;
  let profile = $state<MemberProfile | null>(null);
  let views = $state<GuildRaidView[]>([]);
  let tierItems = $state<number | null>(null);

  $effect(() => {
    void fetchMemberProfile(guildId, 'me')
      .then((p) => (profile = p))
      .catch(() => (profile = null));
    void fetchLootNights(guildId)
      .then((l) => (tierItems = new Set(l.nights.flatMap((n) => n.decisions.map((d) => d.award_id))).size))
      .catch(() => (tierItems = null));
  });
  $effect(() => {
    const ids = (raidsPage?.rows ?? []).flatMap((r) => (r.raid_id == null ? [] : [r.raid_id])).slice(0, 2);
    void Promise.all(ids.map((id) => fetchGuildRaid(guildId, id)))
      .then((list) => (views = list))
      .catch(() => (views = []));
  });

  const keys = $derived(new Set(myRow === null ? [] : [myRow.character_key]));
  const nights = $derived(views.map((v) => myNight(v, keys)).filter((n): n is MyNight => n !== null));
  const last = $derived(nights[0] ?? null);
  const before = $derived(nights[1] ?? null);
  const tz = $derived(next?.guild_timezone ?? viewerTz);
  const day = (iso: string) =>
    new Date(iso).toLocaleDateString('en-GB', {
      weekday: 'short',
      day: 'numeric',
      month: 'short',
      timeZone: tz,
    });
  const ordinal = (n: number) => `${n}${n === 1 ? 'st' : n === 2 ? 'nd' : n === 3 ? 'rd' : 'th'}`;
  const ratingChange = $derived(
    last?.rating == null || before?.rating == null ? null : Math.round(last.rating - before.rating),
  );
  const runOf = $derived(profile?.attendance == null ? null : streak(profile.attendance.nights));
  const att = $derived(profile?.attendance?.last_30 ?? myRow?.attendance_30 ?? null);
  const myPerNight = $derived(
    profile === null || att === null || att.present === 0 ? null : profile.loot.count / att.present,
  );
  const guildRate = $derived(tierItems === null ? null : guildPerNight(roster, tierItems));
  const signed = (n: number) => (n > 0 ? `+${n}` : n < 0 ? `−${-n}` : '±0');
  const mins = (ms: number | null) =>
    ms === null
      ? ''
      : `kill ${Math.floor(ms / 60000)}:${String(Math.floor(ms / 1000) % 60).padStart(2, '0')}`;
</script>

{#if part === 'week'}
  <Section title="Your week" panel frame="framed" testid="guild-your-week">
    {#if next !== null}<NextRaidRow
        {next}
        durationMinutes={raid?.duration_minutes ?? 180}
        {viewerTz}
        {now}
      />{/if}
    {#if you}<div data-testid="guild-next-raid-you">{@render you()}</div>{/if}
    <div class="flex flex-wrap gap-x-10 gap-y-4 pt-4" data-testid="guild-your-figures">
      {#if myRow?.item_level != null}
        <Figure
          value={String(myRow.item_level)}
          label="Item level"
          caption={standing === null
            ? undefined
            : `${standing.spec} ${standing.class} ${ordinal(standing.rank_by_item_level)} of ${standing.same_spec_count}`}
        />
      {/if}
      {#if last?.rating != null}
        <Figure
          value={String(Math.round(last.rating))}
          label="Rating"
          caption={ratingChange === null
            ? undefined
            : `${signed(ratingChange)} vs ${day(before?.startsAt ?? '')}`}
        />
      {/if}
      {#if last?.percentile != null}
        <Figure
          value={String(Math.floor(last.percentile))}
          label="Parse %"
          caption={`${last.boss}, ${day(last.startsAt)}`}
          color={percentileToken(last.percentile)}
        />
      {/if}
      {#if att !== null}
        <Figure
          value={`${att.present}/${att.nights}`}
          label="Nights"
          tone="kill"
          caption={runOf === null ? undefined : `streak ${runOf}`}
        />
      {/if}
      {#if profile !== null}<Figure
          value={String(profile.loot.count)}
          label="Items"
          caption="this tier"
        />{/if}
    </div>
  </Section>
{:else if part === 'nights'}
  {#if nights.length > 0}
    <Section title={`Your ${last?.boss ?? 'nights'}`} panel flush testid="guild-your-nights">
      <StatTable
        columns={[
          { id: 'value', label: (last?.metric ?? 'dps').toUpperCase() },
          { id: 'rating', label: 'Rating' },
          { id: 'pct', label: 'Parse %' },
          { id: 'deaths', label: 'Deaths' },
          { id: 'vs', label: `${(last?.metric ?? 'dps').toUpperCase()} vs last` },
        ]}
        nameHead="Night"
        bar={false}
        dense
        label="Your nights"
      >
        {#each nights as n, i (n.raidId)}
          {@const prev = nights[i + 1]}
          <StatRow
            bar={false}
            figures={[
              n.value === null
                ? { value: '–', muted: true }
                : { value: Math.round(n.value).toLocaleString('en-US') },
              n.rating === null ? { value: '–', muted: true } : { value: String(Math.round(n.rating)) },
              n.percentile === null
                ? { value: '–', muted: true }
                : { value: String(Math.floor(n.percentile)), color: percentileToken(n.percentile) },
              { value: String(n.deaths), tone: n.deaths > 0 ? 'wipe' : undefined },
              prev?.value == null || n.value === null
                ? { value: '–', muted: true }
                : {
                    value: signed(Math.round(n.value - prev.value)),
                    tone: n.value >= prev.value ? 'kill' : 'wipe',
                  },
            ]}
          >
            {#snippet name()}{day(n.startsAt)}{/snippet}
            {#snippet sub()}{n.boss} {mins(n.killMs)}{/snippet}
          </StatRow>
        {/each}
      </StatTable>
    </Section>
  {/if}
{:else if profile !== null && profile.loot.count > 0}
  <Section title="Your loot" panel testid="guild-your-loot">
    {#each profile.loot.items as item (item.item_id + item.awarded_at)}
      <Row title={item.name ?? 'Item'} sub={day(item.awarded_at)} open>
        {#snippet lead()}<ItemIcon
            icon={item.icon === null ? undefined : (lootIconOf({ icon: item.icon }) ?? undefined)}
            quality={item.quality ?? 4}
            {build}
            size={32}
          />{/snippet}
      </Row>
    {/each}
    {#if myPerNight !== null && guildRate !== null}
      <div class="flex gap-x-10 pt-3">
        <Figure value={myPerNight.toFixed(2)} label="Yours per night" size="sm" />
        <Figure value={guildRate.toFixed(2)} label="Guild per night" size="sm" />
      </div>
    {/if}
  </Section>
{/if}
