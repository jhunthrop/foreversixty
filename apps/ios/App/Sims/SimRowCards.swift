import DesignSystem
import Models
import SwiftUI

/// One Server runs row (spec F5.1): a done row opens the result; a failed row's one control is
/// `Run again`; queued and running rows are not links.
struct ServerRunRowView: View {
    let face: SimRowFace
    let now: Date
    let build: String?
    let progress: SimProgress?
    let rerun: RerunStatus
    let offline: Bool
    let mayLinkToPlans: Bool
    let actions: SimsActions

    var body: some View {
        if face.state.isDone {
            Button { actions.open(face.row) } label: { content(chevron: true) }
                .buttonStyle(RowPressStyle())
                .accessibilityHint(Copy.simsOpenHint)
        } else {
            content(chevron: false)
        }
    }

    private func content(chevron: Bool) -> some View {
        HStack(alignment: .top, spacing: 12) {
            SimSpecMark(icon: face.icon, build: build, size: SimsLayout.serverIcon)
            VStack(alignment: .leading, spacing: 2) {
                SimRowTitle(face: face)
                Text(face.meta).font(ScaledType.meta).foregroundStyle(FSColor.muted)
                    .fixedSize(horizontal: false, vertical: true)
                SimStateLine(face: face, now: now).padding(.top, 2)
                    .accessibilityElement(children: .combine)
                if let progress, face.state.isInFlight { SimProgressBar(progress: progress) }
                if face.state == .failed { failedControls }
            }
            .frame(maxWidth: .infinity, alignment: .leading)
            if chevron { SimChevron() }
        }
        .padding(.horizontal, 16)
        .padding(.vertical, 12)
        .frame(minHeight: SimsLayout.serverRowHeight, alignment: .top)
        .contentShape(Rectangle())
        .accessibilityElement(children: .contain)
        .accessibilityLabel(SimsAccessibility.serverRow(face, now: now))
    }

    @ViewBuilder
    private var failedControls: some View {
        Button(rerunTitle) { actions.runAgain(face.row.simId ?? "") }
            .buttonStyle(BrandButtonStyle(.secondary, height: SimsLayout.runAgainHeight))
            .disabled(rerun.isLocked || offline)
            .fixedSize(horizontal: false, vertical: true)
            .accessibilityLabel("\(rerunTitle), \(face.title)")
            .padding(.top, 8)
        if offline {
            Text(Copy.simsOffline).font(ScaledType.meta).foregroundStyle(FSColor.muted)
        } else if rerun == .failed {
            Text(Copy.simsRunAgainFailed).font(ScaledType.meta).foregroundStyle(FSColor.wipe)
                .accessibilityAddTraits(.isStaticText)
        } else if rerun == .lapsed {
            HStack(alignment: .firstTextBaseline, spacing: 6) {
                Text(Copy.simsPremiumLine).font(ScaledType.meta).foregroundStyle(FSColor.wipe)
                if mayLinkToPlans {
                    Button(Copy.youSeePlans, action: actions.openPlans)
                        .font(ScaledType.meta.weight(.semibold)).underline().foregroundStyle(FSColor.gold)
                        .frame(minHeight: SimsLayout.runAgainHeight)
                }
            }
        }
    }

    private var rerunTitle: String {
        switch rerun {
        case .pending: Copy.simsRunAgainPending
        case .started: Copy.simsRunAgainDone
        default: Copy.simsRunAgainRow
        }
    }
}

/// One Your sims row: the headline (or the state while in flight), and where it ran.
struct MySimRowView: View {
    let face: SimRowFace
    let now: Date
    let build: String?
    let open: (SimRow) -> Void

    var body: some View {
        if face.state.isDone {
            Button { open(face.row) } label: { content(chevron: true) }
                .buttonStyle(RowPressStyle())
                .accessibilityHint(Copy.simsOpenHint)
        } else {
            content(chevron: false)
        }
    }

    private func content(chevron: Bool) -> some View {
        HStack(alignment: .top, spacing: 12) {
            SimSpecMark(icon: face.icon, build: build, size: SimsLayout.mineIcon)
            VStack(alignment: .leading, spacing: 2) {
                SimRowTitle(face: face)
                SimStateText(face: face, now: now)
                Text(SimsCopy.origin(face.row)).font(ScaledType.meta).foregroundStyle(FSColor.muted)
            }
            .frame(maxWidth: .infinity, alignment: .leading)
            if chevron { SimChevron() }
        }
        .padding(.horizontal, 16)
        .padding(.vertical, 12)
        .frame(minHeight: SimsLayout.mineRowHeight, alignment: .top)
        .contentShape(Rectangle())
        .accessibilityElement(children: .combine)
    }
}

struct SimChevron: View {
    var body: some View {
        Image(systemName: "chevron.right").font(.system(size: 14, weight: .semibold))
            .foregroundStyle(FSColor.muted).frame(height: 36).accessibilityHidden(true)
    }
}

/// The row labels VoiceOver reads (spec F5 Accessibility).
enum SimsAccessibility {
    static func serverRow(_ face: SimRowFace, now: Date) -> String {
        var parts = [face.title, face.specName]
        switch face.state {
        case .done, .doneUnseen:
            parts.append(face.headline)
            parts.append(SimsCopy.finishedLine(face.row, now: now))
            if face.isUnseen { parts.append(Copy.simsUnread) }
        case .failed:
            parts.append(SimsCopy.failedLine(face.row))
            parts.append(SimAge.line(face, now: now))
        default:
            parts.append(Copy.simsRunning + SimsCopy.stateSuffix(face.row, state: face.state, now: now))
        }
        return parts.filter { !$0.isEmpty }.joined(separator: ", ")
    }
}
