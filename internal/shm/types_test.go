package shm

import (
	"testing"
	"unsafe"
)

func TestStructSizes(t *testing.T) {
	tests := []struct {
		name string
		got  uintptr
		want uintptr
	}{
		{"Position", unsafe.Sizeof(Position{}), 8},
		{"UnitFinderEntry", unsafe.Sizeof(UnitFinderEntry{}), 8},
		{"ForceData", unsafe.Sizeof(ForceData{}), 32},
		{"RegionData", unsafe.Sizeof(RegionData{}), 1068},
		{"BulletData", unsafe.Sizeof(BulletData{}), 80},
		{"PlayerData", unsafe.Sizeof(PlayerData{}), 5788},
		{"UnitData", unsafe.Sizeof(UnitData{}), 336},
		{"Event", unsafe.Sizeof(Event{}), 12},
		{"Command", unsafe.Sizeof(Command{}), 12},
		{"Shape", unsafe.Sizeof(Shape{}), 40},
		{"UnitCommand", unsafe.Sizeof(UnitCommand{}), 24},
		{"GameInstance", unsafe.Sizeof(GameInstance{}), 12},
		{"GameTable", unsafe.Sizeof(GameTable{}), 96},
		{"GameData", unsafe.Sizeof(GameData{}), 33_017_048},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.got != tt.want {
				t.Errorf("sizeof(%s) = %d, want %d", tt.name, tt.got, tt.want)
			}
		})
	}
}
