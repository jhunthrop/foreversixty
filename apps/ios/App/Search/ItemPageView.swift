import DesignSystem
import DesignTokens
import Models
import SwiftUI

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
/// sheet. Back only; no share (no web item page).
struct ItemPageScreen: View {
    let route: ItemPageRoute
    let dependencies: SearchDependencies
    /// A spec row's chevron: that spec's list in the BiS tab.
    let openList: (BisSelection) -> Void

    @State private var store: ItemPageStore

    init(route: ItemPageRoute, dependencies: SearchDependencies, openList: @escaping (BisSelection) -> Void) {
        self.route = route
        self.dependencies = dependencies
        self.openList = openList
        _store = State(initialValue: ItemPageStore(route: route, dependencies: dependencies))
    }

    var body: some View {
        ItemPageView(
            state: store.state,
            openSpec: { row in openList(ItemPageLogic.selection(row.use, viewerFaction: dependencies.context().character?.faction)) },
            retry: { Task { await store.load() } }
        )
        .navigationTitle("")
        .navigationBarTitleDisplayMode(.inline)
        .toolbarBackground(FSColor.bg, for: .navigationBar)
        .task { await store.load() }
    }
}

/// The page's body, pure over its state.
struct ItemPageView: View {
    let state: ItemPageState
    let openSpec: (ItemSpecRow) -> Void
    let retry: () -> Void

    @State private var anchor: ItemPageAnchor = .item
    @State private var specsExpanded: Bool

    init(state: ItemPageState, specsExpanded: Bool = false, openSpec: @escaping (ItemSpecRow) -> Void = { _ in }, retry: @escaping () -> Void = {}) {
        self.state = state
        self.openSpec = openSpec
        self.retry = retry
        _specsExpanded = State(initialValue: specsExpanded)
    }
    @Environment(\.accessibilityReduceMotion) private var reduceMotion

    var body: some View {
        ScrollViewReader { proxy in
            ScrollView {
                LazyVStack(alignment: .leading, spacing: 16, pinnedViews: state.isKnown ? [.sectionHeaders] : []) {
                    head
                    if !state.isKnown {
                        Text(Copy.itemPageUnknown).font(ScaledType.body).foregroundStyle(FSColor.muted)
                            .fixedSize(horizontal: false, vertical: true)
                    } else {
                        if let pick = state.pick { ItemPickCardView(card: pick) }
                        Section {
                            sections
                        } header: {
                            chips(proxy)
                        }
                    }
                }
                .padding(.horizontal, 16)
                .padding(.bottom, 32)
            }
            .background(FSColor.bg)
        }
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
                } else if state.isKnown {
                    SkeletonBar(width: 140, height: 12)
                }
            }
        }
        .padding(.top, 8)
        .accessibilityElement(children: .combine)
        .accessibilityAddTraits(.isHeader)
    }

    // MARK: Strip

    private func chips(_ proxy: ScrollViewProxy) -> some View {
        ScrollView(.horizontal) {
            HStack(spacing: 8) {
                ForEach(ItemPageAnchor.allCases, id: \.self) { chip in
                    SearchChip(title: chip.title, isOn: chip == anchor) {
                        anchor = chip
                        if reduceMotion {
                            proxy.scrollTo(chip, anchor: .top)
                        } else {
                            withAnimation { proxy.scrollTo(chip, anchor: .top) }
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
    private var sections: some View {
        tooltip.id(ItemPageAnchor.item)
        sourcesSection.id(ItemPageAnchor.source).padding(.top, 6)
        specsSection.id(ItemPageAnchor.bisFor).padding(.top, 6)
    }

    @ViewBuilder
    private var tooltip: some View {
        switch state.tooltip {
        case .loading:
            SkeletonBar(width: nil, height: 220)
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
                SkeletonBar(width: nil, height: 120)
            case .failed:
                failure(Copy.itemPageUsesFailed)
            case let .loaded(rows) where rows.isEmpty:
                Text(Copy.itemPageNoSources).font(ScaledType.body).foregroundStyle(FSColor.muted)
            case let .loaded(rows):
                SearchCard {
                    ForEach(rows) { ItemSourceRowView(row: $0) }
                }
                Text(Copy.itemPageSourcesFoot).font(ScaledType.meta).foregroundStyle(FSColor.muted)
                    .fixedSize(horizontal: false, vertical: true)
            }
        }
    }

    private var specsSection: some View {
        VStack(alignment: .leading, spacing: 8) {
            MicroLabel(Copy.itemPageBisFor).accessibilityAddTraits(.isHeader)
            switch state.specs {
            case .loading:
                SkeletonBar(width: nil, height: 120)
            case .failed:
                // The sources region already carries the error and its Try again.
                EmptyView()
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
                Text(Copy.itemPageBisForFoot).font(ScaledType.meta).foregroundStyle(FSColor.muted)
                    .fixedSize(horizontal: false, vertical: true)
            }
        }
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

/// One source: its kind's tile, the boss, vendor or quest, and `{place} · {kind} · {chance}`.
struct ItemSourceRowView: View {
    let row: ItemSourceRow

    var body: some View {
        HStack(alignment: .top, spacing: 12) {
            Image(systemName: Self.glyph(row.kind))
                .font(.system(size: 14, weight: .semibold))
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
        .frame(maxWidth: .infinity, minHeight: SearchLayout.rowMinHeight, alignment: .leading)
        .accessibilityElement(children: .ignore)
        .accessibilityLabel(row.accessibilityLabel)
    }

    static func glyph(_ kind: String) -> String {
        switch kind {
        case "raid", "dungeon": "shield.lefthalf.filled"
        case "vendor": "bag"
        case "quest": "exclamationmark.bubble"
        case "crafted": "hammer"
        case "rep": "flag"
        case "pvp": "flag.2.crossed"
        default: "globe"
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

/// A skeleton block at its real size.
struct SkeletonBar: View {
    let width: CGFloat?
    let height: CGFloat

    var body: some View {
        RoundedRectangle(cornerRadius: FSRadius.control).fill(FSColor.cardTop)
            .frame(maxWidth: width ?? .infinity, minHeight: height, maxHeight: height)
            .frame(width: width)
            .accessibilityHidden(true)
    }
}
