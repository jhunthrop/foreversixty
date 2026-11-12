<!-- web/src/components/guild/report/ReportAttendance.svelte -->
<!-- Attendance in the report's rail (web layout spec 2026-11-05 §7.2 "Raid report"): the night's
     summary sentence for everyone; for officers who was not in the log, late or joined by name, then every raider's row with the night every count
     reads, and under a fix who set it and what the log said, muted. The fix is a small select
     (Present, Absent, Bench, Excused; bench counts as attended) and `Back to the log` clears it.
     While a fix saves the select shows it; a refusal keeps the row and says so under it. A member
     reads the sentence only: their own night is in Your night. -->
<script lang="ts">
  import type { Snippet } from 'svelte';
  import type {
    AttendanceOverrideStatus,
    RaidAttendanceRow,
    RaidNight,
  } from '../../../lib/guild/attendance-contract';
  import { attendanceSummary } from '../../../lib/guild/attendance-copy';
  import type { AttendanceFix } from '../../../lib/guild/night-actions.svelte';
  import { fixOptions, fixRows, FIX_STATUSES, nightNames } from '../../../lib/guild/report-attendance-view';
  import { reportNightCopy as copy } from '../../../lib/guild/report-night-copy';
  import { classTone, LINK, titleColor } from '../../../lib/ui/primitives';
  import Crest from '../../ui/Crest.svelte';
  import Disclosure from '../../ui/Disclosure.svelte';
  import Row from '../../ui/Row.svelte';
  import Section from '../../ui/Section.svelte';
  import Select from '../../ui/Select.svelte';

  let {
    log,
    night,
    rows,
    officer,
    readOnly,
    gateId,
    fixing = {},
    lines = {},
    onFix,
  }: {
    /** The log the night was read from (the match line), above the summary. */
    log?: Snippet;
    night: RaidNight;
    rows: readonly RaidAttendanceRow[];
    officer: boolean;
    /** Officer-read-only: the selects stay, disabled, and name the gate sentence. */
    readOnly: boolean;
    gateId: string;
    /** The fix each row is saving. */
    fixing?: Readonly<Record<string, AttendanceFix>>;
    /** A refused fix's line, by character. */
    lines?: Readonly<Record<string, string>>;
    onFix: (characterKey: string, fix: AttendanceFix) => void;
  } = $props();

  const models = $derived(officer ? fixRows(rows) : []);
  /** The fix rows open on their own once a fix exists or a refusal needs reading; else the
   *  officer opens them. */
  const needsOpen = $derived(
    models.some((model) => model.auto !== null) ||
      Object.keys(lines).length > 0 ||
      Object.keys(fixing).length > 0,
  );
  let openedByHand = $state(false);
  const options = fixOptions();
  const isStatus = (value: string): value is AttendanceOverrideStatus =>
    (FIX_STATUSES as readonly string[]).includes(value);
</script>

<Section title={copy.attendance.head} anchor="attendance" panel testid="raid-attendance">
  {#if log}<div class="mb-2">{@render log()}</div>{/if}
  <p class="text-text text-[14px]" data-testid="raid-attendance-summary">{attendanceSummary(night)}</p>
  {#if officer}
    {#each nightNames(rows) as group (group.key)}
      <p class="text-muted mt-1 text-[13px] leading-[18px]" data-testid="raid-attendance-names">
        {group.label}:
        {#each group.rows as row, index (row.character_key)}<span
            class="font-semibold"
            style:color={titleColor(classTone(row.class))}>{row.name}</span
          >{index < group.rows.length - 1 ? ', ' : ''}{/each}
      </p>
    {/each}
  {/if}
  {#if models.length > 0}
    <Disclosure
      summary={copy.attendance.fix}
      bind:open={() => openedByHand || needsOpen, (open) => (openedByHand = open)}
      testid="raid-attendance-fix-open"
    >
      {#each models as model (model.row.character_key)}
        {@const key = model.row.character_key}
        {@const saving = key in fixing}
        {@const shown = saving ? (fixing[key] ?? '') : model.value}
        <div aria-busy={saving} data-testid="raid-attendance-row" data-fixed={model.auto !== null}>
          <Row
            title={model.row.name}
            titleTone={classTone(model.row.class)}
            selected={model.row.own}
            open
            density="slim"
          >
            {#snippet lead()}<Crest cls={model.row.class} size={24} />{/snippet}
            {#snippet subContent()}
              <span class="text-text block" data-testid="raid-attendance-outcome">{model.effective}</span>
              {#if model.auto !== null}
                <span class="block" data-testid="raid-attendance-auto">{model.auto}</span>
                <button
                  type="button"
                  class="{LINK} relative min-h-11 text-[13px] md:min-h-0"
                  disabled={readOnly || saving}
                  aria-describedby={readOnly ? gateId : undefined}
                  onclick={() => onFix(key, null)}
                  data-testid="raid-attendance-back">{copy.attendance.backToLog}</button
                >
              {/if}
              {#if lines[key] !== undefined}
                <span class="text-wipe block" role="alert" data-testid="raid-attendance-line"
                  >{lines[key]}</span
                >
              {/if}
            {/snippet}
            {#snippet trail()}
              <Select
                id={`raid-attendance-fix-${key}`}
                class="w-[120px]"
                size="dense"
                align="end"
                label={copy.attendance.fixLabel(model.row.name)}
                placeholder={copy.attendance.fixPlaceholder}
                value={shown}
                {options}
                disabled={readOnly}
                busy={saving}
                describedBy={readOnly ? gateId : undefined}
                testid={`raid-attendance-fix-${key}`}
                onchange={(event) => {
                  const value = (event.currentTarget as HTMLSelectElement).value;
                  if (isStatus(value) && value !== model.value) onFix(key, value);
                }}
              />
            {/snippet}
          </Row>
        </div>
      {/each}
    </Disclosure>
  {/if}
</Section>
