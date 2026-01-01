package homeAssistant

import (
	"encoding/hex"
	"fmt"
)

func (ha *HomeAssistanceMqttClient) getPositionTopicName(devicePort byte) string {
	return fmt.Sprintf("%s/%d/position", ha.mqttBaseTopic, devicePort)
}

func (ha *HomeAssistanceMqttClient) getSetPositionTopic(devicePort byte) string {
	return fmt.Sprintf("%s/cmnd/%d/position", ha.mqttBaseTopic, devicePort)
}

func (ha *HomeAssistanceMqttClient) getDirectionTopicName(devicePort byte) string {
	return fmt.Sprintf("%s/%d/direction", ha.mqttBaseTopic, devicePort)
}

func (ha *HomeAssistanceMqttClient) getGetStateTopicName(devicePort byte) string {
	return fmt.Sprintf("%s/%d/state", ha.mqttBaseTopic, devicePort)
}

func (ha *HomeAssistanceMqttClient) getDiscoveryTopic(devicePort byte) string {
	//<discovery_prefix>/<component>/[<node_id>/]<object_id>/config
	entity_type := "cover"
	if !ha.doorStatusSupported {
		entity_type = "button"
	}
	return fmt.Sprintf("homeassistant/%s/halsecur/%s/config", entity_type, ha.getUniqueObjectId(devicePort))
}

func (ha *HomeAssistanceMqttClient) getUniqueObjectId(devicePort byte) string {
	deviceMacStr := hex.EncodeToString(ha.deviceMac[:])
	return fmt.Sprintf("%s%d", deviceMacStr, devicePort)
}

func (ha *HomeAssistanceMqttClient) getAvailabilityTopic(devicePort byte) string {
	return fmt.Sprintf("%s/%d/availability", ha.mqttBaseTopic, devicePort)
}
