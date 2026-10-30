import DesignSystem
import SwiftUI

/// The BiS tab's list (spec F3.1), drawn from plain state so snapshot tests render each state.
struct BisView: View {
    let phase: BisPhase
    let context: BisContext
    let catalog: SpecCatalog
    @Binding var searchText: String
    var onOpenPicker: () -> Void = {}
    var onTryAgain: () -> Void = {}
    var onOpenItem: (ItemTooltipContent) -> Void = { _ in }
    var now: Date = .now

    var body: some View {
        ScrollView {
            VStack(alignment: .leading, spacing: 14) {
                titleRow
                chips
                content
            }
            .padding(.horizontal, Theme.Spacing.md)
            .padding(.vertical, Theme.Spacing.sm)
        }
        .background(Theme.Palette.background)
        .safeAreaInset(edge: .top, spacing: 0) {
            if let offlineSince { OfflineStrip(savedAt: offlineSince, now: now) }
        }
    }

    private var offlineSince: Date? {
        switch phase {
        case let .loaded(_, since), let .noList(since): since
        default: nil
        }
    }

    private var titleRow: some View {
        HStack(alignment: .center, spacing: Theme.Spacing.sm + Theme.Spacing.xs) {
            Text(Copy.bisTitle.uppercased())
                .font(Theme.Typography.screenTitle)
                .foregroundStyle(Theme.Palette.strong)
                .accessibilityAddTraits(.isHeader)
            if context.isYours { Tag(text: Copy.bisYours, color: Theme.Palette.accent) }
        }
    }

    private var chips: some View {
        let selection = context.selection
        return WrapLayout(spacing: Theme.Spacing.sm) {
            BisChip(label: specChipLabel, isOpen: false, action: onOpenPicker) {
                if let spec = catalog.spec(selection.spec) { SpecIcon(icon: spec.icon, size: 28) }
            }
            BisChip(label: BisLabels.range(selection.range), action: onOpenPicker) { EmptyView() }
            BisChip(label: BisLabels.faction(selection.faction), action: onOpenPicker) {
                FactionMark(faction: selection.faction, size: 24)
            }
        }
        .accessibilityElement(children: .ignore)
        .accessibilityLabel(Copy.bisA11yPicker(
            spec: specChipLabel,
            range: BisLabels.range(selection.range),
            faction: BisLabels.faction(selection.faction)
        ))
        .accessibilityAddTraits(.isButton)
        .accessibilityAction { onOpenPicker() }
    }

    private var specChipLabel: String {
        let spec = catalog.spec(context.selection.spec)?.name ?? context.selection.spec
        return Copy.bisSpecChip(spec: spec, className: context.className(context.selection.classSlug))
    }

    @ViewBuilder
    private var content: some View {
        switch phase {
        case .loading:
            BisSkeleton()
        case let .loaded(list, _):
            if let line = list.setDpsLine { SetDpsLine(line: line) }
            let rows = list.filtered(by: searchText)
            if rows.isEmpty {
                Notice(text: Copy.bisSearchEmpty)
            } else {
                BisRowsCard(rows: rows, list: list, onOpenItem: onOpenItem)
            }
        case .noList:
            Notice(text: Copy.homeAnswerNoList)
        case .notCached:
            Notice(text: Copy.bisNotCached)
        case .failed:
            VStack(alignment: .leading, spacing: Theme.Spacing.md) {
                Text(Copy.bisFailed)
                    .font(Theme.Typography.bodyRole)
                    .foregroundStyle(Theme.Palette.wipe)
                    .accessibilityAddTraits(.isStaticText)
                SecondaryButton(title: Copy.commonTryAgain, action: onTryAgain)
            }
            .padding(.top, Theme.Spacing.sm)
        }
    }
}

/// `453.6 DPS at best in slot for your level range · Updated 29 Oct 2026`, the figure in mono.
private struct SetDpsLine: View {
    let line: String

    var body: some View {
        Text(attributed)
            .font(Theme.Typography.metaRole)
            .foregroundStyle(Theme.Palette.muted)
    }

