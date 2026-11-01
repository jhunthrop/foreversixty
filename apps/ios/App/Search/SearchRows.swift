import DesignSystem
import DesignTokens
import Models
import SwiftUI

/// The row sizes of spec F12.2.
enum SearchLayout {
    static let itemIcon: CGFloat = 40
    static let rowIcon: CGFloat = 36
    static let linkTile: CGFloat = 32
    static let rowMinHeight: CGFloat = 56
    static let chipHeight: CGFloat = 36
    static let fieldHeight: CGFloat = 44
    static let target: CGFloat = 44
    static let sectionGap: CGFloat = 22
}

/// What a row ends with: nothing, a chevron (it opens in the app), or the external-link glyph
/// (it opens the website).
enum SearchRowTrailing {
    case none, chevron, external
}

/// Where a result goes. The one place that says which screens are still the website's:
/// guides until F9, characters and rankings until F11. Drop a case from `website` as its lane lands.
enum SearchDestination {
    case item, character, guide, talent, rankings, bis

    private static let website: Set<SearchDestination> = [.character, .guide, .rankings]

    init(_ target: SearchTarget) {
        switch target {
        case .item: self = .item
        case .character: self = .character
        case .guide: self = .guide
        case .talent: self = .talent
        }
    }

    var opensWebsite: Bool { Self.website.contains(self) }

    var trailing: SearchRowTrailing {
        if opensWebsite { return .external }
        switch self {
        case .guide, .character, .rankings, .bis: return .chevron
        case .item, .talent: return .none
        }
    }
}

/// A result row's frame: icon, a name over a meta line that wraps, then its trailing glyph. The
/// row reads as one element and its whole height is the 44 pt target.
struct SearchRow<Icon: View, Title: View>: View {
    let meta: String?
    var trailing: SearchRowTrailing = .none
    let action: () -> Void
    @ViewBuilder let icon: () -> Icon
    @ViewBuilder let title: () -> Title

    var body: some View {
        Button(action: action) {
            HStack(spacing: 12) {
                icon()
                VStack(alignment: .leading, spacing: 2) {
                    title()
                    if let meta, !meta.isEmpty {
                        Text(meta).font(ScaledType.meta).foregroundStyle(FSColor.muted)
                            .fixedSize(horizontal: false, vertical: true)
                    }
                }
                Spacer(minLength: 8)
                switch trailing {
                case .none: EmptyView()
                case .chevron: glyph("chevron.right")
                case .external: glyph("arrow.up.right")
                }
            }
            .padding(.horizontal, 14)
            .padding(.vertical, 8)
            .frame(maxWidth: .infinity, minHeight: SearchLayout.rowMinHeight, alignment: .leading)
            .contentShape(.rect)
        }
        .buttonStyle(RowPressStyle())
        .accessibilityElement(children: .combine)
    }

    private func glyph(_ name: String) -> some View {
        Image(systemName: name).font(.system(size: 13, weight: .semibold))
            .foregroundStyle(FSColor.muted).accessibilityHidden(true)
    }
}

/// A row's name: wraps, never cut. At accessibility sizes it wraps only between words, and a
/// word wider than the row (a long character name) shrinks to fit rather than break inside it.
struct SearchRowName: View {
    let text: String
    var color: Color = FSColor.strong
    @Environment(\.dynamicTypeSize) private var typeSize

    var body: some View {
        if typeSize.isAccessibilitySize {
            WordWrapText(text: text, font: ScaledType.bodyStrong, color: color)
        } else {
            Text(text).font(ScaledType.bodyStrong).foregroundStyle(color).fixedSize(horizontal: false, vertical: true)
        }
    }
}

/// Words laid out with WrapLayout: lines break only between words; a word that cannot fit a
/// line on its own scales down instead of splitting (lane F15 ruling: never a break in a name).
struct WordWrapText: View {
    let text: String
    let font: Font
    let color: Color
    @ScaledMetric(relativeTo: .body) private var gap: CGFloat = 4

    private static let minimumScale = 0.5

    var body: some View {
        WrapLayout(spacing: gap) {
            ForEach(Array(text.split(separator: " ").enumerated()), id: \.offset) { _, word in
                Text(word).font(font).foregroundStyle(color).lineLimit(1).minimumScaleFactor(Self.minimumScale)
            }
        }
        .accessibilityElement(children: .ignore)
        .accessibilityLabel(text)
    }
}

/// A talent's square icon in the 1 pt icon line, as the talent tree draws it.
struct SearchTalentIcon: View {
    let icon: String
    let name: String
    let build: String
    var size: CGFloat = SearchLayout.rowIcon

    var body: some View {
        let shape = RoundedRectangle(cornerRadius: FSRadius.pill)
        GameIconImage(icon: icon, build: build, fallback: String(name.prefix(2)))
            .id("\(build)/\(icon)")
            .frame(width: size, height: size)
            .clipShape(shape)
            .overlay(shape.stroke(GameColors.iconLine, lineWidth: 1))
            .accessibilityHidden(true)
    }
}

/// The rows of every group, one view per kind (spec F12.2 row shapes).
@MainActor
enum SearchRows {
    static func item(_ item: SearchItem, build: String, action: @escaping () -> Void) -> some View {
        SearchRow(meta: SearchWords.itemMeta(item), action: action) {
            ItemIconView(icon: item.icon, quality: item.quality, build: build, size: SearchLayout.itemIcon)
        } title: {
            SearchRowName(text: item.name, color: ItemQuality(value: item.quality).text)
        }
    }

