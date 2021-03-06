package telemetry

import (
	"context"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	log "github.com/sirupsen/logrus"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/xaque208/znet/internal/inventory"
	"github.com/xaque208/znet/internal/lights"
	"github.com/xaque208/znet/pkg/iot"
)

type Server struct {
	UnimplementedTelemetryServer
	inventory  inventory.Inventory
	keeper     thingKeeper
	lights     *lights.Lights
	iotServer  *iot.Server
	seenThings map[string]time.Time
}

type thingKeeper map[string]map[string]string

// NewServer returns a new Server.
func NewServer(inv inventory.Inventory, lig *lights.Lights) (*Server, error) {
	s := &Server{
		inventory:  inv,
		keeper:     make(thingKeeper),
		lights:     lig,
		seenThings: make(map[string]time.Time),
	}

	go func(s *Server) {
		for {
			// Make a copy
			tMap := make(map[string]time.Time)
			for k, v := range s.seenThings {
				tMap[k] = v
			}

			// Expire the old entries
			for k, v := range tMap {
				if time.Since(v) > (300 * time.Second) {
					log.WithFields(log.Fields{
						"device": k,
					}).Info("expiring")

					airHeatindex.Delete(prometheus.Labels{"device": k})
					airHumidity.Delete(prometheus.Labels{"device": k})
					airTemperature.Delete(prometheus.Labels{"device": k})
					thingWireless.Delete(prometheus.Labels{"device": k})
					waterTemperature.Delete(prometheus.Labels{"device": k})

					delete(s.seenThings, k)
					delete(s.keeper, k)
				}
			}

			time.Sleep(30 * time.Second)
		}
	}(s)

	return s, nil
}

// storeThingLabel records the received key/value pair for the given node ID.
func (l *Server) storeThingLabel(nodeID string, key, value string) {
	if len(l.keeper) == 0 {
		l.keeper = make(thingKeeper)
	}

	if _, ok := l.keeper[nodeID]; !ok {
		l.keeper[nodeID] = make(map[string]string)
	}

	if key != "" && value != "" {
		l.keeper[nodeID][key] = value
	}
}

func (l *Server) nodeLabels(nodeID string) map[string]string {
	if nodeLabelMap, ok := l.keeper[nodeID]; ok {
		return nodeLabelMap
	}

	return map[string]string{}
}

// hasLabels checks to see if the keeper has all of the received labels for the given node ID.
func (l *Server) hasLabels(nodeID string, labels []string) bool {
	nodeLabels := l.nodeLabels(nodeID)

	nodeHasLabel := func(nodeLabels map[string]string, label string) bool {

		for key := range nodeLabels {
			if key == label {
				return true
			}
		}

		return false
	}

	for _, label := range labels {
		if !nodeHasLabel(nodeLabels, label) {
			return false
		}
	}

	return true
}

func (l *Server) findMACs(macs []string) ([]*inventory.NetworkHost, []*inventory.NetworkID, error) {
	var keepHosts []*inventory.NetworkHost
	var keepIds []*inventory.NetworkID

	networkHosts, err := l.inventory.ListNetworkHosts()
	if err != nil {
		return nil, nil, err
	}

	for i := range networkHosts {
		x := proto.Clone(&(networkHosts)[i]).(*inventory.NetworkHost)

		if x.MacAddress != nil {
			for _, m := range x.MacAddress {
				for _, mm := range macs {
					if strings.EqualFold(m, mm) {
						keepHosts = append(keepHosts, x)
					}
				}
			}
		}
	}

	networkIDs, err := l.inventory.ListNetworkIDs()
	if err != nil {
		return nil, nil, err
	}

	for i := range networkIDs {
		x := proto.Clone(&(networkIDs)[i]).(*inventory.NetworkID)

		if x.MacAddress != nil {
			for _, m := range x.MacAddress {
				for _, mm := range macs {
					if strings.EqualFold(m, mm) {
						keepIds = append(keepIds, x)
					}
				}
			}
		}
	}

	return keepHosts, keepIds, nil
}

