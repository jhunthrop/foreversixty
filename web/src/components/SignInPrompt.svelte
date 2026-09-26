<!-- web/src/components/SignInPrompt.svelte -->
<!-- What a signed-out visitor sees wherever /logs needs an account: one line saying what
     signing in is for, and a real button. A bare "Sign in" text link inside a sentence read
     as fine print, and the only button was the one in the header.

     `compact` is for the second and third prompt on the same page (the companion pairing
     box and the upload form on /logs): the "Your reports" panel above them already carries
     the button, and three identical buttons down one page read as a broken state. A
     compact prompt keeps the same line and the same Battle.net link, as text. -->
<script lang="ts">
  import { battlenetStartUrl } from '../lib/account/api';
  import { SECONDARY_BUTTON_FIXED } from '../lib/planner/styles';

  let {
    line,
    next = '/account?signed_in=1',
    testid,
    compact = false,
  }: { line: string; next?: string; testid: string; compact?: boolean } = $props();
</script>

<div class="flex flex-col items-start gap-3" data-testid={testid}>
  <p class="text-[14px]">
    {line}
    {#if compact}
      <a class="text-nav underline" href={battlenetStartUrl(next)}>Sign in with Battle.net</a>
      or
      <a href="/login">use an email link.</a>
    {:else}
      <a href="/login">Use an email link instead.</a>
    {/if}
  </p>
  {#if !compact}
    <a
      class="{SECONDARY_BUTTON_FIXED} border-line-warm-strong text-strong px-4"
      href={battlenetStartUrl(next)}
    >
      Sign in with Battle.net
    </a>
  {/if}
</div>
