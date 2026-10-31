import Observation
import Routing

/// Where a sim link lands on Home's stack (spec §6.4): Sims, then the result over it.
@MainActor
@Observable
final class SimsRouter {
    /// Sims is pushed on Home.
    var showsList = false
    /// A result to push over Sims (a `/app/sim/{id}` link or a `sim.finished` push).
    var result: SimResultRoute?

    static func handles(_ link: DeepLink) -> Bool {
        if case .sim = link { return true }
        return false
    }

    /// Opens a sim link on a fresh stack, so back goes to Sims and then Home.
    func open(_ link: DeepLink) {
        guard case let .sim(id) = link else { return }
        showsList = true
        result = SimResultRoute(id: id, row: nil)
    }
}
