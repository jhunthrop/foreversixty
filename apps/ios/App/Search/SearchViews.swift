import DesignSystem
import DesignTokens
import Models
import SwiftUI

/// The glass field (spec F12.1): 44 pt, magnifier, placeholder, clear button, then gold `Cancel`.
struct SearchFieldBar: View {
    /// The placeholder keeps one line at the default size by scaling down this far at most.
    private static let placeholderScale = 0.7

    @Binding var query: String
    var focused: FocusState<Bool>.Binding
    let cancel: () -> Void

    @Environment(\.dynamicTypeSize) private var typeSize

    var body: some View {
        HStack(spacing: 12) {
            HStack(spacing: 8) {
                Image(systemName: "magnifyingglass").foregroundStyle(FSColor.muted).accessibilityHidden(true)
                // The placeholder is drawn under the field so it scales to one line rather than
                // being cut (spec §9). Vertical, so a long query wraps; Return still searches: a
                // typed line break is dropped and the keyboard closes.
                ZStack(alignment: .leading) {
                    if query.isEmpty {
                        // At accessibility sizes the long placeholder would be cut: `Search`.
                        Text(typeSize.isAccessibilitySize ? Copy.searchButton : Copy.searchPlaceholder)
                            .font(ScaledType.body).foregroundStyle(FSColor.muted)
                            .lineLimit(1)
                            .minimumScaleFactor(Self.placeholderScale)
                            .accessibilityHidden(true)
                    }
                    TextField(text: $query, prompt: Text(verbatim: ""), axis: .vertical) { Text(Copy.searchPlaceholder) }
                        .lineLimit(1...3)
                        .onChange(of: query) { _, text in
                            guard text.contains("\n") else { return }
                            query = text.replacingOccurrences(of: "\n", with: "")
                            focused.wrappedValue = false
                        }
                        .font(ScaledType.body)
                        .foregroundStyle(FSColor.strong)
                        .focused(focused)
                        .textInputAutocapitalization(.never)
                        .autocorrectionDisabled()
                        .submitLabel(.search)
                }
                if !query.isEmpty {
                    Button { query = "" } label: {
                        Image(systemName: "xmark.circle.fill").foregroundStyle(FSColor.muted)
                            .frame(width: SearchLayout.target, height: SearchLayout.target)
                    }
                    .accessibilityLabel(Copy.searchClear)
                }
            }
            .padding(.leading, 14)
            .padding(.trailing, query.isEmpty ? 14 : 0)
            .padding(.vertical, 4)
            .frame(minHeight: SearchLayout.fieldHeight)
            // The field's own surface, so it reads as a field wherever the glass material is
            // dropped (spec §7: glass, placeholder muted).
            .background { RingedCapsule(fill: FSColor.raised, ring: FSColor.lineSoft) }
            .glassEffect(.regular, in: .capsule)
            Button(Copy.searchCancel, action: cancel)
                .font(ScaledType.bodyStrong)
                .foregroundStyle(FSColor.gold)
                .frame(minHeight: SearchLayout.target)
        }
    }
}

/// F12's frame: the field and the chips on top, the scrolling body under them.
struct SearchChrome<Content: View>: View {
    @Binding var query: String
    @Binding var scope: SearchScope
    var focused: FocusState<Bool>.Binding
    let cancel: () -> Void
    @ViewBuilder let content: () -> Content

    var body: some View {
        VStack(spacing: 0) {
            VStack(spacing: 4) {
                SearchFieldBar(query: $query, focused: focused, cancel: cancel)
                SearchScopeChips(scope: $scope)
            }
            .padding(.horizontal, 16)
            .padding(.top, 8)
            ScrollView { content() }
                .scrollDismissesKeyboard(.interactively)
        }
        .background(FSColor.bg)
    }
}

/// The scope chips (spec F12.1): 36 pt pills, one on; each is a 44 pt target.
struct SearchScopeChips: View {
    @Binding var scope: SearchScope

    var body: some View {
        ScrollView(.horizontal) {
            HStack(spacing: SearchChip.spacing) {
                ForEach(SearchScope.allCases, id: \.self) { chip in
                    SearchChip(title: SearchWords.scope(chip), isOn: chip == scope) { scope = chip }
                }
            }
        }
        .scrollIndicators(.hidden)
    }
}

/// A 36 pt pill inside a 44 pt target (spec §7: 1 pt line; on, gold ring and text over the 10% wash).
struct SearchChip: View {
    let title: String
    let isOn: Bool
    let action: () -> Void

    /// Five chips fit across 402 pt at the default size (spec F12.1); smaller phones scroll.
    static let sidePadding: CGFloat = 12
    static let spacing: CGFloat = 6

