<!-- web/src/components/guild/planner/RaidPlanner.svelte -->
<!-- The raid planner (web layout spec 2026-11-05 §7.2 "Raid planner"): "Who is in, in what comp, and
     what do I swap?", built only from the ui/ primitives. Page split: the head (zone, guild time and
     the viewer's, the count, the state Tag, the officer buttons, the guild tabs), the control bar
     (group by Status, Role or Class), then the viewer's own sign-up, the one Plate of sign-up Rows
     and, for officers, who has not signed up with Ping on Discord. Rail: comp against the target
     (officers edit it), Bench and swaps, the assignments per encounter, buff and debuff cover, the
     notes. Phone keeps that order
     in one column. A done raid is GuildRaid's report path, never this page. -->
<script lang="ts">
  import EmptyState from '../../ui/EmptyState.svelte';
  import {
    fetchGuildHome,
    fetchGuildReadiness,
    type GuildRaidView,
    type GuildRosterRow,
  } from '../../../lib/guild/api';
  import { untrack } from 'svelte';
  import { plannerFacts, type PlannerReads } from '../../../lib/guild/planner-facts';
  import { fetchMyChecklist, fetchRecentAttendance } from '../../../lib/guild/overview-api';
  import { ATTENDANCE_NIGHTS } from './planner-columns';
  import type { GateFacts } from '../../../lib/guild/plan-gate';
  import { focusOnMount } from '../../../lib/guild/focus-now';
  import {
    eyebrowOf,
    kindCountsRoles,
    kindOf,
    notesEmptyOf,
    raidCopy,
    SEP,
    zoneLineOf,
  } from '../../../lib/guild/raid-copy';
  import { raidCount, raidWhen, type RaidPageFlags } from '../../../lib/guild/raid-page';
  import type { RaidPageController } from '../../../lib/guild/raid-page-controller.svelte';
  import {
    GROUP_PARAM,
    defaultGrouping,
    groupingFrom,
    groupingOptions,
    groupingSearch,
    groupingSummary,
    plannerStatusTag,
    stateCountsLine,
    type PlannerGrouping,
  } from '../../../lib/guild/raid-planner-view';
  import { raidQuiet } from '../../../lib/guild/planned-row';

  import type { RaidStatusAction } from '../../../lib/guild/raids-api';
  import { clock, weekdayDayMonth } from '../../../lib/dates';
  import CompTally from '../overview/CompTally.svelte';
  import GuildTabs from '../GuildTabs.svelte';
  import RaidDoneBand from '../RaidDoneBand.svelte';
  import RaidForm from '../RaidForm.svelte';
  import RaidOfficerRow from '../RaidOfficerRow.svelte';
  import RaidSignupStrip from '../RaidSignupStrip.svelte';
  import ZoneKeyNote from '../ZoneKeyNote.svelte';
  import Button from '../../ui/Button.svelte';
  import ControlBar from '../../ui/ControlBar.svelte';
  import Page from '../../ui/Page.svelte';
  import PageHeader from '../../ui/PageHeader.svelte';
  import Section from '../../ui/Section.svelte';
  import Tag from '../../ui/Tag.svelte';
  import CompTargetForm from './CompTargetForm.svelte';
  import PlannerAssignments from './PlannerAssignments.svelte';
  import PlannerBench from './PlannerBench.svelte';
  import PlannerNotSigned from './PlannerNotSigned.svelte';
  import RaidDiscordPost from './RaidDiscordPost.svelte';
  import PlannerRoster from './PlannerRoster.svelte';

  let {
    c,
    raid,
    flags,
    now,
    viewerTimezone,
    locale,
    guildName,
    guildPath,
    settingsHref,
    gateFacts,
    gateId,
    onStatus,
    onCloseEdit,
  }: {
    c: RaidPageController;
    raid: GuildRaidView;
    flags: RaidPageFlags;
    now: Date;
    viewerTimezone: string;
    locale: string;
    guildName: string;
    guildPath: string;
    settingsHref: string;
    gateFacts: GateFacts;
    gateId: string;
    onStatus: (action: RaidStatusAction, post?: boolean) => Promise<boolean>;
    onCloseEdit: () => void;
  } = $props();

  const officer = $derived(c.role === 'officer');
  const kind = $derived(kindOf(raid));
  const roles = $derived(kindCountsRoles(kind));
  const when = $derived(raidWhen(raid, now, viewerTimezone, locale));
  const count = $derived(raidCount(raid));
  const guildTimezone = $derived(raid.guild_timezone ?? viewerTimezone);
  const status = $derived(plannerStatusTag(flags.status));
  const titleTone = $derived(c.guild?.faction ? { faction: c.guild.faction } : ('strong' as const));
  /** No gold on a cancelled raid or one already under way. */
  const quiet = $derived(raidQuiet(raid, now));
  const target = $derived(raid.comp_target ?? null);

  // svelte-ignore state_referenced_locally
  let grouping = $state<PlannerGrouping>(
    typeof location === 'undefined' ? defaultGrouping(raid) : groupingFrom(location.search, raid),
  );
  let editingComp = $state(false);

  /** The guild's roster (the /home read the page already made, from cache) and, for officers with
   *  the Raid guild tools, this raid's readiness rows: each sign-up's figures. */
  let roster = $state<readonly GuildRosterRow[]>([]);
  let reads = $state<PlannerReads>({ readiness: null, attendance: null, checklist: null });
  $effect(() => {
    const guildId = c.guild?.id;
    if (guildId === undefined) return;
    const raidId = raid.id;
    const ended = flags.ended;
    void untrack(async () => {
      const home = await fetchGuildHome(guildId).catch(() => null);
      if (home === null) return;
      roster = home.roster;
      if (ended) return;
      if (!officer) {
        const checklist = await fetchMyChecklist(guildId, raidId).catch(() => null);
        reads = { ...reads, checklist };
        return;
      }
      const [attendance, readiness] = await Promise.all([
        fetchRecentAttendance(guildId, ATTENDANCE_NIGHTS, true).catch(() => null),
        home.viewer?.raid_guild_tools === true
          ? fetchGuildReadiness(guildId, raidId)
              .then((page) => page.rows)
              .catch(() => null)
          : Promise.resolve(null),
      ]);
      reads = { ...reads, attendance, readiness };
    });
  });
  const facts = $derived(roster.length === 0 ? null : plannerFacts(roster, reads));

  /** The URL is the truth: a reload, a share and Back restore the grouping (spec 3.4). */
  function chooseGrouping(param: string, value: string): void {
    if (param !== GROUP_PARAM) return;
    grouping = groupingFrom(`?${GROUP_PARAM}=${value}`, raid);
    const search = groupingSearch(location.search, grouping, raid);
    history.replaceState(history.state, '', `${location.pathname}${search}${location.hash}`);
  }

  const nextPingClock = (at: Date): string => clock(at, viewerTimezone, { locale });
  /** Owner ruling 2026-11-09: the raid calendar and the sign-ups are free for any claimed guild's
   *  officers (they feed the free Discord bot); the plan gates only the raid tools below them. */
  const CALENDAR_FREE = true;
  const postedWhen = (at: Date): string =>
    `${weekdayDayMonth(at, viewerTimezone)}, ${clock(at, viewerTimezone, { locale })}`;
