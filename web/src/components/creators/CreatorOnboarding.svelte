<!-- web/src/components/creators/CreatorOnboarding.svelte -->
<!-- /creators/onboarding (design/specs/2026-11-08-creator-program.md 3c): accepting the owner's invite,
     and Stripe's return_url (?done=1) and refresh_url (?refresh=1). One Section, one job per state.
     The invite token is sent in a POST body, never put in another URL; the Stripe link is made on
     each press and the page goes straight to it. -->
<script lang="ts">
  import { AccountError, fetchMeOnce } from '../../lib/account/api';
  import { acceptInvite, fetchCreatorMe, invitePreview, onboardingLink } from '../../lib/creators/api';
  import { onboardingCopy as copy } from '../../lib/creators/copy';
  import { button } from '../../lib/ui/button';
  import { LINK } from '../../lib/ui/primitives';
  import SignInPrompt from '../SignInPrompt.svelte';
  import Section from '../ui/Section.svelte';
  import Skeleton from '../ui/Skeleton.svelte';

  type View =
    | { kind: 'loading' }
    | { kind: 'signed-out'; next: string; invite: boolean }
    | { kind: 'invite'; token: string; name: string }
    | { kind: 'bad-invite' }
    | { kind: 'not-creator' }
    | { kind: 'ready' }
    | { kind: 'continue'; done: boolean };

  let {
    search = '',
    go = (url: string) => window.location.assign(url),
  }: { search?: string; go?: (url: string) => void } = $props();

  let view = $state<View>({ kind: 'loading' });
  let busy = $state(false);
  let error = $state('');

  const params = $derived(
    new URLSearchParams(search || (typeof window === 'undefined' ? '' : window.location.search)),
  );

  async function toStripe(): Promise<void> {
    busy = true;
    error = '';
    try {
      go(await onboardingLink());
    } catch {
      error = copy.failed;
      busy = false;
    }
  }

  async function load(): Promise<void> {
    const token = params.get('invite') ?? '';
    const me = await fetchMeOnce().catch(() => null);
    if (me === null) {
      const next = token
        ? `/creators/onboarding?invite=${encodeURIComponent(token)}`
        : '/creators/onboarding';
      view = { kind: 'signed-out', next, invite: token !== '' };
      return;
    }
    if (token !== '') {
      // The token leaves the address bar at once, so no later page's Referer or log carries it.
      const clean = new URL(window.location.href);
      clean.searchParams.delete('invite');
      window.history.replaceState(window.history.state, '', clean.pathname + clean.search + clean.hash);
      try {
        view = { kind: 'invite', token, name: await invitePreview(token) };
      } catch {
        view = { kind: 'bad-invite' };
      }
      return;
    }
    const done = params.get('done') === '1';
    try {
      const mine = await fetchCreatorMe(done);
      if (mine.stripe_status === 'ready') {
        view = { kind: 'ready' };
        return;
      }
      view = { kind: 'continue', done };
      if (params.get('refresh') === '1') await toStripe();
    } catch (err) {
      view =
        err instanceof AccountError && err.status === 403
          ? { kind: 'not-creator' }
          : { kind: 'continue', done };
    }
  }

  $effect(() => {
    void load();
  });

  async function accept(token: string): Promise<void> {
    busy = true;
    error = '';
    try {
      await acceptInvite(token);
      await toStripe();
    } catch {
      view = { kind: 'bad-invite' };
      busy = false;
    }
  }

  const title = $derived(
    view.kind === 'invite' || view.kind === 'bad-invite' ? copy.inviteTitle : copy.stripeTitle,
  );
</script>

<Section {title} testid="onboarding-section">
  <div class="flex flex-col gap-3" aria-live="polite">
    {#if view.kind === 'loading'}
      <Skeleton lines={3} testid="onboarding-loading" />
    {:else if view.kind === 'signed-out'}
      <SignInPrompt
        line={view.invite ? copy.inviteSignedOut : copy.notCreator}
        next={view.next}
        testid="onboarding-signin"
      />
    {:else if view.kind === 'invite'}
      {@const token = view.token}
      <p data-testid="onboarding-join">{copy.inviteJoin(view.name)}</p>
      <p class="text-muted text-[14px]">{copy.stripeLead}</p>
      <div>
        <button
          type="button"
          class={button('primary')}
          disabled={busy}
          aria-busy={busy}
          onclick={() => void accept(token)}
          data-testid="onboarding-accept">{busy ? copy.accepting : copy.accept}</button
        >
      </div>
    {:else if view.kind === 'bad-invite'}
      <p data-testid="onboarding-bad-invite">{copy.badInvite}</p>
    {:else if view.kind === 'not-creator'}
      <p data-testid="onboarding-not-creator">{copy.notCreator}</p>
    {:else if view.kind === 'ready'}
      <p data-testid="onboarding-ready">{copy.doneReady}</p>
      <p><a class={LINK} href="/creators/stats" data-testid="onboarding-stats">{copy.statsLink} ›</a></p>
    {:else}
      <p data-testid="onboarding-continue-line">{view.done ? copy.doneNotReady : copy.stripeLead}</p>
      <div>
        <button
          type="button"
          class={button('primary')}
          disabled={busy}
          aria-busy={busy}
          onclick={() => void toStripe()}
          data-testid="onboarding-continue">{busy ? copy.continuing : copy.continue}</button
        >
      </div>
    {/if}
    {#if error !== ''}<p class="text-wipe text-[14px]" role="alert" data-testid="onboarding-error">
        {error}
      </p>{/if}
  </div>
</Section>
