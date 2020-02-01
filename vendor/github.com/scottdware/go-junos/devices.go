package junos

import (
	"encoding/xml"
	"fmt"
	"regexp"

	"github.com/scottdware/go-rested"
)

// Devices contains a list of managed devices.
type Devices struct {
	XMLName xml.Name `xml:"devices"`
	Devices []Device `xml:"device"`
}

// A Device contains information about each individual device.
type Device struct {
	ID               int    `xml:"key,attr"`
	Family           string `xml:"deviceFamily"`
	Version          string `xml:"OSVersion"`
	Platform         string `xml:"platform"`
	Serial           string `xml:"serialNumber"`
	IPAddress        string `xml:"ipAddr"`
	Name             string `xml:"name"`
	ConnectionStatus string `xml:"connectionStatus"`
	ManagedStatus    string `xml:"managedStatus"`
}

// addDeviceIPXML is the XML we send (POST) for adding a device by IP address.
var addDeviceIPXML = `
<discover-devices>
    <ipAddressDiscoveryTarget>
        <ipAddress>%s</ipAddress>
    </ipAddressDiscoveryTarget>
    <sshCredential>
        <userName>%s</userName>
        <password>%s</password>
    </sshCredential>
    <manageDiscoveredSystemsFlag>true</manageDiscoveredSystemsFlag>
    <usePing>true</usePing>
</discover-devices>
`

// addDeviceHostXML is the XML we send (POST) for adding a device by hostname.
var addDeviceHostXML = `
<discover-devices>
    <hostNameDiscoveryTarget>
        <hostName>%s</hostName>
    </hostNameDiscoveryTarget>
    <sshCredential>
        <userName>%s</userName>
        <password>%s</password>
    </sshCredential>
    <manageDiscoveredSystemsFlag>true</manageDiscoveredSystemsFlag>
    <usePing>true</usePing>
</discover-devices>
`

// getDeviceID returns the ID of a managed device.
func (s *Space) getDeviceID(device interface{}) (int, error) {
	var err error
	var deviceID int
	ipRegex := regexp.MustCompile(`(\d{1,3}\.\d{1,3}\.\d{1,3}\.\d{1,3})`)
	devices, err := s.Devices()
	if err != nil {
		return 0, err
	}

	switch device.(type) {
	case int:
		deviceID = device.(int)
	case string:
		if ipRegex.MatchString(device.(string)) {
			for _, d := range devices.Devices {
				if d.IPAddress == device.(string) {
					deviceID = d.ID
				}
			}
		}
		for _, d := range devices.Devices {
			if d.Name == device.(string) {
				deviceID = d.ID
			}
		}
	}

	return deviceID, nil
}

// AddDevice adds a new managed device to Junos Space, and returns the Job ID.
func (s *Space) AddDevice(host, user, password string) (int, error) {
	r := rested.NewRequest()
	r.BasicAuth(s.User, s.Password)
	headers := map[string]string{
		"Content-Type": contentDiscoverDevices,
	}
	var job jobID
	var addDevice, xmlBody string
	ipRegex := regexp.MustCompile(`(\d{1,3}\.\d{1,3}\.\d{1,3}\.\d{1,3})`)

	if ipRegex.MatchString(host) {
		addDevice = addDeviceIPXML
	}

	addDevice = addDeviceHostXML

	uri := fmt.Sprintf("https://%s/api/space/device-management/discover-devices", s.Host)
	xmlBody = fmt.Sprintf(addDevice, host, user, password)

	resp := r.Send("post", uri, []byte(xmlBody), headers, nil)
	if resp.Error != nil {
		return 0, resp.Error
	}

	err := xml.Unmarshal(resp.Body, &job)
	if err != nil {
		return 0, err
	}

	return job.ID, nil
}

// Devices queries the Junos Space server and returns all of the information
// about each device that is managed by Space.
func (s *Space) Devices() (*Devices, error) {
	r := rested.NewRequest()
	r.BasicAuth(s.User, s.Password)
	var devices Devices
	uri := fmt.Sprintf("https://%s/api/space/device-management/devices", s.Host)

	resp := r.Send("get", uri, nil, nil, nil)
	if resp.Error != nil {
		return nil, resp.Error
	}

	err := xml.Unmarshal(resp.Body, &devices)
	if err != nil {
		return nil, err
	}

	return &devices, nil
}

// RemoveDevice removes a device from Junos Space. You can specify the device ID, name
// or IP address.
func (s *Space) RemoveDevice(device interface{}) error {
	r := rested.NewRequest()
	r.BasicAuth(s.User, s.Password)
	var err error
	deviceID, err := s.getDeviceID(device)
	if err != nil {
		return err
	}

	if deviceID != 0 {
		uri := fmt.Sprintf("https://%s/api/space/device-management/devices/%d", s.Host, deviceID)

		resp := r.Send("delete", uri, nil, nil, nil)
		if resp.Error != nil {
			return resp.Error
		}
	}

	return nil
}

// Resync synchronizes the device with Junos Space. Good to use if you make a lot of
// changes outside of Junos Space such as adding interfaces, zones, etc.
func (s *Space) Resync(device interface{}) (int, error) {
	r := rested.NewRequest()
	r.BasicAuth(s.User, s.Password)
	headers := map[string]string{
		"Content-Type": contentResync,
	}
	var job jobID
	deviceID, err := s.getDeviceID(device)
	if err != nil {
		return 0, err
	}

	uri := fmt.Sprintf("https://%s/api/space/device-management/devices/%d/exec-resync", s.Host, deviceID)

	resp := r.Send("post", uri, nil, headers, nil)
	if resp.Error != nil {
		return 0, resp.Error
	}

	err = xml.Unmarshal(resp.Body, &job)
	if err != nil {
		return 0, err
	}

	return job.ID, nil
}
