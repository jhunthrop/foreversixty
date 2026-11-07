<!-- web/src/components/report/EmptyLine.svelte -->
<!-- A report region's one empty or state line ("Nobody died.", "Nothing in this window.", the
     trash rating line): muted 14/24 with the 16 px inset the mock library's `.empty` gives it, in a
     Plate or under a Section head (`section` takes back the head's own 12 px). -->
<script lang="ts">
  let {
    text,
    section = false,
    role,
    testid,
  }: { text: string; section?: boolean; role?: 'alert' | 'status'; testid?: string } = $props();
</script>

<p class={['text-muted p-4 text-[14px] leading-6', section && '-mt-3']} {role} data-testid={testid}>{text}</p>
