import SwiftUI

/// Presents the editor and the link pages. Applied at the root and again inside a page, so the
/// editor opens over whichever is on top (Edit a copy from a linked viewer).
struct PlannerHost: ViewModifier {
    @Bindable var router: PlannerRouter
    /// The `Plan a new build` row's character, for Your builds.
    let newBuild: () -> PlannerNewBuild?
    var insidePage = false

    func body(content: Content) -> some View {
        content
            .fullScreenCover(item: editor) { launch in
                PlannerEditorScreen(launch: launch, dependencies: router.dependencies)
            }
            .fullScreenCover(item: page) { page in
                NavigationStack {
                    pageView(page)
                        .toolbar {
                            ToolbarItem(placement: .topBarLeading) {
                                Button { router.page = nil } label: {
                                    Label(Copy.tabHome, systemImage: "chevron.left").labelStyle(.titleAndIcon)
                                }
                            }
                        }
                }
                .modifier(PlannerHost(router: router, newBuild: newBuild, insidePage: true))
            }
    }

    /// The root presents the editor only while no page is up; a page presents it over itself.
    private var editor: Binding<PlannerLaunch?> {
        Binding(
            get: { (router.page != nil) == insidePage ? router.editor : nil },
            set: { router.editor = $0 }
        )
    }

    private var page: Binding<PlannerPage?> {
        Binding(get: { insidePage ? nil : router.page }, set: { router.page = $0 })
    }

    @ViewBuilder
    private func pageView(_ page: PlannerPage) -> some View {
        switch page {
        case let .shared(id): SharedBuildScreen(id: id, dependencies: router.dependencies, router: router)
        case .builds: YourBuildsScreen(dependencies: router.dependencies, router: router, newBuild: newBuild())
        }
    }
}

extension PlannerNewBuild {
    /// The current character's class and level for `Plan a new build`; nil without a class.
    init?(className slug: String?, race: String?, level: Int?) {
        guard let slug, !slug.isEmpty else { return nil }
        self.init(
            classSlug: slug, className: PlannerCopy.className(slug, names: .empty),
            raceSlug: race.map(TalentLinks.raceSlug), level: level ?? PlannerRules.maxLevel
        )
    }
}
