import Foundation
import Observation
import Routing
import SwiftUI

/// A planner page a link opens over the tabs: the shared build viewer or Your builds.
enum PlannerPage: Hashable, Identifiable {
    case shared(String)
    case builds

    var id: String {
        switch self {
        case let .shared(id): "shared-\(id)"
        case .builds: "builds"
        }
    }
}

/// Where the planner opens (lane F13): the editor as a full-screen task over everything, and the
/// pages a link opens. The Talents screen, a guide, Your builds and the viewer all ask it, so
/// there is one editor presenter however it is reached. Everything is behind the `planner` switch.
@MainActor
@Observable
final class PlannerRouter {
    var editor: PlannerLaunch?
    var page: PlannerPage?
    let dependencies: PlannerDependencies
    @ObservationIgnored private let isEnabled: @MainActor () -> Bool

    init(dependencies: PlannerDependencies, isEnabled: @escaping @MainActor () -> Bool) {
        self.dependencies = dependencies
        self.isEnabled = isEnabled
    }

    /// The `planner` switch, read live (the app config can turn it on after launch).
    var isOn: Bool { isEnabled() }

    static func handles(_ link: DeepLink) -> Bool {
        switch link {
        case .planner, .builds, .sharedBuild: true
        default: false
        }
    }

    func openEditor(_ launch: PlannerLaunch) {
        guard isOn else { return }
        editor = launch
    }

    /// `/app/planner/{class}` (F11 rankings and the guides call it), `/app/builds`, `/b/{id}`.
    /// Returns false when the switch is off, so the caller can open `/b/` in the browser.
    @discardableResult
    func open(_ link: DeepLink) -> Bool {
        guard isOn else { return false }
        switch link {
        case let .planner(className, code, race):
            editor = PlannerLaunch(
                classSlug: className, raceSlug: race, source: code.map(PlannerLaunch.Source.code) ?? .draft, origin: .link
            )
        case .builds:
            editor = nil
            page = .builds
        case let .sharedBuild(id):
            editor = nil
            page = .shared(id)
        default:
            return false
        }
        return true
    }

    /// Edit a copy from the viewer: the editor replaces the page.
    func editCopy(_ summary: PlannerBuildSummary) {
        openEditor(PlannerLaunch(
            classSlug: summary.classSlug, raceSlug: summary.raceSlug, source: .copy(order: summary.build.pointOrder), origin: .link
        ))
    }
}
