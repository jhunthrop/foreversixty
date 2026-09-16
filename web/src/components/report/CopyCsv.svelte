<!-- web/src/components/report/CopyCsv.svelte -->
<!-- The Copy CSV button every table shares: the table hands over how to build its lines,
     the button puts them on the clipboard and says whether it managed. -->
<script lang="ts">
  import { toCsv } from '../../lib/report/csv';

  let { lines, label = 'Copy this table as CSV' }: { lines: () => string[][]; label?: string } = $props();

  let copied = $state('');
  async function copy(): Promise<void> {
    try {
      await navigator.clipboard.writeText(toCsv(lines()));
      copied = 'Copied';
    } catch {
      copied = 'Copy failed';
    }
    setTimeout(() => (copied = ''), 1500);
  }
</script>

<div class="flex justify-end">
  <button
    type="button"
    class="text-nav inline-flex min-h-11 items-center text-[12px] font-bold tracking-[0.06em] uppercase md:min-h-9"
    title={label}
    data-testid="copy-csv"
    onclick={() => void copy()}>{copied === '' ? 'Copy CSV' : copied}</button
  >
</div>
