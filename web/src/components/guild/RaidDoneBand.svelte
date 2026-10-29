<!-- web/src/components/guild/RaidDoneBand.svelte -->
<!-- The band's lines on a raid with a log, or one past its end with none (attendance spec 3.1):
     the night line (replacing the count line), the match line, or `Ended · no log yet` with its
     hint and, for officers, `Match a log`, which opens the picker in place under the row. -->
<script lang="ts">
  import { tick } from 'svelte';
  import type { GuildRaidView } from '../../lib/guild/api';
  import { attendanceCopy, cameRatio, endedHint, nightFacts } from '../../lib/guild/attendance-copy';
  import { PHONE_HIT_AREA } from '../../lib/guild/field-styles';
  import { focusOnMount } from '../../lib/guild/focus-now';
  import { MatchPicker, PLAN_REQUIRED, type MatchOutcome } from '../../lib/guild/match-picker.svelte';
  import { raidEndsAt } from '../../lib/guild/planned-row';
  import { RAID_BUTTON } from '../../lib/guild/raid-buttons';
  import { SEP } from '../../lib/guild/raid-copy';
  import { RAID_ENDED_LINE_ID } from '../../lib/guild/raid-page';
  import type { MatchCandidates } from '../../lib/guild/attendance-contract';
  import Figures from './Figures.svelte';
  import RaidMatchLine from './RaidMatchLine.svelte';
  import RaidMatchPicker from './RaidMatchPicker.svelte';

  const UPLOAD_HREF = '/upload';
  const MATCH_A_LOG_ID = 'raid-match-a-log';
  const OPEN_NIGHT_ID = 'raid-open-night';

  let {
    raid,
    now,
    guildPath,
    viewerTimezone,
    guildTimezone,
    locale,
    officer,
    readOnly,
    gateId,
    loadCandidates,
    onMatch,
    onClear,
    picker = $bindable(null),
  }: {
    raid: GuildRaidView;
    now: Date;
    guildPath: string;
    viewerTimezone: string;
    guildTimezone: string;
    locale: string;
    officer: boolean;
    readOnly: boolean;
    gateId: string;
    loadCandidates: () => Promise<MatchCandidates>;
    onMatch: (reportId: string) => Promise<MatchOutcome>;
    onClear: () => Promise<MatchOutcome>;
    /** The open picker, or null; a test seeds it. */
    picker?: MatchPicker | null;
  } = $props();

  const night = $derived(raid.night);
  const endsAt = $derived(raidEndsAt(raid));
  const open = $derived(raid.status === 'planned' || raid.status === 'locked');
  const showEnded = $derived(raid.match === null && open && now.getTime() >= endsAt.getTime());
  const hint = $derived(endedHint(endsAt, now));
  /** Bar fill: those who came of those who signed In or Late; no bar before anyone signed. */
  const fill = $derived(night === null || night.signed === 0 ? null : (night.came / night.signed) * 100);
  const picking = $derived(picker !== null);
  /** `Not saved. This guild's officer tools have ended.`, where the picker was. */
  let refusal = $state<string | null>(null);
  let pickerHost = $state<HTMLElement | undefined>();

  // The picker is for officers with tools and a raid still without a log: it closes when either
  // stops being true (the plan ended, or the poll matched the log from elsewhere). A poll match
  // under the officer's focus hands it to the match line's link, never to the page body.
  $effect(() => {
    if (picker === null) return;
    const matchedElsewhere = raid.match !== null;
    if (!readOnly && !matchedElsewhere) return;
    const hadFocus = pickerHost?.contains(document.activeElement) ?? false;
    picker = null;
    if (matchedElsewhere && hadFocus) void focusAfterTick(OPEN_NIGHT_ID);
  });

  async function focusAfterTick(id: string): Promise<void> {
    await tick();
    document.getElementById(id)?.focus();
  }

  function openPicker(): void {
    refusal = null;
    const next = new MatchPicker({ load: loadCandidates, match });
    picker = next;
    void next.start();
  }

  /** A plan refusal closes the picker; the gate sentence is the officer row's, the line is ours. */
  async function match(reportId: string): Promise<MatchOutcome> {
    const outcome = await onMatch(reportId);
    if (!outcome.ok && outcome.code === PLAN_REQUIRED) refusal = outcome.message;
    return outcome;
  }

  async function closePicker(): Promise<void> {
    picker = null;
    await focusAfterTick(MATCH_A_LOG_ID);
  }

  /** A matched log swaps the ended line for the match line; focus goes to its link. */
  async function matched(): Promise<void> {
    picker = null;
    await focusAfterTick(OPEN_NIGHT_ID);
  }

  async function cleared(): Promise<MatchOutcome> {
    const outcome = await onClear();
    if (outcome.ok) await focusAfterTick(MATCH_A_LOG_ID);
    return outcome;
  }
</script>

{#if night !== null}
  <!-- One flow: the bar sits on the count's line and the clauses wrap whole, so the phone line is
       two lines at most (spec 7). -->
  <p class="leading-[22px]" data-testid="raid-night-line">
    {#if fill !== null}
      <span
        class="bg-line-soft relative mr-3 inline-block h-2 w-[120px] overflow-hidden rounded-full align-middle md:w-[160px]"
        aria-hidden="true"
        ><span class="bg-muted absolute inset-y-0 left-0" style:width={`${fill}%`}></span></span
      >
    {/if}<span class="text-strong numeric text-[15px] whitespace-nowrap">{cameRatio(night)}</span><span
      class="text-muted text-[13px]"
      >{#each nightFacts(night) as fact (fact)}{SEP}<Figures text={fact} />{/each}</span
    >
  </p>
{/if}

{#if raid.match !== null}
  <RaidMatchLine
    match={raid.match}
    {guildPath}
    {viewerTimezone}
    {guildTimezone}
    {locale}
    {officer}
    {readOnly}
    {gateId}
    onClear={cleared}
  />
{:else if showEnded}
  <div class="flex flex-col gap-2 md:flex-row md:items-center md:gap-4" data-testid="raid-ended-line">
    <p id={RAID_ENDED_LINE_ID} class="text-text text-[14px]">
      {attendanceCopy.ended.noLog}<span class="text-muted"
        >{#if hint.kind === 'soon'}{hint.text}{:else}{hint.lead}<a
            href={UPLOAD_HREF}
            class="text-text underline {PHONE_HIT_AREA}">{hint.link}</a
          >{hint.tail}{/if}</span
      >
    </p>
    {#if officer && !picking}
      <button
        type="button"
        id={MATCH_A_LOG_ID}
        class="{RAID_BUTTON} w-full justify-center md:w-auto"
        disabled={readOnly}
        aria-describedby={readOnly ? gateId : undefined}
        onclick={openPicker}
        data-testid="raid-match-open">{attendanceCopy.ended.matchALog}</button
      >
    {/if}
  </div>
  {#if refusal !== null}
    <p
      class="text-wipe text-[13px]"
      role="alert"
      tabindex="-1"
      use:focusOnMount
      data-testid="raid-match-refused"
    >
      {refusal}
    </p>
  {/if}
{/if}

{#if officer && picker !== null}
  <div class="contents" bind:this={pickerHost}>
    <RaidMatchPicker
      {picker}
      zoneName={raid.zone_name}
      {guildPath}
      {viewerTimezone}
      {guildTimezone}
      {locale}
      onMatched={() => void matched()}
      onCancel={() => void closePicker()}
    />
  </div>
{/if}
