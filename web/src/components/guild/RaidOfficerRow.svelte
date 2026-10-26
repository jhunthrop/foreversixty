<!-- web/src/components/guild/RaidOfficerRow.svelte -->
<!-- The raid page's officer row (raids spec 3.2 line 6, 4 "Officer buttons", "Cancel confirm",
     "Copy the sign-up list"): Edit, Lock or Unlock, Copy the sign-up list and Cancel raid, all
     outline buttons. Officer-read-only keeps every button but Copy disabled, with the gate sentence
     under the row. Cancel raid swaps the row for RaidCancelConfirm; Keep it returns focus here. -->
<script lang="ts">
  import { tick } from 'svelte';
  import type { GuildRaidView } from '../../lib/guild/api';
  import { browserStorage } from '../../lib/guild/planned-reserve';
  import { copyList, raidCopy } from '../../lib/guild/raid-copy';
  import { splitAtLink } from '../../lib/guild/raid-form';
  import { rememberPostCancel, rememberedPostCancel } from '../../lib/guild/raid-page';
  import type { OfficerLine } from '../../lib/guild/raid-page-controller.svelte';
  import type { RaidStatusAction } from '../../lib/guild/raids-api';
  import { SECONDARY_BUTTON_FIXED } from '../../lib/planner/styles';
  import GuildPlanGate from './GuildPlanGate.svelte';
  import RaidCancelConfirm from './RaidCancelConfirm.svelte';

  /** `Copied` stays this long (spec 4). */
  const COPIED_MS = 4000;
  const BUTTON = `${SECONDARY_BUTTON_FIXED} border-line-warm text-text px-4 active:bg-card-top disabled:cursor-not-allowed disabled:opacity-50`;

  let {
    raid,
    readOnly,
    busy = null,
    line = null,
    gateId,
    settingsHref,
    guildTimezone,
    onEdit,
    onStatus,
    writeClipboard = (text: string) => navigator.clipboard.writeText(text),
    storage = browserStorage(),
  }: {
    raid: GuildRaidView;
    /** Officer-read-only (gate spec 3.2). */
    readOnly: boolean;
    busy?: RaidStatusAction | null;
    line?: OfficerLine | null;
    gateId: string;
    settingsHref: string;
    /** The copied list reads in guild time. */
    guildTimezone: string;
    onEdit: () => void;
    /** Resolves false on a refusal (the line says why). */
    onStatus: (action: RaidStatusAction, postToDiscord?: boolean) => Promise<boolean>;
    writeClipboard?: (text: string) => Promise<void>;
    storage?: Pick<Storage, 'getItem' | 'setItem'> | null;
  } = $props();

  let confirming = $state(false);
  let postCancel = $state(rememberedPostCancel(storage));
  let copied = $state(false);
  let copyFailed = $state<string | null>(null);
  let copiedTimer: ReturnType<typeof setTimeout> | undefined;
  let cancelButton = $state<HTMLButtonElement | null>(null);

  const gateLead = splitAtLink(raidCopy.gate, raidCopy.gateLink).before;
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

  function setPostCancel(post: boolean): void {
    postCancel = post;
    rememberPostCancel(storage, post);
  }

  async function copy(): Promise<void> {
    const text = copyList(raid, guildTimezone);
    try {
      await writeClipboard(text);
      copyFailed = null;
      copied = true;
      clearTimeout(copiedTimer);
      copiedTimer = setTimeout(() => (copied = false), COPIED_MS);
    } catch {
      copyFailed = text;
    }
  }

  function focusNow(node: HTMLElement): void {
    node.focus();
  }

  function selectAll(node: HTMLTextAreaElement): void {
    node.select();
  }
</script>

<div class="flex flex-col gap-2" data-testid="raid-officer-row">
  {#if confirming}
    <RaidCancelConfirm
      {postCancel}
      busy={busy === 'cancel'}
      buttonClass={BUTTON}
      onPostChange={setPostCancel}
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
      <button
        type="button"
        class={BUTTON}
        disabled={writesDisabled}
        aria-describedby={describedBy}
        onclick={() => void onStatus(lockAction)}
        data-testid="raid-officer-lock">{lockLabel}</button
      >
      <button type="button" class={BUTTON} onclick={() => void copy()} data-testid="raid-officer-copy"
        >{copied ? raidCopy.officerRow.copied : raidCopy.officerRow.copy}</button
      >
      <button
        type="button"
        class={BUTTON}
        bind:this={cancelButton}
        disabled={writesDisabled}
        aria-describedby={describedBy}
        onclick={openConfirm}
        data-testid="raid-officer-cancel">{raidCopy.officerRow.cancel}</button
      >
    </div>
  {/if}
  {#if readOnly}
    <div class="md:text-right"><GuildPlanGate tab="raids" {settingsHref} sentence={gateLead} /></div>
  {/if}
  {#if statusLine !== null}
    {#key statusLine}
      <p
        class="text-wipe text-[13px] md:text-right"
        role="alert"
        tabindex="-1"
        use:focusNow
        data-testid="raid-officer-error"
      >
        {statusLine.message}
      </p>
    {/key}
  {/if}
  {#if copyFailed !== null}
    <p class="text-wipe text-[13px]" role="alert" data-testid="raid-copy-failed">{raidCopy.copyFailed}</p>
    <textarea
      class="border-line-warm bg-raised text-text rounded-control w-full border p-2 font-mono text-[12px]"
      rows="6"
      readonly
      aria-label={raidCopy.officerRow.copy}
      use:selectAll
      data-testid="raid-copy-text">{copyFailed}</textarea
    >
  {/if}
</div>
