<!-- web/src/components/report/CopyCsv.svelte -->
<!-- The Copy CSV button every table shares: the table hands over how to build its lines,
     the button puts them on the clipboard and says whether it managed. -->
<script lang="ts">
  import { toCsv } from '../../lib/report/csv';
  import { LINK } from '../../lib/ui/primitives';

  let {
    lines,
    label = 'Copy this table as CSV',
    inHead = false,
  }: {
    lines: () => string[][];
    label?: string;
    /** In a Section head's action slot: one quiet text line, 44 px only on a coarse pointer. */
    inHead?: boolean;
  } = $props();

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
    class={inHead
      ? `${LINK} inline-flex items-center text-[14px] [@media(pointer:coarse)]:-my-3 [@media(pointer:coarse)]:py-3`
      : `${LINK} inline-flex min-h-11 items-center text-[14px] md:min-h-9`}
    title={label}
    data-testid="copy-csv"
    onclick={() => void copy()}>{copied === '' ? 'Copy CSV' : copied}</button
  >
</div>
