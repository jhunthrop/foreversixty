import DesignSystem
import Models
import SwiftUI

/// The words and marks one row draws, worked out once for both sections.
struct SimRowFace {
    let row: SimRow
    let state: ServerRunState
    let title: String
    let specName: String
    let kind: String
    let icon: String?

    init(_ row: SimRow, specs: SpecCatalog, now: Date) {
        self.row = row
        state = ServerRuns.state(row, now: now)
        let spec = specs.spec(row.spec ?? "")
        specName = spec?.name ?? row.spec ?? ""
        icon = spec?.icon
        kind = SimResultModel.kindLabel(row.kind ?? "run")
        let named = row.title ?? ""
        title = named.isEmpty ? specName : named
    }

    /// `{Spec} · {Kind}` with a title, `{Kind}` without (the site's row meta).
    var meta: String { (row.title ?? "").isEmpty ? kind : "\(specName) · \(kind)" }

    var headline: String { row.headline ?? "" }

    var isUnseen: Bool { state == .doneUnseen }
}

/// The 36 or 32 pt spec icon, or the tile alone while the spec is unknown.
struct SimSpecMark: View {
    let icon: String?
    let build: String?
    let size: CGFloat

    var body: some View {
        if let icon, let build {
            SpecIcon(icon: icon, build: build, size: size)
        } else {
            RoundedRectangle(cornerRadius: FSRadius.pill).fill(FSColor.cardTop).frame(width: size, height: size)
        }
    }
}

/// The row title with the unseen gold dot before it.
struct SimRowTitle: View {
    let face: SimRowFace

    var body: some View {
        HStack(spacing: 8) {
            if face.isUnseen {
                Circle().fill(FSColor.gold).frame(width: 8, height: 8).accessibilityHidden(true)
            }
            Text(face.title).font(ScaledType.bodyStrong).foregroundStyle(FSColor.strong)
                .lineLimit(SimsLayout.titleLines)
        }
    }
}

/// The left half of the state line: Starting, Running with its estimate, the headline, or Failed.
struct SimStateText: View {
    let face: SimRowFace
    let now: Date

    var body: some View {
        // Wraps to its full height at every text size; never cut by a tight row.
        text.fixedSize(horizontal: false, vertical: true)
    }

    @ViewBuilder
    private var text: some View {
        switch face.state {
        case .queued:
            HStack(spacing: 6) {
                Circle().fill(FSColor.gold).frame(width: 8, height: 8).accessibilityHidden(true)
                Text(Copy.simsQueued).foregroundStyle(FSColor.strong)
            }
            .font(ScaledType.meta)
        case .running, .runningOverdue, .runningNoEstimate:
            let running = Text(Copy.simsRunning).foregroundStyle(FSColor.gold)
            let suffix = Text(SimsCopy.stateSuffix(face.row, state: face.state, now: now)).foregroundStyle(FSColor.muted)
            Text("\(running)\(suffix)").font(ScaledType.meta)
        case .done, .doneUnseen:
            Text(face.headline).font(SimsLayout.isFigure(face.headline) ? ScaledType.mono : ScaledType.bodyStrong)
                .foregroundStyle(FSColor.strong)
        case .failed:
            Text(SimsCopy.failedLine(face.row)).font(ScaledType.meta).foregroundStyle(FSColor.wipe)
        }
    }
}

/// The right half: Started, Finished or Ended, with its age.
enum SimAge {
    static func line(_ face: SimRowFace, now: Date) -> String {
        switch face.state {
        case .done, .doneUnseen: SimsCopy.finishedLine(face.row, now: now)
        case .failed: Copy.simsEnded(age: SimsCopy.age(face.row.finishedAt ?? face.row.createdAt, now: now))
        default: Copy.simsStarted(age: SimsCopy.age(face.row.createdAt, now: now))
        }
    }
}

/// State left and age right; at the accessibility sizes the age stacks under the state.
struct SimStateLine: View {
    let face: SimRowFace
    let now: Date
    @Environment(\.dynamicTypeSize) private var typeSize

    var body: some View {
        let age = Text(SimAge.line(face, now: now)).font(ScaledType.monoMeta).foregroundStyle(FSColor.muted)
            .fixedSize(horizontal: false, vertical: true)
        if typeSize.isAccessibilitySize {
            VStack(alignment: .leading, spacing: 2) { SimStateText(face: face, now: now); age }
        } else {
            ViewThatFits(in: .horizontal) {
                HStack(alignment: .firstTextBaseline) { SimStateText(face: face, now: now); Spacer(minLength: 8); age }
                VStack(alignment: .leading, spacing: 2) { SimStateText(face: face, now: now); age }
            }
        }
    }
}

/// A bulk run's 8 pt bar and `40 of 64 combinations` (only when the progress read counts them).
struct SimProgressBar: View {
    let progress: SimProgress

    var body: some View {
        if let fraction = ServerRuns.fraction(progress) {
            VStack(alignment: .leading, spacing: 4) {
                GeometryReader { proxy in
                    ZStack(alignment: .leading) {
                        Capsule().fill(FSColor.lineSoft)
                        Capsule().fill(FSColor.gold).frame(width: proxy.size.width * fraction)
                    }
                }
                .frame(height: 8)
                Text(Copy.simsCombos(done: progress.combosDone, total: progress.combosTotal))
                    .font(ScaledType.monoMeta).foregroundStyle(FSColor.muted)
                    .fixedSize(horizontal: false, vertical: true)
            }
            .padding(.top, 8)
            .accessibilityElement(children: .ignore)
            .accessibilityValue(Copy.simsCombos(done: progress.combosDone, total: progress.combosTotal))
        }
    }
}

enum SimsLayout {
    static let serverRowHeight: CGFloat = 88
    static let mineRowHeight: CGFloat = 72
    static let serverIcon: CGFloat = 36
    static let mineIcon: CGFloat = 32
    static let runAgainHeight: CGFloat = 44
    static let titleLines = 2

    /// A headline that is a plain figure (`1,043 DPS`) sets in mono; a sentence does not.
    static func isFigure(_ headline: String) -> Bool {
        headline.first.map { $0.isNumber } ?? false
    }
}
