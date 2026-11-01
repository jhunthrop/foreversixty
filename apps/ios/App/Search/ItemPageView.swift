import DesignSystem
import DesignTokens
import Models
import SwiftUI
import UIKit

/// The item page's three regions, which the strip chips scroll to.
enum ItemPageAnchor: CaseIterable, Hashable {
    case item, source, bisFor

    var title: String {
        switch self {
        case .item: Copy.itemPageChipItem
        case .source: Copy.itemPageChipSource
        case .bisFor: Copy.itemPageChipBisFor
        }
    }
}

/// F12.4 and F12.5, pushed from Search, from F2.6's `More about this item` and from a BiS row's
/// sheet. Back only, titled `Search` from search and `Back` elsewhere; no share (no web item page).
struct ItemPageScreen: View {
    let route: ItemPageRoute
    let dependencies: SearchDependencies
    let backTitle: String
    /// A spec row's chevron: that spec's list in the BiS tab.
    let openList: (BisSelection) -> Void

    @State private var store: ItemPageStore

    init(route: ItemPageRoute, dependencies: SearchDependencies, backTitle: String = Copy.itemPageBackPlain, openList: @escaping (BisSelection) -> Void) {
        self.route = route
        self.dependencies = dependencies
        self.backTitle = backTitle
        self.openList = openList
        _store = State(initialValue: ItemPageStore(route: route, dependencies: dependencies))
    }

    var body: some View {
        ItemPageView(
            state: store.state,
            openSpec: { row in openList(ItemPageLogic.selection(row.use, viewerFaction: dependencies.context().character?.faction)) },
            retry: { Task { await store.load() } }
        )
        .itemPageBar(backTitle: backTitle)
        .task { await store.load() }
    }
}

extension View {
    /// The item page's bar: no title, and a back control that names where it returns to (the
    /// iOS 26 bar draws a bare chevron otherwise). The edge swipe still goes back.
    func itemPageBar(backTitle: String) -> some View {
        modifier(ItemPageBar(backTitle: backTitle))
    }
}

private struct ItemPageBar: ViewModifier {
    let backTitle: String
    @Environment(\.dismiss) private var dismiss

    func body(content: Content) -> some View {
        content
            .navigationTitle("")
            .navigationBarTitleDisplayMode(.inline)
            .navigationBarBackButtonHidden(true)
            .toolbar {
                ToolbarItem(placement: .topBarLeading) {
                    Button { dismiss() } label: {
                        Label(backTitle, systemImage: "chevron.left").labelStyle(.titleAndIcon)
                    }
                    .frame(minHeight: SearchLayout.target)
                }
            }
            .toolbarBackground(FSColor.bg, for: .navigationBar)
            .background(SwipeBackKeeper())
    }
}

/// A hidden system back button also turns the edge swipe off; this turns it back on.
private struct SwipeBackKeeper: UIViewControllerRepresentable {
    func makeUIViewController(context _: Context) -> Keeper { Keeper() }
    func updateUIViewController(_: Keeper, context _: Context) {}

    final class Keeper: UIViewController {
        override func viewDidAppear(_ animated: Bool) {
            super.viewDidAppear(animated)
            navigationController?.interactivePopGestureRecognizer?.isEnabled = true
            navigationController?.interactivePopGestureRecognizer?.delegate = nil
        }
    }
}

/// The page's body, pure over its state.
struct ItemPageView: View {
    let state: ItemPageState
    let openSpec: (ItemSpecRow) -> Void
    let retry: () -> Void

    @State private var anchor: ItemPageAnchor
    @State private var position: ScrollPosition
    @State private var specsExpanded: Bool
    @Environment(\.accessibilityReduceMotion) private var reduceMotion

    /// Rows a skeleton holds, so the loaded page does not move (spec §11).
    private static let skeletonSources = 2
    private static let tooltipHeight: CGFloat = 260
    private static let pickHeight: CGFloat = 66

    init(
        state: ItemPageState, initialAnchor: ItemPageAnchor = .item, specsExpanded: Bool = false,
        openSpec: @escaping (ItemSpecRow) -> Void = { _ in }, retry: @escaping () -> Void = {}
    ) {
        self.state = state
        self.openSpec = openSpec
        self.retry = retry
        _anchor = State(initialValue: initialAnchor)
        _position = State(initialValue: initialAnchor == .item ? ScrollPosition(edge: .top) : ScrollPosition(id: initialAnchor, anchor: .top))
        _specsExpanded = State(initialValue: specsExpanded)
    }

