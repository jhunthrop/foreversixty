<!-- web/src/components/guild/GuildBotRemoveQuestion.svelte -->
<!-- The remove question (bot spec §3.1): it replaces the actions row of the installed state and of
     the picker, names the server, and says where raids go next. `Remove` and `Cancel` are outline
     buttons; focus goes to `Cancel` (the section's focus queue) and Escape acts as `Cancel` (the
     section's keydown). `disabled` is the frozen claim; `busy` is the call in flight. -->
<script lang="ts">
  import { guildBotCopy } from '../../lib/guild/discord-bot-copy';
  import { SECONDARY_BUTTON_FIXED } from '../../lib/planner/styles';

  const OUTLINE = `${SECONDARY_BUTTON_FIXED} border-line-warm text-text px-4`;
  const DIM = 'cursor-not-allowed opacity-50';

  let {
    server,
    webhook,
    busy,
    disabled,
    onRemove,
    onCancel,
  }: {
    server: string;
    /** A webhook channel is connected: raids fall back to it. */
    webhook: boolean;
    busy: boolean;
    disabled: boolean;
    onRemove: () => void;
    onCancel: () => void;
  } = $props();

  const off = $derived(busy || disabled);
</script>

<p class="text-[14px]" data-testid="bot-remove-question">{guildBotCopy.remove.question(server, webhook)}</p>
<div class="flex flex-wrap gap-3">
  <button
    class="{OUTLINE} {disabled ? DIM : ''}"
    disabled={off}
    aria-busy={busy}
    onclick={onRemove}
    data-testid="bot-remove-confirm"
  >
    {busy ? guildBotCopy.installed.removing : guildBotCopy.remove.confirm}
  </button>
  <button
    class="{OUTLINE} {disabled ? DIM : ''}"
    disabled={off}
    onclick={onCancel}
    data-testid="bot-remove-cancel"
  >
    {guildBotCopy.remove.cancel}
  </button>
</div>
