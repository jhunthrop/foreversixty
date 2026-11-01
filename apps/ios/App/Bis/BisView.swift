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
    var onOpenItem: (ItemTooltipRequest) -> Void = { _ in }
    var build: String?
    var now: Date = .now
    /// Lane F15: the new-data strip and its Update.
    var newData: NewDataState = .none
    var onUpdate: () -> Void = {}

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
        .background(FSColor.bg)
        .safeAreaInset(edge: .top, spacing: 0) {
            VStack(spacing: 0) {
                if let offlineSince { OfflineStrip(text: Copy.offlineStrip(time: OfflineStamp.format(offlineSince, now: now))) }
                NewDataStrip(newData, action: onUpdate)
            }
        }
    }

    private var offlineSince: Date? {
        switch phase {
        case let .loaded(_, since), let .noList(since): since
        default: nil
        }
    }

    private var titleRow: some View {
        HStack(alignment: .firstTextBaseline, spacing: Theme.Spacing.sm + Theme.Spacing.xs) {
            Text(Copy.bisTitle.uppercased())
                .font(ScaledType.title)
                .foregroundStyle(FSColor.strong)
                .accessibilityAddTraits(.isHeader)
            if context.isYours { TagView(text: Copy.bisYours, color: FSColor.gold) }
        }
    }

    private var chips: some View {
        let selection = context.selection
        return WrapLayout(spacing: Theme.Spacing.sm) {
            BisChip(label: specChipLabel, isOpen: false, action: onOpenPicker) {
                if let spec = catalog.spec(selection.spec) { SpecIcon(icon: spec.icon, build: build ?? "", size: 28) }
            }
            BisChip(label: context.rangeLabel(selection.range, now: now), action: onOpenPicker) { EmptyView() }
            BisChip(label: BisLabels.faction(selection.faction), action: onOpenPicker) {
                FactionMark(faction: selection.faction, size: 24)
            }
        }
        .accessibilityElement(children: .ignore)
        .accessibilityLabel(Copy.bisA11yPicker(
            spec: specChipLabel,
            range: context.rangeLabel(selection.range, now: now),
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
                BisRowsCard(rows: rows, list: list, context: context, build: build, onOpenItem: onOpenItem)
            }
        case .noList:
            StateNotice(text: Copy.homeAnswerNoList, color: FSColor.muted)
        case .notCached:
            StateNotice(text: Copy.bisNotCached, color: FSColor.muted, onTryAgain: onTryAgain)
        case .failed:
            StateNotice(text: Copy.bisFailed, color: FSColor.wipe, onTryAgain: onTryAgain)
        }
    }
}

/// `453.6 DPS at best in slot for your level range · Updated 29 Oct 2026`, the figure in mono.
private struct SetDpsLine: View {
    let line: String

    var body: some View {
        Text(attributed)
            .font(ScaledType.meta)
            .foregroundStyle(FSColor.muted)
    }

    private var attributed: AttributedString {
        var text = AttributedString(line)
        let figureEnd = line.firstIndex(of: " ").map { line.distance(from: line.startIndex, to: $0) } ?? 0
        if figureEnd > 0 {
            let range = text.startIndex..<text.index(text.startIndex, offsetByCharacters: figureEnd)
            text[range].font = ScaledType.mono
            text[range].foregroundColor = FSColor.strong
        }
        return text
    }
}

/// The one quiet card of 17 rows.
private struct BisRowsCard: View {
    let rows: [BisRow]
    let list: BisList
    let context: BisContext
    let build: String?
    let onOpenItem: (ItemTooltipRequest) -> Void

    var body: some View {
        VStack(spacing: 0) {
            ForEach(Array(rows.enumerated()), id: \.element.id) { index, row in
                if index > 0 { Rectangle().fill(FSColor.lineSoft).frame(height: 1) }
                switch row.content {
                case let .pick(pick):
                    Button {
                        if let build { onOpenItem(BisTooltip.request(row: row, pick: pick, list: list, context: context, build: build)) }
                    } label: {
                        BisPickRow(pick: pick, row: row, unit: list.unit, build: build ?? "")
                    }
                    .buttonStyle(RowPressStyle())
                case let .empty(reason):
                    BisEmptyRow(reason: reason)
                }
            }
        }
        .background(FSColor.raised, in: .rect(cornerRadius: FSRadius.panel))
        .overlay(RoundedRectangle(cornerRadius: FSRadius.panel).stroke(FSColor.lineSoft))
        .clipShape(.rect(cornerRadius: FSRadius.panel))
    }
}

struct BisPickRow: View {
    let pick: BisPick
    let row: BisRow
    let unit: BisUnit
    let build: String
    @Environment(\.dynamicTypeSize) private var typeSize

    static let minHeight: CGFloat = 72

