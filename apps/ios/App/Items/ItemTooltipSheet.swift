import DesignSystem
import Models
import SwiftUI

// MERGE NOTE (lanes F2 and F3): F2's full sheet replaces this file at integration (coordinator
// ruling, 2026-10-31). The entry point is F2's: `ItemTooltipSheet(request:catalog:)`, with the
// pure body `ItemTooltipContent(request:phase:retry:)`. F3's body here is minimal on purpose: no
// stats, weapon line, Binds, Unique, armor type or set lines.

/// The item tooltip as a sheet (spec F2.6): opens at the medium detent, grows to large. It reads
/// only the build's static files, so it needs no API request.
struct ItemTooltipSheet: View {
    let request: ItemTooltipRequest
    let catalog: ItemCatalog

    @State private var phase: Phase = .loading

    enum Phase: Equatable {
        case loading
        case loaded(ItemTooltipModel)
        case failed
    }

    var body: some View {
        ItemTooltipContent(request: request, phase: phase, retry: { Task { await load() } })
            .presentationDetents([.medium, .large])
            .presentationDragIndicator(.visible)
            .task { await load() }
    }

    private func load() async {
        phase = .loading
        do {
            let sources = try await catalog.sources(for: request)
            phase = .loaded(.build(request, sources: sources))
        } catch {
            phase = .failed
        }
    }
}

/// The sheet's body, pure over its phase so the snapshot tests render it directly.
struct ItemTooltipContent: View {
    let request: ItemTooltipRequest
    let phase: ItemTooltipSheet.Phase
    var retry: () -> Void = {}

    var body: some View {
        ScrollView {
            VStack(alignment: .leading, spacing: Theme.Spacing.md) {
                switch phase {
                case .loading:
                    header(request.fallback)
                    Text(Copy.itemsLoading).font(Theme.Typography.metaRole).foregroundStyle(Theme.Palette.muted)
                case let .loaded(model):
                    header(model.look, slotKey: model.slotKey)
                    if let item = model.item { panel(item) }
                    if let source = model.source, !source.isEmpty {
                        section(Copy.tooltipSource) {
                            Text(source).font(Theme.Typography.bodyRole).foregroundStyle(Theme.Palette.strong)
                        }
                    }
                    section(Copy.tooltipNextBest) { nextBest(model.alternatives) }
                    if let footer = model.pickFooter {
                        Text(footer).font(Theme.Typography.metaRole).foregroundStyle(Theme.Palette.muted)
                    }
                case .failed:
                    header(request.fallback)
                    Text(Copy.itemsFailed).font(Theme.Typography.bodyRole).foregroundStyle(Theme.Palette.wipe)
                    SecondaryButton(title: Copy.commonTryAgain, action: retry)
                }
            }
            .padding(Theme.Spacing.md)
        }
        .background(Theme.Palette.sheet)
    }

    private func header(_ look: HomeItemLook, slotKey: String? = nil) -> some View {
        let quality = look.quality ?? 1
        let qualityName = ItemWords.quality(quality)
        return HStack(spacing: Theme.Spacing.sm + Theme.Spacing.xs) {
            ItemIcon(icon: look.icon, quality: quality, size: 56)
            VStack(alignment: .leading, spacing: 2) {
                Text(look.name ?? "")
                    .font(Theme.Typography.section)
                    .foregroundStyle(GamePalette.qualityText(quality))
                Text(slotKey.map { Copy.bisQualitySlot(quality: qualityName, slot: ItemWords.slot($0)) } ?? qualityName)
                    .font(Theme.Typography.metaRole)
                    .foregroundStyle(Theme.Palette.muted)
            }
        }
        .accessibilityElement(children: .combine)
    }

    private func panel(_ item: Item) -> some View {
        VStack(alignment: .leading, spacing: 2) {
            Text(item.name).foregroundStyle(GamePalette.qualityText(item.quality))
            Text(ItemWords.slot(item.slot)).foregroundStyle(Theme.Palette.strong)
            if item.armor > 0 {
                Text(Copy.tooltipArmor(n: item.armor)).foregroundStyle(Theme.Palette.strong)
            }
            if !item.effectText.isEmpty {
                Text(item.effectText).foregroundStyle(Theme.Palette.gameGreen)
            }
            if item.requiredLevel > 0 {
                Text(Copy.tooltipRequires(n: item.requiredLevel)).foregroundStyle(Theme.Palette.strong)
            }
            Text(Copy.tooltipItemLevel(n: item.itemLevel)).foregroundStyle(Theme.Palette.itemLevel)
        }
        .font(Theme.Typography.bodyRole)
        .frame(maxWidth: .infinity, alignment: .leading)
        .padding(.vertical, Theme.Spacing.sm + Theme.Spacing.xs)
        .padding(.horizontal, 14)
        .background(Theme.Palette.tooltipPanel, in: .rect(cornerRadius: Theme.Radius.control))
        .overlay(RoundedRectangle(cornerRadius: Theme.Radius.control).stroke(Theme.Palette.iconLine))
        .accessibilityElement(children: .combine)
    }

    @ViewBuilder
    private func nextBest(_ alternatives: [ItemTooltipAlternative]) -> some View {
        if alternatives.isEmpty {
            Text(Copy.tooltipNothingElse).font(Theme.Typography.metaRole).foregroundStyle(Theme.Palette.muted)
        } else {
            ForEach(alternatives) { alternative in
                let quality = alternative.look.quality ?? 1
                HStack(spacing: Theme.Spacing.sm + Theme.Spacing.xs) {
                    ItemIcon(icon: alternative.look.icon, quality: quality, size: 36)
                    VStack(alignment: .leading, spacing: 2) {
                        Text(alternative.look.name ?? "")
                            .font(Theme.Typography.bodyRole)
                            .foregroundStyle(GamePalette.qualityText(quality))
                        Text(alternative.meta).font(Theme.Typography.metaRole).foregroundStyle(Theme.Palette.muted)
                    }
                }
                .accessibilityElement(children: .combine)
            }
        }
    }

    private func section(_ title: String, @ViewBuilder _ body: () -> some View) -> some View {
        VStack(alignment: .leading, spacing: Theme.Spacing.sm) {
            Text(title.uppercased())
                .font(Theme.Typography.micro)
                .tracking(1.5)
                .foregroundStyle(Theme.Palette.muted)
                .accessibilityAddTraits(.isHeader)
            body()
        }
    }
}