func (l *Server) ReportNetworkID(ctx context.Context, request *inventory.NetworkID) (*inventory.Empty, error) {
	log.WithFields(log.Fields{
		"name":                       request.Name,
		"ip_address":                 request.IpAddress,
		"reporting_source":           request.ReportingSource,
		"reporting_source_interface": request.ReportingSourceInterface,
	}).Trace("inventory.NetworkID report")

	if request.Name == "" {
		return &inventory.Empty{}, fmt.Errorf("unable to fetch inventory.NetworkID with empty name")
	}

	hosts, ids, err := l.findMACs(request.MacAddress)
	if err != nil {
		return &inventory.Empty{}, err
	}

	// do nothing if a host matches
	if len(hosts) > 0 {
		for _, x := range ids {
			err = l.inventory.UpdateTimestamp(x.Dn, "networkHost")
			if err != nil {
				log.Error(err)
			}
		}
		return &inventory.Empty{}, nil
	}

	now := time.Now()

	// update the lastSeen for nettworkIds
	if len(ids) > 0 {
		log.Debugf("ids found for report: %+v", ids)
		for _, id := range ids {
			if id.Dn != "" {
				x := &inventory.NetworkID{
					Dn:                       id.Dn,
					IpAddress:                request.IpAddress,
					MacAddress:               request.MacAddress,
					ReportingSource:          request.ReportingSource,
					ReportingSourceInterface: request.ReportingSourceInterface,
					LastSeen:                 timestamppb.New(now),
				}

				_, err = l.inventory.UpdateNetworkID(x)
				if err != nil {
					return &inventory.Empty{}, err
				}
			}
		}
	}

	log.Debugf("existing mac not found: %+v", request.MacAddress)

	x := &inventory.NetworkID{
		Name:                     request.Name,
		IpAddress:                request.IpAddress,
		MacAddress:               request.MacAddress,
		ReportingSource:          request.ReportingSource,
		ReportingSourceInterface: request.ReportingSourceInterface,
		LastSeen:                 timestamppb.New(now),
	}

	_, err = l.inventory.FetchNetworkID(request.Name)
	if err != nil {
		_, err = l.inventory.CreateNetworkID(x)
		if err != nil {
			return &inventory.Empty{}, err
		}
	}

	return &inventory.Empty{}, nil
}

func (l *Server) ReportIOTDevice(ctx context.Context, request *inventory.IOTDevice) (*inventory.Empty, error) {

	var err error

	if request.DeviceDiscovery == nil {
		return nil, fmt.Errorf("unable to receive IOTDevice with nil DeviceDiscovery")
	}

	log.WithFields(log.Fields{
		"component": request.DeviceDiscovery.Component,
		"node_id":   request.DeviceDiscovery.NodeId,
		"object_id": request.DeviceDiscovery.ObjectId,
		"endpoint":  request.DeviceDiscovery.Endpoint,
		"message":   string(request.DeviceDiscovery.Message),
	}).Trace("inventory.IOTDevice report")

	discovery := request.DeviceDiscovery

	if discovery.ObjectId != "" {
		telemetryIOTReport.WithLabelValues(discovery.ObjectId, discovery.Component).Inc()
	}

	switch discovery.Component {
	case "zigbee2mqtt":
		err = l.handleZigbeeReport(request)
		if err != nil {
			return &inventory.Empty{}, err
		}
	}

	switch discovery.ObjectId {
	case "wifi":
		err = l.handleWifiReport(request)
		if err != nil {
			return &inventory.Empty{}, err
		}
	case "air":
		err = l.handleAirReport(request)
		if err != nil {
			return &inventory.Empty{}, err
		}
	case "water":
		err = l.handleWaterReport(request)
		if err != nil {
			return &inventory.Empty{}, err
		}
	case "led1", "led2":
		err = l.handleLEDReport(request)
		if err != nil {
			return &inventory.Empty{}, err
		}
	default:
		telemetryIOTUnhandledReport.WithLabelValues(discovery.ObjectId, discovery.Component).Inc()
	}

	return &inventory.Empty{}, nil
}

func (l *Server) SetIOTServer(iotServer *iot.Server) error {
	if l.iotServer != nil {
		log.Debugf("replaceing iotServer on telemetryServer")
	}

	l.iotServer = iotServer

	return nil
}

