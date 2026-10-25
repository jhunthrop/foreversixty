<!-- web/src/components/account/discord/DiscordPasteForm.svelte -->
<!-- Paste (or replace) the webhook: the labelled h-11 input, the reserved error line, the two
     helpers, the events and the one filled button. The typed URL lives only in the parent's
     state; it is never put in an attribute other than the input's own value. -->
<script lang="ts">
  import type { DiscordEvents } from '../../../lib/account/api';
  import { notificationsCopy } from '../../../lib/account/notifications-copy';
  import { LINK_BUTTON } from '../../../lib/account/discord-styles';
  import { PRIMARY_BUTTON_FIXED } from '../../../lib/planner/styles';
  import DiscordEventRows from './DiscordEventRows.svelte';

  let {
    webhook = $bindable(),
    events = $bindable(),
    error = null,
    offerPlans = false,
    saving = false,
    replacing = false,
    onSubmit,
    onCancel,
  }: {
    webhook: string;
    events: DiscordEvents;
    error?: string | null;
    offerPlans?: boolean;
    saving?: boolean;
    replacing?: boolean;
    onSubmit: () => void;
    onCancel: () => void;
  } = $props();

  function submit(event: SubmitEvent): void {
    event.preventDefault();
    onSubmit();
  }
</script>

<form class="flex flex-col gap-4" onsubmit={submit} novalidate>
  {#if !replacing}<p class="text-muted text-[13px]">{notificationsCopy.intro}</p>{/if}
  <div class="flex flex-col gap-2">
    <label class="label text-muted" for="discord-webhook">{notificationsCopy.fieldLabel}</label>
    <input
      type="text"
      name="discord-webhook"
      id="discord-webhook"
      class="bg-raised rounded-control text-text h-11 w-full border px-3 text-[15px] {error !== null
        ? 'border-wipe'
        : 'border-line-warm'}"
      inputmode="url"
      autocomplete="off"
      autocapitalize="off"
      autocorrect="off"
      spellcheck="false"
      enterkeyhint="done"
      placeholder={notificationsCopy.placeholder}
      aria-describedby="discord-webhook-help discord-webhook-error"
      aria-invalid={error !== null}
      disabled={saving}
      bind:value={webhook}
    />
    <div
      class="text-wipe min-h-[21px] text-[13px]"
      id="discord-webhook-error"
      data-testid="discord-webhook-error"
    >
      {#if error !== null}
        <span role="alert">{error}</span>
        {#if offerPlans}<a class="text-text underline" href="/premium">{notificationsCopy.seePlans}</a>{/if}
      {/if}
    </div>
    <div id="discord-webhook-help" class="text-muted flex flex-col gap-1 text-[13px]">
      <p>{notificationsCopy.howHelper}</p>
      <p>{notificationsCopy.trustHelper}</p>
    </div>
  </div>
  <DiscordEventRows
    {events}
    disabled={saving}
    busy={saving}
    onToggle={(next) => {
      events = next;
      return true;
    }}
  />
  <div class="flex flex-wrap items-center gap-x-4">
    <button
      type="submit"
      class="{PRIMARY_BUTTON_FIXED} px-5"
      disabled={saving}
      aria-busy={saving}
      data-testid="discord-save"
    >
      {saving ? notificationsCopy.saving : notificationsCopy.save}
    </button>
    {#if replacing}
      <button
        type="button"
        class={LINK_BUTTON}
        disabled={saving}
        onclick={onCancel}
        data-testid="discord-cancel"
      >
        {notificationsCopy.cancel}
      </button>
    {/if}
  </div>
</form>
