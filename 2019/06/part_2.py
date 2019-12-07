import sys

def _find_path_to_node(orbits, node_name, current_node, current_path, visited_nodes):
    if node_name == current_node:
        return current_path
    if current_node not in orbits:
        return None
    for node in orbits[current_node]:
        if node in visited_nodes:
            continue
        local_path = current_path + [node]
        local_visited = visited_nodes + [node]
        path = _find_path_to_node(orbits, node_name, node, local_path, local_visited)
        if path is not None:
            return path
    return None

def find_path_to_node(orbits, node_name):
    start = 'YOU'
    path = [start]
    if node_name == start:
        return path
    for node in orbits[start]:
        local_path = path + [node]
        path = _find_path_to_node(orbits, node_name, node, local_path, [start, node])
        if path is not None:
            return path
    return None


def main(filename):
    with open(filename, 'r') as f:
        orbits = {}
        lines = f.readlines()

        for line in lines:
            center, orbital = [x.strip() for x in line.split(')')]
            orbits[center] = []
            orbits[orbital] = []

        nodes = []
        for line in lines:
            center, orbital = [x.strip() for x in line.split(')')]
            nodes += [center, orbital]
            orbits[center] += [orbital]
            orbits[orbital] += [center]
        

        path = find_path_to_node(orbits, 'SAN')
        if path is None:
            print('Graph not connected. No path available between {} and {}'.format('YOU', 'SAN'))
            return
        print(len(path) - 3)

if __name__ == "__main__":
    main(sys.argv[1])