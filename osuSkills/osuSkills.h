#include <stdbool.h>
#ifdef __cplusplus
extern "C" {
#endif

struct CalculationResult
{
    char* filepath;
    int circles;
    int sliderspinners;
    int mods;
    char* name;
    double ar;
    double cs;
    double stamina;
    double tenacity;
    double agility;
    double precision;
    double reading;
    double memory;
    double accuracy;
    double reaction;
    bool failed;
};

int ReloadFormulaVars();
struct CalculationResult BiCycleCalculateBeatmapSkills(char* filepath, int mods);

#ifdef __cplusplus
}
#endif