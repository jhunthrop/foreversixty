<!-- web/src/components/SignInPrompt.svelte -->
<!-- What a signed-out visitor sees wherever /logs needs an account: one line saying what
     signing in is for, and a real button. A bare "Sign in" text link inside a sentence read
     as fine print, and the only button was the one in the header. -->
<script lang="ts">
  import { battlenetStartUrl } from '../lib/account/api';
  import { SECONDARY_BUTTON_FIXED } from '../lib/planner/styles';

  let { line, next = '/logs', testid }: { line: string; next?: string; testid: string } = $props();
</script>

<div class="flex flex-col items-start gap-3" data-testid={testid}>
  <p class="text-[14px]">
    {line}
    <a href="/login">Use an email link instead.</a>
  </p>
  <a class="{SECONDARY_BUTTON_FIXED} border-line-warm-strong text-strong px-4" href={battlenetStartUrl(next)}>
    Sign in with Battle.net
  </a>
</div>