func (l *Server) handleZigbeeReport(request *inventory.IOTDevice) error {
	if request == nil {
		return fmt.Errorf("unable to read zigbee report from nil request")
	}

	discovery := request.DeviceDiscovery

	msg, err := iot.ReadZigbeeMessage(discovery.ObjectId, discovery.Message, discovery.Endpoint...)
	if err != nil {
		return err
	}

	if msg == nil {
		return nil
	}

	now := time.Now()

	switch reflect.TypeOf(msg).String() {
	case "iot.ZigbeeBridgeState":
		m := msg.(iot.ZigbeeBridgeState)
		switch m {
		case iot.Offline:
			telemetryIOTBridgeState.WithLabelValues().Set(float64(0))
		case iot.Online:
			telemetryIOTBridgeState.WithLabelValues().Set(float64(1))
		}

	case "iot.ZigbeeBridgeLog":
		m := msg.(iot.ZigbeeBridgeLog)

		if m.Message == nil {
			return fmt.Errorf("unhandled iot.ZigbeeBridgeLog type: %s", m.Type)
		}

		messageTypeName := reflect.TypeOf(m.Message).String()

		switch messageTypeName {
		case "string":
			if strings.HasPrefix(m.Message.(string), "Update available") {
				return l.handleZigbeeDeviceUpdate(m)
			}
		case "iot.ZigbeeBridgeMessageDevices":
			return l.handleZigbeeDevices(m.Message.(iot.ZigbeeBridgeMessageDevices))
		default:
			return fmt.Errorf("unhandled iot.ZigbeeBridgeLog: %s", messageTypeName)
		}

	case "iot.ZigbeeBridgeMessageDevices":
		m := msg.(iot.ZigbeeBridgeMessageDevices)

		return l.handleZigbeeDevices(m)
	case "iot.ZigbeeMessage":
		m := msg.(iot.ZigbeeMessage)

		if m.Battery > 0 {
			telemetryIOTBatteryPercent.WithLabelValues(request.DeviceDiscovery.ObjectId, request.DeviceDiscovery.Component).Set(float64(m.Battery))
		}

		if m.LinkQuality > 0 {
			telemetryIOTLinkQuality.WithLabelValues(request.DeviceDiscovery.ObjectId, request.DeviceDiscovery.Component).Set(float64(m.LinkQuality))
		}

		x := &inventory.ZigbeeDevice{
			Name:     request.DeviceDiscovery.ObjectId,
			LastSeen: timestamppb.New(now),
		}

		result, err := l.inventory.FetchZigbeeDevice(x.Name)
		if err != nil {
			log.Error(err)
			_, err = l.inventory.CreateZigbeeDevice(x)
			if err != nil {
				return err
			}
		}

		if m.Action != "" {
			action := &iot.Action{
				Event:  m.Action,
				Device: x.Name,
				Zone:   result.IotZone,
			}

			err = l.lights.ActionHandler(action)
			if err != nil {
				log.Error(err)
			}
		}
	}

	return nil
}

func (l *Server) handleZigbeeDevices(m iot.ZigbeeBridgeMessageDevices) error {
	now := time.Now()

	for _, d := range m {
		x := &inventory.ZigbeeDevice{
			Name:     d.FriendlyName,
			LastSeen: timestamppb.New(now),
			// IeeeAddr:        d.IeeeAddr,
			Type:            d.Type,
			SoftwareBuildId: d.SoftwareBuildID,
			DateCode:        d.DateCode,
			Model:           d.Model,
			Vendor:          d.Vendor,
			// Description      : d.Description,
			ManufacturerName: d.ManufacturerName,
			PowerSource:      d.PowerSource,
			ModelId:          d.ModelID,
			// HardwareVersion:  d.HardwareVersion,
		}

		if x.Name == "Coordinator" {
			continue
		}

		_, err := l.inventory.FetchZigbeeDevice(x.Name)
		if err != nil {
			log.Error(err)
			createResult, err := l.inventory.CreateZigbeeDevice(x)
			if err != nil {
				return err
			}

			log.WithFields(log.Fields{
				"name":   createResult.Name,
				"vendor": createResult.Vendor,
				"model":  createResult.Model,
				"zone":   createResult.IotZone,
			}).Debug("createResult")
		}
	}

	return nil
}

