package notify

import "testing"

func TestChannelID(t *testing.T) {
	t.Run("Two cids are different", func(t *testing.T) {
		cidgen := new(cIDGen)
		cid1 := cidgen.generate(0)
		cid2 := cidgen.generate(0)

		if cid1 == cid2 {
			t.Errorf("The both cids `%s` and `%s` are the same", cid1, cid2)
		}
	})

	t.Run("Two cids from different cIDGen are different", func(t *testing.T) {
		cid1 := new(cIDGen).generate(0)
		cid2 := new(cIDGen).generate(0)

		if cid1 == cid2 {
			t.Errorf("The both cids `%s` and `%s` are the same", cid1, cid2)
		}
	})

	t.Run("cid.uid", func(t *testing.T) {
		cid := new(cIDGen).generate(1)

		if got := cid.uid(); got != 1 {
			t.Errorf("cid.uid() returned %d, expected 1", got)
		}
	})

	t.Run("invalid cid no :", func(t *testing.T) {
		cid := channelID("foobar")

		if got := cid.uid(); got != 0 {
			t.Errorf("cid.uid() returned %d, expected 0", got)
		}
	})

	t.Run("invalid cid uid not a number", func(t *testing.T) {
		cid := channelID("foo:bar:blub")

		if got := cid.uid(); got != 0 {
			t.Errorf("cid.uid() returned %d, expected 0", got)
		}
	})

}
