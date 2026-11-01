import Foundation
import Models

/// What the stepper strip's second line says (spec F13.1, F13.2): the talent's rank text, a
/// refusal in `--color-wipe`, or the note after a guide build loads.
enum PlannerLine: Equatable, Sendable {
    case refusal(String)
    case note(String)
}

/// What a tap, a long press or a strip button did, so the screen plays the right haptic.
enum PlannerEvent: Equatable, Sendable {
    case added, removed, refused
}

/// The editor's whole state (spec F13.1), a value: the class's index, the point order, the level
/// and race, the undo history, the talent the strip shows and its line. Every change returns the
/// event it caused; nothing here touches the screen, the disk or the network.
struct PlannerState: Equatable, Sendable {
    /// Undo steps back through the last 100 changes (spec F13.1 Interactions).
    static let undoLimit = 100

    let index: PlannerIndex
    private(set) var order: [Int]
    private(set) var level: Int
    private(set) var raceSlug: String?
    /// The orders before each change, oldest first, at most `undoLimit`.
    private(set) var history: [[Int]] = []
    /// The talent the strip shows: the last one touched, else the first talent.
    private(set) var focusID: Int?
    private(set) var line: PlannerLine?

    init(index: PlannerIndex, order: [Int] = [], level: Int = PlannerRules.maxLevel, raceSlug: String? = nil, note: String? = nil) {
        self.index = index
        self.order = order
        self.level = min(PlannerRules.maxLevel, max(level, PlannerRules.minimumLevel(spent: order.count)))
        self.raceSlug = raceSlug
        line = note.map(PlannerLine.note)
    }

    // MARK: Derived

    var ranks: [Int: Int] { PlannerRules.ranks(order) }
    var spent: Int { order.count }
    var left: Int { max(0, PlannerRules.pointsAvailable(atLevel: level) - spent) }
    /// Points per tree, summed from the drawn build (never typed).
    var totals: [Int] { PlannerRules.pointsPerTree(index, order: order) }
    var canUndo: Bool { !history.isEmpty }

    /// The strip's talent: the focus, else the first talent of the first tree.
    var focus: Talent? { focusID.flatMap { index.byID[$0] } ?? index.trees.first?.talents.first }

    func rank(of id: Int) -> Int { order.count { $0 == id } }

    // MARK: Changes

    /// A tap or `+`: one point, or the refusal that changes nothing but the line.
    mutating func add(_ id: Int) -> PlannerEvent {
        focusID = id
        if let reason = PlannerRules.add(index, order: order, talent: id, level: level).reason {
            line = .refusal(reason)
            return .refused
        }
        record()
        order.append(id)
        line = nil
        return .added
    }

    /// A long press or `−`: the talent's last point, or the refusal.
    mutating func remove(_ id: Int) -> PlannerEvent {
        focusID = id
        if let reason = PlannerRules.canRemovePoint(index, order: order, talent: id).reason {
            line = .refusal(reason)
            return .refused
        }
        record()
        order = PlannerRules.withoutLastPoint(order, id)
        line = nil
        return .removed
    }

    /// Reset: one undoable step back to no points.
    mutating func reset() {
        guard !order.isEmpty else { return }
        record()
        order = []
        line = nil
    }

    /// Steps back one change; a no-op with nothing to undo.
    mutating func undo() {
        guard let previous = history.popLast() else { return }
        order = previous
        line = nil
    }

    /// The Level sheet: 10 to 60, never below what the build already spends.
    mutating func setLevel(_ level: Int) {
        self.level = min(PlannerRules.maxLevel, max(level, PlannerRules.minimumLevel(spent: spent)))
    }

    mutating func setRace(_ slug: String) { raceSlug = slug }

    /// The keyboard's arrows: the next or previous talent of the class, tree by tree.
    mutating func moveFocus(by step: Int) {
        let all = index.allTalents
        guard !all.isEmpty else { return }
        let current = focus.flatMap { focus in all.firstIndex { $0.id == focus.id } } ?? 0
        focusID = all[(current + step + all.count) % all.count].id
    }

    mutating func focus(_ id: Int) { focusID = id }

    /// A share refused with a field's sentence (400): it goes under the editor.
    mutating func show(refusal: String) { line = .refusal(refusal) }

    private mutating func record() {
        history.append(order)
        if history.count > Self.undoLimit { history.removeFirst(history.count - Self.undoLimit) }
    }
}

/// The editor's draft on the phone, per class (spec F13.1 Draft).
struct PlannerDraft: Codable, Equatable, Sendable {
    let treeVersion: String
    let raceSlug: String?
    let level: Int
    let order: [Int]
}

/// Where drafts live. A protocol so the model's tests keep them in memory.
protocol PlannerDraftStore: Sendable {
    func load(classSlug: String) -> PlannerDraft?
    func save(_ draft: PlannerDraft, classSlug: String)
}

/// UserDefaults survives a force quit; one small JSON value per class.
struct UserDefaultsPlannerDraftStore: PlannerDraftStore, @unchecked Sendable {
    private static let prefix = "gg.foreversixty.planner.draft."
    let defaults: UserDefaults

