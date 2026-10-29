<!-- web/src/components/guild/RaidForm.svelte -->
<!-- The raid form (raids spec 3.5): create on the Raids tab, edit on the raid page. Fields in the
     guild's wall clock; the For you line renders it in the viewer's zone. Every rule lives in
     lib/guild/raid-form.ts; this file draws the fields and the 6.1 lines under them. The Raid
     select takes focus on open; Escape and Cancel close. -->
<script lang="ts">
  import { viewerLocale } from '../../lib/guild/raid-time';
  import { onMount } from 'svelte';
  import type { GuildRaidDiscord, GuildRaidView, GuildZoneView } from '../../lib/guild/api';
  import {
    CHECK_ROW,
    DISABLED_LOOK,
    FIELD_CHROME,
    INLINE_LINK,
    SELECT_CHROME,
    SWITCH_CHROME,
    SWITCH_ON,
    TEXTAREA_CHROME,
  } from '../../lib/guild/field-styles';
  import { raidCopy } from '../../lib/guild/raid-copy';
  import {
    NOTES_MAX,
    TITLE_MAX,
    canRepeat,
    createDefaults,
    discordFormPart,
    discordHelper,
    discordSettingsHref,
    durationOptions,
    editDefaults,
    forYou,
    notesCounter,
    raidPayload,
    repeatHelper,
    splitAtLink,
    startOf,
    validateRaidForm,
    type RaidFormErrors,
    type RaidFormField,
    type RaidFormMode,
    type RaidFormValues,
    type RaidSubmitResult,
  } from '../../lib/guild/raid-form';
  import { timeZoneOptions } from '../../lib/guild/raid-zones';
  import type { RaidInput } from '../../lib/guild/raids-api';
  import { PRIMARY_BUTTON_FIXED } from '../../lib/planner/styles';
  import RaidFormDiscord from './RaidFormDiscord.svelte';
  import SelectChevron from '../ui/SelectChevron.svelte';

  /** The block is only drawn with a channel; this stands in while there is none. */
  const CLOSED_DISCORD: GuildRaidDiscord = {
    available: false,
    connected: false,
    post_default: false,
    remind_default: true,
    channel_name: null,
  };

  interface Props {
    mode: RaidFormMode;
    zones: readonly GuildZoneView[];
    /** The guild's `raid_timezone`; `null` before one is set. */
    guildTimezone: string | null;
    viewerTimezone: string;
    lastRaid?: GuildRaidView | null;
    /** The raid being edited (edit mode). */
    raid?: GuildRaidView;
    /** Planned and locked raids now held, for the cap. */
    plannedCount: number;
    /** The guild channel from the raid view; `null` when no raid has carried it yet. */
    discord: GuildRaidDiscord | null;
    settingsHref: string;
    onSubmit: (payload: RaidInput) => Promise<RaidSubmitResult>;
    onCancel: () => void;
    /** The clock the defaults and checks read; a test pins it. */
    now?: Date;
    locale?: string;
  }

  let {
    mode,
    zones,
    guildTimezone,
    viewerTimezone,
    lastRaid = null,
    raid,
    plannedCount,
    discord,
    settingsHref,
    onSubmit,
    onCancel,
    now,
    locale = viewerLocale(),
  }: Props = $props();

  function clock(): Date {
    return now ?? new Date();
  }

  function initialValues(): RaidFormValues {
    if (mode === 'edit' && raid !== undefined) return editDefaults(raid, guildTimezone ?? viewerTimezone);
    return createDefaults({ last: lastRaid, now: clock(), guildTimezone, viewerTimezone, discord });
  }

  let values = $state<RaidFormValues>(initialValues());
  let errors = $state<RaidFormErrors>({});
  let busy = $state(false);
  /** The API said the guild's officer tools ended: the form stays to show why, but cannot resend. */
  let planEnded = $state(false);
  let zoneSelect = $state<HTMLSelectElement | null>(null);

  const id = (field: string): string => `raid-form-${mode}-${field}`;
  const errorId = (field: RaidFormField | 'form'): string => `${id(field)}-error`;
  const chosenZone = $derived(zones.find((zone) => zone.key === values.zone) ?? null);
  const zoneGroups = $derived(timeZoneOptions(clock(), [values.timezone, viewerTimezone]));
  const forYouText = $derived(forYou(values, viewerTimezone, clock(), locale));
  const notesLeft = $derived(notesCounter(values.notes.length));
  const repeatLine = $derived(mode === 'create' ? repeatHelper(plannedCount) : null);
  const repeatAllowed = $derived(canRepeat(plannedCount));
  const submitLabel = $derived.by(() => {
    if (mode === 'create') return busy ? raidCopy.form.submitCreating : raidCopy.form.submitCreate;
    return busy ? raidCopy.form.submitEditing : raidCopy.form.submitEdit;
  });
  const discordPart = $derived(discordFormPart(discord));
  const discordLines = $derived(
    discordHelper({
      values,
      discord: discord ?? CLOSED_DISCORD,
      post: raid?.discord_post ?? null,
      mode,
      startsAt: startOf(values),
      now: clock(),
      timezone: values.timezone,
    }),
  );
  const discordLine = splitAtLink(raidCopy.form.connectDiscord, raidCopy.form.connectDiscordLink);
  const keepStartsAt = $derived(mode === 'edit' ? raid?.starts_at : undefined);

  onMount(() => zoneSelect?.focus());

  function set<K extends keyof RaidFormValues>(key: K, value: RaidFormValues[K]): void {
    values = { ...values, [key]: value };
    const cleared = key === 'zone' || key === 'date' || key === 'time' ? { [key]: undefined } : {};
    errors = { ...errors, ...cleared, form: planEnded ? errors.form : undefined };
  }

  function focusFirstError(found: RaidFormErrors): void {
    const first = (['zone', 'date', 'time'] as const).find((field) => found[field] !== undefined);
    if (first !== undefined) document.getElementById(id(first))?.focus();
  }

  async function submit(event: SubmitEvent): Promise<void> {
    event.preventDefault();
    if (busy || planEnded) return;
    const found = validateRaidForm(values, { now: clock(), plannedCount, mode, keepStartsAt });
    errors = found;
    if (Object.keys(found).length > 0) {
      focusFirstError(found);
      return;
    }
    busy = true;
    try {
      const result = await onSubmit(raidPayload(values, mode, discordPart === 'box'));
      if (!result.ok) {
        errors = { [result.field ?? 'form']: result.message };
        planEnded = result.code === 'guild_plan_required';
      }
    } catch {
      errors = { form: raidCopy.form.errors.failed };
    } finally {
      busy = false;
    }
  }

  function onKeydown(event: KeyboardEvent): void {
    if (event.key !== 'Escape') return;
    event.preventDefault();
    onCancel();
  }
