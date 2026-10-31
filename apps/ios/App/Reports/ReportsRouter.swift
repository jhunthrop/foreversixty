import Observation
import Routing

/// A report to push: its id.
struct ReportRoute: Hashable, Identifiable {
    let id: String
}

/// Where a report link lands on Home's stack (spec §3.4): Your reports, then the report over it.
@MainActor
@Observable
final class ReportsRouter {
    /// Your reports is pushed on Home.
    var showsList = false
    /// A report to push over Your reports (a `/app/report/{id}` link).
    var report: ReportRoute?

    static func handles(_ link: DeepLink) -> Bool {
        if case .report = link { return true }
        return false
    }

    /// Opens a report link on a fresh stack, so back goes to Your reports and then Home.
    func open(_ link: DeepLink) {
        guard case let .report(id) = link else { return }
        showsList = true
        report = ReportRoute(id: id)
    }
}