    init(defaults: UserDefaults = .standard) { self.defaults = defaults }

    func load(classSlug: String) -> PlannerDraft? {
        defaults.data(forKey: Self.prefix + classSlug).flatMap { try? JSONDecoder().decode(PlannerDraft.self, from: $0) }
    }

    func save(_ draft: PlannerDraft, classSlug: String) {
        guard let data = try? JSONEncoder().encode(draft) else { return }
        defaults.set(data, forKey: Self.prefix + classSlug)
    }
}

/// How the editor opens: the class, what to load, and where it was opened from.
struct PlannerLaunch: Identifiable, Equatable, Sendable {
    enum Source: Equatable, Sendable {
        /// The class's draft, else an empty build.
        case draft
        /// An FS1 code (a guide's build, a rankings link) loaded as one legal order.
        case code(String)
        /// A shared build's points, as a copy (Edit a copy).
        case copy(order: [Int])
    }

    enum Origin: Equatable, Sendable {
        case talents, guide, link
    }

    let id = UUID()
    let classSlug: String
    var raceSlug: String?
    var level: Int?
    var source: Source = .draft
    var origin: Origin = .talents

    static func == (lhs: PlannerLaunch, rhs: PlannerLaunch) -> Bool { lhs.id == rhs.id }
}

extension PlannerState {
    /// The editor's opening state for a launch (spec F13.1, F13.5 Edit a copy, the guide row):
    /// a code or a copy is loaded as one legal order on the active trees, else the draft, else
    /// an empty build. A copy or a draft that no longer validates on these trees is rebuilt from
    /// its ranks, so a changed tree never leaves an illegal order in the editor.
    static func opening(
        _ launch: PlannerLaunch, index: PlannerIndex, draft: PlannerDraft?, legalRaces: [String]
    ) -> PlannerState {
        let defaultLevel = launch.level ?? PlannerRules.maxLevel
        switch launch.source {
        case let .code(code):
            if let decoded = FS1Talents(code: code), decoded.classSlug == index.file.classSlug {
                let order = PlannerRules.orderFromRanks(index, treeRanks: decoded.treeRanks).order
                let race = pick([launch.raceSlug, decoded.raceSlug, draft?.raceSlug], legal: legalRaces)
                return PlannerState(index: index, order: order, level: defaultLevel, raceSlug: race, note: loadedNote(spent: order.count))
            }
        case let .copy(order):
            let legal = legalOrder(order, index: index)
            return PlannerState(index: index, order: legal, level: defaultLevel, raceSlug: pick([launch.raceSlug, draft?.raceSlug], legal: legalRaces))
        case .draft:
            break
        }
        guard let draft else {
            return PlannerState(index: index, level: defaultLevel, raceSlug: pick([launch.raceSlug], legal: legalRaces))
        }
        return PlannerState(
            index: index, order: legalOrder(draft.order, index: index), level: draft.level,
            raceSlug: pick([draft.raceSlug, launch.raceSlug], legal: legalRaces)
        )
    }

    /// The order itself when it validates, else one legal order reaching its ranks.
    static func legalOrder(_ order: [Int], index: PlannerIndex) -> [Int] {
        guard !PlannerRules.validate(index, order: order).isEmpty else { return order }
        let ranks = PlannerRules.ranks(order)
        return PlannerRules.orderFromRanks(index, treeRanks: PlannerRules.treeRanks(index, ranks: ranks)).order
    }

    /// The first candidate the class may be, else the first legal race; nil with no race data.
    static func pick(_ candidates: [String?], legal: [String]) -> String? {
        candidates.compactMap(\.self).first { legal.contains($0) } ?? legal.first
    }

    /// The web's band note: the band the loaded points reach (`bandRangeLabel(bandForLevel(...))`).
    static func loadedNote(spent: Int) -> String? {
        TalentGuide.band(forLevel: spent + TalentRules.baseLevel).map { Copy.planLoadedNote(band: TalentGuide.label(band: $0)) }
    }

    /// F6's rank sheet for one talent of the drawn build: its current and next rank text and the
    /// requirements still unmet (the strip's text is clamped; the sheet has it all).
    func sheet(talentID: Int) -> TalentSheetModel? {
        guard let talent = index.byID[talentID], let treeIndex = index.treeOf[talentID] else { return nil }
        let ranks = ranks
        let rank = ranks[talentID] ?? 0
        let text = { (position: Int) in talent.ranks.indices.contains(position) ? talent.ranks[position].description : nil }
        return TalentSheetModel(
            talent: talent, rank: rank,
            state: PlannerRules.cellState(talent, index: index, order: order, ranks: ranks, left: left),
            current: rank > 0 ? text(rank - 1) : nil,
            next: rank < talent.maxRank ? text(rank) : nil,
            requirements: TalentsPresenter.requirements(talent, in: index.trees[treeIndex], ranks: ranks),
            guideLine: nil, isOver: false
        )
    }

    /// The draft this state saves.
    var draft: PlannerDraft {
        PlannerDraft(treeVersion: index.file.build, raceSlug: raceSlug, level: level, order: order)
    }
}
