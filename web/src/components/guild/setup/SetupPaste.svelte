<!-- web/src/components/guild/setup/SetupPaste.svelte -->
<!-- A line an officer pastes into guild chat or Discord (guild onboarding spec §5): its label, the
     exact text in the shared Code face, wrapping whole at any width, and the shared CopyButton. -->
<script lang="ts">
  import { setupCopy } from '../../../lib/guild/setup-copy';
  import CopyButton from '../../ui/CopyButton.svelte';
  import Code from '../../ui/Code.svelte';

  let { label, text, testid }: { label: string; text: string; testid: string } = $props();
</script>

<div class="flex flex-col gap-1.5" data-testid={testid}>
  <span class="text-muted text-[13px]">{label}</span>
  <div class="flex flex-col items-start gap-2 sm:flex-row sm:items-center">
    <p class="min-w-0 flex-1 [overflow-wrap:anywhere]"><Code {text} testid={`${testid}-text`} /></p>
    <CopyButton {text} label={setupCopy.invite.copy} tone="ghost" testid={`${testid}-copy`} />
  </div>
</div>
