from input import stream as stream

def get_proccesed_markers_before_unique_group(unique_markers: int) -> int:
    for i in range(unique_markers, len(stream)):
        char_group = stream[i - unique_markers:i]
        if len(set(char_group)) == unique_markers:
            return i

print(f"First part: {get_proccesed_markers_before_unique_group(4)}")

print(f"Second part: {get_proccesed_markers_before_unique_group(14)}")

    