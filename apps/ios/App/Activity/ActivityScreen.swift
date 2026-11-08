import Auth
import DesignSystem
import Models
import SwiftUI

/// The Activity tab (flows spec 4.4, mocks ios-activity, -empty, -out): every notification event
/// newest first, by day, unread rows dotted, the group filters (All preselected) and Mark all read.
/// Each row opens its answer (the F4 table); the list is the push inbox.
struct ActivityScreen: View {
    let session: SessionStore?
    let build: String
    let onUnread: (Int) -> Void
    let open: (IntentDestination) -> Void
    @State private var model: ActivityModel
    @Environment(\.scenePhase) private var scenePhase

    init(
        service: any ActivityService, session: SessionStore?, build: String,
        onUnread: @escaping (Int) -> Void, open: @escaping (IntentDestination) -> Void
    ) {
        self.session = session
        self.build = build
        self.onUnread = onUnread
        self.open = open
        _model = State(initialValue: ActivityModel(service: service))
    }

    var body: some View {
        NavigationStack {
            Group {
                if session?.phase.isSignedIn == false {
                    ActivitySignedOutView(signingIn: session?.phase == .signingIn) {
                        Task { await session?.signIn(deviceName: UIDevice.current.name) }
                    }
                } else {
                    ActivityView(
                        phase: model.phase, sections: ActivitySections(model.rows, filter: model.filter, now: Date()),
                        filter: model.filter, build: build,
                        actions: ActivityActions(
                            select: { model.filter = $0 },
                            markAllRead: { Task { await model.markAllRead() } },
                            open: { row in
                                let link = URL(string: row.link)
                                PushLanding.destination(kind: row.kind, link: link, data: [:]).map(open)
                            },
                            more: { Task { await model.loadMore() } },
                            retry: { Task { await model.load() } }
                        )
                    )
                    .refreshable { await model.load() }
                }
            }
            .toolbarBackground(FSColor.bg, for: .navigationBar)
            .tabRootChrome()
        }
        .task {
            model.onUnread = onUnread
            await model.load()
        }
        .onChange(of: scenePhase) { _, phase in
            if phase == .active { Task { await model.load() } }
        }
    }
}

/// What the view hands back.
struct ActivityActions {
    var select: (ActivityFilter) -> Void = { _ in }
    var markAllRead: () -> Void = {}
    var open: (ActivityRow) -> Void = { _ in }
    var more: () -> Void = {}
    var retry: () -> Void = {}
}

struct ActivityView: View {
    let phase: ActivityModel.Phase
    let sections: ActivitySections
    let filter: ActivityFilter
    let build: String
    var actions = ActivityActions()

    /// `.ap section { margin-top: 28px }`.
    static let sectionGap: CGFloat = 28
    /// The unread dot (`.act-row.unread::before`): 6 pt, 10 pt left of the row.
    static let dotSize: CGFloat = 6
    static let dotInset: CGFloat = 10

    var body: some View {
        ScrollView {
            VStack(alignment: .leading, spacing: 0) {
                PageHeader(title: Copy.tabActivity)
                filters
                switch phase {
                case .loading:
                    ProgressView().frame(maxWidth: .infinity).padding(.top, FSSpace.space6)
                case .failed:
                    FSButton(Copy.commonTryAgain, action: actions.retry).padding(.top, FSSpace.space6)
                case .loaded where sections.sections.isEmpty:
                    SmallText(Copy.activityEmpty).padding(.top, FSSpace.space4)
                case .loaded:
                    ForEach(Array(sections.sections.enumerated()), id: \.element.id) { index, section in
                        daySection(section, first: index == 0)
                    }
                    Color.clear.frame(height: 1).onAppear(perform: actions.more)
                }
            }
            .padding(.top, FSSpace.space2)
            .padding(.horizontal, FSSpace.gutterPhone)
            .padding(.bottom, FSSpace.space6)
            .parityContent()
        }
        .background(FSColor.bg)
    }

    private var filters: some View {
        ScrollView(.horizontal, showsIndicators: false) {
            HStack(spacing: FSSpace.space2) {
                ForEach(ActivityFilter.allCases) { item in
                    Button(item.title) { actions.select(item) }
                        .buttonStyle(FSButtonStyle(item == filter ? .primary : .secondary, fullWidth: false, small: true))
                        .accessibilityAddTraits(item == filter ? .isSelected : [])
                }
            }
            .padding(.top, FSSpace.space1)
            .padding(.bottom, FSSpace.space2)
        }
        .parityRegion(.controls)
    }

    private func daySection(_ section: ActivitySections.Section, first: Bool) -> some View {
        FSSection(section.day.title) {
            if first, section.rows.contains(where: { !$0.row.read }) {
                Button(Copy.activityMarkAllRead, action: actions.markAllRead)
                    .font(.custom(FSFont.bodySemibold, size: FSType.small.size, relativeTo: .footnote))
                    .foregroundStyle(FSColor.strong)
            }
        } content: {
            VStack(spacing: 0) {
                ForEach(section.rows) { item in row(item) }
            }
        }
        .padding(.top, Self.sectionGap)
    }

    private func row(_ item: ActivitySections.Row) -> some View {
        Button { actions.open(item.row) } label: {
            FSRow(title: item.row.title, sub: item.row.body, detail: item.when, inset: 0) {
                ItemIconView(icon: item.row.icon, quality: nil, build: build, size: FSSize.icon40)
                    .parityItem(.image)
            }
            .overlay(alignment: .topLeading) {
                if !item.row.read {
                    Circle().fill(FSColor.gold)
                        .frame(width: Self.dotSize, height: Self.dotSize)
                        .offset(x: -Self.dotInset, y: 24)
                        .accessibilityHidden(true)
                }
            }
            .accessibilityValue(item.row.read ? "" : Copy.activityUnread)
        }
        .buttonStyle(RowPressStyle())
    }
}

/// Activity signed out (mock ios-activity-out): what arrives here, and Sign in.
struct ActivitySignedOutView: View {
    let signingIn: Bool
    let signIn: () -> Void

    var body: some View {
        ScrollView {
            VStack(alignment: .leading, spacing: 0) {
                PageHeader(title: Copy.tabActivity) { SmallText(Copy.activitySignedOut) }
                FSButton(Copy.signinButton, kind: .primary, action: signIn)
                    .disabled(signingIn)
                    .padding(.top, FSSpace.space5)
            }
            .padding(.top, FSSpace.space2)
            .padding(.horizontal, FSSpace.gutterPhone)
            .parityContent()
        }
        .background(FSColor.bg)
    }
}
