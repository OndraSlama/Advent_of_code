from copy import deepcopy
from input import cranes as cranes
from input import instructions as instructions

cranes_first = deepcopy(cranes)


def move_crane(cranes, instruction, number_of_cranes = 1) :
    from_column, to_column = instruction
    cranes_to_move = [cranes[from_column - 1].pop(0) for _ in range(number_of_cranes)]
    cranes[to_column - 1] = cranes_to_move + cranes[to_column - 1]

def get_top_cranes(cranes) -> str:
    return "".join([crane[0] for crane in cranes])

for instruction in instructions:
    for _ in range(instruction[0]):
        move_crane(cranes_first, instruction[1])


print(f"First part: {get_top_cranes(cranes_first)}")

cranes_second = deepcopy(cranes)

for instruction in instructions:
    move_crane(cranes_second, instruction[1], number_of_cranes=instruction[0])

print(f"Second part: {get_top_cranes(cranes_second)}")

    