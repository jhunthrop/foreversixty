// api/internal/directory/optout_test.go
package directory

import (
	"context"
	"errors"
	"testing"
)

func insertDirectoryRow(t *testing.T, s *Service, id int64, name, guild string) {
	t.Helper()
	if _, err := s.Pool.Exec(context.Background(),
		`insert into character_directory (region, bnet_character_id, name, slug, guild_bnet_id, guild_name, roster_seen_at)
		 values ('us', $1, $2, $3, 77, $4, now())`, id, name, bnetSlug(name), guild); err != nil {
		t.Fatal(err)
	}
}

// A code in the character's public guild note, sent by a guildmate's addon for that guild,
// removes the character, and no later roster read stores it again.
func TestOptOutByGuildNoteCode(t *testing.T) {
	pool := testPool(t)
	b := newBlizzard(t)
	s := b.service(pool)
	o := &OptOuts{Pool: pool, Client: s.Client}
	ctx := context.Background()
	insertDirectoryRow(t, s, 101, "Vela Thornbrook", "Ember Wardens")

	code, err := o.StartCode(ctx, "us", "Vela Thornbrook")
	if err != nil || !ValidCode(code.Code) {
		t.Fatalf("code = %+v, %v", code, err)
	}
	if state, _ := o.Status(ctx, "us", "vela thornbrook", code.Code); state != OptOutWaiting {
		t.Fatalf("state = %q, want waiting", state)
	}
	// Without the code a stranger learns nothing.
	if state, _ := o.Status(ctx, "us", "Vela Thornbrook", "FS-ABCDEF"); state != OptOutNone {
		t.Fatalf("state without the code = %q, want none", state)
	}
	// Another guild's addon, or a wrong name, removes nothing.
	if n, err := o.VerifyNoteCodes(ctx, "us", "Other Guild", []NoteCode{{Name: "Vela Thornbrook", Code: code.Code}}); err != nil || n != 0 {
		t.Fatalf("other guild removed %d, %v", n, err)
	}
	if n, _ := o.VerifyNoteCodes(ctx, "us", "Ember Wardens", []NoteCode{{Name: "Orrin Ashmantle", Code: code.Code}}); n != 0 {
		t.Fatal("a code under another member's name removed the character")
	}
	if n, err := o.VerifyNoteCodes(ctx, "us", "ember wardens", []NoteCode{{Name: "Vela Thornbrook", Code: code.Code}}); err != nil || n != 1 {
		t.Fatalf("removed %d, %v; want 1", n, err)
	}
	if _, ok := readRow(t, s, 101); ok {
		t.Fatal("the character stayed in the directory")
	}
	if state, _ := o.Status(ctx, "us", "Vela Thornbrook", code.Code); state != OptOutRemoved {
		t.Fatalf("state = %q, want removed", state)
	}
	if state, _ := o.Status(ctx, "us", "Vela Thornbrook", ""); state != OptOutNone {
		t.Fatalf("removed state without the code = %q, want none", state)
	}

	seedGuildRow(t, s, "Ember Wardens", "pvp")
	b.guild("ember-wardens", "Ember Wardens", "PvP", 77, rosterMember{id: 101, name: "Vela Thornbrook", level: 41, classID: 8})
	b.publicCharacter(101, "Vela Thornbrook", "PvP", 77, "ember-wardens", "Ember Wardens")
	if _, err := s.Pass(ctx); err != nil {
		t.Fatal(err)
	}
	if _, ok := readRow(t, s, 101); ok {
		t.Fatal("a roster read stored an opted-out character again")
	}
	if b.called("/profile/wow/character/vela-thornbrook") != 0 {
		t.Fatal("an opted-out character's profile was read")
	}
}

// A signed-in account removes its own Battle.net character at once, and no one else's.
func TestOptOutByAccount(t *testing.T) {
	pool := testPool(t)
	b := newBlizzard(t)
	s := b.service(pool)
	o := &OptOuts{Pool: pool}
	ctx := context.Background()
	insertDirectoryRow(t, s, 101, "Vela Thornbrook", "Ember Wardens")
	var owner, other int64
	for _, id := range []*int64{&owner, &other} {
		if err := pool.QueryRow(ctx, `insert into users default values returning id`).Scan(id); err != nil {
			t.Fatal(err)
		}
	}
	if _, err := pool.Exec(ctx,
		`insert into characters (key, region, ruleset, name, user_id, bnet_character_id, source)
		 values ('us/pvp/vela-thornbrook', 'us', 'pvp', 'Vela Thornbrook', $1, 101, 'bnet')`, owner); err != nil {
		t.Fatal(err)
	}
	if err := o.RemoveOwned(ctx, other, "us", "Vela Thornbrook"); !errors.Is(err, ErrNotYours) {
		t.Fatalf("another account: %v, want ErrNotYours", err)
	}
	if err := o.RemoveOwned(ctx, owner, "us", "Vela Thornbrook"); err != nil {
		t.Fatal(err)
	}
	if _, ok := readRow(t, s, 101); ok {
		t.Fatal("the character stayed in the directory")
	}
}

// A name the directory does not hold is looked up on Blizzard, so it can be kept out; a name
// nobody has is refused.
func TestOptOutCodeForANameTheDirectoryLacks(t *testing.T) {
	pool := testPool(t)
	b := newBlizzard(t)
	s := b.service(pool)
	o := &OptOuts{Pool: pool, Client: s.Client}
	ctx := context.Background()
	b.publicCharacter(202, "Sel Duskwater", "Normal", 0, "", "")
	if _, err := o.StartCode(ctx, "us", "Sel Duskwater"); err != nil {
		t.Fatal(err)
	}
	// An unknown name is answered like a held one (no existence probe), and nothing is saved.
	if c, err := o.StartCode(ctx, "us", "Nobody Here"); err != nil || !ValidCode(c.Code) {
		t.Fatalf("unknown name: %+v, %v, want a code", c, err)
	}
	var saved int
	if err := pool.QueryRow(ctx, `select count(*) from directory_optout_codes where slug = 'nobody-here'`).Scan(&saved); err != nil || saved != 0 {
		t.Fatalf("unknown name saved %d codes (%v)", saved, err)
	}
	if _, err := o.StartCode(ctx, "kr", "Sel Duskwater"); !errors.Is(err, ErrBadCharacter) {
		t.Fatalf("kr: %v, want ErrBadCharacter", err)
	}
}
