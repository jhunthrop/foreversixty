<!-- web/src/components/guild/setup/SetupPaste.svelte -->
<!-- A line an officer pastes into guild chat or Discord (guild onboarding spec §5): its label, the
     exact text in the code field, and Copy, which says Copied for a moment. -->
<script lang="ts">
  import { setupCopy } from '../../../lib/guild/setup-copy';
  import { COPIED_LONG_MS } from '../../../lib/ui/copy-feedback';
  import Button from '../../ui/Button.svelte';
  import CodeField from '../../ui/CodeField.svelte';

  let { label, text, testid }: { label: string; text: string; testid: string } = $props();

  let copied = $state(false);
  let timer: ReturnType<typeof setTimeout> | undefined;

  $effect(() => () => clearTimeout(timer));

  async function copy(): Promise<void> {
    try {
      await navigator.clipboard.writeText(text);
      copied = true;
      clearTimeout(timer);
      timer = setTimeout(() => (copied = false), COPIED_LONG_MS);
    } catch {
      copied = false;
    }
  }
</script>

<div class="flex flex-col gap-1.5" data-testid={testid}>
  <span class="text-muted text-[13px]">{label}</span>
  <div class="flex flex-col gap-2 sm:flex-row sm:items-start">
    <div class="min-w-0 flex-1">
      <CodeField value={text} {label} readonly testid={`${testid}-text`} />
    </div>
    <Button size="sm" onclick={() => void copy()} testid={`${testid}-copy`}
      >{copied ? setupCopy.invite.copied : setupCopy.invite.copy}</Button
    >
  </div>
</div>
