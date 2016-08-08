package quickfix

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type resendStateTestSuite struct {
	SessionSuite
}

func TestResendStateTestSuite(t *testing.T) {
	suite.Run(t, new(resendStateTestSuite))
}

func (s *resendStateTestSuite) SetupTest() {
	s.Init()
	s.session.State = resendState{}
}

func (s *resendStateTestSuite) TestIsLoggedOn() {
	s.True(s.session.IsLoggedOn())
}

func (s *resendStateTestSuite) TestTimeoutPeerTimeout() {
	s.mockApp.On("ToAdmin")
	s.session.Timeout(s.session, peerTimeout)

	s.mockApp.AssertExpectations(s.T())
	s.State(pendingTimeout{resendState{}})
}

func (s *resendStateTestSuite) TestTimeoutUnchangedIgnoreLogonLogoutTimeout() {
	tests := []event{logonTimeout, logoutTimeout}

	for _, event := range tests {
		s.session.Timeout(s.session, event)
		s.State(resendState{})
	}
}

func (s *resendStateTestSuite) TestTimeoutUnchangedNeedHeartbeat() {
	s.mockApp.On("ToAdmin")
	s.session.Timeout(s.session, needHeartbeat)

	s.mockApp.AssertExpectations(s.T())
	s.State(resendState{})
}