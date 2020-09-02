package lights

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/amimof/huego"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	log "github.com/sirupsen/logrus"
	"github.com/xaque208/rftoy/rftoy"

	"github.com/xaque208/znet/internal/inventory"
	"github.com/xaque208/znet/internal/timer"
	"github.com/xaque208/znet/pkg/events"
	"github.com/xaque208/znet/pkg/iot"
)

// Lights holds the information necessary to communicate with lighting
// equipment, and the configuration to add a bit of context.
type Lights struct {
	RFToy      *rftoy.RFToy
	HUE        *huego.Bridge
	inv        *inventory.Inventory
	config     Config
	mqttClient mqtt.Client
}

// NewLights creates and returns a new Lights object based on the received
// configuration.
func NewLights(config Config, inv *inventory.Inventory, mqttClient mqtt.Client) *Lights {
	return &Lights{
		HUE:        huego.New(config.Hue.Endpoint, config.Hue.User),
		RFToy:      &rftoy.RFToy{Address: config.RFToy.Endpoint},
		inv:        inv,
		config:     config,
		mqttClient: mqttClient,
	}
}

// Subscriptions returns the data for mapping event names with functions.
func (l *Lights) Subscriptions() map[string][]events.Handler {
	s := events.NewSubscriptions()

	namedEvents := []string{
		"PreSunset",
		"SolarEvent",
		"TimerExpired",
		"NamedTimer",
	}

	for _, e := range namedEvents {
		s.Subscribe(e, l.eventHandler)
	}

	s.Subscribe("Click", l.clickHandler)

	return s.Table
}

func (l *Lights) clickHandler(name string, payload events.Payload) error {
	log.Tracef("Lights.clickHandler: %s : %+v", name, string(payload))

	var e iot.Click

	err := json.Unmarshal(payload, &e)
	if err != nil {
		return fmt.Errorf("failed to unmarshal %T: %s", e, err)
	}

	for _, room := range l.config.Rooms {
		alert := false
		dim := false
		off := false
		on := false
		toggle := false

		if room.Name == e.Zone {
			log.WithFields(log.Fields{
				"room_name": room.Name,
				"name":      e.Zone,
			}).Trace("using room")

			switch e.Count {
			case "single":
				off = true
			case "double":
				on = true
			case "triple":
				toggle = true
			case "long":
				dim = true
			case "many":
				alert = true
			default:
				log.Warnf("e: %+v", e)
			}

			if toggle {
				l.Toggle(room.Name)
			}

			if off {
				l.Off(room.Name)
			}

			if on {
				l.On(room.Name)
			}

			if alert {
				l.Alert(room.Name)
			}

			if dim {
				l.Dim(room.Name, 100)
			}
		}
	}

	return nil
}

func (l *Lights) eventHandler(name string, payload events.Payload) error {
	log.Tracef("Lights.eventHandler: %s : %+v", name, string(payload))

	var e timer.NamedTimer

	err := json.Unmarshal(payload, &e)
	if err != nil {
		return fmt.Errorf("failed to unmarshal %T: %s", e, err)
	}

	for _, room := range l.config.Rooms {
		for _, o := range room.On {
			if o == e.Name {
				l.On(room.Name)
			}
		}

		for _, o := range room.Off {
			if o == e.Name {
				l.Off(room.Name)
			}
		}

		for _, o := range room.Dim {
			if o == e.Name {
				l.Dim(room.Name, 100)
			}
		}

		for _, o := range room.Alert {
			if o == e.Name {
				l.Alert(room.Name)
			}
		}
	}

	return nil
}

// getLight calls the Hue bridge and looks for a light, that, when normalized,
// matches the name received.
func (l *Lights) getLight(lightName string) (*huego.Light, error) {
	lights, err := l.HUE.GetLights()
	if err != nil {
		log.Error(err)
	}

	log.Tracef("lights: %+v", lights)

	for _, g := range lights {
		flatName := strings.ToLower(strings.ReplaceAll(g.Name, " ", "_"))

		if lightName == flatName {
			return &g, nil
		}

	}

	return &huego.Light{}, fmt.Errorf("light %s not found", lightName)
}

// GetGroup calls the Hue bridge and looks for a group, that, when normalized,
// matches the name received.
func (l *Lights) getGroup(groupName string) (*huego.Group, error) {
	groups, err := l.HUE.GetGroups()
	if err != nil {
		log.Error(err)
	}

	log.Tracef("found HUE groups: %+v", groups)

	for _, g := range groups {
		flatName := strings.ToLower(strings.ReplaceAll(g.Name, " ", "_"))

		if groupName == flatName {
			return &g, nil
		}
	}

	return &huego.Group{}, fmt.Errorf("group %s not found", groupName)
}