    private var attributed: AttributedString {
        var text = AttributedString(line)
        let figureEnd = line.firstIndex(of: " ").map { line.distance(from: line.startIndex, to: $0) } ?? 0
        if figureEnd > 0 {
            let range = text.startIndex..<text.index(text.startIndex, offsetByCharacters: figureEnd)
            text[range].font = Theme.Typography.figure
            text[range].foregroundColor = Theme.Palette.strong
        }
        return text
    }
}

/// The one quiet card of 17 rows.
private struct BisRowsCard: View {
    let rows: [BisRow]
    let list: BisList
    let onOpenItem: (ItemTooltipContent) -> Void

    var body: some View {
        VStack(spacing: 0) {
            ForEach(Array(rows.enumerated()), id: \.element.id) { index, row in
                if index > 0 { Rectangle().fill(Theme.Palette.lineSoft).frame(height: 1) }
                switch row.content {
                case let .pick(pick):
                    Button { onOpenItem(BisTooltip.content(row: row, pick: pick, list: list)) } label: {
                        BisPickRow(pick: pick, row: row, unit: list.unit)
                    }
                    .buttonStyle(RowPressStyle())
                case let .empty(reason):
                    BisEmptyRow(reason: reason)
                }
            }
        }
        .background(Theme.Palette.raised, in: .rect(cornerRadius: Theme.Radius.panel))
        .overlay(RoundedRectangle(cornerRadius: Theme.Radius.panel).stroke(Theme.Palette.lineSoft))
        .clipShape(.rect(cornerRadius: Theme.Radius.panel))
    }
}

struct BisPickRow: View {
    let pick: BisPick
    let row: BisRow
    let unit: BisUnit
    @Environment(\.dynamicTypeSize) private var typeSize

    static let minHeight: CGFloat = 72

    var body: some View {
        // From the accessibility sizes the tag leaves the name's line, so the name keeps the width.
        let tagInline = !typeSize.isAccessibilitySize
        HStack(alignment: .top, spacing: Theme.Spacing.sm + Theme.Spacing.xs) {
            ItemIcon(icon: pick.item.icon, quality: pick.item.quality, size: 44)
            VStack(alignment: .leading, spacing: 2) {
                HStack(alignment: .top, spacing: Theme.Spacing.sm) {
                    Text(pick.item.name)
                        .font(Theme.Typography.bodyRole)
                        .foregroundStyle(GamePalette.qualityText(pick.item.quality))
                        .frame(maxWidth: .infinity, alignment: .leading)
                        .fixedSize(horizontal: false, vertical: true)
                    if pick.isWorn, tagInline { wornTag }
                }
                meta(pick.slotSource)
                meta(pick.next)
                if pick.isUnmodelled { meta(Copy.homeUpgradesNotModelled) }
                if pick.isWorn, !tagInline { wornTag.padding(.top, Theme.Spacing.xs) }
            }
        }
        .padding(.vertical, Theme.Spacing.sm + Theme.Spacing.xs)
        .padding(.horizontal, Theme.Spacing.md)
        .frame(maxWidth: .infinity, minHeight: Self.minHeight, alignment: .leading)
        .contentShape(.rect)
        .accessibilityElement(children: .ignore)
        .accessibilityLabel(accessibilityLabel)
        .accessibilityHint(Copy.bisA11yShowsItem)
        .accessibilityAddTraits(.isButton)
    }

    private var wornTag: some View {
        Tag(text: Copy.bisYouWear, color: Theme.Palette.kill, systemImage: "checkmark")
    }

    private func meta(_ text: String) -> some View {
        Text(text)
            .font(Theme.Typography.metaRole)
            .foregroundStyle(Theme.Palette.muted)
            .fixedSize(horizontal: false, vertical: true)
    }

    /// `Head. Blue Dragonscale Helm, Epic, item level 61. Crafted: Leatherworking. Next: …`
    private var accessibilityLabel: String {
        var parts = [row.slotLabel + "."]
        if let item = pick.item.item {
            parts.append(Copy.bisA11yItem(item: pick.item.name, quality: ItemWords.quality(item.quality), n: item.itemLevel))
        } else {
            parts.append(pick.item.name + ".")
        }
        if !pick.source.isEmpty { parts.append(pick.source + ".") }
        parts.append(pick.next + ".")
        if pick.isUnmodelled { parts.append(Copy.homeUpgradesNotModelled + ".") }
        if pick.isWorn { parts.append(Copy.bisYouWear + ".") }
        return parts.joined(separator: " ")
    }
}

