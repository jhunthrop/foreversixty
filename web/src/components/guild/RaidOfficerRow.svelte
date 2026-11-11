<!-- web/src/components/guild/RaidOfficerRow.svelte -->
<!-- The raid page's officer row (raids spec 3.2 line 6, 4 "Officer buttons", "Cancel confirm",
     "Copy the sign-up list"): Edit, Lock or Unlock, Copy the sign-up list and Cancel raid, all
     outline buttons. Officer-read-only keeps every button but Copy disabled, with the gate sentence
     under the row. Cancel raid swaps the row for RaidCancelConfirm; Keep it returns focus here.
     An ended raid with no log keeps Edit alone (attendance spec 3.1): fixing the time helps the
     match; the band's `Match a log` is the other control. A done raid draws no officer row. -->
<script lang="ts">
  import { tick } from 'svelte';
  import type { GuildRaidView } from '../../lib/guild/api';
  import { TEXTAREA_CHROME } from '../../lib/guild/field-styles';
  import { focusOnMount } from '../../lib/guild/focus-now';
  import { cancelActionOf, kindOf, raidCopy } from '../../lib/guild/raid-copy';
  import { signupListText } from '../../lib/guild/raid-comp';
  import type { OfficerLine } from '../../lib/guild/raid-page-types';
  import { RAID_BUTTON as BUTTON, RAID_GHOST_BUTTON } from '../../lib/guild/raid-buttons';
  import type { RaidStatusAction } from '../../lib/guild/raids-api';
  import type { GateFacts } from '../../lib/guild/plan-gate';
  import { CopyState } from '../../lib/ui/copy-state.svelte';
  import GuildPlanGate from './GuildPlanGate.svelte';
  import RaidCancelConfirm from './RaidCancelConfirm.svelte';

  let {
    raid,
    readOnly,
    busy = null,
    line = null,
    gateId,
    gateFacts,
    guildTimezone,
    onEdit,
    onStatus,
    ended = false,
    writeClipboard = undefined,
  }: {
    raid: GuildRaidView;
    /** Officer-read-only (gate spec 3.2). */
    readOnly: boolean;
    busy?: RaidStatusAction | null;
    line?: OfficerLine | null;
    gateId: string;
    /** The gate sentence's facts (spec 2026-11-05 §3.2). */
    gateFacts: GateFacts;
    /** The copied list reads in guild time. */
    guildTimezone: string;
    onEdit: () => void;
    /** Resolves false on a refusal (the line says why). */
    onStatus: (action: RaidStatusAction, postToDiscord?: boolean) => Promise<boolean>;
    /** Past its end with no log: Edit only. */
    ended?: boolean;
    writeClipboard?: (text: string) => Promise<void>;
  } = $props();

  let confirming = $state(false);
  /** Ticked when the guild's channel posts raids by default (spec 4 "Cancel confirm"). */
  let postCancel = $state(false);
  // svelte-ignore state_referenced_locally
  const clip = new CopyState(writeClipboard);
  $effect(() => () => clip.dispose());
  let cancelButton = $state<HTMLButtonElement | null>(null);

  /** A cancelled or done raid takes no edit, lock or cancel. */
  const closed = $derived(raid.status === 'cancelled' || raid.status === 'done');
  const lockAction = $derived<RaidStatusAction>(raid.status === 'locked' ? 'unlock' : 'lock');
  const lockLabel = $derived.by(() => {
    if (lockAction === 'unlock') return raidCopy.officerRow.unlock;
    return busy === 'lock' ? raidCopy.officerRow.locking : raidCopy.officerRow.lock;
  });
  const writesDisabled = $derived(readOnly || closed || busy !== null);
  const describedBy = $derived(readOnly ? gateId : undefined);
  const statusLine = $derived(
    line !== null && (line.control === 'lock' || line.control === 'unlock' || line.control === 'cancel')
      ? line
      : null,
  );

  function openConfirm(): void {
    postCancel = raid.discord.post_default;
    confirming = true;
  }

  async function keep(): Promise<void> {
    confirming = false;
    await tick();
    cancelButton?.focus();
  }

  async function confirmCancel(): Promise<void> {
    if (await onStatus('cancel', postCancel)) confirming = false;
  }

  function copy(): Promise<void> {
    return clip.copy(signupListText(raid, guildTimezone));
  }

  function selectAll(node: HTMLTextAreaElement): void {
    node.select();
  }
</script>

<div class="flex flex-col gap-2" data-testid="raid-officer-row">
  {#if confirming}
    <RaidCancelConfirm
      kind={kindOf(raid)}
      {postCancel}
      busy={busy === 'cancel'}
      onPostChange={(post) => (postCancel = post)}
      onConfirm={() => void confirmCancel()}
      onKeep={() => void keep()}
    />
  {:else}
    <div class="flex flex-wrap gap-2 md:justify-end">
      <button
        type="button"
        id="raid-officer-edit"
        class={BUTTON}
        disabled={writesDisabled}
        aria-describedby={describedBy}
        onclick={onEdit}
        data-testid="raid-officer-edit">{raidCopy.officerRow.edit}</button
      >
      {#if !ended}
        <button
          type="button"
          class={BUTTON}
          disabled={writesDisabled}
          aria-describedby={describedBy}
          onclick={() => void onStatus(lockAction)}
          data-testid="raid-officer-lock">{lockLabel}</button
        >
        <button type="button" class={BUTTON} onclick={() => void copy()} data-testid="raid-officer-copy"
          >{clip.copied() ? raidCopy.officerRow.copied : raidCopy.officerRow.copy}</button
        >
        <button
          type="button"
          class={RAID_GHOST_BUTTON}
          bind:this={cancelButton}
          disabled={writesDisabled}
          aria-describedby={describedBy}
          onclick={openConfirm}
          data-testid="raid-officer-cancel">{cancelActionOf(kindOf(raid))}</button
        >
      {/if}
    </div>
  {/if}
  {#if readOnly}
    <div class="md:text-right"><GuildPlanGate tab="raids" facts={gateFacts} /></div>
  {/if}
  {#if statusLine !== null}
    {#key statusLine}
      <p
        class="text-wipe text-[13px] md:text-right"
        role="alert"
        tabindex="-1"
        use:focusOnMount
        data-testid="raid-officer-error"
      >
        {statusLine.message}
      </p>
    {/key}
  {/if}
  {#if clip.failed()}
    <p class="text-wipe text-[13px]" role="alert" data-testid="raid-copy-failed">{raidCopy.copyFailed}</p>
    <textarea
      class="{TEXTAREA_CHROME} font-mono text-[12px]"
      rows="6"
      readonly
      aria-label={raidCopy.officerRow.copy}
      use:selectAll
      data-testid="raid-copy-text">{clip.text}</textarea
    >
  {/if}
</div>
