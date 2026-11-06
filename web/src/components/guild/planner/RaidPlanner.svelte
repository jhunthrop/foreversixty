<!-- web/src/components/guild/planner/RaidPlanner.svelte -->
<!-- The raid planner (web layout spec 2026-11-05 §7.2 "Raid planner"): "Who is in, in what comp, and
     what do I swap?", built only from the ui/ primitives. Page split: the head (zone, guild time and
     the viewer's, the count, the state Tag, the officer buttons, the guild tabs), the control bar
     (group by Status, Role or Class), then the viewer's own sign-up, the one Plate of sign-up Rows
     and, for officers, who has not signed up with Ping on Discord. Rail: comp against the target
     (officers edit it), Bench and swaps, buff and debuff cover, the notes. Phone keeps that order
     in one column. A done raid is GuildRaid's report path, never this page. -->
<script lang="ts">
  import type { GuildRaidView } from '../../../lib/guild/api';
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
  import { clockText } from '../../../lib/guild/raid-time';
  import type { RaidStatusAction } from '../../../lib/guild/raids-api';
  import BuffCover from '../overview/BuffCover.svelte';
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
  import PlannerBench from './PlannerBench.svelte';
  import PlannerNotSigned from './PlannerNotSigned.svelte';
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

  /** The URL is the truth: a reload, a share and Back restore the grouping (spec 3.4). */
  function chooseGrouping(param: string, value: string): void {
    if (param !== GROUP_PARAM) return;
    grouping = groupingFrom(`?${GROUP_PARAM}=${value}`, raid);
    const search = groupingSearch(location.search, grouping, raid);
    history.replaceState(history.state, '', `${location.pathname}${search}${location.hash}`);
  }

  const nextPingClock = (at: Date): string => clockText(at, viewerTimezone, locale);
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
  <!-- One block, so the shell's gap sits under the header, not between the back link and it. -->
  <div>
    <div class="px-[var(--page-inset)]">
      <a
        href={`${guildPath}#raids`}
        class="text-muted inline-flex min-h-11 items-center text-[13px] underline"
        data-testid="raid-back-link">{raidCopy.backLink(guildName)}</a
      >
    </div>
    <PageHeader title={raid.title} tone={titleTone} {meta} testid="raid-head">
      {#snippet actions()}
        {#if officer && !c.editing}
          <RaidOfficerRow
            {raid}
            ended={flags.ended}
            readOnly={!c.officerTools}
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
  </div>
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

<Page layout="split" {header} {bar} testid="raid-page">
  {#if flags.cancelled && c.raceLine !== null}
    <p class="text-wipe text-[13px]" role="alert" tabindex="-1" use:focusOnMount data-testid="raid-race-line">
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
  {#if c.pollPaused}<p class="text-muted text-[13px]" data-testid="raid-poll-paused">{raidCopy.poll}</p>{/if}
  <RaidSignupStrip
    {raid}
    locked={flags.locked}
    ended={flags.ended}
    cancelled={flags.cancelled}
    onSignup={c.signup}
    currentKey={c.currentKey}
  />
  <Section title={raidCopy.subheads.signups} meta={count.ratio} testid="raid-signups">
    <PlannerRoster
      {raid}
      {grouping}
      {officer}
      officerTools={c.officerTools}
      {gateId}
      closed={flags.cancelled}
      busyKey={c.rosterBusy}
      line={c.line}
      onSetState={(row, state) => c.setRosterState(row, state)}
    />
  </Section>
  {#if officer}
    <PlannerNotSigned
      {raid}
      {now}
      clock={nextPingClock}
      officerTools={c.officerTools}
      {gateId}
      closed={flags.cancelled || flags.ended}
      busyKey={c.rosterBusy}
      pinging={c.planner.pinging}
      pingAnswer={c.planner.pingAnswer}
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
        officerTools={c.officerTools}
        closed={flags.cancelled}
        busyKey={c.rosterBusy}
        onSetState={(row, state) => c.setRosterState(row, state)}
      />
    {/if}
    {#if raid.cover && raid.cover.length > 0}<BuffCover cover={raid.cover} />{/if}
    <Section title={raidCopy.subheads.notes} id="raid-notes" testid="raid-notes">
      {#if raid.notes_html.trim() === ''}
        <p class="text-muted text-[14px]" data-testid="raid-notes-empty">{notesEmptyOf(kind, officer)}</p>
      {:else}
        <div
          class="prose-forever [&_a]:text-gold [&_li]:marker:text-muted text-[14px] leading-[1.6] [&_ol]:list-decimal [&_ul]:list-disc"
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
