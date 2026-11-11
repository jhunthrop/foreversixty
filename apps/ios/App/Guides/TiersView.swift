import DesignSystem
import Models
import SwiftUI

/// F9.4 under the segments, as drawn from a phase: the role and faction chips, the callout for
/// your spec, then the file's rows in the file's order (never sorted, gapped or lettered here).
struct TiersView: View {
    let phase: StaticPhase<TiersFile>
    let choice: TierChoice
    /// The current character's guide key; nil with no character (no callout, no Yours).
    let yourSpec: String?
    var select: (TierChoice) -> Void = { _ in }
    var open: (TierRow) -> Void = { _ in }
    var retry: () -> Void = {}

    var body: some View {
        VStack(alignment: .leading, spacing: 12) {
            TierChips(choice: choice, select: select)
            switch phase {
            case .loading: TiersSkeleton()
            case .notCached: GuidesNotCached(text: Copy.guidesListNotCached)
            case .failed: LoadError(Copy.tiersError, retryLabel: Copy.commonTryAgain, retry: retry)
            case let .loaded(file, _): loaded(TiersLogic.list(file, choice: choice), build: file.build)
            }
        }
    }

    private func loaded(_ list: TierList, build: String) -> some View {
        VStack(alignment: .leading, spacing: 12) {
            if let callout = TiersLogic.callout(list, yourSpec: yourSpec) {
                TierCalloutCard(callout: callout, build: build)
            }
            let rows = TiersLogic.rows(list, yourSpec: yourSpec)
            if let unit = rows.first?.row.unit, !unit.isEmpty {
                // The figures' column head, as the web's table head.
                Label12(unit, color: FSColor.mutedGround)
                    .frame(maxWidth: .infinity, alignment: .trailing)
                    .padding(.trailing, GuidesLayout.cardPadding)
            }
            // The web's anatomy: the tier letter once, at the head of its group.
            let heads = Set(rows.indices.filter { $0 == 0 || rows[$0 - 1].letter != rows[$0].letter }.map { rows[$0].id })
            GuideCardRows(items: rows) { model in
                Button { open(model.row) } label: { TierRowView(model: model, build: build, showsLetter: heads.contains(model.id)) }
                    .buttonStyle(RowPressStyle())
                    .accessibilityLabel(model.a11y)
                    .accessibilityHint(Copy.tiersA11yOpens)
            }
        }
    }
}

/// `DPS`, `Tank`, `Healer` and the faction chip with its mark (40 pt chips on the control radius).
struct TierChips: View {
    let choice: TierChoice
    let select: (TierChoice) -> Void

    var body: some View {
        WrapLayout(spacing: 8) {
            ForEach(TierRole.allCases) { role in
                chip(role.title, isOn: role == choice.role) { select(choice.with(role: role)) }
            }
            Menu {
                ForEach(TierFaction.allCases) { faction in
                    Button(faction.title) { select(choice.with(faction: faction)) }
                }
            } label: {
                BisChip(label: choice.faction.title, action: {}) { FactionMark(faction: choice.faction.rawValue, size: 18) }
                    .allowsHitTesting(false)
            }
        }
        .haptic(.selection, trigger: choice)
    }

    private func chip(_ title: String, isOn: Bool, action: @escaping () -> Void) -> some View {
        Button(action: action) {
            Text(title)
                .font(ScaledType.bodyStrong)
                .foregroundStyle(isOn ? FSColor.gold : FSColor.strong)
                .padding(.horizontal, 16)
                .frame(minHeight: BisChip<EmptyView>.height)
                .chipSurface(fill: isOn ? FSColor.gold.opacity(0.12) : GameColors.controlFill, line: isOn ? FSColor.gold : FSColor.lineCard)
                .frame(minHeight: 44)
        }
        .buttonStyle(.plain)
        .accessibilityAddTraits(isOn ? .isSelected : [])
    }
}

/// The callout (the one `--surface-card`): your spec's place, in the file's own sentences.
struct TierCalloutCard: View {
    let callout: TierCallout
    let build: String

    var body: some View {
        VStack(alignment: .leading, spacing: 10) {
            HStack(alignment: .top, spacing: 12) {
                TierSpecIcon(row: callout.row, build: build, size: GuidesLayout.rowIcon)
                VStack(alignment: .leading, spacing: 4) {
                    Label12(Copy.tiersYours, color: FSColor.gold)
                    Text(callout.row.title).font(ScaledType.lead).foregroundStyle(FSColor.strong)
                        .fixedSize(horizontal: false, vertical: true)
                    Text(callout.row.detail).font(ScaledType.meta).foregroundStyle(FSColor.nav)
                        .fixedSize(horizontal: false, vertical: true)
                }
                .frame(maxWidth: .infinity, alignment: .leading)
            }
            .padding(GuidesLayout.cardPadding)
            .frame(maxWidth: .infinity, alignment: .leading)
            .primaryCard()
            .accessibilityElement(children: .combine)
            if let note = callout.note {
                Text(note).font(ScaledType.meta).foregroundStyle(FSColor.mutedGround).fixedSize(horizontal: false, vertical: true)
            }
        }
    }
}

/// The spec icon of a tier row; a spec with no icon draws its class crest, never a letter tile.
struct TierSpecIcon: View {
    let row: TierRow
    /// The tiers file's build: icons load from it.
    let build: String
    let size: CGFloat

    var body: some View {
        if let icon = row.icon {
            SpecIcon(icon: icon, build: build, size: size)
        } else {
            ClassCrest(className: row.classSlug, size: size)
        }
    }
}

