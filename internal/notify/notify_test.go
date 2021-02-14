package notify

import "testing"

func TestApplause(t *testing.T) {
	backend := new(backendMock)
	closed := make(chan struct{})
	defer close(closed)

	n := New(backend, new(applauserMock), 1, closed)

	t.Run("no applause send", func(t *testing.T) {
		a, b, err := n.receiceApplause()
		if err != nil {
			t.Fatalf("got unexpected error %v", err)
		}

		if a != 0 || b != 100 {
			t.Errorf("receiceApplause returned %d, %d, expect 0, 100", a, b)
		}
	})

	t.Run("one applause send", func(t *testing.T) {
		backend.a = 5

		a, b, err := n.receiceApplause()
		if err != nil {
			t.Fatalf("got unexpected error %v", err)
		}

		if a != 5 || b != 100 {
			t.Errorf("receiceApplause returned %d, %d, expect 5, 100", a, b)
		}
	})

}

type backendMock struct {
	a int
}

func (m *backendMock) SendNotify([]byte) error {
	return nil
}

func (m *backendMock) ReceiveNotify(closing <-chan struct{}) (mail string, err error) {
	return "", nil
}

func (m *backendMock) AddApplause(userID int) error {
	return nil
}

func (m *backendMock) GetApplause(since int64) (int, error) {
	return m.a, nil
}

type applauserMock struct{}

func (a *applauserMock) ApplauseConfig() (waitTime int, base int) {
	return 5, 100
}
