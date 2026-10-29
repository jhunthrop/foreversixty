package discordbot

import (
	"strings"
	"testing"

	"github.com/jhunthrop/foreversixty/api/internal/spectable"
)

func TestGuideCopyMatchesTheSpec(t *testing.T) {
	g := guideCopy
	for key, want := range map[string]string{
		"bis.author": g.BisAuthor, "bis.title.none": g.BisTitleNone, "bis.title.class": g.BisTitleClass,
		"bis.title.ready": g.BisTitleReady, "bis.lead.unlinked": g.BisLeadUnlinked, "bis.lead.none": g.BisLeadNone,
		"bis.lead.class": g.BisLeadClass, "bis.lead.ready": g.BisLeadReady, "bis.lead.slot": g.BisLeadSlot,
		"bis.placeholder.class": g.BisPlaceholderClass, "bis.placeholder.specLocked": g.BisPlaceholderSpecLocked,
		"bis.placeholder.spec": g.BisPlaceholderSpec, "bis.placeholder.range": g.BisPlaceholderRange,
		"bis.show": g.BisShow, "bis.useMain": g.BisUseMain, "bis.change": g.BisChange,
		"char.lead.unlinked": g.CharLeadUnlinked, "char.lead.none": g.CharLeadNone, "char.unknown": g.CharUnknown,
		"char.someoneElse": g.CharSomeoneElse, "char.typeName": g.CharTypeName,
		"modal.char.title": g.ModalCharTitle, "modal.attendance.title": g.ModalAttendanceTitle,
		"modal.label": g.ModalLabel, "modal.placeholder": g.ModalPlaceholder,
		"signup.placeholder.raid": g.SignupPlaceholderRaid, "signup.placeholder.character": g.SignupPlaceholderCharacter,
		"attendance.placeholder": g.AttendancePlaceholder, "attendance.unknown": g.AttendanceUnknown,
		"readiness.placeholder": g.ReadinessPlaceholder, "report.placeholder": g.ReportPlaceholder,
		"card.stale": g.CardStale, "card.failure": g.CardFailure,
	} {
		if got := guideString(t, "`"+key+"`", 0); got != want {
			t.Errorf("%s = %q, spec %q", key, want, got)
		}
	}
	checkBoldRows(t, g)
	checkIndexedRows(t, g)
}

// checkBoldRows compares the rows the spec writes with bold segments: the stored string keeps
// Discord's ** markers, and without them it is the spec's text.
func checkBoldRows(t *testing.T, g guideStrings) {
	t.Helper()
	for _, c := range []struct{ key, have, bold string }{
		{"bis.lead.main", g.BisLeadMain, "**{Name}**"},
		{"bis.nudge.spec", g.BisNudgeSpec, "**Choose a spec first.**"},
		{"bis.nudge.both", g.BisNudgeBoth, "**Choose a class and a spec first.**"},
	} {
		segments := spectable.Cell(t, guideSpecPath, guideSpecHeading, "`"+c.key+"`")
		want := segments[0]
		if c.key != "bis.lead.main" {
			want = segments[0] + segments[1]
		}
		if got := strings.ReplaceAll(c.have, "**", ""); got != want {
			t.Errorf("%s = %q, spec %q", c.key, got, want)
		}
		if !strings.Contains(c.have, c.bold) {
			t.Errorf("%s lacks bold %q", c.key, c.bold)
		}
	}
}

func checkIndexedRows(t *testing.T, g guideStrings) {
	t.Helper()
	for i := range g.BisRangeDesc {
		if want := guideString(t, "`bis.range.desc`", i); g.BisRangeDesc[i] != want {
			t.Errorf("range desc %d = %q", i, g.BisRangeDesc[i])
		}
	}
	for i := range g.BisRole {
		if want := guideString(t, "`bis.role`", i); g.BisRole[i] != want {
			t.Errorf("role %d = %q", i, g.BisRole[i])
		}
	}
	for _, c := range []struct {
		key  string
		idx  int
		have string
	}{
		{"bis.lead.unknownSpec", 0, g.BisLeadUnknownSpec}, {"bis.lead.unknownSpec", 1, g.BisLeadUnknownSpecNoClass},
		{"char.select", 0, g.CharSelect}, {"char.select", 1, g.CharSelectDesc},
		{"signup.raid.option", 0, g.SignupRaidLabel}, {"signup.raid.option", 1, g.SignupRaidDesc},
		{"signup.raid.option", 2, g.SignupRaidDescNoSize},
	} {
		if want := guideString(t, "`"+c.key+"`", c.idx); c.have != want {
			t.Errorf("%s[%d] = %q, spec %q", c.key, c.idx, c.have, want)
		}
	}
}

func TestGuideCopyHasNoBannedWords(t *testing.T) {
	g := guideCopy
	all := []string{
		g.BisAuthor, g.BisLeadUnlinked, g.BisLeadNone, g.BisLeadMain, g.BisLeadClass, g.BisLeadReady,
		g.BisLeadSlot, g.BisLeadUnknownSpec, g.BisNudgeSpec, g.BisNudgeBoth, g.CharLeadUnlinked,
		g.CharLeadNone, g.CharUnknown, g.CardStale, g.CardFailure, g.AttendanceUnknown,
	}
	all = append(all, g.BisRangeDesc[:]...)
	for _, s := range all {
		low := strings.ToLower(s)
		if strings.Contains(low, "soon") || strings.Contains(s, "!") || strings.Contains(low, "unlock") {
			t.Errorf("banned word in %q", s)
		}
	}
}
