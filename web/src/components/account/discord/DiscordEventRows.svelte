<!-- web/src/components/account/discord/DiscordEventRows.svelte -->
<!-- "Post when": the two event switches and the nested private-logs switch (disabled, value kept,
     while the log switch is off), each a SettingsRow with a Switch. The private row carries the
     note as its sub. A press shows its new position at once, at 60 %, until `onToggle` answers
     whether the change took; a refusal puts the switch back. -->
<script lang="ts">
  import type { DiscordEvents } from '../../../lib/account/api';
  import { notificationsCopy } from '../../../lib/account/notifications-copy';
  import SettingsRow from '../../ui/SettingsRow.svelte';
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
    /** The block is not live (switched off, lapsed): greyed rows, switches off the tab order. */
    inactive?: boolean;
    busy?: boolean;
    showNote?: boolean;
    onToggle: (next: DiscordEvents) => Promise<boolean> | boolean;
  } = $props();

  const rows = [
    {
      key: 'sim_finished',
      label: notificationsCopy.eventSimFinished,
      testid: 'discord-event-sim',
      glyph: 'server',
    },
    {
      key: 'upload_parsed',
      label: notificationsCopy.eventUploadParsed,
      testid: 'discord-event-upload',
      glyph: 'chart',
    },
    {
      key: 'upload_private',
      label: notificationsCopy.eventUploadPrivate,
      testid: 'discord-event-private',
      glyph: 'lock',
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

<div class="flex flex-col" role="group" aria-labelledby="{uid}-legend" aria-busy={busy}>
  <p
    id="{uid}-legend"
    class="text-muted border-line-soft border-t px-5 pt-4 pb-2 text-[11px] leading-none font-bold tracking-[.14em] uppercase"
  >
    {notificationsCopy.eventsLabel}
  </p>
  {#each rows as row (row.key)}
    <SettingsRow
      label={row.label}
      glyph={row.glyph}
      nested={'child' in row}
      sub={'child' in row && showNote ? notificationsCopy.eventsNote : undefined}
      disabled={inactive || ('child' in row && !events.upload_parsed)}
      testid="{row.testid}-row"
    >
      {#snippet control(labelId)}
        <Switch
          checked={shown(row.key)}
          labelledby={labelId}
          disabled={isDisabled(row.key)}
          saving={pending?.key === row.key}
          onchange={(next) => change(row.key, next)}
          testid={row.testid}
        />
      {/snippet}
    </SettingsRow>
  {/each}
</div>