    var body: some View {
        // From the accessibility sizes the tag leaves the name's line, so the name keeps the width.
        let tagInline = !typeSize.isAccessibilitySize
        HStack(alignment: .top, spacing: Theme.Spacing.sm + Theme.Spacing.xs) {
            ItemIconView(icon: pick.item.icon, quality: pick.item.quality, build: build, size: 44)
            VStack(alignment: .leading, spacing: 2) {
                HStack(alignment: .top, spacing: Theme.Spacing.sm) {
                    Text(pick.item.name)
                        .font(ScaledType.body)
                        .foregroundStyle(ItemQuality(value: pick.item.quality).text)
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
        TagView(text: Copy.bisYouWear, color: FSColor.kill, systemImage: "checkmark")
    }

    private func meta(_ text: String) -> some View {
        Text(text)
            .font(ScaledType.meta)
            .foregroundStyle(FSColor.muted)
            .fixedSize(horizontal: false, vertical: true)
    }

    /// `Head. Blue Dragonscale Helm, Epic, item level 61. Crafted: Leatherworking. Next: …`
    private var accessibilityLabel: String {
        var parts = [row.slotLabel + "."]
        if let item = pick.item.item {
            parts.append(Copy.bisA11yItem(item: pick.item.name, quality: ItemQuality(value: item.quality).name, n: item.itemLevel))
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
            RoundedRectangle(cornerRadius: FSRadius.pill)
                .strokeBorder(FSColor.line, style: StrokeStyle(lineWidth: 1, dash: [3, 3]))
                .frame(width: 44, height: 44)
                .accessibilityHidden(true)
            Text(reason)
                .font(ScaledType.meta)
                .foregroundStyle(FSColor.muted)
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
                if index > 0 { Rectangle().fill(FSColor.lineSoft).frame(height: 1) }
                HStack(alignment: .top, spacing: Theme.Spacing.sm + Theme.Spacing.xs) {
                    RoundedRectangle(cornerRadius: FSRadius.pill).fill(FSColor.cardTop).frame(width: 44, height: 44)
                    VStack(alignment: .leading, spacing: 6) {
                        Capsule().fill(FSColor.cardTop).frame(width: 170, height: 12)
                        Capsule().fill(FSColor.cardTop).frame(width: 220, height: 10)
                        Capsule().fill(FSColor.cardTop).frame(width: 190, height: 10)
                    }
                    Spacer(minLength: 0)
                }
                .padding(.vertical, Theme.Spacing.sm + Theme.Spacing.xs)
                .padding(.horizontal, Theme.Spacing.md)
                .frame(minHeight: BisPickRow.minHeight)
            }
        }
        .background(FSColor.raised, in: .rect(cornerRadius: FSRadius.panel))
        .accessibilityElement(children: .ignore)
        .accessibilityLabel(Copy.bisLoading)
        .onAppear { AccessibilityNotification.Announcement(Copy.bisLoading).post() }
    }
}

private struct Notice: View {
    let text: String

    var body: some View {
        Text(text)
            .font(ScaledType.body)
            .foregroundStyle(FSColor.text)
            .frame(maxWidth: .infinity, alignment: .leading)
            .padding(.top, Theme.Spacing.sm)
    }
}

/// A state of the list region (not saved, no list, error): its sentence, 24 below the chips, and
/// `Try again` where a retry can help (spec §5.0).
private struct StateNotice: View {
    let text: String
    let color: Color
    var onTryAgain: (() -> Void)?

    var body: some View {
        VStack(alignment: .leading, spacing: Theme.Spacing.md) {
            Text(text)
                .font(ScaledType.body)
                .foregroundStyle(color)
                .frame(maxWidth: .infinity, alignment: .leading)
            if let onTryAgain { SecondaryButton(title: Copy.commonTryAgain, fullWidth: true, action: onTryAgain) }
        }
        // 24 below the chips: the stack's 14 plus 10.
        .padding(.top, 10)
    }
}

/// The tooltip request of a list row, in F2's entry point's terms (`ItemTooltipSheet(request:catalog:)`).
enum BisTooltip {
    static func request(row: BisRow, pick: BisPick, list: BisList, context: BisContext, build: String) -> ItemTooltipRequest {
        let selection = context.selection
        let item = pick.item.item
        return ItemTooltipRequest(
            itemID: pick.item.id,
            fallback: HomeItemLook(id: pick.item.id, name: pick.item.name, quality: item?.quality, icon: item?.icon, itemLevel: item?.itemLevel),
            build: build,
            classSlug: selection.classSlug,
            spec: selection.spec,
            specName: list.specLabel,
            band: selection.range.band,
            bandLabel: list.rangeLabel,
            preset: selection.range.preset,
            faction: selection.faction,
            sourceLine: pick.source,
            wornID: pick.runnersUp.first(where: \.isWorn)?.item.id ?? (pick.isWorn ? pick.item.id : nil)
        )
    }
}