</script>

{#snippet meta()}
  <Tag tone={status.tone} testid="raid-status-chip">{status.text}</Tag>
  {#if kind !== 'raid'}<span data-testid="raid-kind">{eyebrowOf(kind)}</span>{/if}
  {#if raid.title !== raid.zone_name}<span data-testid="raid-title-zone">{zoneLineOf(raid)}</span>{/if}
  <span data-testid="raid-when"
    >{when.date}{SEP}{when.start} to {when.end}{#if when.yourTime !== null}{SEP}{when.yourTime}{/if}{SEP}<span
      class={quiet ? 'text-muted' : 'text-gold'}>{when.countdown}</span
    ></span
  >
  {#if raid.night === null}
    <span data-testid="raid-count"
      ><span class="text-strong">{count.line}</span>{#if count.line !== ''}{SEP}{stateCountsLine(
          raid.counts,
        )}{/if}</span
    >
  {/if}
{/snippet}

{#snippet header()}
  <PageHeader title={raid.title} tone={titleTone} {meta} testid="raid-head">
    {#snippet crumb()}<a href={`${guildPath}#raids`} data-testid="raid-back-link"
        >{raidCopy.backLink(guildName)}</a
      >{/snippet}
    {#snippet actions()}
      {#if officer && !c.editing}
        <RaidOfficerRow
          {raid}
          ended={flags.ended}
          readOnly={!CALENDAR_FREE}
          busy={c.busy}
          line={c.line}
          {gateId}
          {gateFacts}
          {guildTimezone}
          onEdit={() => c.openEdit()}
          {onStatus}
        />
      {/if}
    {/snippet}
    {#snippet tabs()}
      <GuildTabs active="raids" role={officer ? 'officer' : 'member'} hrefBase={guildPath} />
    {/snippet}
  </PageHeader>
{/snippet}

{#snippet bar()}
  <ControlBar
    label="Sign-up grouping"
    groups={[
      {
        label: 'Group by',
        param: GROUP_PARAM,
        kind: 'segmented',
        options: groupingOptions(raid),
        value: grouping,
      },
    ]}
    summary={groupingSummary(grouping)}
    onchange={chooseGrouping}
    id="raid-planner-bar"
    testid="raid-planner-bar"
  />
{/snippet}

<!-- One block: the guild shell spaces its children, and the head, bar and page sit flush as the mock draws them. -->
<div>
  <Page layout="split" panels {header} {bar} testid="raid-page">
    {#if flags.cancelled && c.raceLine !== null}
      <p
        class="text-wipe text-[13px]"
        role="alert"
        tabindex="-1"
        use:focusOnMount
        data-testid="raid-race-line"
      >
        {c.raceLine}
      </p>
    {/if}
    {#if c.editing}
      <RaidForm
        mode="edit"
        zones={c.zones}
        guildTimezone={raid.guild_timezone}
        {viewerTimezone}
        {raid}
        plannedCount={0}
        discord={raid.discord}
        {settingsHref}
        onSubmit={async (payload) => {
          const result = await c.saveEdit(payload);
          if (result.ok) onCloseEdit();
          return result;
        }}
        onCancel={onCloseEdit}
        {now}
        {locale}
      />
    {/if}
    {#if flags.ended || raid.match !== null}
      <RaidDoneBand
        {raid}
        {now}
        {guildPath}
        {viewerTimezone}
        {guildTimezone}
        {locale}
        {officer}
        readOnly={!c.officerTools}
        {gateId}
        loadCandidates={c.loadCandidates}
        onMatch={c.matchLog}
        onClear={c.clearMatch}
      />
    {/if}
    {#if raid.key_item != null}<ZoneKeyNote item={raid.key_item} testid="raid-key" />{/if}
    {#if c.pollPaused}<p class="text-muted text-[13px]" data-testid="raid-poll-paused">
        {raidCopy.poll}
      </p>{/if}
    <RaidSignupStrip
      {raid}
      locked={flags.locked}
      ended={flags.ended}
      cancelled={flags.cancelled}
      onSignup={c.signup}
      currentKey={c.currentKey}
    />
    <Section title={raidCopy.subheads.signups} panel meta={count.ratio} testid="raid-signups">
      <PlannerRoster
        {raid}
        {grouping}
        {officer}
        officerTools={CALENDAR_FREE}
        {gateId}
        closed={flags.cancelled}
        busyKey={c.rosterBusy}
        line={c.line}
        {facts}
        onSetState={(row, state) => c.setRosterState(row, state)}
      />
    </Section>
    {#if officer && raid.discord.connected && !flags.cancelled && !flags.ended && raid.status !== 'done'}
      <RaidDiscordPost
        {raid}
        when={postedWhen}
        officerTools={c.officerTools}
        posting={c.planner.posting}
        answer={c.planner.postAnswer}
        onPost={() => void c.planner.postNow()}
      />
    {/if}
    {#if officer}
      <PlannerNotSigned
        {raid}
        {now}
        clock={nextPingClock}
        officerTools={c.officerTools}
        closed={flags.cancelled || flags.ended}
        busyKey={c.rosterBusy}
        pinging={c.planner.pinging}
        pingAnswer={c.planner.pingAnswer}
        {facts}
        onPing={() => void c.planner.ping()}
        onSetState={(row, state) => c.setRosterState(row, state)}
      />
    {/if}

    {#snippet rail()}
      {#if target !== null && roles}
        <CompTally {raid} {target}>
          {#snippet edit()}
            {#if officer && editingComp}
              <CompTargetForm
                {target}
                disabled={!c.officerTools}
                busy={c.planner.compBusy}
                fields={c.planner.compFields}
                line={c.planner.compLine}
                onSave={c.planner.saveComp}
                onReset={c.planner.resetComp}
                onRefuse={(fields) => c.planner.refuseComp(fields)}
                onDone={() => (editingComp = false)}
              />
            {:else if officer}
              <Button
                size="sm"
                disabled={!c.officerTools || flags.cancelled}
                onclick={() => (editingComp = true)}
                testid="raid-comp-edit">Change the target</Button
              >
            {/if}
          {/snippet}
        </CompTally>
      {/if}
      {#if officer && roles}
        <PlannerBench
          {raid}
          officerTools={CALENDAR_FREE}
          closed={flags.cancelled}
          busyKey={c.rosterBusy}
          onSetState={(row, state) => c.setRosterState(row, state)}
        />
      {/if}
      {#if raid.assignments}
        <PlannerAssignments
          {raid}
          {officer}
          officerTools={c.officerTools}
          closed={flags.cancelled || flags.ended}
          actions={c.assignments}
        />
      {/if}
      <Section title={raidCopy.subheads.notes} panel id="raid-notes" testid="raid-notes">
        {#if raid.notes_html.trim() === ''}
          <EmptyState message={notesEmptyOf(kind, officer)} testid="raid-notes-empty" />
        {:else}
          <div
            class="prose-forever [&_a]:text-gold [&_li]:marker:text-muted text-[14px] leading-[1.6] [&_ol]:list-decimal [&_ul]:list-disc [&>:last-child]:mb-0"
            data-testid="raid-notes-prose"
          >
            <!-- The API renders and sanitises the notes (paragraphs, bold, italic, lists, links). -->
            <!-- eslint-disable-next-line svelte/no-at-html-tags -->
            {@html raid.notes_html}
          </div>
        {/if}
      </Section>
    {/snippet}
  </Page>
</div>
