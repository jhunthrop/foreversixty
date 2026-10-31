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
                    .font(ScaledType.bodyStrong)
                    .foregroundStyle(FSColor.strong)
                    .fixedSize(horizontal: false, vertical: true)
                Image(systemName: "chevron.down")
                    .font(.system(size: 11, weight: .bold))
                    .foregroundStyle(FSColor.muted)
                    .accessibilityHidden(true)
            }
            .padding(.leading, Theme.Spacing.sm)
            .padding(.trailing, Theme.Spacing.sm + Theme.Spacing.xs)
            .padding(.vertical, Theme.Spacing.xs)
            .frame(minHeight: Self.height)
            .pillSurface(fill: GameColors.controlFill, line: isOpen ? FSColor.gold : FSColor.lineCard)
        }
        .buttonStyle(.plain)
    }
}

/// `Choose a list` (spec F3.2): class, spec, level range and faction. Every choice applies at
/// once and the list redraws under the sheet; the sheet stays until dismissed.
struct BisPickerSheet: View {
    let context: BisContext
    let catalog: SpecCatalog
    /// The build the spec icons load from; nil before the first config read.
    let build: String?
    let onSelect: (BisSelection) -> Void
    @Environment(\.dynamicTypeSize) private var typeSize

    private var selection: BisSelection { context.selection }

    /// Three crests across; two from the accessibility sizes, so a class name never breaks inside
    /// the word.
    private var classColumns: Int { typeSize.isAccessibilitySize ? 2 : 3 }

    var body: some View {
        ScrollView {
            VStack(alignment: .leading, spacing: 22) {
                Text(Copy.bisPickerTitle.uppercased())
                    .font(ScaledType.sec)
                    .foregroundStyle(FSColor.strong)
                    .accessibilityAddTraits(.isHeader)
                section(Copy.bisPickerClass) { classGrid }
                section(Copy.bisPickerSpec) { specs }
                section(Copy.bisPickerRange) { ranges }
                section(Copy.bisPickerFaction) { factions }
                if let back = context.backTarget {
                    SecondaryButton(title: Copy.bisPickerBack(name: back.name), fullWidth: true) { onSelect(back.selection) }
                }
            }
            .padding(Theme.Spacing.md)
        }
        .background(GameColors.sheet)
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
        LazyVGrid(columns: Array(repeating: GridItem(.flexible(), spacing: Theme.Spacing.sm), count: classColumns), spacing: Theme.Spacing.sm) {
            ForEach(classSlugs, id: \.self) { slug in
                let isCurrent = slug == selection.classSlug
                Button { onSelect(selection.withClass(slug, catalog: catalog)) } label: {
                    VStack(spacing: 6) {
                        ClassCrest(className: slug, size: 44)
                        Text(context.className(slug))
                            .font(ScaledType.bodyStrong)
                            .foregroundStyle(ClassColor.of(slug))
                            .lineLimit(1)
                            .minimumScaleFactor(0.7)
                    }
                    .frame(maxWidth: .infinity, minHeight: 88)
                    .background(isCurrent ? GameColors.pressedWash : .clear, in: .rect(cornerRadius: FSRadius.control))
                    .overlay(RoundedRectangle(cornerRadius: FSRadius.control).stroke(isCurrent ? FSColor.gold : .clear))
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
                        SpecIcon(icon: spec.icon, build: build ?? "", size: 32)
                        VStack(alignment: .leading, spacing: 0) {
                            Text(spec.name).font(ScaledType.bodyStrong).foregroundStyle(FSColor.strong)
                            Text(BisLabels.role(spec.role)).font(ScaledType.meta).foregroundStyle(FSColor.muted)
                        }
                    }
                    .padding(.horizontal, Theme.Spacing.sm + 2)
                    .frame(minHeight: 48)
                    .background(isCurrent ? GameColors.pressedWash : GameColors.controlFill, in: .rect(cornerRadius: FSRadius.control))
                    .overlay(RoundedRectangle(cornerRadius: FSRadius.control).stroke(isCurrent ? FSColor.gold : FSColor.lineCard))
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
                pill(context.rangeLabel(range), isCurrent: range == selection.range) { onSelect(selection.withRange(range)) }
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
                    .font(isCurrent ? ScaledType.bodyStrong : ScaledType.body)
                    .foregroundStyle(isCurrent ? FSColor.goldHover : FSColor.text)
            }
            .padding(.horizontal, Theme.Spacing.md)
            .frame(minHeight: 48)
            .pillSurface(
                fill: isCurrent ? GameColors.pressedWash : GameColors.controlFill,
                line: isCurrent ? FSColor.gold : FSColor.lineCard
            )
        }
        .buttonStyle(.plain)
        .accessibilityAddTraits(isCurrent ? .isSelected : [])
    }

    private func section(_ title: String, @ViewBuilder _ body: () -> some View) -> some View {
        VStack(alignment: .leading, spacing: Theme.Spacing.sm + 2) {
            Text(title.uppercased())
                .font(ScaledType.micro)
                .tracking(1.5)
                .foregroundStyle(FSColor.muted)
                .accessibilityAddTraits(.isHeader)
            body()
        }
    }
}
