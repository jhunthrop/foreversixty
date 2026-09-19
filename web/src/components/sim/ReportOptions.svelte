<!-- web/src/components/sim/ReportOptions.svelte -->
<!-- Design 5.4: the report title and the finish notification, shown above the save form
     once a result is on screen. A pure render -- SimView.svelte keeps the `Notifier`
     itself and the `$effect` that actually raises the notification, since an `$effect`
     needs component scope bound to `store.result`; this component only draws the two
     controls and reports what the player did with them. -->
<script lang="ts">
  import { simCopy } from '../../lib/sim/copy';

  let {
    title,
    notifierAvailable,
    notifyWanted,
    ontitlechange,
    onnotifychange,
  }: {
    /** `store.reportTitle`: whatever the player typed, or the settings clause. */
    title: string;
    /** Whether the browser has a Notification API to ask at all (`browserNotifier() !== null`). */
    notifierAvailable: boolean;
    notifyWanted: boolean;
    ontitlechange: (value: string) => void;
    onnotifychange: (wanted: boolean) => void;
  } = $props();
</script>

<div class="mx-[18px] flex flex-wrap items-end gap-3 md:mx-0">
  <label class="flex min-w-0 flex-1 flex-col gap-1 md:max-w-[420px]">
    <span class="label text-muted">{simCopy.reportTitleLabel}</span>
    <input
      type="text"
      class="border-line-warm rounded-control bg-raised text-text h-11 w-full border px-3 text-[14px]"
      value={title}
      onchange={(event) => ontitlechange(event.currentTarget.value)}
      data-testid="sim-report-title"
    />
  </label>
  {#if notifierAvailable}
    <label class="flex min-h-11 items-center gap-2 text-[13px]">
      <input
        type="checkbox"
        class="accent-gold h-5 w-5"
        checked={notifyWanted}
        onchange={(event) => onnotifychange(event.currentTarget.checked)}
        data-testid="sim-notify"
      />
      <span class="text-muted">{simCopy.notifyLabel}</span>
    </label>
  {/if}
</div>
