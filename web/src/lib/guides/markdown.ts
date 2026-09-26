// Renders one guide section's raw markdown to an HTML string, independent of Astro's whole-
// entry `render(entry)`. [spec].astro needs this because a spec guide's markdown carries no
// MDX (not installed, and this worktree may never `npm install`) -- so a Svelte island or an
// Astro component cannot live inside the markdown itself. Splitting the raw body into its
// nine `## ` sections (sections.ts's own splitSpecSections) and rendering each one separately
// through this is what lets the page plant a real component between two of them while every
// section keeps the exact heading markup (id, tag) Astro's own pipeline would have produced,
// since this is the very processor that pipeline uses internally.
import { createSatteriMarkdownProcessor } from '@astrojs/markdown-satteri';

type Processor = Awaited<ReturnType<typeof createSatteriMarkdownProcessor>>;
let processor: Promise<Processor> | null = null;

function sharedProcessor(): Promise<Processor> {
  if (processor === null) processor = createSatteriMarkdownProcessor();
  return processor;
}

export async function renderGuideMarkdown(markdown: string): Promise<string> {
  if (markdown.trim() === '') return '';
  const { render } = await sharedProcessor();
  const { code } = await render(markdown);
  return code;
}
