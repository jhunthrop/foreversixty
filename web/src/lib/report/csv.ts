// web/src/lib/report/csv.ts
// One CSV writer for every table's Copy button: every cell quoted, quotes doubled, so a
// name with a comma or a quote survives a spreadsheet's import.
export function toCsv(lines: readonly (readonly string[])[]): string {
  return lines.map((row) => row.map((cell) => `"${cell.replaceAll('"', '""')}"`).join(',')).join('\n');
}