    var body: some View {
        Button(action: action) {
            Text(title)
                .font(ScaledType.segment)
                .foregroundStyle(isOn ? FSColor.gold : FSColor.text)
                .lineLimit(1)
                .padding(.horizontal, Self.sidePadding)
                .padding(.vertical, 8)
                .frame(minHeight: SearchLayout.target)
                .background {
                    RingedCapsule(fill: isOn ? GameColors.pressedWash : GameColors.controlFill, ring: isOn ? FSColor.gold : FSColor.line)
                        .padding(.vertical, (SearchLayout.target - SearchLayout.chipHeight) / 2)
                }
                .contentShape(.capsule)
        }
        .buttonStyle(.plain)
        .accessibilityAddTraits(isOn ? .isSelected : [])
    }
}

/// A capsule with a 1 pt ring, drawn as the ring colour showing round an inset fill rather than
/// a stroke: a stroked capsule renders stray bars at its ends. `fill` may be translucent; the
/// page ground sits under it.
struct RingedCapsule: View {
    let fill: Color
    let ring: Color

    var body: some View {
        ZStack {
            Capsule().fill(ring)
            Capsule().fill(FSColor.bg).padding(1)
            Capsule().fill(fill).padding(1)
        }
    }
}

/// A guide by its class and spec slugs (`/app/guides/{class}/{spec}`).
struct SearchGuideLink: Equatable {
    let className: String
    let spec: String
}

/// The `For {name}` card's character: the BiS link, the guide link when the spec is known.
struct SearchForCharacter: Equatable {
    let name: String
    let guideTitle: String?
    let guideURL: URL?
    /// The guide's class and spec slugs, for the in-app guide (lane F15).
    let guideLink: SearchGuideLink?

    init?(_ context: SearchContext, specs: SpecCatalog) {
        guard let character = context.character else { return nil }
        name = character.name
        let spec = character.className.flatMap { slug in character.spec.flatMap { specs.spec(classSlug: slug, named: $0) } }
        if let spec, let specName = character.spec {
            guideTitle = Copy.searchForCharacterGuide(spec: specName, className: SearchWords.className(spec.classSlug))
            let slug = SearchLinks.guideSpec(specKey: spec.spec, className: spec.classSlug)
            guideURL = SearchLinks.guide(className: spec.classSlug, spec: slug)
            guideLink = SearchGuideLink(className: spec.classSlug, spec: slug)
        } else {
            guideTitle = nil
            guideURL = nil
            guideLink = nil
        }
    }
}

/// What a tap on F12 asks the screen to do.
struct SearchActions {
    var openItem: (SearchItem) -> Void = { _ in }
    var openCharacter: (CharacterSearchRow) -> Void = { _ in }
    var openGuide: (SearchGuideHit) -> Void = { _ in }
    var openTalent: (SearchTalent) -> Void = { _ in }
    var openRecent: (RecentSearch) -> Void = { _ in }
    var showAll: (SearchGroupKind) -> Void = { _ in }
    var clearRecent: () -> Void = {}
    var retryIndex: () -> Void = {}
    var retryCharacters: () -> Void = {}
    var openMyBis: () -> Void = {}
    var openURL: (URL) -> Void = { _ in }
    /// Lane F15: a guide by its class and spec slugs, F9 in the app when `guides` is on.
    var openGuideLink: (String, String) -> Void = { _, _ in }
    /// Lane F15: the `Rankings` link, F11 in the app when its switch is on.
    var openRankings: () -> Void = {}
}

/// F12's body under the field and chips, pure over its inputs so the snapshot tests draw it.
struct SearchContentView: View {
    let hasQuery: Bool
    let query: String
    let scope: SearchScope
    let index: SearchIndexPhase
    let results: SearchResults
    let characters: CharacterSearchPhase
    let recent: [RecentSearch]
    let forCharacter: SearchForCharacter?
    let viewer: SearchViewer
    let build: String
    var actions = SearchActions()
    /// Lane F15: the characters you watch, between `For {name}` and `Recent`.
    var watching: [CharacterSearchRow] = []
    /// Lane F15: characters and Rankings open in the app (a chevron, not the external glyph).
    var nativeRankings = false
    /// Lane F15: guides open in the app (F9) rather than the website.
    var nativeGuides = false

    private var guideTrailing: SearchRowTrailing { nativeGuides ? .chevron : SearchDestination.guide.trailing }

    private var characterTrailing: SearchRowTrailing { nativeRankings ? .chevron : SearchDestination.character.trailing }

    var body: some View {
        LazyVStack(alignment: .leading, spacing: SearchLayout.sectionGap) {
            if hasQuery { resultsBody } else { emptyBody }
        }
        .padding(.horizontal, 16)
        .padding(.top, 8)
        .padding(.bottom, 32)
    }

    // MARK: F12.1 Empty field

