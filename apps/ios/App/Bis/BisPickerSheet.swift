import DesignSystem
import Models
import SwiftUI

/// A BiS chip (spec F3.1): 40 tall, pill, the 4.5% fill, the card border and a chevron.
struct BisChip<Leading: View>: View {
    let label: String
    var isOpen = false
    let action: () -> Void
    @ViewBuilder let leading: () -> Leading

    static var height: CGFloat { 40 }

    var body: some View {
        Button(action: action) {
            HStack(spacing: Theme.Spacing.sm) {
                leading()
                Text(label)
                    .font(Theme.Typography.bodyStrong)
                    .foregroundStyle(Theme.Palette.strong)
                    .fixedSize(horizontal: false, vertical: true)
                Image(systemName: "chevron.down")
                    .font(.system(size: 11, weight: .bold))
                    .foregroundStyle(Theme.Palette.muted)
                    .accessibilityHidden(true)
            }
            .padding(.leading, Theme.Spacing.sm)
            .padding(.trailing, Theme.Spacing.sm + Theme.Spacing.xs)
            .padding(.vertical, Theme.Spacing.xs)
            .frame(minHeight: Self.height)
            .pillSurface(fill: Theme.Palette.controlFill, line: isOpen ? Theme.Palette.accent : Theme.Palette.lineCard)
        }
        .buttonStyle(.plain)
    }
}

/// `Choose a list` (spec F3.2): class, spec, level range and faction. Every choice applies at
/// once and the list redraws under the sheet; the sheet stays until dismissed.
struct BisPickerSheet: View {
    let context: BisContext
    let catalog: SpecCatalog
    let onSelect: (BisSelection) -> Void

    private var selection: BisSelection { context.selection }

    var body: some View {
        ScrollView {
            VStack(alignment: .leading, spacing: 22) {
                Text(Copy.bisPickerTitle.uppercased())
                    .font(Theme.Typography.section)
                    .foregroundStyle(Theme.Palette.strong)
                    .accessibilityAddTraits(.isHeader)
                section(Copy.bisPickerClass) { classGrid }
                section(Copy.bisPickerSpec) { specs }
                section(Copy.bisPickerRange) { ranges }
                section(Copy.bisPickerFaction) { factions }
                if let back = context.backTarget {
                    SecondaryButton(title: Copy.bisPickerBack(name: back.name)) { onSelect(back.selection) }
                }
            }
            .padding(Theme.Spacing.md)
        }
        .background(Theme.Palette.sheet)
        .presentationDetents([.large])
        .presentationDragIndicator(.visible)
        .sensoryFeedback(.selection, trigger: selection)
    }

    private var classSlugs: [String] {
        if !context.classes.isEmpty { return context.classes.map(\.slug) }
        var seen: [String] = []
        for spec in catalog.specs where !seen.contains(spec.classSlug) { seen.append(spec.classSlug) }
        return seen
    }

    private var classGrid: some View {
        LazyVGrid(columns: Array(repeating: GridItem(.flexible(), spacing: Theme.Spacing.sm), count: 3), spacing: Theme.Spacing.sm) {
            ForEach(classSlugs, id: \.self) { slug in
                let isCurrent = slug == selection.classSlug
                Button { onSelect(selection.withClass(slug, catalog: catalog)) } label: {
                    VStack(spacing: 6) {
                        ClassCrest(classSlug: slug, size: 44)
                        Text(context.className(slug))
                            .font(Theme.Typography.bodyStrong)
                            .foregroundStyle(GamePalette.classColor(slug))
                    }
                    .frame(maxWidth: .infinity, minHeight: 88)
                    .background(isCurrent ? Theme.Palette.pressedWash : .clear, in: .rect(cornerRadius: Theme.Radius.control))
                    .overlay(RoundedRectangle(cornerRadius: Theme.Radius.control).stroke(isCurrent ? Theme.Palette.accent : .clear))
                }
                .buttonStyle(.plain)
                .accessibilityLabel(context.className(slug))
                .accessibilityAddTraits(isCurrent ? .isSelected : [])
            }
        }
    }

    private var specs: some View {
        WrapLayout(spacing: Theme.Spacing.sm) {
            ForEach(catalog.specs(of: selection.classSlug)) { spec in
                let isCurrent = spec.spec == selection.spec
                Button { onSelect(selection.withSpec(spec.spec)) } label: {
                    HStack(spacing: Theme.Spacing.sm + 2) {
                        SpecIcon(icon: spec.icon, size: 32)
                        VStack(alignment: .leading, spacing: 0) {
                            Text(spec.name).font(Theme.Typography.bodyStrong).foregroundStyle(Theme.Palette.strong)
                            Text(BisLabels.role(spec.role)).font(Theme.Typography.metaRole).foregroundStyle(Theme.Palette.muted)
                        }
                    }
                    .padding(.horizontal, Theme.Spacing.sm + 2)
                    .frame(minHeight: 48)
                    .background(isCurrent ? Theme.Palette.pressedWash : Theme.Palette.controlFill, in: .rect(cornerRadius: Theme.Radius.control))
                    .overlay(RoundedRectangle(cornerRadius: Theme.Radius.control).stroke(isCurrent ? Theme.Palette.accent : Theme.Palette.lineCard))
                }
                .buttonStyle(.plain)
                .accessibilityElement(children: .combine)
                .accessibilityAddTraits(isCurrent ? .isSelected : [])
            }
        }
    }

    private var ranges: some View {
        WrapLayout(spacing: Theme.Spacing.sm) {
            ForEach(BisRange.all, id: \.self) { range in
                pill(BisLabels.range(range), isCurrent: range == selection.range) { onSelect(selection.withRange(range)) }
            }
        }
    }

    private var factions: some View {
        WrapLayout(spacing: Theme.Spacing.sm) {
            ForEach(BisSelection.factions, id: \.self) { faction in
                pill(BisLabels.faction(faction), isCurrent: faction == selection.faction, mark: faction) {
                    onSelect(selection.withFaction(faction))
                }
            }
        }
    }

    private func pill(_ label: String, isCurrent: Bool, mark: String? = nil, action: @escaping () -> Void) -> some View {
        Button(action: action) {
            HStack(spacing: Theme.Spacing.sm) {
                if let mark { FactionMark(faction: mark, size: 22) }
                Text(label)
                    .font(isCurrent ? Theme.Typography.bodyStrong : Theme.Typography.bodyRole)
                    .foregroundStyle(isCurrent ? Theme.Palette.accentHover : Theme.Palette.text)
            }
            .padding(.horizontal, Theme.Spacing.md)
            .frame(minHeight: 48)
            .pillSurface(
                fill: isCurrent ? Theme.Palette.pressedWash : Theme.Palette.controlFill,
                line: isCurrent ? Theme.Palette.accent : Theme.Palette.lineCard
            )
        }
        .buttonStyle(.plain)
        .accessibilityAddTraits(isCurrent ? .isSelected : [])
    }

    private func section(_ title: String, @ViewBuilder _ body: () -> some View) -> some View {
        VStack(alignment: .leading, spacing: Theme.Spacing.sm + 2) {
            Text(title.uppercased())
                .font(Theme.Typography.micro)
                .tracking(1.5)
                .foregroundStyle(Theme.Palette.muted)
                .accessibilityAddTraits(.isHeader)
            body()
        }
    }
}