func (l *Lights) Toggle(groupName string) error {

	log.Debugf("toggle: %s", groupName)

	result, err := l.inv.ListZigbeeDevices()
	if err != nil {
		return err
	}

	if result != nil {
		log.Debugf("result: %s", *result)
		for _, d := range *result {
			if d.IotZone != groupName {
				continue
			}

			log.Debugf("match: %s", d)

			topic := fmt.Sprintf("zigbee2mqtt/%s/set", d.Name)
			message := map[string]string{
				"state": "TOGGLE",
			}

			m, err := json.Marshal(message)
			if err != nil {
				log.Error(err)
				continue
			}

			l.mqttClient.Publish(topic, byte(0), false, string(m))

		}
	}

	return nil
}

// On turns off the Hue lights for a room.
func (l *Lights) On(groupName string) {
	room, err := l.config.Room(groupName)
	if err != nil {
		log.Error(err)
	}

	g, err := l.getGroup(groupName)
	if err != nil {
		log.Error(err)
		var light *huego.Light

		light, err = l.getLight(groupName)
		if err != nil {
			log.Error(err)
		} else {
			log.WithFields(log.Fields{
				"name": light.Name,
			}).Debug("turning on light")

			err = light.On()
			if err != nil {
				log.Error(err)
			}
		}
	} else {
		log.WithFields(log.Fields{
			"group": g.Name,
		}).Debug("turning on light group")

		err = g.On()
		if err != nil {
			log.Error(err)
		}
	}

	if len(room.IDs) > 0 {
		log.WithFields(log.Fields{
			"ids": room.IDs,
		}).Debug("turning on rftoy ids")

		for _, i := range room.IDs {
			err := l.RFToy.On(i)
			if err != nil {
				log.Error(err)
			}
		}
	}
}

// Off turns off the Hue lights for a room.
func (l *Lights) Off(groupName string) {
	room, err := l.config.Room(groupName)
	if err != nil {
		log.Error(err)
	}

	// try the light by group first
	g, err := l.getGroup(groupName)
	if err != nil {
		log.Error(err)
		var light *huego.Light

		// then try to get just the light
		light, err = l.getLight(groupName)
		if err != nil {
			log.Error(err)
		} else {
			log.WithFields(log.Fields{
				"name": light.Name,
			}).Debug("turning off light")

			err = light.Off()
			if err != nil {
				log.Error(err)
			}
		}

	} else {
		log.WithFields(log.Fields{
			"group": g.Name,
		}).Debug("turning off light group")

		err = g.Off()
		if err != nil {
			log.Error(err)
		}
	}

	if len(room.IDs) > 0 {
		log.WithFields(log.Fields{
			"ids": room.IDs,
		}).Debug("turning off rftoy ids")

		for _, i := range room.IDs {
			err := l.RFToy.Off(i)
			if err != nil {
				log.Error(err)
			}
		}
	}
}

// Dim modifies the brightness of a light group.
func (l *Lights) Dim(groupName string, brightness int32) {
	room, err := l.config.Room(groupName)
	if err != nil {
		log.Error(err)
	}

	groups, err := l.HUE.GetGroups()
	if err != nil {
		log.Error(err)
	}

	for _, g := range groups {
		for _, i := range room.HueIDs {
			if g.ID == i {
				log.WithFields(log.Fields{
					"name":  g.Name,
					"state": g.State,
				}).Debug("setting group brightness")

				err := g.Bri(uint8(brightness))
				if err != nil {
					log.Error(err)
				}
			}
		}
	}
}

// Alert blinks all lights in the given light group.
func (l *Lights) Alert(groupName string) {
	g, err := l.getGroup(groupName)
	if err != nil {
		log.Error(err)

		// then try to get just the light
		light, err := l.getLight(groupName)
		if err != nil {
			log.Error(err)
		} else {
			log.WithFields(log.Fields{
				"name": light.Name,
			}).Debug("alerting light")

			err := light.Alert("select")
			if err != nil {
				log.Error(err)
			}
		}

	} else {
		log.WithFields(log.Fields{
			"group": g.Name,
		}).Debug("alerting light group")

		err := g.Alert("select")
		if err != nil {
			log.Error(err)
		}
	}
}
