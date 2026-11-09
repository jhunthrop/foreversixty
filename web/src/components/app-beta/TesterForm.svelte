<!-- web/src/components/app-beta/TesterForm.svelte -->
<!-- The tester signup on /app (mock web-app, design/mocks/ws_app_beta.py): the job line, then the
     form (Email, Your phone as a Segmented iPhone / Android defaulting to the visitor's phone, the
     one primary, the consent line, the failure line under the button), or, once sent, the one line
     that says the address is on the list. What a press comes to is lib/app-beta/form.ts. -->
<script lang="ts">
  import { onMount } from 'svelte';
  import { appBetaCopy as copy, PLATFORM_NAMES } from '../../lib/app-beta/copy';
  import { submitSignup } from '../../lib/app-beta/form';
  import {
    DEFAULT_PLATFORM,
    detectPlatform,
    MAX_EMAIL_LENGTH,
    type Platform,
  } from '../../lib/app-beta/platform';
  import { button } from '../../lib/ui/button';
  import { FIELD_LABEL } from '../../lib/ui/field';
  import Field from '../ui/Field.svelte';
  import Segmented from '../ui/Segmented.svelte';

  const PLATFORMS: readonly { value: Platform; label: string }[] = [
    { value: 'ios', label: PLATFORM_NAMES.ios },
    { value: 'android', label: PLATFORM_NAMES.android },
  ];

  let email = $state('');
  let platform = $state<Platform>(DEFAULT_PLATFORM);
  /** The visitor picked a phone themselves: the detected one never overrides it. */
  let picked = false;
  let sending = $state(false);
  let emailError = $state('');
  let failure = $state('');
  let done = $state('');

  // After hydration only: the server render cannot see the phone, so it draws the default.
  onMount(() => {
    if (!picked) platform = detectPlatform(navigator.userAgent, navigator.maxTouchPoints);
  });

  async function onSubmit(event: SubmitEvent): Promise<void> {
    event.preventDefault();
    sending = true;
    emailError = '';
    failure = '';
    const outcome = await submitSignup(email, platform);
    sending = false;
    if (outcome.kind === 'done') done = outcome.message;
    else if (outcome.kind === 'invalid') emailError = outcome.message;
    else failure = outcome.message;
  }
</script>

<div class="flex max-w-[440px] flex-col gap-3" data-testid="app-beta-form-body">
  <p class="text-muted text-[15px] leading-6">{copy.formJob}</p>
  {#if done !== ''}
    <p class="text-strong text-[15px] leading-6" role="status" data-testid="app-beta-done">{done}</p>
  {:else}
    <form class="flex flex-col gap-4" onsubmit={onSubmit} novalidate data-testid="app-beta-form">
      <Field
        id="app-beta-email"
        name="email"
        type="email"
        label={copy.emailLabel}
        placeholder={copy.emailPlaceholder}
        autocomplete="email"
        maxlength={MAX_EMAIL_LENGTH}
        required
        error={emailError || undefined}
        bind:value={email}
        testid="app-beta-email"
      />
      <div class="flex flex-col gap-[6px]">
        <span class={FIELD_LABEL} aria-hidden="true">{copy.platformLabel}</span>
        <Segmented
          options={PLATFORMS}
          value={platform}
          label={copy.platformLabel}
          onChange={(next) => {
            picked = true;
            platform = next;
          }}
          testid="app-beta-platform"
        />
      </div>
      <div class="flex flex-col items-start gap-2">
        <button
          type="submit"
          class={button('primary')}
          disabled={sending}
          aria-busy={sending}
          data-testid="app-beta-submit">{sending ? copy.submitting : copy.submit}</button
        >
        <p class="text-small text-muted m-0" data-testid="app-beta-consent">{copy.consent}</p>
        {#if failure !== ''}<p class="text-wipe text-small m-0" role="alert" data-testid="app-beta-failed">
            {failure}
          </p>{/if}
      </div>
    </form>
  {/if}
</div>
