"""Functions to prevent a nuclear meltdown."""


def is_criticality_balanced(temperature, neutrons_emitted):
   return 800>temperature and neutrons_emitted>500 and temperature * neutrons_emitted < 500000

def reactor_efficiency(voltage, current, theoretical_max_power):
    per = (( voltage * current ) / theoretical_max_power ) * 100
    if 80<=per:
        return "green"
    elif 60<=per and per<80:
        return "orange"
    elif 30<=per and per<80:
        return "red"
    else:
        return "black"

def fail_safe(temperature, neutrons_produced_per_second, threshold):
    power = temperature * neutrons_produced_per_second

    if power < 0.9 * threshold:
        return "LOW"
    elif 0.9 * threshold <= power <= 1.1 * threshold:
        return "NORMAL"
    else:
        return "DANGER"