    static func guide(
        _ hit: SearchGuideHit, build: String, trailing: SearchRowTrailing = SearchDestination.guide.trailing,
        action: @escaping () -> Void
    ) -> some View {
        SearchRow(meta: SearchWords.guideMeta(hit.guide), trailing: trailing, action: action) {
            SpecIcon(icon: hit.guide.icon, build: build, size: SearchLayout.rowIcon)
        } title: {
            SearchRowName(text: hit.title, color: ClassColor.of(hit.guide.className))
        }
    }

    static func talent(_ talent: SearchTalent, build: String, action: @escaping () -> Void) -> some View {
        SearchRow(meta: SearchWords.talentMeta(talent), action: action) {
            SearchTalentIcon(icon: talent.icon, name: talent.name, build: build)
        } title: {
            SearchRowName(text: talent.name)
        }
    }

    /// The watch star (F11) joins this row when `/v1/me/watching` exists; until then no star.
    static func character(
        _ row: CharacterSearchRow, isYours: Bool, trailing: SearchRowTrailing = SearchDestination.character.trailing,
        action: @escaping () -> Void
    ) -> some View {
        SearchRow(meta: SearchWords.characterMeta(row), trailing: trailing, action: action) {
            ClassCrest(className: row.className, size: SearchLayout.rowIcon)
        } title: {
            if isYours {
                // `You` beside the name while it fits, else under it; the name is never cut.
                ViewThatFits(in: .horizontal) {
                    HStack(spacing: 8) { characterName(row); youTag }
                    VStack(alignment: .leading, spacing: 4) { characterName(row); youTag }
                }
            } else {
                characterName(row)
            }
        }
    }

    private static func characterName(_ row: CharacterSearchRow) -> some View {
        SearchRowName(text: row.name, color: ClassColor.of(row.className))
    }

    private static var youTag: some View { TagView(text: Copy.searchCharacterYou, color: FSColor.gold) }

    static func recent(_ entry: RecentSearch, build: String, action: @escaping () -> Void) -> some View {
        SearchRow(meta: SearchWords.kind(entry.target), trailing: SearchDestination(entry.target).trailing, action: action) {
            recentIcon(entry, build: build)
        } title: {
            SearchRowName(text: entry.title, color: recentColor(entry))
        }
    }

    @ViewBuilder
    private static func recentIcon(_ entry: RecentSearch, build: String) -> some View {
        switch entry.target {
        case .item: ItemIconView(icon: entry.icon, quality: entry.quality, build: build, size: SearchLayout.rowIcon)
        case .character: ClassCrest(className: entry.className, size: SearchLayout.rowIcon)
        case .guide: SpecIcon(icon: entry.icon ?? "", build: build, size: SearchLayout.rowIcon)
        case .talent: SearchTalentIcon(icon: entry.icon ?? "", name: entry.title, build: build)
        }
    }

    private static func recentColor(_ entry: RecentSearch) -> Color {
        switch entry.target {
        case .item: ItemQuality(value: entry.quality).text
        case .character, .guide: ClassColor.of(entry.className)
        case .talent: FSColor.strong
        }
    }

    /// A `For {name}` link: a 32 pt glyph tile, the label, its trailing glyph.
    static func link(_ title: String, systemImage: String, trailing: SearchRowTrailing, action: @escaping () -> Void) -> some View {
        SearchRow(meta: nil, trailing: trailing, action: action) {
            Image(systemName: systemImage)
                .font(.system(size: 14, weight: .semibold))
                .foregroundStyle(FSColor.gold)
                .frame(width: SearchLayout.linkTile, height: SearchLayout.linkTile)
                .background(FSColor.cardTop, in: .rect(cornerRadius: FSRadius.pill))
                .overlay(RoundedRectangle(cornerRadius: FSRadius.pill).stroke(FSColor.lineSoft))
                .accessibilityHidden(true)
        } title: {
            SearchRowName(text: title)
        }
    }

    /// A group's last row when it has more: `Show all {n}` in gold.
    static func showAll(_ total: Int, action: @escaping () -> Void) -> some View {
        Button(action: action) {
            Text(Copy.searchShowAll(n: total)).font(ScaledType.bodyStrong).foregroundStyle(FSColor.gold)
                .padding(.horizontal, 14)
                .frame(maxWidth: .infinity, minHeight: SearchLayout.target, alignment: .leading)
                .contentShape(.rect)
        }
        .buttonStyle(RowPressStyle())
    }
}

/// A quiet card of rows split by soft lines.
struct SearchCard<Content: View>: View {
    @ViewBuilder let content: () -> Content

    var body: some View {
        VStack(spacing: 0) {
            Group(subviews: content()) { rows in
                ForEach(Array(rows.enumerated()), id: \.offset) { index, row in
                    if index > 0 { FSColor.lineSoft.frame(height: 1).padding(.leading, 14) }
                    row
                }
            }
        }
        .clipShape(.rect(cornerRadius: FSRadius.panel))
        .quietCard()
    }
}

/// A section: the micro label (with an optional trailing action), then its card.
struct SearchSection<Content: View>: View {
    let title: String
    var trailing: (label: String, action: () -> Void)?
    @ViewBuilder let content: () -> Content

    var body: some View {
        VStack(alignment: .leading, spacing: 8) {
            HStack {
                MicroLabel(title).accessibilityAddTraits(.isHeader)
                Spacer()
                if let trailing {
                    Button(trailing.label, action: trailing.action)
                        .font(ScaledType.bodyStrong)
                        .foregroundStyle(FSColor.gold)
                        .frame(minHeight: SearchLayout.target)
                }
            }
            .frame(minHeight: trailing == nil ? nil : SearchLayout.target)
            content()
        }
    }
}
