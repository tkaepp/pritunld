package connection

import (
	"sync"
	"time"

	"github.com/pritunl/pritunl-client-electron/service/utils"
	"github.com/sirupsen/logrus"
)

type State struct {
	conn               *Connection
	startTime          time.Time
	stop               bool
	lastStopCheckStack string
	lastStopCheckTime  time.Time
	deadline           bool
	delay              bool
	automatic          bool
	closed             bool
	closeWaiters       []chan bool
	closeWaitersLock   sync.Mutex
	tempPaths          []string
}

func (s *State) Fields() logrus.Fields {
	return logrus.Fields{
		"state_time":       s.startTime,
		"state_stop":       s.stop,
		"state_deadline":   s.deadline,
		"state_delay":      s.delay,
		"state_automatic":  s.automatic,
		"state_closed":     s.closed,
		"state_temp_paths": s.tempPaths,
	}
}

func (s *State) Init(opts Options) (err error) {
	s.deadline = opts.Deadline
	s.delay = opts.Delay
	s.automatic = opts.Automatic
	s.startTime = time.Now()
	s.tempPaths = []string{}

	go s.stopWatch()

	return
}

func (s *State) PreStart() {
	if s.delay {
		time.Sleep(3 * time.Second)
	}
}

func (s *State) stopWatch() {
	for {
		time.Sleep(1 * time.Second)
		if s.closed {
			return
		}

		if time.Since(s.lastStopCheckTime) > 1*time.Minute {
			logrus.WithFields(s.conn.Fields(logrus.Fields{
				"last_stop_check": s.lastStopCheckTime.Format(
					"2006-01-02 15:04:05"),
				"trace": s.lastStopCheckStack,
			})).Info("state: Detected dead state")
			s.lastStopCheckTime = time.Now()
		}
	}
}

func (s *State) Stop() {
	if s.stop {
		logrus.WithFields(s.conn.Fields(nil)).Info(
			"state: Profile already in stop")
		return
	}

	s.stop = true
}

func (s *State) IsStop() bool {
	trace := utils.GetStackTrace()

	s.lastStopCheckTime = time.Now()
	s.lastStopCheckStack = trace

	if Shutdown || s.stop {
		return true
	}
	return false
}

func (s *State) IsAutomatic() bool {
	return s.automatic
}

func (s *State) SetConnecting() {
	logrus.WithFields(logrus.Fields{
		"profile_id":       s.conn.Profile.Id,
		"mode":             s.conn.Profile.Mode,
		"dynamic_firewall": s.conn.Profile.DynamicFirewall,
		"device_auth":      s.conn.Profile.DeviceAuth,
		"disable_gateway":  s.conn.Profile.DisableGateway,
		"disable_dns":      s.conn.Profile.DisableDns,
		"geo_sort":         s.conn.Profile.GeoSort,
		"force_connect":    s.conn.Profile.ForceConnect,
		"force_dns":        s.conn.Profile.ForceDns,
		"sso_auth":         s.conn.Profile.SsoAuth,
		"reconnect":        s.conn.Profile.Reconnect,
	}).Info("profile: Connecting")

	s.conn.Data.Status = Connecting

	return
}

func (s *State) AddPath(pth string) {
	s.tempPaths = append(s.tempPaths, pth)
}

func (s *State) CloseWait() {
	waiter := make(chan bool, 3)

	s.closeWaitersLock.Lock()
	if !s.closed {
		s.closeWaitersLock.Unlock()
		return
	}
	s.closeWaiters = append(s.closeWaiters, waiter)
	s.closeWaitersLock.Unlock()

	<-waiter
	time.Sleep(50 * time.Millisecond)
}

func (s *State) Close() {
	s.closeWaitersLock.Lock()
	if s.closed {
		s.closeWaitersLock.Unlock()
		return
	}
	s.closed = true

	for _, waiter := range s.closeWaiters {
		waiter <- true
	}
	s.closeWaiters = []chan bool{}

	s.closeWaitersLock.Unlock()
}