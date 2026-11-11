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
                        filter: model.filter, build: build, older: model.hasOlder ? (model.loadingMore ? .loading : .ready) : nil,
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
                    .bandBar()
                }
            }
            // Round 5 on iOS (m-ios-activity): the tab's large title on the band, the chips under it.
            .navigationTitle(AppTab.activity.title)
            .toolbarTitleDisplayMode(.large)
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
    /// Show older at the list's end (i1): nil with no older page.
    var older: OlderState?
    var actions = ActivityActions()

    enum OlderState: Equatable { case ready, loading }

    var body: some View {
        ScrollView {
            BandPage {
                // Nothing to filter with no rows at all: no chips under the bar.
                if !isEmpty { ActivityFilterChips(filter: filter, select: actions.select) }
            } content: {
                switch phase {
                case .loading:
                    ProgressView().frame(maxWidth: .infinity).padding(.top, FSSpace.space6)
                case .failed:
                    PanelBlock {
                        Button(action: actions.retry) { PanelRow(Copy.commonTryAgain, last: true) }.buttonStyle(RowPressStyle())
                    }
                case .loaded where sections.sections.isEmpty:
                    ActivityNote(text: Copy.activityEmpty).parityRegion(ParityRegionName("section-empty"))
                case .loaded:
                    ForEach(Array(sections.sections.enumerated()), id: \.element.id) { index, section in
                        daySection(section, first: index == 0)
                    }
                    if let older {
                        // The same words as Android, left-aligned under the list as the web's pager link
                        // (round 4, row 7), a 44 pt target; disabled while the page loads.
                        Button(action: actions.more) {
                            Text(Copy.activityMore).type(FSType.bodyStrong).foregroundStyle(FSColor.gold)
                                .frame(minHeight: FSTarget.min)
                                .contentShape(.rect)
                        }
                        .buttonStyle(RowPressStyle())
                        .disabled(older == .loading)
                        .parityItem(.button)
                        .frame(maxWidth: .infinity, alignment: .leading)
                    }
                }
            }
        }
        .tabRootScroll()
        .screenGround()
    }

    private var isEmpty: Bool {
        if case .loaded = phase { return sections.sections.isEmpty && filter == .all }
        return false
    }

    private func daySection(_ section: ActivitySections.Section, first: Bool) -> some View {
        let markAll = first && section.rows.contains { !$0.row.read }
        return PanelSection(section.day.title, action: markAll ? (Copy.activityMarkAllRead, actions.markAllRead) : nil) {
            ForEach(Array(section.rows.enumerated()), id: \.element.id) { index, item in
                Button { actions.open(item.row) } label: {
                    ActivityRowView(row: item.row, line: item.row.body, meta: item.when, build: build, last: index == section.rows.count - 1)
                }
                .buttonStyle(RowPressStyle())
            }
        }
    }
}

/// Activity's group chips on the band (round 5): 36 pt faces in 44 pt targets, the selected one in
/// `teal` with a `tealEdge` hairline as the Segmented (round 4, row 9), scrolling sideways under a
/// trailing fade.
struct ActivityFilterChips: View {
    let filter: ActivityFilter
    let select: (ActivityFilter) -> Void

    static let face: CGFloat = 36
    static let radius: CGFloat = 8
    /// How far the row fades out at its trailing edge, so a cut chip reads as more to scroll.
    static let fade: CGFloat = FSSpace.space6

    var body: some View {
        ScrollView(.horizontal, showsIndicators: false) {
            HStack(spacing: FSSpace.space2) {
                ForEach(ActivityFilter.allCases) { item in
                    Button(item.title) { select(item) }
                        .buttonStyle(ChipStyle(on: item == filter))
                        .accessibilityAddTraits(item == filter ? .isSelected : [])
                }
            }
            .padding(.trailing, FSSpace.space4)
            .padding(.top, 2)
            .padding(.bottom, FSSpace.space2)
        }
        .mask {
            HStack(spacing: 0) {
                Color.black
                LinearGradient(colors: [.black, .clear], startPoint: .leading, endPoint: .trailing).frame(width: Self.fade)
            }
        }
        .parityRegion(.controls)
        .padding(.leading, FSSpace.space4)
    }

    private struct ChipStyle: ButtonStyle {
        let on: Bool

        func makeBody(configuration: Configuration) -> some View {
            configuration.label
                .type(FSType.chip)
                .foregroundStyle(FSColor.strong)
                .padding(.horizontal, FSSpace.space4)
                .frame(height: ActivityFilterChips.face)
                .background(on ? FSColor.teal : .clear, in: .rect(cornerRadius: ActivityFilterChips.radius))
                .overlay { if configuration.isPressed { TealPress.veil.clipShape(.rect(cornerRadius: ActivityFilterChips.radius)) } }
                .overlay(RoundedRectangle(cornerRadius: ActivityFilterChips.radius).strokeBorder(on ? FSColor.tealEdge : FSColor.rule, lineWidth: FSStroke.hairline))
                // The 36 pt face in its 44 pt target.
                .padding(.vertical, (FSTarget.min - ActivityFilterChips.face) / 2)
                .contentShape(.rect)
                .parityItem(.button)
        }
    }
}

/// One line on a panel (Activity empty, signed out).
struct ActivityNote: View {
    let text: String

    var body: some View {
        PanelBlock {
            Text(text).type(FSType.small).foregroundStyle(FSColor.text).padding(.vertical, FSSpace.space2)
        }
    }
}

/// Activity signed out (mock ios-activity-out): what arrives here, and Sign in.
struct ActivitySignedOutView: View {
    let signingIn: Bool
    let signIn: () -> Void

    var body: some View {
        ScrollView {
            BandPage {
                EmptyView()
            } content: {
                PanelBlock {
                    Text(Copy.activityOut).type(FSType.small).foregroundStyle(FSColor.text).padding(.vertical, FSSpace.space2)
                    FSButton(Copy.signinButton, kind: .primary, action: signIn)
                        .disabled(signingIn)
                        .padding(.vertical, FSSpace.space3)
                }
                .parityRegion(ParityRegionName("section-activity"))
            }
        }
        .screenGround()
    }
}
