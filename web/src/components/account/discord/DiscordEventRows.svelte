<!-- web/src/components/account/discord/DiscordEventRows.svelte -->
<!-- "Post when": the two event checkboxes (the anonymize row's markup), the indented private-logs
     modifier (disabled, value kept, while the log box is off) and the note under it.
     `onToggle` answers whether the change took; a refusal puts the box back. -->
<script lang="ts">
  import type { DiscordEvents } from '../../../lib/account/api';
  import { notificationsCopy } from '../../../lib/account/notifications-copy';

  let {
    events,
    disabled = false,
    busy = false,
    showNote = true,
    onToggle,
  }: {
    events: DiscordEvents;
    disabled?: boolean;
    busy?: boolean;
    showNote?: boolean;
    onToggle: (next: DiscordEvents) => Promise<boolean> | boolean;
  } = $props();

  const rows = [
    { key: 'sim_finished', label: notificationsCopy.eventSimFinished, testid: 'discord-event-sim' },
    { key: 'upload_parsed', label: notificationsCopy.eventUploadParsed, testid: 'discord-event-upload' },
    {
      key: 'upload_private',
      label: notificationsCopy.eventUploadPrivate,
      testid: 'discord-event-private',
      child: true,
    },
  ] as const;

  /** The private box modifies the log box: it cannot be changed while that one is off. */
  const isDisabled = (key: keyof DiscordEvents): boolean =>
    disabled || (key === 'upload_private' && !events.upload_parsed);

  async function change(event: Event, key: keyof DiscordEvents): Promise<void> {
    const box = event.currentTarget as HTMLInputElement;
    const accepted = await onToggle({ ...events, [key]: box.checked });
    if (!accepted) box.checked = events[key];
  }
</script>

<fieldset class="m-0 flex flex-col border-0 p-0" aria-busy={busy}>
  <legend class="label text-muted mb-1 p-0">{notificationsCopy.eventsLabel}</legend>
  {#each rows as row (row.key)}
    {@const off = isDisabled(row.key)}
    <label
      class="flex min-h-11 items-center gap-3 text-[14px] {'child' in row ? 'ml-7' : ''} {off
        ? 'text-muted'
        : ''}"
    >
      <input
        type="checkbox"
        checked={events[row.key]}
        disabled={off}
        onchange={(event) => change(event, row.key)}
        data-testid={row.testid}
      />
      {row.label}
    </label>
  {/each}
  {#if showNote}<p class="text-muted text-[13px]">{notificationsCopy.eventsNote}</p>{/if}
</fieldset>