</script>

{#snippet fieldError(field: RaidFormField)}
  {#if errors[field] !== undefined}
    <p id={errorId(field)} class="text-wipe text-[13px]" data-testid={`raid-form-error-${field}`}>
      {errors[field]}
    </p>
  {/if}
{/snippet}

{#snippet label(field: string, text: string)}
  <label class="quiet-label" for={id(field)}>{text}</label>
{/snippet}

<form
  class="bg-raised border-line rounded-panel flex flex-col gap-4 border p-4 shadow-[inset_0_1px_0_rgba(229,185,85,.35)] md:p-5"
  novalidate
  aria-busy={busy}
  onsubmit={(event) => void submit(event)}
  onkeydown={onKeydown}
  data-testid={`raid-form-${mode}`}
>
  <div class="grid grid-cols-1 gap-4 lg:grid-cols-2">
    <div class="flex flex-col gap-2">
      {@render label('zone', raidCopy.form.raid)}
      <span class="relative block">
        <select
          id={id('zone')}
          bind:this={zoneSelect}
          class={SELECT_CHROME}
          value={values.zone}
          disabled={busy}
          aria-invalid={errors.zone !== undefined}
          aria-describedby={errors.zone !== undefined ? errorId('zone') : undefined}
          onchange={(event) => set('zone', event.currentTarget.value)}
          data-testid="raid-form-zone"
        >
          {#if values.zone === ''}<option value="" disabled>{raidCopy.placeholders.chooseRaid}</option>{/if}
          {#each zones as zone (zone.key)}<option value={zone.key}>{zone.name}</option>{/each}
        </select>
        <SelectChevron />
      </span>
      {@render fieldError('zone')}
    </div>
    <div class="flex flex-col gap-2">
      {@render label('title', raidCopy.form.title)}
      <input
        id={id('title')}
        type="text"
        class={FIELD_CHROME}
        maxlength={TITLE_MAX}
        placeholder={chosenZone?.name ?? ''}
        value={values.title}
        disabled={busy}
        oninput={(event) => set('title', event.currentTarget.value)}
        data-testid="raid-form-title"
      />
    </div>
  </div>

  <div class="grid grid-cols-1 gap-4 lg:grid-cols-3">
    <div class="flex flex-col gap-2">
      {@render label('date', raidCopy.form.date)}
      <input
        id={id('date')}
        type="date"
        class={FIELD_CHROME}
        value={values.date}
        disabled={busy}
        aria-invalid={errors.date !== undefined}
        aria-describedby={errors.date !== undefined ? errorId('date') : undefined}
        oninput={(event) => set('date', event.currentTarget.value)}
        data-testid="raid-form-date"
      />
      {@render fieldError('date')}
    </div>
    <div class="flex flex-col gap-2">
      {@render label('time', raidCopy.form.time)}
      <input
        id={id('time')}
        type="time"
        step="900"
        class={FIELD_CHROME}
        value={values.time}
        disabled={busy}
        aria-invalid={errors.time !== undefined}
        aria-describedby={errors.time !== undefined ? errorId('time') : undefined}
        oninput={(event) => set('time', event.currentTarget.value)}
        data-testid="raid-form-time"
      />
      {@render fieldError('time')}
    </div>
    <div class="flex flex-col gap-2">
      {@render label('duration', raidCopy.form.duration)}
      <span class="relative block">
        <select
          id={id('duration')}
          class={SELECT_CHROME}
          value={String(values.duration)}
          disabled={busy}
          onchange={(event) => set('duration', Number(event.currentTarget.value))}
          data-testid="raid-form-duration"
        >
          {#each durationOptions() as option (option.value)}<option value={option.value}
              >{option.label}</option
            >{/each}
        </select>
        <SelectChevron />
      </span>
    </div>
  </div>

  <div class="flex flex-col gap-2">
    {@render label('timezone', raidCopy.form.guildZone)}
    <span class="relative block">
      <select
        id={id('timezone')}
        class={SELECT_CHROME}
        value={values.timezone}
        disabled={busy}
        aria-describedby={id('for-you')}
        onchange={(event) => set('timezone', event.currentTarget.value)}
        data-testid="raid-form-timezone"
      >
        {#each zoneGroups as group (group.region)}
          <optgroup label={group.region}>
            {#each group.options as option (option.value)}<option value={option.value}>{option.label}</option
              >{/each}
          </optgroup>
        {/each}
      </select>
      <SelectChevron />
    </span>
    <p id={id('for-you')} class="text-muted text-[13px]" data-testid="raid-form-for-you">{forYouText}</p>
  </div>

  <div class="flex flex-col gap-2">
    {@render label('notes', raidCopy.form.notes)}
    <div class="relative">
      <textarea
        id={id('notes')}
        rows="6"
        class={TEXTAREA_CHROME}
        maxlength={NOTES_MAX}
        value={values.notes}
        disabled={busy}
        aria-describedby={id('notes-helper')}
        oninput={(event) => set('notes', event.currentTarget.value)}
        data-testid="raid-form-notes"></textarea>
      {#if notesLeft !== null}
        <span
          class="text-muted numeric pointer-events-none absolute right-3 bottom-2 text-[12px]"
          data-testid="raid-form-notes-counter">{notesLeft}</span
        >
      {/if}
    </div>
    <p id={id('notes-helper')} class="text-muted text-[12px]">{raidCopy.form.markdownHelper}</p>
  </div>

  {#if mode === 'create'}
    <div class="flex flex-col gap-1">
      <label class="{CHECK_ROW} {repeatAllowed ? '' : DISABLED_LOOK}" data-testid="raid-form-repeat-label">
        <input
          type="checkbox"
          class="{SWITCH_CHROME} {SWITCH_ON}"
          role="switch"
          aria-checked={values.repeat && repeatAllowed}
          checked={values.repeat && repeatAllowed}
          disabled={busy || !repeatAllowed}
          aria-describedby={repeatLine === null ? undefined : id('repeat-helper')}
          onchange={(event) => set('repeat', event.currentTarget.checked)}
          data-testid="raid-form-repeat"
        />
        {raidCopy.form.repeat}
      </label>
      {#if repeatLine !== null}
        <p id={id('repeat-helper')} class="text-muted text-[13px]" data-testid="raid-form-repeat-helper">
          {repeatLine}
        </p>
      {/if}
    </div>
  {/if}

  {#if discordPart === 'box'}
    <RaidFormDiscord {values} helper={discordLines} {busy} {set} />
  {:else if discordPart === 'line'}
    <p class="text-muted text-[13px]" data-testid="raid-form-discord-line">
      {discordLine.before}<a
        class={INLINE_LINK}
        href={discordSettingsHref(settingsHref)}
        data-testid="raid-form-settings-link">{raidCopy.form.connectDiscordLink}</a
      >{discordLine.after}
    </p>
  {/if}

  <div class="flex flex-col gap-1">
    <div class="flex flex-wrap items-center gap-4">
      <button
        type="submit"
        class="{PRIMARY_BUTTON_FIXED} focus-visible:outline-gold px-4 focus-visible:outline-2 focus-visible:outline-offset-2 active:bg-[#cfa64a]"
        disabled={busy || planEnded}
        aria-describedby={errors.form !== undefined ? errorId('form') : undefined}
        data-testid="raid-form-submit">{submitLabel}</button
      >
      <button
        type="button"
        class="text-text focus-visible:outline-gold h-11 text-[13px] underline focus-visible:outline-2"
        disabled={busy}
        onclick={onCancel}
        data-testid="raid-form-cancel">{raidCopy.form.cancel}</button
      >
    </div>
    <div class="min-h-5 text-[13px] leading-5" data-testid="raid-form-status">
      {#if errors.form !== undefined}<p id={errorId('form')} class="text-wipe" role="alert">
          {errors.form}
        </p>{/if}
    </div>
  </div>
</form>