    var body: some View {
        ScrollView {
            if state.isKnown {
                LazyVStack(alignment: .leading, spacing: 16, pinnedViews: [.sectionHeaders]) {
                    head
                    pickSlot
                    Section {
                        tooltip.id(ItemPageAnchor.item)
                        sourcesSection.id(ItemPageAnchor.source).padding(.top, 6)
                        specsSection.id(ItemPageAnchor.bisFor).padding(.top, 6)
                    } header: {
                        chips
                    }
                }
                .scrollTargetLayout()
                .padding(.horizontal, 16)
                .padding(.bottom, 32)
            } else {
                Text(Copy.itemPageUnknown)
                    .font(ScaledType.bodyStrong).foregroundStyle(FSColor.strong)
                    .fixedSize(horizontal: false, vertical: true)
                    .frame(maxWidth: .infinity, alignment: .leading)
                    .padding(.horizontal, 16)
                    .padding(.top, 16)
            }
        }
        .scrollPosition($position)
        .onScrollTargetVisibilityChange(idType: ItemPageAnchor.self, threshold: 0.3) { visible in
            if let first = ItemPageAnchor.allCases.first(where: visible.contains) { anchor = first }
        }
        .background(FSColor.bg)
    }

    // MARK: Head

    private var head: some View {
        HStack(alignment: .center, spacing: 14) {
            ItemIconView(icon: state.look.icon, quality: state.look.quality, build: state.build, size: 64)
            VStack(alignment: .leading, spacing: 4) {
                if let name = state.look.name {
                    Text(name)
                        .font(ScaledType.h2)
                        .foregroundStyle(ItemQuality(value: state.look.quality).text)
                        .fixedSize(horizontal: false, vertical: true)
                } else {
                    SkeletonBar(width: 200, height: 24)
                }
                if let meta = state.meta {
                    Text(meta).font(ScaledType.meta).foregroundStyle(FSColor.muted).fixedSize(horizontal: false, vertical: true)
                } else {
                    SkeletonBar(width: 140, height: 12)
                }
            }
        }
        .padding(.top, 8)
        .accessibilityElement(children: .combine)
        .accessibilityAddTraits(.isHeader)
    }

    @ViewBuilder
    private var pickSlot: some View {
        if let pick = state.pick {
            ItemPickCardView(card: pick)
        } else if state.reservesPick {
            SkeletonBar(width: nil, height: Self.pickHeight)
        }
    }

    // MARK: Strip

    private var chips: some View {
        ScrollView(.horizontal) {
            HStack(spacing: 6) {
                ForEach(ItemPageAnchor.allCases, id: \.self) { chip in
                    SearchChip(title: chip.title, isOn: chip == anchor) {
                        anchor = chip
                        if reduceMotion {
                            position.scrollTo(id: chip, anchor: .top)
                        } else {
                            withAnimation { position.scrollTo(id: chip, anchor: .top) }
                        }
                    }
                }
            }
        }
        .scrollIndicators(.hidden)
        .frame(maxWidth: .infinity, alignment: .leading)
        .background(FSColor.bg)
    }

    // MARK: Sections

    @ViewBuilder
    private var tooltip: some View {
        switch state.tooltip {
        case .loading:
            SkeletonBar(width: nil, height: Self.tooltipHeight)
        case let .loaded(model):
            ItemTooltipPanel(model: model)
        case .failed:
            failure(Copy.itemsFailed)
        }
    }

    private var sourcesSection: some View {
        VStack(alignment: .leading, spacing: 8) {
            MicroLabel(Copy.itemPageSources).accessibilityAddTraits(.isHeader)
            switch state.sources {
            case .loading:
                skeletonRows(Self.skeletonSources)
            case .failed:
                failure(Copy.itemPageUsesFailed)
            case let .loaded(rows) where rows.isEmpty:
                Text(Copy.itemPageNoSources).font(ScaledType.body).foregroundStyle(FSColor.muted)
            case let .loaded(rows):
                SearchCard {
                    ForEach(rows) { ItemSourceRowView(row: $0) }
                }
                if let footnote = ItemPageLogic.chanceFootnote(rows) {
                    Text(footnote).font(ScaledType.meta).foregroundStyle(FSColor.muted)
                }
            }
        }
    }

