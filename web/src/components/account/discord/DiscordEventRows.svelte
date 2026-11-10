<!-- web/src/components/account/discord/DiscordEventRows.svelte -->
<!-- The event switches (premium spec 2026-11-08 section 7.6): a slim row with its Switch for each
     event, and the nested private-logs switch (disabled, value kept, while the log switch is off)
     with the note as its sub. A press shows its new position at once, at 60 %, until `onToggle`
     answers whether the change took; a refusal puts the switch back. -->
<script lang="ts">
  import type { DiscordEvents } from '../../../lib/account/api';
  import { notificationsCopy } from '../../../lib/account/notifications-copy';
  import Row from '../../ui/Row.svelte';
  import Switch from '../../ui/Switch.svelte';

  const uid = $props.id();

  let {
    events,
    disabled = false,
    inactive = false,
    busy = false,
    showNote = true,
    onToggle,
  }: {
    events: DiscordEvents;
    disabled?: boolean;
    /** The block is not live (switched off, lapsed): the switches are disabled (45 %, off the tab order). */
    inactive?: boolean;
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

  /** The switch just pressed, with the position it asked for, until the save answers. */
  let pending = $state<{ key: keyof DiscordEvents; value: boolean } | null>(null);

  const shown = (key: keyof DiscordEvents): boolean => (pending?.key === key ? pending.value : events[key]);

  /** The private switch modifies the log switch: it cannot be changed while that one is off. */
  const isDisabled = (key: keyof DiscordEvents): boolean =>
    inactive || (disabled && pending?.key !== key) || (key === 'upload_private' && !events.upload_parsed);

  async function change(key: keyof DiscordEvents, value: boolean): Promise<void> {
    pending = { key, value };
    try {
      await onToggle({ ...events, [key]: value });
    } finally {
      pending = null;
    }
  }
</script>

<div aria-busy={busy ? 'true' : undefined}>
  {#each rows as row (row.key)}
    <div
      class:pl-6={'child' in row}
      data-nested={'child' in row ? 'true' : undefined}
      data-testid="{row.testid}-row"
    >
      <Row
        title={row.label}
        sub={'child' in row && showNote ? notificationsCopy.eventsNote : undefined}
        density="slim"
        open
      >
        {#snippet titleContent()}<span id="{uid}-{row.key}" class="text-text font-normal">{row.label}</span
          >{/snippet}
        {#snippet trail()}
          <!-- mr-2 takes back the Switch's -mr-2 hit area: the track ends on the panel's text edge. -->
          <span class="mr-2 inline-flex"
            ><Switch
              checked={shown(row.key)}
              labelledby="{uid}-{row.key}"
              disabled={isDisabled(row.key)}
              saving={pending?.key === row.key}
              onchange={(next) => change(row.key, next)}
              testid={row.testid}
            /></span
          >
        {/snippet}
      </Row>
    </div>
  {/each}
</div>
