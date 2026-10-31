import Foundation
import Models
import Networking
import Observation

/// The fight a meters sheet reads (spec F10.3): a boss pull, or a zone's folded trash.
struct MetersTarget: Hashable, Identifiable, Sendable {
    enum Kind: Hashable, Sendable {
        case boss
        /// A trash pull with no zone: its own fight's meters.
        case loneTrash
        /// A zone's trash folded by the server (`scope=trash`, n is the group's key).
        case trash(pulls: Int)
    }

    let reportID: String
    /// The fight index, or the trash group's key.
    let n: Int
    let kind: Kind
    /// The boss name, or the trash group's zone.
    let name: String
    let kill: Bool
    let durationMs: Int64
    let players: Int

    var id: String { "\(reportID)-\(scope ?? "fight")-\(n)" }

    var scope: String? {
        if case .trash = kind { return MeterRows.trashScope }
        return nil
    }

    var isTrash: Bool { kind != .boss }

    /// How many pulls the sheet covers.
    var pulls: Int {
        if case let .trash(pulls) = kind { return pulls }
        return 1
    }

    /// The sheet a Fights row opens; nil for the fight still in progress.
    static func of(_ row: FightRow, reportID: String) -> MetersTarget? {
        switch row {
        case let .boss(fight), let .loneTrash(fight):
            let isBoss = fight.kind == FightKind.encounter
            return MetersTarget(
                reportID: reportID, n: fight.index, kind: isBoss ? .boss : .loneTrash,
                name: isBoss ? fight.name : fight.zone ?? fight.name, kill: fight.kill,
                durationMs: fight.durationMs, players: fight.players.count
            )
        case let .trash(group):
            return MetersTarget(
                reportID: reportID, n: group.key, kind: .trash(pulls: group.fights.count), name: group.zone, kill: false,
                durationMs: group.durationMs, players: 0
            )
        case .inProgress:
            return nil
        }
    }
}

enum MetersPhase: Equatable, Sendable {
    case loading
    case loaded(FightMeters)
    case failed
    /// 404 `flagged`: the report cannot be shown.
    case flagged
}

/// Loads one fight's meters per segment, keeping each answer for the sheet's life.
@MainActor
@Observable
final class MetersStore {
    private(set) var metric: MeterMetric = .damage
    private(set) var phase: MetersPhase = .loading
    var showAll = false

    let target: MetersTarget
    private let service: any ReportsService
    private var answers: [MeterMetric: FightMeters] = [:]

    static let flaggedCode = "flagged"

    init(target: MetersTarget, service: any ReportsService) {
        self.target = target
        self.service = service
    }

    func select(_ metric: MeterMetric) async {
        guard metric != self.metric else { return }
        self.metric = metric
        showAll = false
        await load()
    }

    func load() async {
        let asked = metric
        if let answer = answers[asked] {
            phase = .loaded(answer)
            return
        }
        phase = .loading
        do {
            let answer = try await service.meters(id: target.reportID, n: target.n, metric: asked, scope: target.scope)
            answers[asked] = answer
            if metric == asked { phase = .loaded(answer) }
        } catch {
            guard metric == asked else { return }
            phase = error.code == Self.flaggedCode ? .flagged : .failed
        }
    }
}
