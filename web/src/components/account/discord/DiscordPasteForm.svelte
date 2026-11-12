<!-- web/src/components/account/discord/DiscordPasteForm.svelte -->
<!-- Paste (or replace) the webhook: the three numbered steps (an ordered list, gold numerals), the
     input in the Field recipe (lib/ui/field, the one field: ui/Field cannot carry the URL keyboard
     attributes or the reserved error line with its plans link), the reserved error line directly
     under it, the trust helper, the events as switch rows and the one filled button. The typed URL lives only in the parent's state; it is
     never put in an attribute other than the input's own value. -->
<script lang="ts">
  import type { DiscordEvents } from '../../../lib/account/api';
  import { PREMIUM_LINKS } from '../../../lib/billing/premium-links';
  import { notificationsCopy } from '../../../lib/account/notifications-copy';
  import { button } from '../../../lib/ui/button';
  import { FIELD_CONTROL, FIELD_FOCUS, FIELD_HEIGHT, FIELD_LABEL, fieldEdge } from '../../../lib/ui/field';
  import { ACTION_LINK, ACTION_LINK_CHEVRON } from '../../../lib/ui/link';
  import Icon from '../../ui/Icon.svelte';
  import { ICON_CHEVRON_RIGHT } from '../../../lib/ui/ui-icons';
  import Button from '../../ui/Button.svelte';
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

<form class="flex flex-col gap-4" onsubmit={submit} novalidate>
  <div class="flex flex-col gap-4">
    {#if !replacing}<p class="text-text text-[14px]">{notificationsCopy.intro}</p>{/if}
    <ol class="marker:text-gold m-0 flex list-decimal flex-col gap-2 pl-5 text-[13px] marker:font-bold">
      {#each notificationsCopy.howSteps as step, index (step)}
        <li class="text-muted pl-1" data-testid="discord-step-{index + 1}">{step}</li>
      {/each}
    </ol>
    <div class="flex flex-col">
      <label class="{FIELD_LABEL} mb-[6px]" for="discord-webhook">{notificationsCopy.fieldLabel}</label>
      <input
        type="text"
        name="discord-webhook"
        id="discord-webhook"
        class={[FIELD_CONTROL, FIELD_HEIGHT, FIELD_FOCUS, fieldEdge(error !== null)]}
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
          {#if offerPlans}<a class="text-gold underline" href={PREMIUM_LINKS.notifications}
              >{notificationsCopy.seePlans}</a
            >{/if}
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
  <div class="flex flex-wrap items-center gap-x-3 gap-y-1">
    <Button
      type="submit"
      tone={secondary ? 'secondary' : 'primary'}
      busy={saving}
      busyLabel={notificationsCopy.saving}
      testid="discord-save"
    >
      {notificationsCopy.save}
    </Button>
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
        {notificationsCopy.connectInstead}<Icon icon={ICON_CHEVRON_RIGHT} class={ACTION_LINK_CHEVRON} />
      </button>
    {/if}
  </div>
</form>
