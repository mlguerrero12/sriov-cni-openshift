package sriov

import (
	"net"

	"github.com/containernetworking/plugins/pkg/ns"
	"github.com/containernetworking/plugins/pkg/testutils"
	sriovtypes "github.com/intel/sriov-cni/pkg/types"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/stretchr/testify/mock"
	"github.com/vishvananda/netlink"
)

// Code generated by mockery v1.0.0. DO NOT EDIT.
// MockNetlinkManager is an autogenerated mock type for the NetlinkManager type
type MockNetlinkManager struct {
	mock.Mock
}

// LinkByName provides a mock function with given fields: _a0
func (_m *MockNetlinkManager) LinkByName(_a0 string) (netlink.Link, error) {
	ret := _m.Called(_a0)

	var r0 netlink.Link
	if rf, ok := ret.Get(0).(func(string) netlink.Link); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(netlink.Link)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// LinkSetVfVlan provides a mock function with given fields: _a0, _a1, _a2
func (_m *MockNetlinkManager) LinkSetVfVlan(_a0 netlink.Link, _a1 int, _a2 int) error {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 error
	if rf, ok := ret.Get(0).(func(netlink.Link, int, int) error); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// LinkSetVfHardwareAddr provides a mock function with given fields: _a0, _a1, _a2
func (_m *MockNetlinkManager) LinkSetVfHardwareAddr(_a0 netlink.Link, _a1 int, _a2 net.HardwareAddr) error {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 error
	if rf, ok := ret.Get(0).(func(netlink.Link, int, net.HardwareAddr) error); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// LinkSetHardwareAddr provides a mock function with given fields: _a0, _a1
func (_m *MockNetlinkManager) LinkSetHardwareAddr(_a0 netlink.Link, _a1 net.HardwareAddr) error {
	ret := _m.Called(_a0, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(netlink.Link, net.HardwareAddr) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// LinkSetUp provides a mock function with given fields: _a0
func (_m *MockNetlinkManager) LinkSetUp(_a0 netlink.Link) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(netlink.Link) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// LinkSetDown provides a mock function with given fields: _a0
func (_m *MockNetlinkManager) LinkSetDown(_a0 netlink.Link) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(netlink.Link) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// LinkSetNsFd provides a mock function with given fields: _a0, _a1
func (_m *MockNetlinkManager) LinkSetNsFd(_a0 netlink.Link, _a1 int) error {
	ret := _m.Called(_a0, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(netlink.Link, int) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// LinkSetName provides a mock function with given fields: _a0, _a1
func (_m *MockNetlinkManager) LinkSetName(_a0 netlink.Link, _a1 string) error {
	ret := _m.Called(_a0, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(netlink.Link, string) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// mockPciUtils is an autogenerated mock type for the mockPciUtils type
type mockPciUtils struct {
	mock.Mock
}

// getPciAddress provides a mock function with given fields: ifName, vf
func (_m *mockPciUtils) getPciAddress(ifName string, vf int) (string, error) {
	ret := _m.Called(ifName, vf)

	var r0 string
	if rf, ok := ret.Get(0).(func(string, int) string); ok {
		r0 = rf(ifName, vf)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, int) error); ok {
		r1 = rf(ifName, vf)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// getSriovNumVfs provides a mock function with given fields: ifName
func (_m *mockPciUtils) getSriovNumVfs(ifName string) (int, error) {
	ret := _m.Called(ifName)

	var r0 int
	if rf, ok := ret.Get(0).(func(string) int); ok {
		r0 = rf(ifName)
	} else {
		r0 = ret.Get(0).(int)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(ifName)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// getVFLinkNamesFromVFID provides a mock function with given fields: pfName, vfID
func (_m *mockPciUtils) getVFLinkNamesFromVFID(pfName string, vfID int) ([]string, error) {
	ret := _m.Called(pfName, vfID)

	var r0 []string
	if rf, ok := ret.Get(0).(func(string, int) []string); ok {
		r0 = rf(pfName, vfID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, int) error); ok {
		r1 = rf(pfName, vfID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FakeLink is a dummy netlink struct used during testing
type FakeLink struct {
	netlink.LinkAttrs
}

// type FakeLink struct {
// 	linkAtrrs *netlink.LinkAttrs
// }

func (l *FakeLink) Attrs() *netlink.LinkAttrs {
	return &l.LinkAttrs
}

func (l *FakeLink) Type() string {
	return "FakeLink"
}

var _ = Describe("Sriov", func() {

	Context("Checking SetupVF function", func() {
		var (
			podifName string
			contID    string
			netconf   *sriovtypes.NetConf
		)

		BeforeEach(func() {
			podifName = "net1"
			contID = "dummycid"
			netconf = &sriovtypes.NetConf{
				Master:      "enp175s0f1",
				DeviceID:    "0000:af:06.0",
				VFID:        0,
				HostIFNames: "enp175s6",
				ContIFNames: "net1",
			}
		})

		It("Assuming existing interface", func() {
			var targetNetNS ns.NetNS
			targetNetNS, err := testutils.NewNS()
			defer func() {
				if targetNetNS != nil {
					targetNetNS.Close()
				}
			}()
			Expect(err).NotTo(HaveOccurred())
			mocked := &MockNetlinkManager{}
			fakeMac, err := net.ParseMAC("6e:16:06:0e:b7:e9")

			Expect(err).NotTo(HaveOccurred())

			fakeLink := &FakeLink{netlink.LinkAttrs{
				Index:        1000,
				Name:         "dummylink",
				HardwareAddr: fakeMac,
			}}

			mocked.On("LinkByName", mock.AnythingOfType("string")).Return(fakeLink, nil)
			mocked.On("LinkSetDown", fakeLink).Return(nil)
			mocked.On("LinkSetName", fakeLink, mock.Anything).Return(nil)
			mocked.On("LinkSetNsFd", fakeLink, mock.AnythingOfType("int")).Return(nil)
			mocked.On("LinkSetUp", fakeLink).Return(nil)
			mocked.On("LinkSetVfVlan", mock.Anything, mock.AnythingOfType("int"), mock.AnythingOfType("int")).Return(nil)
			sm := sriovManager{nLink: mocked}
			macAddr, err := sm.SetupVF(netconf, podifName, contID, targetNetNS)
			Expect(err).NotTo(HaveOccurred())
			Expect(macAddr).To(Equal("6e:16:06:0e:b7:e9"))
		})
		Context("Checking ReleaseVF function", func() {
			var (
				podifName string
				contID    string
				netconf   *sriovtypes.NetConf
			)

			BeforeEach(func() {
				podifName = "net1"
				contID = "dummycid"
				netconf = &sriovtypes.NetConf{
					Master:      "enp175s0f1",
					DeviceID:    "0000:af:06.0",
					VFID:        0,
					HostIFNames: "enp175s6",
					ContIFNames: "net1",
				}
			})
			It("Assuming existing interface", func() {
				var targetNetNS ns.NetNS
				targetNetNS, err := testutils.NewNS()
				defer func() {
					if targetNetNS != nil {
						targetNetNS.Close()
					}
				}()
				Expect(err).NotTo(HaveOccurred())
				mocked := &MockNetlinkManager{}
				fakeLink := &FakeLink{netlink.LinkAttrs{Index: 1000, Name: "dummylink"}}

				mocked.On("LinkByName", mock.AnythingOfType("string")).Return(fakeLink, nil)
				mocked.On("LinkSetDown", fakeLink).Return(nil)
				mocked.On("LinkSetName", fakeLink, mock.Anything).Return(nil)
				mocked.On("LinkSetNsFd", fakeLink, mock.AnythingOfType("int")).Return(nil)
				mocked.On("LinkSetUp", fakeLink).Return(nil)
				mocked.On("LinkSetVfVlan", mock.Anything, mock.AnythingOfType("int"), mock.AnythingOfType("int")).Return(nil)
				sm := sriovManager{nLink: mocked}
				err = sm.ReleaseVF(netconf, podifName, contID, targetNetNS)
				Expect(err).NotTo(HaveOccurred())
			})
		})
	})
})
