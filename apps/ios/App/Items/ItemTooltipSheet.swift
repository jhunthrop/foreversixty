import DesignSystem
import Models
import SwiftUI

// MERGE NOTE (lanes F2 and F3): this is F3's minimal item tooltip sheet (spec F2.6). Lane F2
// builds the full one in parallel; the two merge into this one file, F2's body winning. F3
// reads only `ItemTooltipContent` and `ItemTooltipSheet(content:)`.

/// A runner-up as the sheet's `The next best` lists it.
struct ItemTooltipAlternative: Equatable, Sendable, Identifiable {
    let id: Int
    let name: String
    let icon: String?
    let quality: Int
    /// `−1.3 DPS vs the pick`, led by `You wear this · ` for the worn one.
    let meta: String
}

/// Everything the sheet draws, already worded by the screen that opens it.
struct ItemTooltipContent: Equatable, Sendable, Identifiable {
    let id: Int
    let name: String
    /// The build's item facts; nil when the class's items file is not on the phone.
    let item: Item?
    let slotKey: String
    let source: String?
    let nextBest: [ItemTooltipAlternative]
    /// `Best in slot for Head · Elemental Shaman 60 · Raid-ready, Phase 1`, when a pick.
    let pickOf: String?

    var quality: Int { item?.quality ?? 1 }
}

/// The item tooltip sheet (spec F2.6): header, the game-ordered panel, where it comes from,
/// the next best, and the pick footer. Opens at the medium detent.
struct ItemTooltipSheet: View {
    let content: ItemTooltipContent

    var body: some View {
        ScrollView {
            VStack(alignment: .leading, spacing: Theme.Spacing.md) {
                header
                if let item = content.item { panel(item) }
                if let source = content.source, !source.isEmpty {
                    section(Copy.tooltipSource) {
                        Text(source).font(Theme.Typography.bodyRole).foregroundStyle(Theme.Palette.strong)
                    }
                }
                section(Copy.tooltipNextBest) { nextBest }
                if let pickOf = content.pickOf {
                    Text(pickOf).font(Theme.Typography.metaRole).foregroundStyle(Theme.Palette.muted)
                }
            }
            .padding(Theme.Spacing.md)
        }
        .background(Theme.Palette.sheet)
        .presentationDetents([.medium, .large])
        .presentationDragIndicator(.visible)
    }

    private var header: some View {
        HStack(spacing: Theme.Spacing.sm + Theme.Spacing.xs) {
            ItemIcon(icon: content.item?.icon, quality: content.quality, size: 56)
            VStack(alignment: .leading, spacing: 2) {
                Text(content.name)
                    .font(Theme.Typography.section)
                    .foregroundStyle(GamePalette.qualityText(content.quality))
                Text(Copy.bisQualitySlot(quality: ItemWords.quality(content.quality), slot: ItemWords.slot(content.slotKey)))
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
    private var nextBest: some View {
        if content.nextBest.isEmpty {
            Text(Copy.tooltipNothingElse).font(Theme.Typography.metaRole).foregroundStyle(Theme.Palette.muted)
        } else {
            ForEach(content.nextBest) { alternative in
                HStack(spacing: Theme.Spacing.sm + Theme.Spacing.xs) {
                    ItemIcon(icon: alternative.icon, quality: alternative.quality, size: 36)
                    VStack(alignment: .leading, spacing: 2) {
                        Text(alternative.name)
                            .font(Theme.Typography.bodyRole)
                            .foregroundStyle(GamePalette.qualityText(alternative.quality))
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
