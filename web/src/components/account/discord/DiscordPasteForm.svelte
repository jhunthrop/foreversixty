<!-- web/src/components/account/discord/DiscordPasteForm.svelte -->
<!-- Paste (or replace) the webhook: the three numbered steps (24 px numbered tiles), the labelled
     52 px input, the reserved error line directly under it, the trust helper, the events as
     switch rows and the one filled button. The typed URL lives only in the parent's state; it is
     never put in an attribute other than the input's own value. -->
<script lang="ts">
  import type { DiscordEvents } from '../../../lib/account/api';
  import { notificationsCopy } from '../../../lib/account/notifications-copy';
  import { button } from '../../../lib/ui/button';
  import { ACTION_LINK, ACTION_LINK_CHEVRON } from '../../../lib/ui/link';
  import Glyph from '../../ui/Glyph.svelte';
  import Spinner from '../../ui/Spinner.svelte';
  import DiscordEventRows from './DiscordEventRows.svelte';

  let {
    webhook = $bindable(),
    events = $bindable(),
    error = null,
    offerPlans = false,
    saving = false,
    replacing = false,
    secondary = false,
    onSubmit,
    onCancel,
    onConnect,
  }: {
    webhook: string;
    events: DiscordEvents;
    error?: string | null;
    offerPlans?: boolean;
    saving?: boolean;
    replacing?: boolean;
    /** Another block holds the page's one fill (a seat offer, launch pricing spec 5): Save is secondary. */
    secondary?: boolean;
    onSubmit: () => void;
    onCancel: () => void;
    /** Present only where connect is on: swaps back to the Connect view. */
    onConnect?: () => void;
  } = $props();

  function submit(event: SubmitEvent): void {
    event.preventDefault();
    onSubmit();
  }
</script>

<form class="flex flex-col gap-4 pt-1" onsubmit={submit} novalidate>
  <div class="flex flex-col gap-4 px-5">
    {#if !replacing}<p class="text-text text-[14px]">{notificationsCopy.intro}</p>{/if}
    <ol class="m-0 flex list-none flex-col gap-2 p-0 text-[13px]">
      {#each notificationsCopy.howSteps as step, index (step)}
        <li class="text-muted flex items-center gap-3" data-testid="discord-step-{index + 1}">
          <span
            class="bg-tile text-tile-ink rounded-pill inline-flex h-6 w-6 shrink-0 items-center justify-center text-[12px] font-bold tabular-nums"
            aria-hidden="true">{index + 1}</span
          >{step}
        </li>
      {/each}
    </ol>
    <div class="flex flex-col">
      <label class="quiet-label mb-2" for="discord-webhook">{notificationsCopy.fieldLabel}</label>
      <input
        type="text"
        name="discord-webhook"
        id="discord-webhook"
        class="rounded-control text-strong focus-visible:outline-gold h-[52px] w-full border bg-white/[.045] px-3 text-[15px] focus-visible:outline-2 focus-visible:outline-offset-2 {error !==
        null
          ? 'border-wipe'
          : 'border-line-card'}"
        inputmode="url"
        autocomplete="off"
        autocapitalize="off"
        autocorrect="off"
        spellcheck="false"
        enterkeyhint="done"
        placeholder={notificationsCopy.placeholder}
        aria-describedby="discord-webhook-error discord-webhook-help"
        aria-invalid={error !== null}
        disabled={saving}
        bind:value={webhook}
      />
      <div
        class="text-wipe min-h-[21px] pt-1 text-[13px]"
        id="discord-webhook-error"
        data-testid="discord-webhook-error"
      >
        {#if error !== null}
          <span role="alert">{error}</span>
          {#if offerPlans}<a class="text-gold underline" href="/premium">{notificationsCopy.seePlans}</a>{/if}
        {/if}
      </div>
      <p id="discord-webhook-help" class="text-muted text-[13px]">{notificationsCopy.trustHelper}</p>
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
  <div class="flex flex-wrap items-center gap-x-3 gap-y-1 px-5 pb-5">
    <button
      type="submit"
      class={button(secondary ? 'secondary' : 'primary')}
      disabled={saving}
      aria-busy={saving ? 'true' : undefined}
      data-testid="discord-save"
    >
      {#if saving}<Spinner />{/if}{saving ? notificationsCopy.saving : notificationsCopy.save}
    </button>
    {#if replacing}
      <button
        type="button"
        class={button('ghost')}
        disabled={saving}
        onclick={onCancel}
        data-testid="discord-cancel"
      >
        {notificationsCopy.cancel}
      </button>
    {/if}{#if onConnect !== undefined}
      <button
        type="button"
        class="{ACTION_LINK} basis-full"
        onclick={onConnect}
        data-testid="discord-connect-instead"
      >
        {notificationsCopy.connectInstead}<Glyph name="chevron" class={ACTION_LINK_CHEVRON} />
      </button>
    {/if}
  </div>
</form>