private struct BisEmptyRow: View {
    let reason: String

    var body: some View {
        HStack(alignment: .center, spacing: Theme.Spacing.sm + Theme.Spacing.xs) {
            RoundedRectangle(cornerRadius: Theme.Radius.pill)
                .strokeBorder(Theme.Palette.line, style: StrokeStyle(lineWidth: 1, dash: [3, 3]))
                .frame(width: 44, height: 44)
                .accessibilityHidden(true)
            Text(reason)
                .font(Theme.Typography.metaRole)
                .foregroundStyle(Theme.Palette.muted)
                .frame(maxWidth: .infinity, alignment: .leading)
        }
        .padding(.vertical, Theme.Spacing.sm + Theme.Spacing.xs)
        .padding(.horizontal, Theme.Spacing.md)
        .frame(minHeight: BisPickRow.minHeight)
        .accessibilityElement(children: .combine)
    }
}

/// 17 rows at their real height while nothing is saved (spec §5.0 Loading).
private struct BisSkeleton: View {
    static let rows = 17

    var body: some View {
        VStack(spacing: 0) {
            ForEach(0..<Self.rows, id: \.self) { index in
                if index > 0 { Rectangle().fill(Theme.Palette.lineSoft).frame(height: 1) }
                HStack(alignment: .top, spacing: Theme.Spacing.sm + Theme.Spacing.xs) {
                    RoundedRectangle(cornerRadius: Theme.Radius.pill).fill(Theme.Palette.cardTop).frame(width: 44, height: 44)
                    VStack(alignment: .leading, spacing: 6) {
                        Capsule().fill(Theme.Palette.cardTop).frame(width: 170, height: 12)
                        Capsule().fill(Theme.Palette.cardTop).frame(width: 220, height: 10)
                        Capsule().fill(Theme.Palette.cardTop).frame(width: 190, height: 10)
                    }
                    Spacer(minLength: 0)
                }
                .padding(.vertical, Theme.Spacing.sm + Theme.Spacing.xs)
                .padding(.horizontal, Theme.Spacing.md)
                .frame(minHeight: BisPickRow.minHeight)
            }
        }
        .background(Theme.Palette.raised, in: .rect(cornerRadius: Theme.Radius.panel))
        .accessibilityElement(children: .ignore)
        .accessibilityLabel(Copy.bisLoading)
        .onAppear { AccessibilityNotification.Announcement(Copy.bisLoading).post() }
    }
}

private struct Notice: View {
    let text: String

    var body: some View {
        Text(text)
            .font(Theme.Typography.bodyRole)
            .foregroundStyle(Theme.Palette.text)
            .frame(maxWidth: .infinity, alignment: .leading)
            .padding(.top, Theme.Spacing.sm)
    }
}

/// The tooltip content of a list row.
enum BisTooltip {
    static func content(row: BisRow, pick: BisPick, list: BisList) -> ItemTooltipContent {
        ItemTooltipContent(
            id: pick.item.id,
            name: pick.item.name,
            item: pick.item.item,
            slotKey: row.slot,
            source: pick.source,
            nextBest: pick.runnersUp.prefix(maxNextBest).map { alternative in
                let gap = BisListBuilder.signedDelta(alternative.dpsDelta).map { Copy.tooltipDelta(delta: $0) }
                    ?? Copy.bisNextSameShort(unit: list.unit.word)
                return ItemTooltipAlternative(
                    id: alternative.item.id,
                    name: alternative.item.name,
                    icon: alternative.item.icon,
                    quality: alternative.item.quality,
                    meta: (alternative.isWorn ? Copy.tooltipYouWear : "") + gap
                )
            },
            pickOf: Copy.tooltipPickOf(slot: row.slotLabel, spec: list.specLabel, band: list.rangeLabel, preset: list.presetLabel)
        )
    }

    /// Spec F2.6: up to two alternatives.
    static let maxNextBest = 2
}