    private var specsSection: some View {
        VStack(alignment: .leading, spacing: 8) {
            MicroLabel(Copy.itemPageBisFor).accessibilityAddTraits(.isHeader)
            switch state.specs {
            case .loading:
                skeletonRows(ItemPageLogic.visibleSpecs)
            case .failed:
                failure(Copy.itemPageBisForFailed)
            case let .loaded(rows) where rows.isEmpty:
                Text(Copy.itemPageBisForNone).font(ScaledType.body).foregroundStyle(FSColor.muted)
            case let .loaded(rows):
                let shown = ItemPageLogic.visible(rows, expanded: specsExpanded)
                SearchCard {
                    ForEach(shown.rows) { row in
                        ItemSpecRowView(row: row, build: state.build) { openSpec(row) }
                    }
                    if shown.hidden > 0 {
                        moreSpecs(shown.hidden)
                    }
                }
            }
        }
    }

    private func skeletonRows(_ count: Int) -> some View {
        SearchCard {
            ForEach(0..<count, id: \.self) { _ in
                HStack(spacing: 12) {
                    SkeletonBar(width: SearchLayout.rowIcon, height: SearchLayout.rowIcon)
                    VStack(alignment: .leading, spacing: 8) {
                        SkeletonBar(width: 180, height: 15)
                        SkeletonBar(width: 120, height: 12)
                    }
                    Spacer(minLength: 0)
                }
                .padding(.horizontal, 14)
                .frame(minHeight: SearchLayout.rowMinHeight + 10)
            }
        }
        .accessibilityHidden(true)
    }

    private func moreSpecs(_ hidden: Int) -> some View {
        Button { specsExpanded = true } label: {
            HStack {
                Text(Copy.itemPageMoreSpecs(n: hidden)).font(ScaledType.bodyStrong).foregroundStyle(FSColor.gold)
                Spacer()
                Image(systemName: "chevron.down").font(.system(size: 13, weight: .semibold)).foregroundStyle(FSColor.muted)
                    .accessibilityHidden(true)
            }
            .padding(.horizontal, 14)
            .frame(maxWidth: .infinity, minHeight: SearchLayout.target)
            .contentShape(.rect)
        }
        .buttonStyle(RowPressStyle())
    }

    private func failure(_ text: String) -> some View {
        VStack(alignment: .leading, spacing: 10) {
            Text(text).font(ScaledType.body).foregroundStyle(FSColor.wipe)
            SecondaryButton(title: Copy.commonTryAgain, action: retry)
        }
    }
}

/// `Your Back pick` (spec F12.4): the primary card, the crest, the meta, the `Yours` tag.
struct ItemPickCardView: View {
    let card: YourPickCard

    var body: some View {
        // The tag sits at the trailing edge while the words fit on their lines, else under them.
        ViewThatFits(in: .horizontal) {
            HStack(spacing: 12) {
                crest
                words
                Spacer(minLength: 8)
                tag
            }
            HStack(alignment: .top, spacing: 12) {
                crest
                VStack(alignment: .leading, spacing: 6) { words; tag }
                Spacer(minLength: 0)
            }
        }
        .padding(14)
        .frame(maxWidth: .infinity, alignment: .leading)
        .primaryCard()
        .accessibilityElement(children: .combine)
    }

    private var crest: some View { ClassCrest(className: card.className, size: SearchLayout.rowIcon) }
    private var tag: some View { TagView(text: Copy.itemPageYours, color: FSColor.gold) }

    private var words: some View {
        VStack(alignment: .leading, spacing: 2) {
            Text(card.title).font(ScaledType.bodyStrong).foregroundStyle(FSColor.strong)
                .fixedSize(horizontal: false, vertical: true)
            Text(card.meta).font(ScaledType.meta).foregroundStyle(FSColor.muted)
                .fixedSize(horizontal: false, vertical: true)
        }
    }
}

