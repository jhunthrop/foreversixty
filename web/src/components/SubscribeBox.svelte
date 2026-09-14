<!-- web/src/components/SubscribeBox.svelte -->
<!-- The community panel's email form. client:visible, because it sits below the fold and
     must not compete with the homepage LCP budget. If the API cannot be reached the form
     is replaced by a link to the community Discord, so the panel is never a dead end.

     The submit button stays disabled until this component has actually mounted. All the
     form's submit handling arrives with hydration -- there is no `action` and no `method`,
     and submit() is what calls preventDefault() -- so between "scrolled into view" and
     "chunk executed" a click would run the browser's own submit instead: a GET back to the
     current page that navigates away and throws away what the user typed. Disabling the
     button is the honest state, because before hydration the form genuinely cannot do
     anything, and it says so to assistive tech rather than looking ready. It also stops the
     email field from ever becoming a query parameter should the input later gain a `name`
     (today it has none, so it is not a successful control and the native submit carries no
     data -- but that is one attribute away from being an address in the URL bar, browser
     history and the next request's Referer). -->
<script lang="ts">
  import { API_BASE_URL } from '../lib/planner/config';
  import { SECONDARY_BUTTON_FIXED } from '../lib/planner/styles';
  import { onMount } from 'svelte';
  import { subscribeMessageFor, type SubscribeMessage } from '../lib/subscribe';

  let { fallbackHref }: { fallbackHref: string } = $props();

  let email = $state('');
  let sending = $state(false);
  let result = $state<SubscribeMessage | null>(null);
  /** False through SSR and until this island mounts; gates the submit button. */
  let hydrated = $state(false);

  onMount(() => {
    hydrated = true;
  });

  async function submit(event: SubmitEvent): Promise<void> {
    event.preventDefault();
    sending = true;
    result = null;
    try {
      const response = await fetch(`${API_BASE_URL}/v1/subscribe`, {
        method: 'POST',
        headers: { 'content-type': 'application/json', accept: 'application/json' },
        body: JSON.stringify({ email }),
      });
      let apiMessage: string | null = null;
      try {
        const envelope = (await response.json()) as { error?: { message?: string } | null };
        apiMessage = envelope.error?.message ?? null;
      } catch {
        apiMessage = null;
      }
      result = subscribeMessageFor(response.status, apiMessage);
    } catch {
      result = subscribeMessageFor(0, null);
    }
    sending = false;
    if (result.kind === 'done') email = '';
  }
</script>

<form class="flex flex-col gap-2" onsubmit={submit} data-testid="subscribe">
  <label class="label text-muted" for="subscribe-email">Email updates</label>
  <div class="flex flex-wrap gap-2">
    <input
      id="subscribe-email"
      type="email"
      required
      autocomplete="email"
      bind:value={email}
      placeholder="you@example.com"
      class="border-line-warm rounded-control bg-raised text-text placeholder:text-muted h-11 min-w-[8rem] flex-1 border px-3 text-[14px]"
    />
    <button
      type="submit"
      disabled={sending || !hydrated}
      class="{SECONDARY_BUTTON_FIXED} border-line-warm-strong text-text px-4"
    >
      {sending ? 'Sending' : 'Subscribe'}
    </button>
  </div>
  <p class="text-muted text-[12px]">One email when something changes. Confirm the address first.</p>
  {#if result}
    <p role="status" class="text-[13px]" data-testid="subscribe-result">{result.message}</p>
    {#if result.kind === 'offline'}
      <a href={fallbackHref} class="text-[13px] font-semibold" data-testid="subscribe-fallback">
        Join the Discord instead
      </a>
    {/if}
  {/if}
</form>
