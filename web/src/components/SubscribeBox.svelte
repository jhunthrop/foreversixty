<!-- web/src/components/SubscribeBox.svelte -->
<!-- The community panel's email form. client:visible, because it sits below the fold and
     must not compete with the homepage LCP budget. If the API cannot be reached the form
     is replaced by a mailto link, so the panel is never a dead end. -->
<script lang="ts">
  import { API_BASE_URL } from '../lib/planner/config';
  import { subscribeMessageFor, type SubscribeMessage } from '../lib/subscribe';

  let { mailto }: { mailto: string } = $props();

  let email = $state('');
  let sending = $state(false);
  let result = $state<SubscribeMessage | null>(null);

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
      class="border-line-warm rounded-control bg-raised text-text placeholder:text-muted h-11 min-w-[200px] flex-1 px-3 text-[14px]"
    />
    <button
      type="submit"
      disabled={sending}
      class="border-line-warm-strong rounded-control text-text inline-flex h-11 items-center px-4 text-[12px] font-bold tracking-[0.06em] uppercase"
    >
      {sending ? 'Sending' : 'Subscribe'}
    </button>
  </div>
  <p class="text-muted text-[12px]">One email when something changes. Confirm the address first.</p>
  {#if result}
    <p role="status" class="text-[13px]" data-testid="subscribe-result">{result.message}</p>
    {#if result.kind === 'offline'}
      <a href={mailto} class="text-[13px] font-semibold" data-testid="subscribe-mailto">
        Email the site instead
      </a>
    {/if}
  {/if}
</form>