func (l *Server) handleZigbeeDeviceUpdate(m iot.ZigbeeBridgeLog) error {
	// zigbee2mqtt/bridge/request/device/ota_update/update
	log.WithFields(log.Fields{
		"device": m.Meta["device"],
		"status": m.Meta["status"],
	}).Debug("upgrade report")

	req := &iot.UpdateRequest{
		Device: m.Meta["device"].(string),
	}

	go func() {
		_, err := l.iotServer.UpdateDevice(context.Background(), req)
		if err != nil {
			log.Error()
		}
	}()

	return nil
}

func (l *Server) handleLEDReport(request *inventory.IOTDevice) error {
	if request == nil {
		return fmt.Errorf("unable to read led report from nil request")
	}

	discovery := request.DeviceDiscovery

	msg, err := iot.ReadMessage("led", discovery.Message, discovery.Endpoint...)
	if err != nil {
		return err
	}

	if msg != nil {
		m := msg.(iot.LEDConfig)

		for i, deviceConnection := range m.Device.Connections {
			if len(deviceConnection) == 2 {
				l.storeThingLabel(discovery.NodeId, "mac", m.Device.Connections[i][1])
			}
		}
	}

	return nil
}

func (l *Server) handleWaterReport(request *inventory.IOTDevice) error {
	if request == nil {
		return fmt.Errorf("unable to read water report from nil request")
	}

	discovery := request.DeviceDiscovery

	msg, err := iot.ReadMessage("water", discovery.Message, discovery.Endpoint...)
	if err != nil {
		return err
	}

	if msg != nil {
		m := msg.(iot.WaterMessage)

		if m.Temperature != nil {
			waterTemperature.WithLabelValues(discovery.NodeId).Set(float64(*m.Temperature))
		}
	}

	return nil
}

func (l *Server) handleAirReport(request *inventory.IOTDevice) error {
	if request == nil {
		return fmt.Errorf("unable to read air report from nil request")
	}

	discovery := request.DeviceDiscovery

	msg, err := iot.ReadMessage("air", discovery.Message, discovery.Endpoint...)
	if err != nil {
		return err
	}

	if msg != nil {
		m := msg.(iot.AirMessage)

		// l.storeThingLabel(discovery.NodeId, "tempcoef", m.TempCoef)

		if m.Temperature != nil {
			airTemperature.WithLabelValues(discovery.NodeId).Set(float64(*m.Temperature))
		}

		if m.Humidity != nil {
			airHumidity.WithLabelValues(discovery.NodeId).Set(float64(*m.Humidity))
		}
		if m.HeatIndex != nil {
			airHeatindex.WithLabelValues(discovery.NodeId).Set(float64(*m.HeatIndex))
		}
	}

	return nil
}

func (l *Server) handleWifiReport(request *inventory.IOTDevice) error {
	if request == nil {
		return fmt.Errorf("unable to read wifi report from nil request")
	}

	discovery := request.DeviceDiscovery

	msg, err := iot.ReadMessage("wifi", discovery.Message, discovery.Endpoint...)
	if err != nil {
		return err
	}

	if msg != nil {
		m := msg.(iot.WifiMessage)

		l.storeThingLabel(discovery.NodeId, "ssid", m.SSID)
		l.storeThingLabel(discovery.NodeId, "bssid", m.BSSID)
		l.storeThingLabel(discovery.NodeId, "ip", m.IP)

		labels := l.nodeLabels(discovery.NodeId)

		if l.hasLabels(discovery.NodeId, []string{"ssid", "bssid", "ip"}) {
			if m.RSSI != 0 {
				thingWireless.With(prometheus.Labels{
					"device": discovery.NodeId,
					"ssid":   labels["ssid"],
					"bssid":  labels["ssid"],
					"ip":     labels["ip"],
				}).Set(float64(m.RSSI))
			}
		}
	}

	return nil
}
