package main

import (
	_ "github.com/octoposprime/op-go-os/pkg/notify"
	object "github.com/octoposprime/op-go-os/pkg/object"
	_ "github.com/octoposprime/op-go-os/tool/logger"
	logger "github.com/octoposprime/op-go-os/tool/logger"
	pb_dlr "github.com/octoposprime/op-proto/pkg/proto/pb/dlr"
	pb_led "github.com/octoposprime/op-proto/pkg/proto/pb/led"
	pb_tick "github.com/octoposprime/op-proto/pkg/proto/pb/tick"
)

func init() {
	logger.GetLoggerInstance().SetDebug(true)
}

func main() {
	dlrPb := GetDlr()
	object.NewDlr(dlrPb.Body.Key.Id)
}

func GetDlr() pb_dlr.Dlr {

	dlrHeader := pb_dlr.DlrHeader{
		Id:         "dlr_header_1",
		IsRepeated: false,
		HookId:     nil,
	}

	dlrKeyTick := pb_dlr.DlrKey_Tick{
		Tick: &pb_tick.Tick{},
	}

	dlrKey := pb_dlr.DlrKey{
		Id:  "dlr_key_1",
		Key: &dlrKeyTick,
	}

	components := make([]*pb_dlr.DlrComponent, 7)
	led1 := pb_led.Led{}
	led2 := pb_led.Led{}
	led3 := pb_led.Led{}
	led4 := pb_led.Led{}
	led5 := pb_led.Led{}
	led6 := pb_led.Led{}
	led7 := pb_led.Led{}

	// Components
	components[0] = &pb_dlr.DlrComponent{
		Id: "dlr_component_1",
		Component: &pb_dlr.DlrComponent_Led{
			Led: &led1,
		},
	}
	components[1] = &pb_dlr.DlrComponent{
		Id: "dlr_component_2",
		Component: &pb_dlr.DlrComponent_Led{
			Led: &led2,
		},
	}
	components[2] = &pb_dlr.DlrComponent{
		Id: "dlr_component_3",
		Component: &pb_dlr.DlrComponent_Led{
			Led: &led3,
		},
	}
	components[3] = &pb_dlr.DlrComponent{
		Id: "dlr_component_4",
		Component: &pb_dlr.DlrComponent_Led{
			Led: &led4,
		},
	}
	components[4] = &pb_dlr.DlrComponent{
		Id: "dlr_component_5",
		Component: &pb_dlr.DlrComponent_Led{
			Led: &led5,
		},
	}
	components[5] = &pb_dlr.DlrComponent{
		Id: "dlr_component_6",
		Component: &pb_dlr.DlrComponent_Led{
			Led: &led6,
		},
	}
	components[6] = &pb_dlr.DlrComponent{
		Id: "dlr_component_7",
		Component: &pb_dlr.DlrComponent_Led{
			Led: &led7,
		},
	}
	// End Components

	conditions := make([]*pb_dlr.DlrCondition, 5)

	conditionComponent0 := make([]*pb_dlr.DlrConditionComponent, 7)
	conditionComponent1 := make([]*pb_dlr.DlrConditionComponent, 7)
	conditionComponent2 := make([]*pb_dlr.DlrConditionComponent, 7)
	conditionComponent3 := make([]*pb_dlr.DlrConditionComponent, 7)
	conditionComponent4 := make([]*pb_dlr.DlrConditionComponent, 7)

	// Condition Component 0
	conditionComponent0[0] = &pb_dlr.DlrConditionComponent{
		Id: "dlr_condition_component_1",
		ComponentState: &pb_dlr.DlrConditionComponent_LedState{
			LedState: pb_led.LedState_LED_STATE_ON,
		},
	}
	conditionComponent0[1] = &pb_dlr.DlrConditionComponent{
		Id: "dlr_condition_component_2",
		ComponentState: &pb_dlr.DlrConditionComponent_LedState{
			LedState: pb_led.LedState_LED_STATE_ON,
		},
	}
	conditionComponent0[2] = &pb_dlr.DlrConditionComponent{
		Id: "dlr_condition_component_3",
		ComponentState: &pb_dlr.DlrConditionComponent_LedState{
			LedState: pb_led.LedState_LED_STATE_ON,
		},
	}
	conditionComponent0[3] = &pb_dlr.DlrConditionComponent{
		Id: "dlr_condition_component_4",
		ComponentState: &pb_dlr.DlrConditionComponent_LedState{
			LedState: pb_led.LedState_LED_STATE_OFF,
		},
	}
	conditionComponent0[4] = &pb_dlr.DlrConditionComponent{
		Id: "dlr_condition_component_5",
		ComponentState: &pb_dlr.DlrConditionComponent_LedState{
			LedState: pb_led.LedState_LED_STATE_OFF,
		},
	}
	conditionComponent0[5] = &pb_dlr.DlrConditionComponent{
		Id: "dlr_condition_component_6",
		ComponentState: &pb_dlr.DlrConditionComponent_LedState{
			LedState: pb_led.LedState_LED_STATE_OFF,
		},
	}
	conditionComponent0[6] = &pb_dlr.DlrConditionComponent{
		Id: "dlr_condition_component_7",
		ComponentState: &pb_dlr.DlrConditionComponent_LedState{
			LedState: pb_led.LedState_LED_STATE_OFF,
		},
	}
	// End Condition Component 0
	// Condition Component 1
	conditionComponent1[0] = &pb_dlr.DlrConditionComponent{
		Id: "dlr_condition_component_1",
		ComponentState: &pb_dlr.DlrConditionComponent_LedState{
			LedState: pb_led.LedState_LED_STATE_OFF,
		},
	}
	conditionComponent1[1] = &pb_dlr.DlrConditionComponent{
		Id: "dlr_condition_component_2",
		ComponentState: &pb_dlr.DlrConditionComponent_LedState{
			LedState: pb_led.LedState_LED_STATE_OFF,
		},
	}
	conditionComponent1[2] = &pb_dlr.DlrConditionComponent{
		Id: "dlr_condition_component_3",
		ComponentState: &pb_dlr.DlrConditionComponent_LedState{
			LedState: pb_led.LedState_LED_STATE_ON,
		},
	}
	conditionComponent1[3] = &pb_dlr.DlrConditionComponent{
		Id: "dlr_condition_component_4",
		ComponentState: &pb_dlr.DlrConditionComponent_LedState{
			LedState: pb_led.LedState_LED_STATE_ON,
		},
	}
	conditionComponent1[4] = &pb_dlr.DlrConditionComponent{
		Id: "dlr_condition_component_5",
		ComponentState: &pb_dlr.DlrConditionComponent_LedState{
			LedState: pb_led.LedState_LED_STATE_ON,
		},
	}
	conditionComponent1[5] = &pb_dlr.DlrConditionComponent{
		Id: "dlr_condition_component_6",
		ComponentState: &pb_dlr.DlrConditionComponent_LedState{
			LedState: pb_led.LedState_LED_STATE_OFF,
		},
	}
	conditionComponent1[6] = &pb_dlr.DlrConditionComponent{
		Id: "dlr_condition_component_7",
		ComponentState: &pb_dlr.DlrConditionComponent_LedState{
			LedState: pb_led.LedState_LED_STATE_OFF,
		},
	}
	// End Condition Component 1
	// Condition Component 2
	conditionComponent2[0] = &pb_dlr.DlrConditionComponent{
		Id: "dlr_condition_component_1",
		ComponentState: &pb_dlr.DlrConditionComponent_LedState{
			LedState: pb_led.LedState_LED_STATE_OFF,
		},
	}
	conditionComponent2[1] = &pb_dlr.DlrConditionComponent{
		Id: "dlr_condition_component_2",
		ComponentState: &pb_dlr.DlrConditionComponent_LedState{
			LedState: pb_led.LedState_LED_STATE_OFF,
		},
	}
	conditionComponent2[2] = &pb_dlr.DlrConditionComponent{
		Id: "dlr_condition_component_3",
		ComponentState: &pb_dlr.DlrConditionComponent_LedState{
			LedState: pb_led.LedState_LED_STATE_OFF,
		},
	}
	conditionComponent2[3] = &pb_dlr.DlrConditionComponent{
		Id: "dlr_condition_component_4",
		ComponentState: &pb_dlr.DlrConditionComponent_LedState{
			LedState: pb_led.LedState_LED_STATE_OFF,
		},
	}
	conditionComponent2[4] = &pb_dlr.DlrConditionComponent{
		Id: "dlr_condition_component_5",
		ComponentState: &pb_dlr.DlrConditionComponent_LedState{
			LedState: pb_led.LedState_LED_STATE_ON,
		},
	}
	conditionComponent2[5] = &pb_dlr.DlrConditionComponent{
		Id: "dlr_condition_component_6",
		ComponentState: &pb_dlr.DlrConditionComponent_LedState{
			LedState: pb_led.LedState_LED_STATE_ON,
		},
	}
	conditionComponent2[6] = &pb_dlr.DlrConditionComponent{
		Id: "dlr_condition_component_7",
		ComponentState: &pb_dlr.DlrConditionComponent_LedState{
			LedState: pb_led.LedState_LED_STATE_ON,
		},
	}
	// End Condition Component 2
	// Condition Component 3
	conditionComponent3[0] = &pb_dlr.DlrConditionComponent{
		Id: "dlr_condition_component_1",
		ComponentState: &pb_dlr.DlrConditionComponent_LedState{
			LedState: pb_led.LedState_LED_STATE_ON,
		},
	}
	conditionComponent3[1] = &pb_dlr.DlrConditionComponent{
		Id: "dlr_condition_component_2",
		ComponentState: &pb_dlr.DlrConditionComponent_LedState{
			LedState: pb_led.LedState_LED_STATE_ON,
		},
	}
	conditionComponent3[2] = &pb_dlr.DlrConditionComponent{
		Id: "dlr_condition_component_3",
		ComponentState: &pb_dlr.DlrConditionComponent_LedState{
			LedState: pb_led.LedState_LED_STATE_OFF,
		},
	}
	conditionComponent3[3] = &pb_dlr.DlrConditionComponent{
		Id: "dlr_condition_component_4",
		ComponentState: &pb_dlr.DlrConditionComponent_LedState{
			LedState: pb_led.LedState_LED_STATE_OFF,
		},
	}
	conditionComponent3[4] = &pb_dlr.DlrConditionComponent{
		Id: "dlr_condition_component_5",
		ComponentState: &pb_dlr.DlrConditionComponent_LedState{
			LedState: pb_led.LedState_LED_STATE_OFF,
		},
	}
	conditionComponent3[5] = &pb_dlr.DlrConditionComponent{
		Id: "dlr_condition_component_6",
		ComponentState: &pb_dlr.DlrConditionComponent_LedState{
			LedState: pb_led.LedState_LED_STATE_OFF,
		},
	}
	conditionComponent3[6] = &pb_dlr.DlrConditionComponent{
		Id: "dlr_condition_component_7",
		ComponentState: &pb_dlr.DlrConditionComponent_LedState{
			LedState: pb_led.LedState_LED_STATE_ON,
		},
	}
	// End Condition Component 3
	// Condition Component 4
	conditionComponent4[0] = &pb_dlr.DlrConditionComponent{
		Id: "dlr_condition_component_1",
		ComponentState: &pb_dlr.DlrConditionComponent_LedState{
			LedState: pb_led.LedState_LED_STATE_OFF,
		},
	}
	conditionComponent4[1] = &pb_dlr.DlrConditionComponent{
		Id: "dlr_condition_component_2",
		ComponentState: &pb_dlr.DlrConditionComponent_LedState{
			LedState: pb_led.LedState_LED_STATE_ON,
		},
	}
	conditionComponent4[2] = &pb_dlr.DlrConditionComponent{
		Id: "dlr_condition_component_3",
		ComponentState: &pb_dlr.DlrConditionComponent_LedState{
			LedState: pb_led.LedState_LED_STATE_ON,
		},
	}
	conditionComponent4[3] = &pb_dlr.DlrConditionComponent{
		Id: "dlr_condition_component_4",
		ComponentState: &pb_dlr.DlrConditionComponent_LedState{
			LedState: pb_led.LedState_LED_STATE_ON,
		},
	}
	conditionComponent4[4] = &pb_dlr.DlrConditionComponent{
		Id: "dlr_condition_component_5",
		ComponentState: &pb_dlr.DlrConditionComponent_LedState{
			LedState: pb_led.LedState_LED_STATE_OFF,
		},
	}
	conditionComponent4[5] = &pb_dlr.DlrConditionComponent{
		Id: "dlr_condition_component_6",
		ComponentState: &pb_dlr.DlrConditionComponent_LedState{
			LedState: pb_led.LedState_LED_STATE_OFF,
		},
	}
	conditionComponent4[6] = &pb_dlr.DlrConditionComponent{
		Id: "dlr_condition_component_7",
		ComponentState: &pb_dlr.DlrConditionComponent_LedState{
			LedState: pb_led.LedState_LED_STATE_OFF,
		},
	}
	// End Condition Component 4

	// Conditions
	condition0 := pb_dlr.DlrCondition{
		Id:         "dlr_condition_1",
		Components: conditionComponent0,
		KeyState: &pb_dlr.DlrCondition_TickState{
			TickState: &pb_tick.TickState{
				Value: 1,
				Unit:  pb_tick.TickUnit_TICK_UNIT_SECOND,
			},
		},
	}

	condition1 := pb_dlr.DlrCondition{
		Id:         "dlr_condition_2",
		Components: conditionComponent1,
		KeyState: &pb_dlr.DlrCondition_TickState{
			TickState: &pb_tick.TickState{
				Value: 2,
				Unit:  pb_tick.TickUnit_TICK_UNIT_SECOND,
			},
		},
	}

	condition2 := pb_dlr.DlrCondition{
		Id:         "dlr_condition_3",
		Components: conditionComponent2,
		KeyState: &pb_dlr.DlrCondition_TickState{
			TickState: &pb_tick.TickState{
				Value: 3,
				Unit:  pb_tick.TickUnit_TICK_UNIT_SECOND,
			},
		},
	}

	condition3 := pb_dlr.DlrCondition{
		Id:         "dlr_condition_4",
		Components: conditionComponent3,
		KeyState: &pb_dlr.DlrCondition_TickState{
			TickState: &pb_tick.TickState{
				Value: 4,
				Unit:  pb_tick.TickUnit_TICK_UNIT_SECOND,
			},
		},
	}

	condition4 := pb_dlr.DlrCondition{
		Id:         "dlr_condition_5",
		Components: conditionComponent4,
		KeyState: &pb_dlr.DlrCondition_TickState{
			TickState: &pb_tick.TickState{
				Value: 5,
				Unit:  pb_tick.TickUnit_TICK_UNIT_SECOND,
			},
		},
	}
	// End Conditions

	conditions[0] = &condition0
	conditions[1] = &condition1
	conditions[2] = &condition2
	conditions[3] = &condition3
	conditions[4] = &condition4

	dlrBody := pb_dlr.DlrBody{
		Id:         "dlr_body_1",
		Key:        &dlrKey,
		Components: components,
		Conditions: conditions,
	}

	dlr := pb_dlr.Dlr{
		Id:     "dlr_1",
		Header: &dlrHeader,
		Body:   &dlrBody,
	}

	return dlr
}
