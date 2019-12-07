import sys

def _find_path_to_node(orbits, node_name, current_node, current_path):
    if node_name == current_node:
        return current_path
    if current_node not in orbits:
        return None
    for node in orbits[current_node]:
        local_path = current_path + [node]
        path = _find_path_to_node(orbits, node_name, node, local_path)
        if path is not None:
            return path
    return None

def find_path_to_node(orbits, node_name):
    start = 'COM'
    path = [start]
    if node_name == start:
        return path
    for node in orbits[start]:
        local_path = path + [node]
        path = _find_path_to_node(orbits, node_name, node, local_path)
        if path is not None:
            return path
    return None


def main(filename):
    with open(filename, 'r') as f:
        orbits = {}
        lines = f.readlines()

        for line in lines:
            center = line.split(')')[0]
            orbits[center] = []

        nodes = []
        for line in lines:
            center, orbital = [x.strip() for x in line.split(')')]
            nodes += [center, orbital]
            orbits[center] += [orbital]
        
        orbit_counter = 0
        for node in set(nodes):
            path = find_path_to_node(orbits, node)
            if path is None:
                print('Graph not connected. No path available between {} and {}'.format('COM', node))
            orbit_counter += len(path) - 1
        
        print(orbit_counter)

        

if __name__ == "__main__":
    main(sys.argv[1])