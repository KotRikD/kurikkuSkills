package osuSkills

import (
	"errors"

	"github.com/KotRikD/kurikkuSkills/structs"
)

/*
#cgo CFLAGS: -I .
#cgo LDFLAGS: -L .
#include "osuSkills.h"
#include <stdio.h>
#include <stdlib.h>

typedef struct CalculationResult (*FPNTR)(char* filepath, int mods);
typedef int (*FPNTR2)(void);

int bridgeFormulaVars(FPNTR2 f)
{
	return f();
}

struct CalculationResult bridgeBiCycleCalculateBeatmapSkills(FPNTR f, char* filepath, int mods)
{
	return f(filepath, mods);
}
*/
import "C"

type CalculationResult struct {
	FilePath       string
	Circles        int
	SliderSpinners int
	Mods           int
	Name           string
	AR             float64
	CS             float64
	Stamina        float64
	Tenacity       float64
	Agility        float64
	Precision      float64
	Reading        float64
	Memory         float64
	Accuracy       float64
	Reaction       float64
}

// LoadVars function which loading config.cfg for default osuSkills values
func LoadVars() {
	f := C.FPNTR2(C.ReloadFormulaVars)
	C.bridgeFormulaVars(f)
}

// CalculateSkills calculating beatmap skills
func CalculateSkills(path string, mods int) (CalculationResult, error) {
	filepath := C.CString(path)
	f2 := C.FPNTR(C.BiCycleCalculateBeatmapSkills)
	calculation := C.bridgeBiCycleCalculateBeatmapSkills(f2, filepath, C.int(mods))
	if calculation.failed {
		return CalculationResult{}, errors.New("filepath")
	}

	// Disable some values for RX/AP score
	if mods&structs.RL > 0 {
		calculation.tenacity = 0
		calculation.stamina = 0
		calculation.accuracy /= 2
	} else if mods&structs.AP > 0 {
		calculation.precision = 0
		calculation.agility = 0
		calculation.accuracy /= 2
	}

	return CalculationResult{
		C.GoString(calculation.filepath),
		int(calculation.circles),
		int(calculation.sliderspinners),
		int(calculation.mods),
		C.GoString(calculation.name),
		float64(calculation.ar),
		float64(calculation.cs),
		float64(calculation.stamina),
		float64(calculation.tenacity),
		float64(calculation.agility),
		float64(calculation.precision),
		float64(calculation.reading),
		float64(calculation.memory),
		float64(calculation.accuracy),
		float64(calculation.reaction),
	}, nil
}