    @ViewBuilder
    private var emptyBody: some View {
        if let forCharacter {
            SearchSection(title: Copy.searchGroupForCharacter(name: forCharacter.name)) {
                SearchCard {
                    SearchRows.link(Copy.searchForCharacterBis(name: forCharacter.name), systemImage: "shield", trailing: SearchDestination.bis.trailing, action: actions.openMyBis)
                    if let title = forCharacter.guideTitle, let link = forCharacter.guideLink {
                        SearchRows.link(title, systemImage: "book.closed", trailing: guideTrailing) { actions.openGuideLink(link.className, link.spec) }
                    }
                    SearchRows.link(Copy.searchForCharacterRankings, systemImage: "trophy", trailing: characterTrailing, action: actions.openRankings)
                }
            }
        }
        // A group with nothing in it is not drawn.
        if !watching.isEmpty {
            SearchSection(title: Copy.searchGroupWatching) {
                SearchCard {
                    ForEach(watching, id: \.key) { row in
                        SearchRows.character(row, isYours: false, trailing: characterTrailing) { actions.openCharacter(row) }
                    }
                }
            }
        }
        if !recent.isEmpty {
            SearchSection(title: Copy.searchGroupRecent, trailing: (Copy.searchRecentClear, actions.clearRecent)) {
                SearchCard {
                    ForEach(recent) { entry in
                        SearchRows.recent(entry, build: build) { actions.openRecent(entry) }
                    }
                }
            }
        }
    }

    // MARK: F12.2 Results

    @ViewBuilder
    private var resultsBody: some View {
        indexStatus
        group(.guides, results.guides) { hit in SearchRows.guide(hit, build: build, trailing: guideTrailing) { actions.openGuide(hit) } }
        group(.talents, results.talents) { talent in SearchRows.talent(talent, build: build) { actions.openTalent(talent) } }
        group(.items, results.items) { item in SearchRows.item(item, build: build) { actions.openItem(item) } }
        group(.characters, results.characters) { row in
            SearchRows.character(row, isYours: viewer.ownKeys.contains(row.key), trailing: characterTrailing) { actions.openCharacter(row) }
        }
        charactersStatus
        if showsNoResults { noResults }
        if characters == .offline { SearchOfflineNote() }
    }

    private var showsNoResults: Bool {
        guard results.isEmpty, index.corpus != nil else { return false }
        return characters != .loading
    }

    @ViewBuilder
    private var indexStatus: some View {
        switch index {
        case .loading:
            Text(Copy.searchIndexLoading).font(ScaledType.meta).foregroundStyle(FSColor.muted)
        case .failed:
            VStack(alignment: .leading, spacing: 10) {
                Text(Copy.searchIndexFailed).font(ScaledType.body).foregroundStyle(FSColor.wipe)
                SecondaryButton(title: Copy.searchTryAgain, action: actions.retryIndex)
            }
        case .ready:
            EmptyView()
        }
    }

    @ViewBuilder
    private var charactersStatus: some View {
        if scope.shows(.characters), characters == .failed {
            VStack(alignment: .leading, spacing: 10) {
                MicroLabel(Copy.searchGroupCharacters)
                Text(Copy.searchCharactersFailed).font(ScaledType.body).foregroundStyle(FSColor.wipe)
                SecondaryButton(title: Copy.searchTryAgain, action: actions.retryCharacters)
            }
        }
    }

    @ViewBuilder
    private func group<Row: Sendable, RowView: View>(
        _ kind: SearchGroupKind, _ group: SearchGroup<Row>, @ViewBuilder row: @escaping (Row) -> RowView
    ) -> some View {
        if !group.rows.isEmpty {
            SearchSection(title: SearchWords.group(kind)) {
                SearchCard {
                    ForEach(Array(group.rows.enumerated()), id: \.offset) { _, value in row(value) }
                    if scope == .all, group.hasMore {
                        SearchRows.showAll(group.total) { actions.showAll(kind) }
                    }
                }
            }
        }
    }

    // MARK: F12.6 No results

    private var noResults: some View {
        VStack(alignment: .leading, spacing: 6) {
            Text(Copy.searchNoResultsTitle(query: query.trimmingCharacters(in: .whitespacesAndNewlines)))
                .font(ScaledType.bodyStrong).foregroundStyle(FSColor.strong)
                .fixedSize(horizontal: false, vertical: true)
            Text(Copy.searchNoResultsBody).font(ScaledType.meta).foregroundStyle(FSColor.muted)
                .fixedSize(horizontal: false, vertical: true)
        }
        .padding(14)
        .frame(maxWidth: .infinity, alignment: .leading)
        .quietCard()
        .accessibilityElement(children: .combine)
    }
}

/// The offline note inside the results (spec F12.6): characters need a connection.
struct SearchOfflineNote: View {
    var body: some View {
        HStack(alignment: .top, spacing: 10) {
            Image(systemName: "wifi.slash").font(.system(size: 14)).accessibilityHidden(true)
            Text(Copy.searchOffline).font(ScaledType.meta).fixedSize(horizontal: false, vertical: true)
            Spacer(minLength: 0)
        }
        .foregroundStyle(FSColor.nav)
        .padding(12)
        .background(GameColors.offlineGround, in: .rect(cornerRadius: FSRadius.control))
        .overlay(RoundedRectangle(cornerRadius: FSRadius.control).stroke(FSColor.line))
        .accessibilityElement(children: .combine)
        .onAppear { AccessibilityNotification.Announcement(Copy.searchOffline).post() }
    }
}