/// A source's tile: one glyph per kind, the same on both apps (a boss is crossed swords, as the
/// board draws it; a vendor a bag; a quest the giver's `!`).
enum SourceGlyph {
    case symbol(String)
    case text(String)

    static func of(_ kind: String) -> SourceGlyph {
        switch kind {
        case "raid", "dungeon": .text("\u{2694}\u{FE0E}")
        case "vendor": .symbol("bag.fill")
        case "quest": .symbol("exclamationmark")
        case "crafted": .symbol("hammer.fill")
        case "rep": .symbol("flag.fill")
        case "pvp": .symbol("flag.2.crossed.fill")
        default: .symbol("globe")
        }
    }
}

/// One source: its kind's tile, the boss, vendor or quest, and `{place} · {kind}`.
struct ItemSourceRowView: View {
    let row: ItemSourceRow

    var body: some View {
        HStack(alignment: .top, spacing: 12) {
            glyph
                .foregroundStyle(FSColor.gold)
                .frame(width: SearchLayout.rowIcon, height: SearchLayout.rowIcon)
                .background(FSColor.cardTop, in: .rect(cornerRadius: FSRadius.pill))
                .overlay(RoundedRectangle(cornerRadius: FSRadius.pill).stroke(FSColor.lineSoft))
                .accessibilityHidden(true)
            VStack(alignment: .leading, spacing: 2) {
                SearchRowName(text: row.title)
                if !row.meta.isEmpty {
                    Text(row.meta).font(ScaledType.meta).foregroundStyle(FSColor.muted).fixedSize(horizontal: false, vertical: true)
                }
            }
            Spacer(minLength: 0)
        }
        .padding(.horizontal, 14)
        .padding(.vertical, 10)
        .frame(maxWidth: .infinity, minHeight: SearchLayout.rowMinHeight + 10, alignment: .leading)
        .accessibilityElement(children: .ignore)
        .accessibilityLabel(row.accessibilityLabel)
    }

    @ViewBuilder
    private var glyph: some View {
        switch SourceGlyph.of(row.kind) {
        case let .symbol(name): Image(systemName: name).font(.system(size: 14, weight: .semibold))
        case let .text(text): Text(text).font(.system(size: 18, weight: .semibold))
        }
    }
}

/// One `Best in slot for` row: spec icon, `{Spec} {Class}` in the class colour, the band, the tag, a chevron.
struct ItemSpecRowView: View {
    let row: ItemSpecRow
    let build: String
    let action: () -> Void

    var body: some View {
        Button(action: action) {
            // The tag stays beside the chevron while the words fit, else it moves under the meta.
            ViewThatFits(in: .horizontal) {
                HStack(spacing: 12) {
                    icon
                    words
                    Spacer(minLength: 8)
                    tag
                    chevron
                }
                HStack(spacing: 12) {
                    icon
                    VStack(alignment: .leading, spacing: 6) { words; tag }
                    Spacer(minLength: 8)
                    chevron
                }
            }
            .padding(.horizontal, 14)
            .padding(.vertical, 8)
            .frame(maxWidth: .infinity, minHeight: SearchLayout.rowMinHeight, alignment: .leading)
            .contentShape(.rect)
        }
        .buttonStyle(RowPressStyle())
        .accessibilityElement(children: .ignore)
        .accessibilityLabel(row.accessibilityLabel)
        .accessibilityAddTraits(.isButton)
    }

    @ViewBuilder
    private var icon: some View {
        if let icon = row.icon {
            SpecIcon(icon: icon, build: build, size: SearchLayout.rowIcon)
        } else {
            ClassCrest(className: row.use.className, size: SearchLayout.rowIcon)
        }
    }

    private var words: some View {
        VStack(alignment: .leading, spacing: 2) {
            SearchRowName(text: row.title, color: ClassColor.of(row.use.className))
            Text(row.meta).font(ScaledType.meta).foregroundStyle(FSColor.muted)
        }
    }

    private var tag: some View {
        TagView(text: row.tag.text, color: row.tag == .pick ? FSColor.kill : FSColor.muted)
    }

    private var chevron: some View {
        Image(systemName: "chevron.right").font(.system(size: 13, weight: .semibold)).foregroundStyle(FSColor.muted)
    }
}