/// One row (68 pt min), the web's anatomy: on a DPS list the tier letter in a left column at its
/// group's head, then rank, the ringed crest, the name in the class colour and race; right, the
/// figure over the gap (DPS) or the metric (tank, healer). No bars and no tier chips. At the
/// accessibility sizes the letter, rank, crest and tag lead a line of their own, and the name,
/// race and figures stack under it at the full width, so a name never breaks inside a word.
struct TierRowView: View {
    let model: TierRowModel
    let build: String
    /// True on the first row of its tier: the letter shows once per group.
    var showsLetter = false
    @Environment(\.dynamicTypeSize) private var typeSize

    /// The letter's column, so every group's rows line up under it.
    static let letterWidth: CGFloat = 24
    /// The tier letter: the display face at 24, as the web's tier column.
    static let letterFont = Font.custom(FSFont.title, size: 24, relativeTo: .title3)

    private var row: TierRow { model.row }
    private var color: Color { TierColors.of(model.letter) }

    var body: some View {
        Group {
            if typeSize.isAccessibilitySize { stacked } else { inline }
        }
        .padding(.horizontal, GuidesLayout.cardPadding)
        .padding(.vertical, 12)
        .frame(minHeight: GuidesLayout.tierRow)
        .background(model.isYours ? GameColors.rowWash : .clear)
        .overlay(alignment: .leading) { if model.isYours { FSColor.gold.frame(width: GuidesLayout.yoursEdge) } }
        .contentShape(Rectangle())
    }

    private var inline: some View {
        HStack(alignment: .center, spacing: 12) {
            letterColumn
            rank
            crest
            VStack(alignment: .leading, spacing: 4) {
                HStack(spacing: 8) {
                    name
                    if model.isYours { yours }
                }
                race
            }
            .frame(maxWidth: .infinity, alignment: .leading)
            trailing(alignment: .trailing)
        }
    }

    private var stacked: some View {
        VStack(alignment: .leading, spacing: 6) {
            HStack(spacing: 12) {
                letterColumn
                rank
                crest
                if model.isYours { yours }
            }
            name
            race
            trailing(alignment: .leading).padding(.top, 4)
        }
        .frame(maxWidth: .infinity, alignment: .leading)
    }

    /// The letter's column on a lettered (DPS) list; nothing on tank and healer lists.
    @ViewBuilder
    private var letterColumn: some View {
        if let letter = model.letter {
            Text(showsLetter ? letter : "")
                .font(Self.letterFont)
                .foregroundStyle(color)
                .frame(width: Self.letterWidth, alignment: .leading)
                .accessibilityHidden(true)
        }
    }

    private var crest: some View {
        ClassCrest(className: row.classSlug, size: GuidesLayout.tierIcon, flat: true).accessibilityHidden(true)
    }

    private var rank: some View {
        Text("\(row.rank)").font(ScaledType.tabularMeta).foregroundStyle(FSColor.muted).frame(minWidth: 20, alignment: .leading)
    }

    private var name: some View {
        Text(row.name).font(ScaledType.name).foregroundStyle(ClassColor.of(row.classSlug))
            .fixedSize(horizontal: false, vertical: true)
    }

    private var race: some View {
        Text(row.race).font(ScaledType.meta).foregroundStyle(FSColor.muted)
    }

    private var yours: some View {
        TagView(text: Copy.tiersYoursTag, color: FSColor.gold)
    }

    @ViewBuilder
    private func trailing(alignment: HorizontalAlignment) -> some View {
        if model.letter != nil {
            VStack(alignment: alignment, spacing: 2) {
                Text(model.metric).font(ScaledType.tabular).foregroundStyle(FSColor.strong)
                Text(model.gap).font(ScaledType.tabularMeta).foregroundStyle(FSColor.muted)
            }
        } else {
            VStack(alignment: alignment, spacing: 2) {
                Text(model.metric).font(ScaledType.tabular).foregroundStyle(FSColor.strong)
                Text(row.unit).font(ScaledType.meta).foregroundStyle(FSColor.muted)
                Text(model.gap).font(ScaledType.tabularMeta).foregroundStyle(FSColor.text)
            }
        }
    }
}

/// The tier colours: the page's own mapping onto the rarity tokens (`TierPage.astro`).
enum TierColors {
    static func of(_ letter: String?) -> Color {
        switch letter {
        case "S": FSColor.rarityLegendary
        case "A": FSColor.rarityEpicText
        case "B": FSColor.rarityRareText
        case "C": FSColor.rarityUncommon
        default: FSColor.muted
        }
    }
}

/// F9.7 for the tier list: the callout and five rows at their real heights.
struct TiersSkeleton: View {
    static let rows = 5

    var body: some View {
        VStack(spacing: 12) {
            SkeletonBar(width: nil, height: 92)
            VStack(spacing: 0) {
                ForEach(0..<Self.rows, id: \.self) { _ in
                    HStack(spacing: 12) {
                        SkeletonBar(width: GuidesLayout.tierIcon, height: GuidesLayout.tierIcon)
                        VStack(alignment: .leading, spacing: 8) { SkeletonBar(width: 160, height: 15); SkeletonBar(width: 90, height: 12) }
                        Spacer(minLength: 0)
                        SkeletonBar(width: GuidesLayout.tierFigure, height: 15)
                    }
                    .padding(.horizontal, GuidesLayout.cardPadding)
                    .frame(height: GuidesLayout.tierRow)
                }
            }
            .quietCard()
        }
        .accessibilityElement(children: .ignore)
        .accessibilityLabel(Copy.tiersLoading)
    }
}
