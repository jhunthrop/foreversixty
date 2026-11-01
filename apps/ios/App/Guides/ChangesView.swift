import DesignSystem
import Models
import SwiftUI

/// F9.5 under the segments: Key dates, then the changelog. Each region fails on its own.
struct ChangesView: View {
    /// Nil while the dates are read.
    let dates: [KeyDateRow]?
    let changelog: StaticPhase<Changelog>
    var open: (ChangelogEntry) -> Void = { _ in }
    var retry: () -> Void = {}

    var body: some View {
        VStack(alignment: .leading, spacing: 0) {
            if let dates, !dates.isEmpty {
                MicroLabel(Copy.changesDates).accessibilityAddTraits(.isHeader).padding(.bottom, 8)
                GuideCardRows(items: dates) { KeyDateRowView(row: $0) }
            }
            GuideSectionLabel(text: Copy.changesList)
            switch changelog {
            case .loading: SkeletonBar(width: nil, height: 240)
            case .notCached: GuidesNotCached(text: Copy.guidesListNotCached)
            case .failed: ReportsErrorCard(text: Copy.changesError, retry: retry)
            case let .loaded(file, _):
                GuideCardRows(items: ChangelogLogic.rows(file)) { row in
                    Button { open(row.entry) } label: { ChangelogRowView(row: row) }
                        .buttonStyle(RowPressStyle())
                }
            }
        }
    }
}

/// A key date: the 56 x 48 date block, the label, the note, and the gold running line.
struct KeyDateRowView: View {
    let row: KeyDateRow

    var body: some View {
        HStack(alignment: .top, spacing: 12) {
            Text(row.key)
                .font(ScaledType.monoMeta.weight(.semibold))
                .foregroundStyle(FSColor.strong)
                .multilineTextAlignment(.center)
                .minimumScaleFactor(0.7)
                .frame(minWidth: GuidesLayout.dateBlock.width, minHeight: GuidesLayout.dateBlock.height)
                .background(FSColor.cardTop, in: .rect(cornerRadius: FSRadius.control))
                .overlay(RoundedRectangle(cornerRadius: FSRadius.control).stroke(FSColor.line))
            VStack(alignment: .leading, spacing: 3) {
                Text(row.label).font(ScaledType.name).foregroundStyle(FSColor.strong)
                if let note = row.note { Text(note).font(ScaledType.meta).foregroundStyle(FSColor.muted) }
                if let running = row.running { Text(running).font(ScaledType.meta).foregroundStyle(FSColor.gold) }
            }
            .fixedSize(horizontal: false, vertical: true)
            .frame(maxWidth: .infinity, alignment: .leading)
        }
        .padding(GuidesLayout.cardPadding)
        .accessibilityElement(children: .combine)
    }
}

/// A changelog row: `10/28` mono muted, the title, the Site or Game tag.
struct ChangelogRowView: View {
    let row: ChangelogRow

    var body: some View {
        HStack(alignment: .top, spacing: 12) {
            Text(row.date).font(ScaledType.monoMeta).foregroundStyle(FSColor.muted).padding(.top, 2)
            Text(row.title).font(ScaledType.bodyStrong).foregroundStyle(FSColor.strong)
                .fixedSize(horizontal: false, vertical: true)
                .frame(maxWidth: .infinity, alignment: .leading)
            TagView(text: row.tag, color: FSColor.muted)
        }
        .padding(GuidesLayout.cardPadding)
        .frame(minHeight: 56)
        .contentShape(Rectangle())
        .accessibilityElement(children: .combine)
    }
}

/// A changelog entry in the guide reader's style: its blocks, then the sources disclosure.
struct ChangelogEntryView: View {
    let entry: ChangelogEntry

    var body: some View {
        VStack(alignment: .leading, spacing: 12) {
            Text(GuideDates.long(entry.date)).font(ScaledType.meta).foregroundStyle(FSColor.muted)
            Text(entry.title).font(ScaledType.h2).foregroundStyle(FSColor.strong)
                .fixedSize(horizontal: false, vertical: true)
                .accessibilityAddTraits(.isHeader)
            VStack(alignment: .leading, spacing: 10) {
                ForEach(Array(entry.blocks.enumerated()), id: \.offset) { _, block in ChangelogBlockView(block: block) }
            }
            .padding(GuidesLayout.cardPadding)
            .frame(maxWidth: .infinity, alignment: .leading)
            .quietCard()
            GuideSources(sources: entry.sources)
        }
        .padding(GuidesLayout.gutter)
        .padding(.bottom, GuidesLayout.tabBarClearance)
    }
}

/// One block: `p` a paragraph, `h` a heading, `li` a list line.
struct ChangelogBlockView: View {
    let block: ChangelogBlock

    private static let heading = "h"
    private static let listItem = "li"
    private static let bullet = "•"

    var body: some View {
        switch block.kind {
        case Self.heading:
            Text(block.text).font(ScaledType.sec).foregroundStyle(FSColor.strong).accessibilityAddTraits(.isHeader)
                .fixedSize(horizontal: false, vertical: true)
        case Self.listItem:
            HStack(alignment: .firstTextBaseline, spacing: 8) {
                Text(Self.bullet).foregroundStyle(FSColor.gold).accessibilityHidden(true)
                Text(block.text).fixedSize(horizontal: false, vertical: true)
            }
            .font(ScaledType.body).foregroundStyle(FSColor.text)
        default:
            Text(block.text).font(ScaledType.body).foregroundStyle(FSColor.text).fixedSize(horizontal: false, vertical: true)
        }
    }
}
